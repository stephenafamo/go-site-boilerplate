package controller

import (
	"net/http"
	"path"
	"github.com/stephenafamo/what-to-do/config"
)

type AssetController struct {
	Controller
}

func (a *AssetController) Index(w http.ResponseWriter, r *http.Request, p interface{}) {	
	AssetBaseFolder := config.GetS("AssetBaseFolder")
	vars := a.GetVars(r)
	if _, ok := vars["path"]; ok != true{
		http.NotFound(w, r )
	}
	filePath := path.Join(AssetBaseFolder, vars["path"])
	http.ServeFile(w, r, filePath)
}
