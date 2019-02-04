package controller

import (
	"net/http"
	"../view"
)

type homeController struct {}

func (h homeController) registerRoutes() {
	http.HandleFunc("/", indexHandle) //根目录的页面
	http.HandleFunc("/login", loginHandle) //根目录下login的页面
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	templates := PopulateTemplates() //返回一个map，每一项由解析的公共模版_base.html加上解析content目录下的一个html模版文件组成

	m := view.IndexViewModeler{}
	v := m.GetViewModel() //得到结构体，这个结构体的元素由html的{{}}调用，以显示值
	templates["index.html"].Execute(w, &v) //选择被解析的公共模版文件base.html加上被解析的index.html文件
}

func loginHandle(w http.ResponseWriter, r *http.Request) {
	templates := PopulateTemplates()
	m := view.LoginViewModeler{}
	v := m.GetViewModel()
	templates["login.html"].Execute(w, &v)
}
