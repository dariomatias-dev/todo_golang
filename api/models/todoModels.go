package models

type CreateTodo struct {
	Title       string `json:"title" binding:"required,min=3,max=100"`
	Description string `json:"description" binding:"required,min=3,max=256"`
}

type UpdateTodo struct {
	Title       *string `json:"title" binding:"omitempty,min=3,max=100"`
	Description *string `json:"description" binding:"omitempty,min=3,max=256"`
}

type UpdateTodoStatus struct {
	Status *bool `json:"status" binding:"required,boolean"`
}
