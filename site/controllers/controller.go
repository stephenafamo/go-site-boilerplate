package controller

import (
	"github.com/gorilla/mux"
	"github.com/stephenafamo/what-to-do/config"
	"net/http"
	"reflect"
	"html/template"
)

type Controller struct {
	Registry map[string]reflect.Type
	Params map[string]interface{}
}

var controllers = make(map[string]reflect.Type)

func init() {
	controllers["Controller"] = reflect.TypeOf(Controller{})
	controllers["IndexController"] = reflect.TypeOf(IndexController{})
	controllers["AssetController"] = reflect.TypeOf(AssetController{})
}

func (c *Controller) Render(w http.ResponseWriter, templateName string, p interface{}) {
	var Templates = template.Must(template.ParseGlob(config.GetS("TemplateDirectory")))
	if err := Templates.ExecuteTemplate(w, templateName, p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Controller) GetVars(r *http.Request) map[string]string {
	return mux.Vars(r)
}

func Get(name string) reflect.Type {
	return controllers[name]
}
