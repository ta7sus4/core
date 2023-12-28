package main

import (
	"goki.dev/gi/v2/gi"
	"goki.dev/gi/v2/gimain"
)

func main() { gimain.Run(app) }

func app() {
	b := gi.NewAppBody("hello")
	gi.NewLabel(b).SetText("Hello, World!")
	b.NewWindow().Run().Wait()
}
