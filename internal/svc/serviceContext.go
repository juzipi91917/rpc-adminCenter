package svc

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rabbitmqx"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/config"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config       config.Config
	DBCli        *gorm.DB
	RedisCli     *redis.Client
	ConsumerPool *rabbitmqx.RabbitPool
	ProducerPool *rabbitmqx.RabbitPool
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		DBCli:        config.RegisterDBCli(c),
		RedisCli:     config.RegisterRedisCli(c),
		ConsumerPool: config.RegisterRabbitMq(c, rabbitmqx.RABBITMQ_TYPE_CONSUME),
		ProducerPool: config.RegisterRabbitMq(c, rabbitmqx.RABBITMQ_TYPE_PUBLISH),
	}
}

func NewJobContext(c config.Config) *ServiceContext {
	config.WatchEtcd(c)
	return &ServiceContext{
		Config:   c,
		DBCli:    config.RegisterDBCli(c),
		RedisCli: config.RegisterRedisCli(c),
	}
}
