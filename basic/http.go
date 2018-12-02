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
		tmpl, _ := layout.Clone()

		_, err = tmpl.Parse(string(content))
		//Named template definitions ({{define ...}} or {{block ...}} statements) in text 
		//define additional templates associated with t and are removed from the definition of t itself
		//在layout的基础上添加content目录下新定义的模版，此处添加了content目录下index.html的模版内容
		//这样做的目的是，首先最开始解析了_base.html这个基础template，其中{{template "content" .}}的内容由之后{{define "content"}}的模版文件决定并保存到layout变量里；
		//之后遍历content目录，该目录下的文件都定义了{{define "content"}}，每次都在_base.html解析后的模版layout中加入content的定义(Parse方法)，存入tmpl。最后存入map，key为该文件名
		//之后若需要修改_base.html里的{{.content}}模版
		//只需要重新写一个文件，放在content目录下，并加入{{define "content"}}，这样在map中就可以根据文件名这个key来加载不同的页面效果

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

// go解析模版文件的两种方式：
// 1.直接用template.ParseFiles解析模版文件，需要加载(ParseFiles)，然后执行(Execute)
// 2.用ParseFiles方法解析字符串，需要先新建(New)，然后加载(Parse)，最后执行(Execute)

func main() {
	http.HandleFunc("/", PopulateTemplatesTest)
	http.ListenAndServe(":8888",nil)
}
