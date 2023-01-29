package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func server(){
	// echoインスタンスを生成
	e := echo.New()
	// Middleware
	// httpリクエストの情報をログに表示
	e.Use(middleware.Logger())
	// パニックを回復し、スタックトレースを表示
	e.Use(middleware.Recover())
	// ルーティング
	e.GET("/", getSample)

	// サーバーをスタートさせる
	// ポート番号は引数で指定できる
	e.Logger.Fatal(e.Start(":8080"))
}

// Get API
func getSample(c echo.Context) error {
	return c.String(http.StatusOK, "Get!")
}
