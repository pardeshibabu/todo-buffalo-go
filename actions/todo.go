package actions

import (
	"crud/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// // TodoIndex default implementation.
// func TodoIndex(c buffalo.Context) error {
// 	return c.Render(http.StatusOK, r.HTML("todo/index.html"))
// }

// // TodoShow default implementation.
// func TodoShow(c buffalo.Context) error {
// 	return c.Render(http.StatusOK, r.HTML("todo/show.html"))
// }

// // TodoAdd default implementation.
// func TodoAdd(c buffalo.Context) error {
// 	return c.Render(http.StatusOK, r.HTML("todo/add.html"))
// }

// TodoIndex default implementation.
func TodoIndex(c buffalo.Context) error {
	// Create an array to receive todos
	todos := []models.Todo{}
	//get all the todos from database
	err := models.DB.All(&todos)
	// handle any error
	if err != nil {
		return c.Render(http.StatusOK, r.JSON(err))
	}
	//return list of todos as json
	return c.Render(http.StatusOK, r.JSON(todos))
}

// TodoShow default implementation.
func TodoShow(c buffalo.Context) error {
	// grab the id url parameter defined in app.go
	id := c.Param("id")
	// create a variable to receive the todo
	todo := models.Todo{}
	// grab the todo from the database
	err := models.DB.Find(&todo, id)
	// handle possible error
	if err != nil {
		return c.Render(http.StatusOK, r.JSON(err))
	}
	//return the data as json
	return c.Render(http.StatusOK, r.JSON(&todo))
}

// TodoAdd default implementation.
func TodoAdd(c buffalo.Context) error {

	//get item from url query
	title := c.Param("title")
	body := c.Param("body")

	//create new instance of todo
	todo := models.Todo{
		Title: title,
		Body:  body,
	}

	// Create a fruit without running validations
	err := models.DB.Create(&todo)

	// handle error
	if err != nil {
		return c.Render(http.StatusOK, r.JSON(err))
	}

	c.Flash().Add("success", "Item was created successfully")
	//return new todo as json
	return c.Render(http.StatusOK, r.JSON(todo.ID))
}
