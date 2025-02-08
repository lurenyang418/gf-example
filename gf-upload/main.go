package main

import (
	_ "gf-upload/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"gf-upload/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
