package main

import (
	_ "GaAdmin/internal/packed"

	_ "GaAdmin/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"GaAdmin/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
