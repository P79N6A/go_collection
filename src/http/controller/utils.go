package controller

import (
	"html/template"
	"io/ioutil"
	"os"
)

func CheckErr(e error, errType string) {
	if e != nil {
		panic(errType + e.Error())
	}
}

func PopulateTemplates() map[string] *template.Template {
	basePath := "/Users/wangzijie/code/golang/src/test/go_learning/http/templates"
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
	return result
}
