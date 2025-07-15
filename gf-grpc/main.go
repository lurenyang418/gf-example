package main

import (
	_ "gf-grpc/internal/packed"

	_ "gf-grpc/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"gf-grpc/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
