package main

import (
    // "fmt"
    "testing"
    // "reflect"
    // "net/http"
    // "net/http/httptest"
    // "github.com/gorilla/mux"
    // "github.com/stephenafamo/site/controllers"
)

func TestAutoLoadControllers (t *testing.T) {

	var expectedResult = routes()
	var something = expectedResult.handler("/{Title}/{Name}", "IndexController", "Index")
	// var something2 = reflect.Indirect(something)

	t.Fatalf("%#v", something)
}