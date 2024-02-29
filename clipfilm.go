package main

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-clipFilm/internal/config"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-clipFilm/internal/pkg"
	bgmserviceServer "codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-clipFilm/internal/server/bgmservice"
	filtercustomserviceServer "codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-clipFilm/internal/server/filtercustomservice"
	fontserviceServer "codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-clipFilm/internal/server/fontservice"
	paperworkserviceServer "codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-clipFilm/internal/server/paperworkservice"
	sellingpointimageserviceServer "codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-clipFilm/internal/server/sellingpointimageservice"
	sellingpointserviceServer "codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-clipFilm/internal/server/sellingpointservice"
	soundpeopleserviceServer "codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-clipFilm/internal/server/soundpeopleservice"
	subtitlestyleserviceServer "codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-clipFilm/internal/server/subtitlestyleservice"
	taskserviceServer "codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-clipFilm/internal/server/taskservice"
	templateserviceServer "codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-clipFilm/internal/server/templateservice"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-clipFilm/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-clipFilm/pb/clipFilm"
	"fmt"

	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	c := config.Init()
	ctx := svc.NewServiceContext(*c)
	// 注册其他服务
	pkg.ServerInit(ctx)
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		clipFilm.RegisterBgmServiceServer(grpcServer, bgmserviceServer.NewBgmServiceServer(ctx))
		clipFilm.RegisterFilterCustomServiceServer(grpcServer, filtercustomserviceServer.NewFilterCustomServiceServer(ctx))
		clipFilm.RegisterTemplateServiceServer(grpcServer, templateserviceServer.NewTemplateServiceServer(ctx))
		clipFilm.RegisterPaperworkServiceServer(grpcServer, paperworkserviceServer.NewPaperworkServiceServer(ctx))
		clipFilm.RegisterTaskServiceServer(grpcServer, taskserviceServer.NewTaskServiceServer(ctx))
		clipFilm.RegisterFontServiceServer(grpcServer, fontserviceServer.NewFontServiceServer(ctx))
		clipFilm.RegisterSellingPointServiceServer(grpcServer, sellingpointserviceServer.NewSellingPointServiceServer(ctx))
		clipFilm.RegisterSellingPointImageServiceServer(grpcServer, sellingpointimageserviceServer.NewSellingPointImageServiceServer(ctx))
		clipFilm.RegisterSubtitleStyleServiceServer(grpcServer, subtitlestyleserviceServer.NewSubtitleStyleServiceServer(ctx))
		clipFilm.RegisterSoundPeopleServiceServer(grpcServer, soundpeopleserviceServer.NewSoundPeopleServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			//wg := new(sync.WaitGroup)
			//go func() {
			//	consumer.RegisterConsumer(ctx, wg)
			//}()

			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
