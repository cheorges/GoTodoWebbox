package models

type Todo struct {
	Id        string
	Item      string
	Completed bool
}

type View struct {
	Todos     []Todo
	Completed float64
}
