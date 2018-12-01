package main

import(
	//"fmt"
	"net/http"
	"html/template"
)

type User struct {
	UserName string
}

type IndexViewModel struct {
	Title string
	User User
}

func printHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world go web."))
}

func TemplateTest(w http.ResponseWriter, r *http.Request) {
	user := User{UserName : "Tom"}
	tpl,_ := template.New("").Parse(`
	<html>
	<head>
	<title>home page</title>
	</head>
	<body>
	<h1>hello, {{.UserName}}!</h1>
	</body>
	</html>`)
	tpl.Execute(w,&user)
}

func TemplateFileTest(w http.ResponseWriter, r *http.Request) {
	user := User{UserName : "John"}
	tpl,_ := template.ParseFiles("templates/hello.html")
	tpl.Execute(w,&user)
}

func TemplateIndexViewModelTest(w http.ResponseWriter, r *http.Request) {
	user := User{UserName : "kevin"}
	v := IndexViewModel{Title : "home page", User : user}
	tpl,_ := template.ParseFiles("templates/index_view_model.html")
	tpl.Execute(w,&v)
}

func main() {
	http.HandleFunc("/",TemplateIndexViewModelTest)
	http.ListenAndServe(":8888",nil)
}
