package route

import "github.com/kouhei-github/fiber-sample-framework/controller"

func (router *Router) GetRouter() {
	// 練習
	router.FiberApp.Get("/", controller.HelloHandler)
	// パスパラメータ取得
	router.FiberApp.Get("/path/:id", controller.PathParamTestHandler)
	// パスパラメータ取得
	router.FiberApp.Get("/query", controller.QueryParamTestHandler)

	// ヘルスチェック
	router.FiberApp.Post("/test", controller.HealthCheckHandler)
}
