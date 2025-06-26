package models

type NewTodo struct {
	Title string `json:"title" validate:"required"`
}

type Todo struct {
	Completed bool   `json:"completed" db:"completed"`
	Id        string `json:"id" db:"id"`
	Title     string `json:"title" db:"title"`
}
