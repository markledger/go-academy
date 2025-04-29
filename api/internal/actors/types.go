package actors

import "api/internal/models"

type RequestStruct struct {
	Action   string
	Task     models.Task
	Response chan ResponseStruct
}

type ResponseStruct struct {
	Data  []models.Task
	Error error
}

var RequestQueue = make(chan RequestStruct)
