package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := NewApp()

	// Wailsアプリケーションの構築と起動設定
	err := wails.Run(&options.App{
		Title:  "構造的数式シミュレーター",
		Width:  800,
		Height: 600,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 59, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app, // ここでApp構造体を登録することで、JS側にバインドされる
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}