package controller

func StartUp() {
	h := homeController{}
	h.registerRoutes()
}
