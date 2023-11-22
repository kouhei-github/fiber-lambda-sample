package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kouhei-github/fiber-sample-framework/route"
	"os"
)

var fiberLambda *fiberadapter.FiberLambda

func init() {
	app := fiber.New()

	// ルーターの設定
	router := &route.Router{FiberApp: app}

	// CORS (Cross Origin Resource Sharing)の設定
	// アクセスを許可するドメイン等を設定します
	router.FiberApp.Use(cors.New(cors.Config{AllowHeaders: "Origin, Content-Type, Accept"}))

	route.LoadRouter(router)

	// fmt.Println(os.Getenv("ENVIRONMENT"))
	// Webサーバー起動時のエラーハンドリング => localhostの時コメントイン必要
	if os.Getenv("ENVIRONMENT") == "local" {
		if err := router.FiberApp.Listen(":8080"); err != nil {
			panic(err)
		}
	} else {
		// fmt.Println("lambda")
		// AWS Lambdaとの連携設定
		fiberLambda = fiberadapter.New(router.FiberApp)
	}
}

// Handler will deal with Fiber working with Lambda
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error

	return fiberLambda.ProxyWithContext(ctx, req)
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(Handler)
}
