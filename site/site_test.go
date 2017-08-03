package main

import (
	"testing"
	"reflect"
	"github.com/stephenafamo/what-to-do/controllers"
)

func TestAutoLoadControllers(t *testing.T) {

	actualResult := controller.Get("IndexController")
	expectedResult := reflect.TypeOf(controller.IndexController{})
	if reflect.DeepEqual(expectedResult, actualResult) != true {
		t.Fatalf("controller is not autoloaded correctly \n expected %#v \n got %#v", expectedResult, actualResult)		
	}

}
