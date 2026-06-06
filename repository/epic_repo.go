package repository

import (
	"context"

	"go.taskmanager.gin_chi_fiber/model"
)

func CreateEpic(epic *model.Epic) error {
	query := `INSERT INTO epics (title, description, status) VALUES ($1, $2, $3) RETURNING id`
	err := Pool.QueryRow(context.Background(), query, epic.Title, epic.Description, epic.Status.String()).Scan(&epic.ID)
	return err
}

func GetEpic(id int) (*model.Epic, error) {
	query := `SELECT id, title, description, status FROM epics WHERE id = $1`
	row := Pool.QueryRow(context.Background(), query, id)
	epic := &model.Epic{}
	var statusText string
	err := row.Scan(&epic.ID, &epic.Title, &epic.Description, &statusText)
	if err != nil {
		return nil, err
	}
	epic.Status = parseStatus(statusText)
	return epic, nil
}

func UpdateEpic(epic *model.Epic) error {
	query := `UPDATE epics SET title=$1, description=$2, status=$3 WHERE id=$4`
	_, err := Pool.Exec(context.Background(), query, epic.Title, epic.Description, epic.Status.String(), epic.ID)
	return err
}

func DeleteEpic(id int) error {
	query := `DELETE FROM epics WHERE id=$1`
	_, err := Pool.Exec(context.Background(), query, id)
	return err
}

// Вспомогательная функция для парсинга статуса из строки БД
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

// GetAllEpics возвращает все эпики из базы данных без подзадач.
func GetAllEpics() ([]model.Epic, error) {
	query := `SELECT id, title, description, status FROM epics ORDER BY id`
	rows, err := Pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var epics []model.Epic
	for rows.Next() {
		var epic model.Epic
		var statusText string
		if err := rows.Scan(&epic.ID, &epic.Title, &epic.Description, &statusText); err != nil {
			return nil, err
		}
		epic.Status = parseStatus(statusText)
		epics = append(epics, epic)
	}
	return epics, rows.Err()
}
