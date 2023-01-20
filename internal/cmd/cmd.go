package cmd

import (
	"context"
	"goframe-erp-v1/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"goframe-erp-v1/internal/controller"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					// 响应处理中间件
					service.Middleware().ResponseHandler,
				)
				group.Bind(
					controller.User,
				)
			})
			s.Run()
			return nil
		},
	}
)
