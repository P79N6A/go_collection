package view

type LoginViewModel struct {
	BasicViewModel
}

type LoginViewModeler struct {}

func (m LoginViewModeler) GetViewModel () LoginViewModel {
	v := LoginViewModel{}
	v.SetTitle("login")
	return v
}
