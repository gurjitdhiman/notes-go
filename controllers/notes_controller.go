package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gurjitdhiman/notes-go/models"
	"github.com/gurjitdhiman/notes-go/storage"
	"github.com/julienschmidt/httprouter"
)

type NotesController struct {
	Storage storage.NotesStorage
}

func (nc *NotesController) IndexHandler(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// Get all notes from storage.
	notes, err := nc.Storage.FindAllNotes()
	if err != nil {
		res.WriteHeader(500)
		res.Write([]byte(err.Error()))
		return
	}

	// Format note to JSON.
	notesJSON, err := json.Marshal(notes)
	if err != nil {
		res.WriteHeader(500)
		res.Write([]byte(err.Error()))
		return
	}
	res.WriteHeader(200)
	res.Write([]byte(notesJSON))
}

func (nc *NotesController) CreateHandler(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// Read note properties.
	title := req.FormValue("title")
	content := req.FormValue("content")
	priority := req.FormValue("priority")

	// Persist note.
	noteData := models.NewNote(title, content, priority)
	err := nc.Storage.InsertNote(noteData)
	if err != nil {
		res.WriteHeader(500)
		res.Write([]byte(err.Error()))
		return
	}

	// Format note to JSON.
	noteJSON, err := json.Marshal(noteData)
	if err != nil {
		res.WriteHeader(500)
		res.Write([]byte(err.Error()))
		return
	}
	res.WriteHeader(200)
	res.Write([]byte(noteJSON))
}

func (nc *NotesController) UpdateHandler(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// Read note id as an integer.
	noteIdStr := params.ByName("id")
	noteId, err := strconv.ParseInt(noteIdStr, 10, 64)
	if err != nil {
		res.WriteHeader(500)
		res.Write([]byte(err.Error()))
		return
	}

	// Read note properties.
	title := req.FormValue("title")
	content := req.FormValue("content")
	priority := req.FormValue("priority")
	noteData := models.NewNote(title, content, priority)

	// Update note values in storage.
	err = nc.Storage.UpdateNote(int(noteId), noteData)
	if err != nil {
		res.WriteHeader(500)
		res.Write([]byte(err.Error()))
		return
	}

	// Format note into JSON.
	noteJSON, err := json.Marshal(noteData)
	if err != nil {
		res.WriteHeader(500)
		res.Write([]byte(err.Error()))
		return
	}
	res.WriteHeader(200)
	res.Write([]byte(noteJSON))
}

func (nc *NotesController) ReadHandler(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// Read note id as an integer.
	noteIdStr := params.ByName("id")
	noteId, err := strconv.ParseInt(noteIdStr, 10, 64)
	if err != nil {
		res.WriteHeader(500)
		res.Write([]byte(err.Error()))
		return
	}

	// Find note from storage.
	note, err := nc.Storage.FindNote(int(noteId))
	if err != nil {
		res.WriteHeader(500)
		res.Write([]byte(err.Error()))
		return
	}

	// Format note into JSON.
	noteJSON, err := json.Marshal(note)
	if err != nil {
		res.WriteHeader(500)
		res.Write([]byte(err.Error()))
		return
	}
	res.WriteHeader(200)
	res.Write([]byte(noteJSON))
}

func (nc *NotesController) DestroyHandler(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// Read note id as an integer.
	noteIdStr := params.ByName("id")
	noteId, err := strconv.ParseInt(noteIdStr, 10, 64)
	if err != nil {
		res.WriteHeader(500)
		res.Write([]byte(err.Error()))
		return
	}

	// Delete note in storage.
	err = nc.Storage.DeleteNote(int(noteId))
	if err != nil {
		res.WriteHeader(500)
		res.Write([]byte(err.Error()))
		return
	}
	res.WriteHeader(200)
	res.Write([]byte("Deleted"))
}
