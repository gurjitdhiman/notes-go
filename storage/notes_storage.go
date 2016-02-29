package storage

import (
	"database/sql"
	"time"

	"errors"

	"github.com/gurjitdhiman/notes-go/models"
)

type NotesStorage interface {
	InsertNote(*models.Note) error
	FindAllNotes() ([]models.Note, error)
	UpdateNote(int, *models.Note) error
	FindNote(int) (models.Note, error)
	DeleteNote(int) error
}

type NotesStorageDB struct {
	DB *sql.DB
}

func (ns *NotesStorageDB) FindAllNotes() ([]models.Note, error) {
	notes := []models.Note{}

	// Find all notes from notes table.
	rows, err := ns.DB.Query("SELECT id,title,content,priority,created_at FROM notes")
	if err != nil {
		return notes, err
	}
	defer rows.Close()
	for rows.Next() {
		note := models.Note{}

		//  Read values into note.
		err := rows.Scan(&note.Id, &note.Title, &note.Content, &note.Priority, &note.CreatedAt)
		if err != nil {
			return []models.Note{}, err
		}

		// Append note to notes slice.
		notes = append(notes, note)
	}
	return notes, nil
}

func (ns *NotesStorageDB) InsertNote(note *models.Note) error {
	return ns.DB.QueryRow("INSERT INTO notes(title, content, priority, created_at) VALUES($1,$2,$3,$4) RETURNING id", note.Title, note.Content, note.Priority, time.Now()).Scan(&note.Id)
}

func (ns *NotesStorageDB) UpdateNote(noteID int, note *models.Note) error {
	return ns.DB.QueryRow("UPDATE notes SET title=$1, content=$2,priority=$3 WHERE id=$4 RETURNING id, title, content, priority, created_at", note.Title, note.Content, note.Priority, noteID).
		Scan(&note.Id, &note.Title, &note.Content, &note.Priority, &note.CreatedAt)
}

func (ns *NotesStorageDB) FindNote(noteID int) (models.Note, error) {
	note := models.Note{}
	err := ns.DB.QueryRow("SELECT id, title, content, priority, created_at FROM notes WHERE id=$1", noteID).Scan(&note.Id, &note.Title, &note.Content, &note.Priority, &note.CreatedAt)
	if err != nil {
		return note, err
	}
	return note, nil
}

func (ns *NotesStorageDB) DeleteNote(noteID int) error {
	// Delete note by note id.
	result, err := ns.DB.Exec("DELETE FROM notes WHERE id=$1", noteID)
	if err != nil {
		return err
	}

	// Check rows affected by query.
	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("Error: Could Not Delete Note")
	}
	return nil
}
