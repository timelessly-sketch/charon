package main

import (
	_ "charon/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"charon/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
