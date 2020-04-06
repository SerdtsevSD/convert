package main

func initializeRoutes() {
	router.GET("/pageConv", pageConv)
	router.GET("/login", Login)
	router.POST("/getRBK", getRBK)
	router.POST("/auth", auth)
}
