package main

import (
	_ "gf-captcha/internal/packed"

	_ "gf-captcha/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"gf-captcha/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
