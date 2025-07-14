package main

import (
	_ "gf-sse/internal/packed"

	_ "gf-sse/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"gf-sse/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
