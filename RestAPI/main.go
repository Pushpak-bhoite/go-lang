package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Buy groceries", Completed: false},
	{ID: "2", Item: "Walk the dog", Completed: true},
	{ID: "3", Item: "Read a book", Completed: false},
	{ID: "4", Item: "Write some Go code", Completed: true},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo
	// we are checking structure of incoming json, does it resembles to our todos struct
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo) // appended newTodo to existing one
	context.IndentedJSON(http.StatusCreated, todos)
}

func getToDoById(id string) (*todo, error) {
	for i, val := range todos {
		if id == val.ID {
			fmt.Printf("val->%v\n ", &todos[i])
			return &todos[i], nil // or &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func getSingleTodo(context *gin.Context) {
	id := context.Param("id")
	fmt.Println("id-->", id)
	res, err := getToDoById(id)
	fmt.Printf("res--> %v\n %T \n", res, res)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, res)
}
func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.GET("/todos/:id", getSingleTodo)
	router.Run("localhost:9090")
}
