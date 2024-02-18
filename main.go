package main

import (
	"github.com/MrBercisk/golang-todos/controller"
	"github.com/MrBercisk/golang-todos/database"
	"github.com/labstack/echo"
)


func main() {
	db := database.InitDb()
	defer db.Close()

	err := db.Ping()
	if err != nil{
		panic(err)
	}
	e := echo.New()

	/* import controller */
	controller.NewGetAllTodosController(e, db)
	controller.NewCreateTodoController(e, db)
	controller.NewDeleteTodoController(e, db)
	controller.NewUpdateTodoController(e, db)
	controller.NewCheckTodoController(e, db)

	e.Start(":8080")
}