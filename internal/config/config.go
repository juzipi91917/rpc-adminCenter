package config

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/mysqlx"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rabbitmqx"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-clipFilm/internal/model/entity"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"go.etcd.io/etcd/client/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type (
	Config struct {
		zrpc.RpcServerConf
		DBX       string
		RedisX    string
		RabbitMqX string
		AliX      string
		TikTokX   string
	}

	DBX struct {
		Uri string `json:"uri"`
	}

	RedisX struct {
		Addr     string `json:"addr"`
		UserName string `json:"username"`
		PassWord string `json:"password"`
		Db       int    `json:"db"`
	}

	RabbitMqX struct {
		Host        string `json:"host"`
		PassWord    string `json:"password"`
		Port        int    `json:"port"`
		User        string `json:"user"`
		VirtualHost string `json:"virtualHost"`
	}
	Ali struct {
		Cloud struct {
			AccessKeyId     string `json:"AccessKeyId"`
			AccessKeySecret string `json:"AccessKeySecret"`
		}
		Oss struct {
			PublicEndpoint  string `json:"PublicEndpoint"`
			PrivateEndpoint string `json:"PrivateEndpoint"`
			Bucket          string `json:"Bucket"`
		}
		NotifyUrl string `json:"NotifyUrl"`
	}
	Tiktok struct {
		Url               string `json:"url"`
		BillboardMusicUrn string `json:"billboard_music_urn"`
	}
)

var (
	configOnce = sync.Once{}

	dBClient     *gorm.DB
	redisClient  *redis.Client
	etcdClient   *clientv3.Client
	producerPool *rabbitmqx.RabbitPool
	consumerPool *rabbitmqx.RabbitPool
)

func Init() (config *Config) {
	config = new(Config)

	configOnce.Do(func() {
		var (
			filePath string
			err      error
		)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		for {
			select {
			case <-ctx.Done():
				panic("read config yaml file failed")
			case <-time.After(time.Nanosecond):
				env := os.Getenv("environ")
				switch env {
				case service.ProMode:
					filePath, err = filepath.Abs("etc/release.clipfilm.yaml")
					if err != nil {
						panic(err)
					}
				case service.TestMode:
					filePath, err = filepath.Abs("etc/test.clipfilm.yaml")
					if err != nil {
						panic(err)
					}
				default:
					filePath, err = filepath.Abs("etc/clipfilm.yaml")
					if err != nil {
						panic(err)
					}
				}

				conf.MustLoad(filePath, config)

				logx.WithContext(ctx).Infof("successfully set config, path: %s, env: %s", filePath, env)
				registerEtcd(ctx, config)

				return
			}
		}
	})

	return
}

func registerEtcd(ctx context.Context, config *Config) {
	var err error
	etcdClient, err = clientv3.New(clientv3.Config{
		Endpoints:   config.Etcd.Hosts,
		DialTimeout: 5 * time.Second,
		Username:    config.Etcd.User,
		Password:    config.Etcd.Pass,
	})
	if err != nil {
		panic(err)
	}

	if _, err = etcdClient.Get(ctx, "ping"); err != nil {
		panic("failed connect etcd client")
	}
	logx.WithContext(ctx).Info("successfully connect etcd client")
}

func RegisterDBCli(config Config) *gorm.DB {
	return initDB(config.Mode, getDB(config))
}

func RegisterRedisCli(config Config) *redis.Client {
	return initRedis(getRedis(config))
}

func RegisterRabbitMq(config Config, classes ...uint) *rabbitmqx.RabbitPool {
	for _, class := range classes {
		switch class {
		case rabbitmqx.RABBITMQ_TYPE_CONSUME:
			return initConsumer(getRabbitMq(config))
		case rabbitmqx.RABBITMQ_TYPE_PUBLISH:
			return initProducer(getRabbitMq(config))
		}
	}

	return nil
}

func getDB(config Config) (dbx *DBX) {
	dbResp, err := etcdClient.Get(context.Background(), config.DBX)
	if err != nil {
		panic(err)
	}

	dbx = &DBX{}
	if err = json.Unmarshal(dbResp.Kvs[0].Value, dbx); err != nil {
		panic(err)
	}

	return
}

func getRedis(config Config) (redisX *RedisX) {
	redisResp, err := etcdClient.Get(context.Background(), config.RedisX)
	if err != nil {
		panic(err)
	}

	redisX = &RedisX{}
	if err = json.Unmarshal(redisResp.Kvs[0].Value, redisX); err != nil {
		panic(err)
	}

	return
}

func GetAli(config Config) (ali *Ali) {
	aliResp, err := etcdClient.Get(context.Background(), config.AliX)
	if err != nil {
		panic(err)
	}
	ali = &Ali{}
	if err = json.Unmarshal(aliResp.Kvs[0].Value, ali); err != nil {
		panic(err)
	}
	return
}

func GetTiktok(config Config) (tiktok *Tiktok) {
	tiktokResp, err := etcdClient.Get(context.Background(), config.TikTokX)
	if err != nil {
		panic(err)
	}
	tiktok = &Tiktok{}
	if err = json.Unmarshal(tiktokResp.Kvs[0].Value, tiktok); err != nil {
		panic(err)
	}
	return
}

func getRabbitMq(config Config) (rabbitMqX *RabbitMqX) {
	rabbitResp, err := etcdClient.Get(context.Background(), config.RabbitMqX)
	if err != nil {
		panic(err)
	}

	rabbitMqX = &RabbitMqX{}
	if err = json.Unmarshal(rabbitResp.Kvs[0].Value, rabbitMqX); err != nil {
		panic(err)
	}

	return
}

func initProducer(rabbitMqX *RabbitMqX) *rabbitmqx.RabbitPool {
	producerPool = rabbitmqx.NewProductPool()
	if err := producerPool.ConnectVirtualHost(rabbitMqX.Host, rabbitMqX.Port, rabbitMqX.User, rabbitMqX.PassWord, rabbitMqX.VirtualHost); err != nil {
		panic(err)
	}
	logx.Info("rabbitMQ producer pool init successfully")

	return producerPool
}

func initConsumer(rabbitMqX *RabbitMqX) *rabbitmqx.RabbitPool {
	consumerPool = rabbitmqx.NewConsumePool()
	if err := consumerPool.ConnectVirtualHost(rabbitMqX.Host, rabbitMqX.Port, rabbitMqX.User, rabbitMqX.PassWord, rabbitMqX.VirtualHost); err != nil {
		panic(err)
	}
	logx.Info("rabbitMQ consumer pool init successfully")

	return consumerPool
}

func initRedis(redisX *RedisX) *redis.Client {
	redisClient = redis.NewClient(&redis.Options{
		Username: redisX.UserName,
		Addr:     redisX.Addr,
		Password: redisX.PassWord, // 密码
		DB:       redisX.Db,       // 默认数据库
	})
	// 设置连接超时时间

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	logx.Infof("successfully connect redis client")

	return redisClient
}

func initDB(env string, db *DBX) *gorm.DB {
	var err error

	dBClient, err = gorm.Open(mysql.Open(db.Uri), mysqlx.DefaultOpts())
	if err != nil {
		panic(err)
	}
	logx.Info("successfully connect mysql client")

	if env != service.ProMode {
		_ = dBClient.Set("gorm:table_options", "COMMENT='背景音乐表'").AutoMigrate(&entity.Bgm{})
		_ = dBClient.Set("gorm:table_options", "COMMENT='背景音乐类型表'").AutoMigrate(&entity.BgmClass{})
		_ = dBClient.Set("gorm:table_options", "COMMENT='自定义滤镜表'").AutoMigrate(&entity.FilterCustom{})
		_ = dBClient.Set("gorm:table_options", "COMMENT='字体表'").AutoMigrate(&entity.Font{})
		_ = dBClient.Set("gorm:table_options", "COMMENT='花字参数表'").AutoMigrate(&entity.SellingPoint{})
		_ = dBClient.Set("gorm:table_options", "COMMENT='花字背景图表'").AutoMigrate(&entity.SellingPointImage{})
		_ = dBClient.Set("gorm:table_options", "COMMENT='字幕描边表'").AutoMigrate(&entity.SubtitleStyle{})
		_ = dBClient.Set("gorm:table_options", "COMMENT='文案字幕表'").AutoMigrate(&entity.Paperwork{})
		_ = dBClient.Set("gorm:table_options", "COMMENT='违禁词表'").AutoMigrate(&entity.ForbiddenWord{})
		_ = dBClient.Set("gorm:table_options", "COMMENT='操作日志表'").AutoMigrate(&entity.OperateLog{})
		_ = dBClient.Set("gorm:table_options", "COMMENT='多音字替换表'").AutoMigrate(&entity.PolyphonicSubstitution{})
		_ = dBClient.Set("gorm:table_options", "COMMENT='配置表'").AutoMigrate(&entity.Conf{})
		_ = dBClient.Set("gorm:table_options", "COMMENT='视频处理任务表'").AutoMigrate(&entity.Task{})
		_ = dBClient.Set("gorm:table_options", "COMMENT='视频处理请求表'").AutoMigrate(&entity.TaskDetail{})
		_ = dBClient.Set("gorm:table_options", "COMMENT='视频处理记录表'").AutoMigrate(&entity.TaskProcess{})
		_ = dBClient.Set("gorm:table_options", "COMMENT='视频处理文案使用表'").AutoMigrate(&entity.TaskPaperwork{})
		_ = dBClient.Set("gorm:table_options", "COMMENT='视频模板表'").AutoMigrate(&entity.Template{})
		_ = dBClient.Set("gorm:table_options", "COMMENT='用户视频模板表'").AutoMigrate(&entity.UserTemplate{})
		_ = dBClient.Set("gorm:table_options", "COMMENT='用户文案表'").AutoMigrate(&entity.UserPaperwork{})
		_ = dBClient.Set("gorm:table_options", "COMMENT='用户bgm表'").AutoMigrate(&entity.UserBgm{})
		_ = dBClient.Set("gorm:table_options", "COMMENT='发音人表'").AutoMigrate(&entity.SoundPeople{})
	}

	return dBClient
}
