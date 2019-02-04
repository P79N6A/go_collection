package view

type BasicViewModel struct {
	Title string
}

func (v *BasicViewModel) SetTitle(title string) {
	v.Title = title
}
