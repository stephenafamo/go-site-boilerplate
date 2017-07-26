package main

import (
    // "fmt"
    "reflect"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/stephenafamo/site/controllers"
)

func customRouter() *stephenRouter {
	router := new( stephenRouter)
	router.Router = *mux.NewRouter()
	return router

}

type stephenRouter struct {
	mux.Router
}

func (c *stephenRouter) handler(path string, controllerName string, method string) {	
	theController := reflect.New(controller.Get(controllerName))
	theMethod := theController.MethodByName(method)
	c.HandleFunc(path, func (w http.ResponseWriter, r *http.Request) {
		theMethod.Call([]reflect.Value{reflect.ValueOf(w), reflect.ValueOf(r)})
	})
}