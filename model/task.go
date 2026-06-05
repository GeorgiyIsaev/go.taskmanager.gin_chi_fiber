package model

// Task – общий интерфейс для задач.
type Task interface {
	GetID() int
	GetTitle() string
	GetDescription() string
	GetStatus() Status
}
