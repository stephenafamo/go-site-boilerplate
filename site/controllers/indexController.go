package controller

import (
	"net/http"
)

type IndexController struct {
	Controller
}

func (i *IndexController) Index(w http.ResponseWriter, r *http.Request, p interface{}) {
	w.WriteHeader(http.StatusOK)
	i.Render(w, "index", i.GetVars(r))
}
