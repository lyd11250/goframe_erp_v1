package redis

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/guid"
	"goframe-erp-v1/internal/consts"
)

type sRedis struct {
	ctx context.Context
}

func Ctx(ctx context.Context) *sRedis {
	return &sRedis{ctx: ctx}
}

func (s *sRedis) Login(userId int64) {
	// 浏览器设置cookie
	request := g.RequestFromCtx(s.ctx)
	token := guid.S()
	request.Cookie.SetCookie("token", token, "", "/", consts.CookieEx)

	// redis缓存
	_ = g.Redis().SetEX(s.ctx, token, userId, consts.RedisEx)
}

func (s *sRedis) Logout() {
	request := g.RequestFromCtx(s.ctx)
	token := request.Cookie.Get("token").String()
	request.Cookie.Remove("token")
	_, _ = g.Redis().Del(s.ctx, token)
}

func (s *sRedis) CheckLogin() int64 {
	request := g.RequestFromCtx(s.ctx)
	if !request.Cookie.Contains("token") {
		return 0
	}
	token := request.Cookie.Get("token").String()
	idVar, err := g.Redis().Get(s.ctx, token)
	if err != nil {
		return 0
	}
	// 续期
	request.Cookie.SetCookie("token", token, "", "/", consts.CookieEx)
	_, _ = g.Redis().Expire(s.ctx, token, consts.RedisEx)
	return idVar.Int64()
}
