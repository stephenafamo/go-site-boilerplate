package controller

import (
	"net/http"
)

type AssetController struct {
	Controller
}

func (i *AssetController) Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/go/src/github.com/stephenafamo/site")
}
