package job

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-clipFilm/internal/common"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-clipFilm/internal/svc"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type TikTokBgm struct {
	Ctx    context.Context
	SvcCtx *svc.ServiceContext
}

func registerDetailCountJob(ctx context.Context, svcCtx *svc.ServiceContext) *TikTokBgm {
	return &TikTokBgm{
		Ctx:    ctx,
		SvcCtx: svcCtx,
	}
}

func (t *TikTokBgm) run() error {
	logx.WithContext(t.Ctx).Infof("clipFilm执行定时任务[同步抖音bgm]开始,startTime:%s", time.Now().String())
	for true {
		accessToken := t.SvcCtx.RedisCli.Get(t.Ctx, "tiktok:awf54a1sqduu1v21:client_token:").Val()
		accessToken = "clt.49b6da1d0284ab82ffd53854a2d97fc0Jn1QEnQ4h5iz7jLhrSRcxpK0vAk3"
		if accessToken == "" {
			logx.WithContext(t.Ctx).Errorf("clipFilm执行定时任务[同步抖音bgm],获取抖音access_token为空")
			break
		}
		res, err := TiktokGetBillboardMusic(t.SvcCtx.TiktokConfig.Url+t.SvcCtx.TiktokConfig.BillboardMusicUrn, accessToken)
		if err != nil {
			logx.WithContext(t.Ctx).Errorf("clipFilm执行定时任务[同步抖音bgm],获取抖音bgm出错,err:%+v", err)
			break
		}
		tikTokBgmRequest := TikTokBgmRequest{}
		err = json.Unmarshal(res, &tikTokBgmRequest)
		if err != nil {
			logx.WithContext(t.Ctx).Errorf("clipFilm执行定时任务[同步抖音bgm],json解析获取抖音bgm出错,err:%+v", err)
			break
		}
		if tikTokBgmRequest.Extra.ErrorCode != 0 || tikTokBgmRequest.Extra.Description != "" {
			logx.WithContext(t.Ctx).Errorf("clipFilm执行定时任务[同步抖音bgm],获取到的数据出错,request:%+v", tikTokBgmRequest)
			break
		}
		//bgmDao := dao.NewBgmDao(t.SvcCtx)
		for _, item := range tikTokBgmRequest.Data.List {
			_, path, err := common.UrlAnalysis(item.ShareUrl)
			if err != nil {
				logx.WithContext(t.Ctx).Errorf("clipFilm执行定时任务[同步抖音bgm],解析url失败,url:%s,err:%+v", item.ShareUrl, err)
				continue
			}
			pathSplit := strings.Split(path, "/")
			if len(pathSplit) == 0 {
				logx.WithContext(t.Ctx).Errorf("clipFilm执行定时任务[同步抖音bgm],解析url的path为空,err:%+v", err)
				continue
			}

			err = t.SvcCtx.OssServer.PutLocalFile(t.SvcCtx.AliConfig.Oss.Bucket, "bgm/"+pathSplit[0], item.ShareUrl)
			if err != nil {
				logx.WithContext(t.Ctx).Errorf("clipFilm执行定时任务[同步抖音bgm],写入oss失败,data:%+v,err:%+v", item, err)
				continue
			}

		}
		fmt.Println(tikTokBgmRequest)
		break
	}
	logx.WithContext(t.Ctx).Infof("clipFilm执行定时任务[同步抖音bgm]开始,endTime:%s", time.Now().String())
	return nil
}

type (
	TikTokBgmRequest struct {
		Data struct {
			List []TikTokBgmList `json:"list"`
		}
		Extra struct {
			Logid          string `json:"logid"`
			Now            int64  `json:"now"`
			ErrorCode      int32  `json:"error_code"`
			Description    string `json:"description"`
			SubErrorCode   int32  `json:"sub_error_code"`
			SubDescription string `json:"sub_description"`
		}
	}
	TikTokBgmList struct {
		Title    string `json:"title"`
		Duration int32  `json:"duration"`
		Author   string `json:"author"`
		UseCount int64  `json:"use_count"`
		ShareUrl string `json:"share_url"`
		Rank     int32  `json:"rank"`
		Cover    string `json:"cover"`
	}
)

func TiktokGetBillboardMusic(uri, accessToken string) (res []byte, err error) {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("access-token", accessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 从响应中读取内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
