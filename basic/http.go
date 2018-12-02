package main

import(
	//"fmt"
	"net/http"
	"html/template"
	"io/ioutil"
	"os"
)

type User struct {
	UserName string
}

type IndexViewModel struct {
	Title string
	User
	Posts []Post
}

type Post struct {
	User
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

func PopulateTemplatesTest(w http.ResponseWriter, r *http.Request) {
	basePath := "templates"
	result := make(map[string] *template.Template)
	//layout := template.Must(template.ParseFiles(basePath + "/_base.html")) Must方法只是使返回值只为template，不返回err
	layout, _ := template.ParseFiles(basePath + "/_base.html")

	dir, err := os.Open(basePath + "/content")
	CheckErr(err, "failed to open content dir: ")

	files, err := dir.Readdir(-1)
	CheckErr(err, "failed to read files in content dir")

	for _, file := range files {
		f, err := os.Open(basePath + "/content/" + file.Name())
		CheckErr(err, "failed to open file: " + file.Name())

		content, err := ioutil.ReadAll(f)
		CheckErr(err, "failed to get content of the file: " + file.Name())

		f.Close()

		//tmpl := template.Must(layout.Clone()) Must方法只是使返回值只为template，不返回err
		tmpl,_ := layout.Clone()

		_, err = tmpl.Parse(string(content))
		CheckErr(err, "failed to parse content of file: " + file.Name())
		result[file.Name()] = tmpl
	}
	u1 := User{UserName : "tom"}
	u2 := User{UserName : "john"}
	p1 := Post{User : u1, Body : "tom is 1st user"}
	p2 := Post{User : u2, Body : "kevin is 2nd user"}
	posts := []Post{p1,p2}
	v := IndexViewModel{Title : "home page", User : u1, Posts : posts}
	templates := result
	templates["index.html"].Execute(w, &v)
}

func CheckErr(e error, errType string) {
	if e != nil {
		panic(errType + e.Error())
	}
}

func main() {
	http.HandleFunc("/", PopulateTemplatesTest)
	http.ListenAndServe(":8888",nil)
}
