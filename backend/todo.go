package main

import (
	"context"
	"time"
)

// Todo represents one todo item
type Todo struct {
	ID          uint64     `json:"id"`
	Username    string     `json:"username"`
	Title       string     `json:"title"`
	IsCompleted bool       `json:"is_completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

// payload for creating todo
type CreateTodoInput struct {
	Username string `json:"username" binding:"required"`
	Title    string `json:"title" binding:"required"`
}

// payload for completing a todo
type CompleteTodoInput struct {
	Username string `json:"username" binding:"required"`
}

// insert a new todo into ClickHouse
func createTodo(input CreateTodoInput) (*Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := uint64(time.Now().UnixNano())
	now := time.Now()

	err := ch.Exec(ctx, `
        INSERT INTO todos (id, username, title, is_completed, created_at, completed_at)
        VALUES (?, ?, ?, ?, ?, ?)
    `,
		id,
		input.Username,
		input.Title,
		uint8(0),
		now,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &Todo{
		ID:          id,
		Username:    input.Username,
		Title:       input.Title,
		IsCompleted: false,
		CreatedAt:   now,
		CompletedAt: nil,
	}, nil
}

// fetch todos by username
func listTodosByUsername(username string) ([]Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := ch.Query(ctx, `
		SELECT id, username, title, is_completed, created_at, completed_at
		FROM todos
		WHERE username = ?
		ORDER BY created_at DESC
	`, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		var (
			id          uint64
			user        string
			title       string
			isCompleted uint8
			createdAt   time.Time
			completedAt *time.Time
		)

		if err := rows.Scan(&id, &user, &title, &isCompleted, &createdAt, &completedAt); err != nil {
			return nil, err
		}

		todos = append(todos, Todo{
			ID:          id,
			Username:    user,
			Title:       title,
			IsCompleted: isCompleted == 1,
			CreatedAt:   createdAt,
			CompletedAt: completedAt,
		})
	}

	return todos, nil
}

// mark todo as completed
func completeTodo(id uint64, username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	now := time.Now()

	// ClickHouse updates = mutations
	err := ch.Exec(ctx, `
		ALTER TABLE todos
		UPDATE is_completed = 1, completed_at = ?
		WHERE id = ? AND username = ?
	`,
		now,
		id,
		username,
	)
	if err != nil {
		return err
	}

	return nil
}
