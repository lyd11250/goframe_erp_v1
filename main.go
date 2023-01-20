package main

import (
	"github.com/gogf/gf/v2/os/gctx"
	"goframe-erp-v1/internal/cmd"
	_ "goframe-erp-v1/internal/logic"
	_ "goframe-erp-v1/internal/packed"
)

func main() {
	cmd.Main.Run(gctx.New())
}
