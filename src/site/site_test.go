package main

import (
	"testing"
)

func TestAutoLoadControllers(t *testing.T) {

	var expectedResult = routes(customRouter())
	expectedResult.handler("/asset", "AssetController", "Index")

	t.Fatalf("%#v", something)
}
