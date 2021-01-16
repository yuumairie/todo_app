package models

import (
	"log"
	"time"
)

// Todo todo定義
type Todo struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
}

// CreateTodo todo作成関数
func (u *User) CreateTodo(content string) (err error) {
	cmd := `insert into todos (
		content, 
		user_id, 
		created_at) values (?,?,?)`

	_, err = Db.Exec(cmd, content, u.ID, time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// GetTodo 指定したIDのtodo取得関数
func GetTodo(id int) (todo Todo, err error) {
	cmd := `select id, 
	content,
	user_id,
	created_at from todos where id = ?`

	todo = Todo{}
	err = Db.QueryRow(cmd, id).Scan(&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt)

	return todo, err
}

// GetTodos todo全行取得関数
func GetTodos() (todos []Todo, err error) {
	cmd := `select id, 
	content,
	user_id,
	created_at from todos`

	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)

		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

// GetTodosByUser 指定したユーザのtodo取得関数
func (u *User) GetTodosByUser() (todos []Todo, err error) {
	cmd := `select id, 
	content,
	user_id,
	created_at from todos where user_id = ?`

	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)

		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()
	return todos, err
}
