package controller

import (
	"net/http"
	"path"
)

type AssetController struct {
	Controller
}

func (a *AssetController) Index(w http.ResponseWriter, r *http.Request, p interface{}) {
	AssetBaseFolder := "/go/src/github.com/stephenafamo/site/assets"
	vars := a.GetVars(r)
	if _, ok := vars["path"]; ok != true{
		http.NotFound(w, r )
	}
	filePath := path.Join(AssetBaseFolder, vars["path"])
	http.ServeFile(w, r, filePath)
}
