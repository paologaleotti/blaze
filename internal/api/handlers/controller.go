package handlers

import "blaze/pkg/models"

type ApiController struct {
	db []*models.Todo
}

func NewApiController() *ApiController {
	return &ApiController{
		db: make([]*models.Todo, 0),
	}
}
