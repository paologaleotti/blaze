package models

type NewTodo struct {
	Title string `json:"title" validate:"required"`
}

type Todo struct {
	Completed bool   `json:"completed"`
	Id        string `json:"id"`
	Title     string `json:"title"`
}
