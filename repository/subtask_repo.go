package repository

import (
	"context"

	"go.taskmanager.gin_chi_fiber/model"
)

func AddSubtask(sub *model.Subtask) error {
	query := `INSERT INTO subtasks (title, description, status, epic_id) VALUES ($1, $2, $3, $4) RETURNING id`
	err := Pool.QueryRow(context.Background(), query, sub.Title, sub.Description, sub.Status.String(), sub.EpicID).Scan(&sub.ID)
	return err
}

func GetSubtask(id int) (*model.Subtask, error) {
	query := `SELECT id, title, description, status, epic_id FROM subtasks WHERE id = $1`
	row := Pool.QueryRow(context.Background(), query, id)
	sub := &model.Subtask{}
	var statusText string
	err := row.Scan(&sub.ID, &sub.Title, &sub.Description, &statusText, &sub.EpicID)
	if err != nil {
		return nil, err
	}
	sub.Status = parseStatus(statusText)
	return sub, nil
}

func UpdateSubtask(sub *model.Subtask) error {
	query := `UPDATE subtasks SET title=$1, description=$2, status=$3 WHERE id=$4`
	_, err := Pool.Exec(context.Background(), query, sub.Title, sub.Description, sub.Status.String(), sub.ID)
	return err
}

func DeleteSubtask(id int) error {
	query := `DELETE FROM subtasks WHERE id=$1`
	_, err := Pool.Exec(context.Background(), query, id)
	return err
}

func GetSubtasksByEpicID(epicID int) ([]model.Subtask, error) {
	query := `SELECT id, title, description, status, epic_id FROM subtasks WHERE epic_id = $1`
	rows, err := Pool.Query(context.Background(), query, epicID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subs []model.Subtask
	for rows.Next() {
		var s model.Subtask
		var statusText string
		if err := rows.Scan(&s.ID, &s.Title, &s.Description, &statusText, &s.EpicID); err != nil {
			return nil, err
		}
		s.Status = parseStatus(statusText)
		subs = append(subs, s)
	}
	return subs, rows.Err()
}
