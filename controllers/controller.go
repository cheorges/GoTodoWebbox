package controllers

import (
	"GoTodoWebbox/db"
	"GoTodoWebbox/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"html/template"
	"math"
	"net/http"
)

var (
	id        int
	item      string
	completed int
	view      = template.Must(template.ParseFiles("./views/index.html"))
)

func Show(w http.ResponseWriter, r *http.Request) {
	statement := db.GetAllTask()
	var todos []models.Todo
	completed := 0

	for _, s := range statement {
		status := s["status"].(bool)
		todos = append(todos, models.Todo{
			Id:        s["_id"].(primitive.ObjectID).Hex(),
			Item:      s["task"].(string),
			Completed: status})
		if status {
			completed++
		}
	}

	data := models.View{
		Todos:     todos,
		Completed: math.Round(float64(completed) / float64(len(todos)) * 100),
	}

	_ = view.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")
	db.InsertOneTask(item)
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	db.DeleteOneTask(id)
	http.Redirect(w, r, "/", 301)
}

func Complete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	db.TaskComplete(id)
	http.Redirect(w, r, "/", 301)
}
