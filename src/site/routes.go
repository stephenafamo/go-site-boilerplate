package main

func routes() *stephenRouter {
	router := customRouter()
	
	router.handler("/", "IndexController", "Index")
	router.handler("/index", "IndexController", "Index")
    router.handler("/{Title}/{Name}", "IndexController", "Index")
    return router
}	