package cmd

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
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
					// 跨域处理中间件
					service.Middleware().CorsHandler,
					// 响应处理中间件
					service.Middleware().ResponseHandler,
					// 权限认证中间件
					service.Middleware().AccessHandler,
				)
				group.Bind(
					controller.User,
					controller.Access,
					controller.Role,
					controller.Supplier,
					controller.Customer,
					controller.Goods,
					controller.Inventory,
					controller.Order,
					controller.File,
				)
			})
			path, err := g.Cfg().Get(ctx, "app.path")
			if err != nil {
				panic(err)
			}
			s.AddStaticPath("/file", path.String())
			err = gtime.SetTimeZone("Asia/Shanghai")
			if err != nil {
				panic(err)
			}
			s.Run()
			return nil
		},
	}
)
