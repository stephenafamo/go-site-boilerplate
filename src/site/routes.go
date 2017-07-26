package main

func routes(router *stephenRouter) *stephenRouter {

	router.handler("/", "IndexController", "Index")
	router.handler("/index", "IndexController", "Index")
	router.handler("/{Title}/{Name}", "IndexController", "Index")
	router.handler("/asset/", "AssetController", "Index")

	return router
}
