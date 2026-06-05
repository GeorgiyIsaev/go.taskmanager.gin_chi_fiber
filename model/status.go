package model

// Status задачи
type Status int

const (
	StatusNew Status = iota
	StatusInProgress
	StatusDone
)

func (s Status) String() string {
	switch s {
	case StatusNew:
		return "Новая"
	case StatusInProgress:
		return "В процессе"
	case StatusDone:
		return "Выполнена"
	default:
		return "Неизвестно"
	}
}
