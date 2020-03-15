// routes.go

package main

func initializeRoutes() {

	// определение роута главной страницы
	router.GET("/pageConv", pageConv)
	router.POST("/getRBK", getRBK)
	router.GET("/", mainPage)
}
