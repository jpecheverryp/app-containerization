package main

import (
	"html/template"
	"log"
	"net/http"
)

type todo struct {
	ID   int
	Text string
	Done bool
}

type application struct {
	store []todo
}

func (app *application) indexPage(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./index.html")

	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, app.store)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *application) createTodo(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err!=nil{
        http.Error(w, "Client Error", http.StatusBadRequest)
        return
    }

    log.Print(r.PostForm)

    newTodo := todo{
        ID: 1,
        Text: r.PostForm.Get("new-todo"),
        Done: false,
    }

    app.store = append(app.store, newTodo)

    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	port := ":8080"

	app := &application{
		store: []todo{
			{
				ID:   0,
				Text: "Code apps",
				Done: true,
			},
			{
				ID:   1,
				Text: "Deploy them",
				Done: false,
			},
		},
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", app.indexPage)
    mux.HandleFunc("POST /new-todo", app.createTodo)

	log.Print("running go server in port ", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
