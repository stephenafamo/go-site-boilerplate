package main

func routes(router *stephenRouter) *stephenRouter {

	router.handler("/", "IndexController", "Index")
	router.handler(`/assets/{path:[a-zA-Z0-9=\-\/.]+}`, "AssetController", "Index")

	return router
}
