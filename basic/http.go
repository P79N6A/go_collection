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
	Posts []Post
}

type Post struct {
	User User
	Body string
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

func TemplateRangeTest(w http.ResponseWriter, r *http.Request) {
	u1 := User{UserName : "tom"}
	u2 := User{UserName : "john"}
	p1 := Post{User : u1, Body : "tom is 1st user"}
	p2 := Post{User : u2, Body : "kevin is 2nd user"}
	posts := []Post{p1,p2}
	v := IndexViewModel{Title : "home page", User : u1, Posts : posts}
	tpl,_ := template.ParseFiles("templates/range.html")
	tpl.Execute(w,&v)
}

func main() {
	http.HandleFunc("/",TemplateRangeTest)
	http.ListenAndServe(":8888",nil)
}
