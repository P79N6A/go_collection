package controller

import (
	"net/http"
	"../view"
)

type homeController struct {}

func (h homeController) registerRoutes() {
	http.HandleFunc("/", indexHandle)
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	templates := PopulateTemplates()
	m := view.IndexViewModeler{}
	v := m.GetViewModel()
	templates["index.html"].Execute(w, &v)
}
