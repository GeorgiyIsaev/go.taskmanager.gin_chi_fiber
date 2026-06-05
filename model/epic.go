package model

// Epic – главная задача, может содержать подзадачи.
type Epic struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	Subtasks    []Subtask `json:"subtasks,omitempty"`
}

func (e Epic) GetID() int             { return e.ID }
func (e Epic) GetTitle() string       { return e.Title }
func (e Epic) GetDescription() string { return e.Description }
func (e Epic) GetStatus() Status      { return e.Status }
