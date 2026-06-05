package model

// Subtask – подзадача, принадлежит определённому Epic.
type Subtask struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      Status `json:"status"`
	EpicID      int    `json:"epicId"`
}

func (s Subtask) GetID() int             { return s.ID }
func (s Subtask) GetTitle() string       { return s.Title }
func (s Subtask) GetDescription() string { return s.Description }
func (s Subtask) GetStatus() Status      { return s.Status }
