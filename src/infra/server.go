package infraHttp

import "example/web-service-gin/src/infra/http/routes"

func InitializeServer() {
	router := routes.GetRoutes()

	router.Run("localhost:8080")
}
