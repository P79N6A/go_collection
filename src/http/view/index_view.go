package view

import(
	"../model"
)

type IndexViewModel struct {
	BasicViewModel
	model.User
	Posts []model.Post
}

type IndexViewModeler struct {}

func (m IndexViewModeler) GetViewModel () IndexViewModel {
	u1 := model.User{UserName: "tom"}
	u2 := model.User{UserName: "jeff"}

	posts := []model.Post {
		model.Post{User: u1, Body: "this is 1st user"},
		model.Post{User: u2, Body: "this is 2nd user"},
	}

	v := IndexViewModel{BasicViewModel{"home page"}, u1, posts}
	return v
}
