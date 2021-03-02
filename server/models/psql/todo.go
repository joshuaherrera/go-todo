package psql

import (
	"database/sql"

	"github.com/joshuaherrera/go-todo/server/models"
)

// TodoModel defines wrapper for sql.DB connection pool
type TodoModel struct {
	DB *sql.DB
}

// All grabs all todos from db
func (m *TodoModel) All() ([]*models.Todo, error) {
	stmt := `SELECT todo_id, description FROM todo`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	todos := []*models.Todo{}

	for rows.Next() {
		t := &models.Todo{}

		err := rows.Scan(&t.TodoID, &t.Description)
			if err != nil {
				return nil, err
		}

		todos = append(todos, t)
	}
	// check for errors in scanning
		if err = rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}
// Get gets one todo record from the DB
func (m *TodoModel) Get(id int) (*models.Todo, error) {
	stmt := `SELECT todo_id, description FROM todo WHERE todo_id = $1`

	row := m.DB.QueryRow(stmt, id)

	t := &models.Todo{}

	err := row.Scan(&t.TodoID, &t.Description)

	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecords
	} else if err != nil {
		return nil, err
	}

	return t, nil
}

// Insert adds a todo record to the DB
func (m *TodoModel) Insert(description string) (*models.Todo, error) {
	stmt := `INSERT INTO todo (description) VALUES ($1) RETURNING *`

	// use QueryRow since we are retuning, otherwise would use Exec
	row := m.DB.QueryRow(stmt, description)

	t := &models.Todo{}

	err := row.Scan(&t.TodoID, &t.Description)

	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecords
	} else if err != nil {
		return nil, err
	}

	return t, nil
}

// Update updates one todo record
func (m *TodoModel) Update(id int, description string) (error) {
	stmt := `UPDATE todo SET description = $1 WHERE todo_id = $2`

	_, err := m.DB.Exec(stmt, description, id)
	if err != nil {
		return  err
	}
	return nil
}

// Delete removes a record from the todo table
func (m *TodoModel) Delete(id int) (error) {
	stmt := `DELETE FROM todo WHERE todo_id = $1`
	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		return  err
	}
	return nil
}