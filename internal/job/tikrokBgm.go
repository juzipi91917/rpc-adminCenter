package job

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/svc"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
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

	logx.WithContext(t.Ctx).Infof("clipFilm执行定时任务[同步抖音bgm]开始,endTime:%s", time.Now().String())
	return nil
}
