package main

import (
	_ "goframe-erp-v1/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"goframe-erp-v1/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
