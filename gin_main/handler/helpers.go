package handler

import "go.taskmanager.gin_chi_fiber/model"

// ----- Запросы и ответы для эпиков -----

type CreateEpicRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreateEpicResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type UpdateEpicRequest struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type GetEpicResponse struct {
	ID          int               `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Status      string            `json:"status"`
	Subtasks    []SubtaskResponse `json:"subtasks"`
}

// ----- Запросы и ответы для подзадач -----

type CreateSubtaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	EpicID      int    `json:"epicId"`
}

type UpdateSubtaskRequest struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type SubtaskResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	EpicID      int    `json:"epicId"`
}

// ----- Вспомогательные функции -----

func parseStatus(s string) model.Status {
	switch s {
	case "Новая":
		return model.StatusNew
	case "В процессе":
		return model.StatusInProgress
	case "Выполнена":
		return model.StatusDone
	default:
		return model.StatusNew
	}
}
