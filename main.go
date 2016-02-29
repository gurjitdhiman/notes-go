package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"

	"github.com/gurjitdhiman/notes-go/controllers"
	"github.com/gurjitdhiman/notes-go/storage"
)

var DB *sql.DB

func initDb() {
	// Open database by database url.
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("Error:", err)
	}

	// Check Database Running Properly.
	err = db.Ping()
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	DB = db
}

func main() {
	initDb()

	router := httprouter.New()
	server := negroni.Classic()

	notesStorage := &storage.NotesStorageDB{DB: DB}
	notesController := &controllers.NotesController{Storage: notesStorage}

	router.GET("/", homeHandler)
	router.GET("/notes", notesController.IndexHandler)
	router.GET("/notes/:id", notesController.ReadHandler)
	router.POST("/notes", notesController.CreateHandler)
	router.PUT("/notes/:id", notesController.UpdateHandler)
	router.DELETE("/notes/:id", notesController.DestroyHandler)

	server.UseHandler(router)
	server.Run(":8080")
}

func homeHandler(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	res.WriteHeader(200)
	res.Write([]byte("Server Running"))
}
