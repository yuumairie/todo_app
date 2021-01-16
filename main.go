package main

import (
	"fmt"
	"todo_app/app/models"
)

func main() {
	// fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.Password = "testtest"
	// fmt.Println(u)

	// u.CreateUser()

	// user, _ := models.GetUser(2)
	// user.CreateTodo("First Todo")

	// fmt.Println(u)

	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	user, _ := models.GetUser(2)
	t, _ := user.GetTodosByUser()
	fmt.Println(t)
	// user.CreateTodo("Second Todo")

	// t, _ := models.GetTodos()
	// fmt.Println(t)
}
