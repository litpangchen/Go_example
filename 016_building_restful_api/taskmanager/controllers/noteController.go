package controllers

import (
	"net/http"
	"encoding/json"
	"go_web_programming/016_building_restful_api/taskmanager/common"
	"go_web_programming/016_building_restful_api/taskmanager/models"
	"gopkg.in/mgo.v2/bson"
	"go_web_programming/016_building_restful_api/taskmanager/data"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var dataResource NoteResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Note data", 500)
		return
	}
	noteModel := dataResource.Data
	note := &models.TaskNote{
		TaskId:      bson.ObjectIdHex(noteModel.TaskId),
		Description: noteModel.Description,
	}
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("notes")
	repo := &data.NoteRepository{C: col}
	repo.Create(note)
	j, err := json.Marshal(note)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func GetNotes(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("notes")
	repo := &data.NoteRepository{C: col}
	notes := repo.GetAll()
	j, err := json.Marshal(NotesResource{Data: notes})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func GetNoteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("notes")
	repo := &data.NoteRepository{C: col}
	note, err := repo.GetById(id)
	if err != nil {
		if err != mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	j, err := json.Marshal(note)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func GetNoteByTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("notes")
	repo := &data.NoteRepository{C: col}
	notes := repo.GetByTaskId(id)
	j, err := json.Marshal(NotesResource{Data: notes})
	if err != nil {
		common.DisplayAppError(w, err, "An expected error", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
	var dataResource NoteResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid note data", 500)
		return
	}
	noteModel := dataResource.Data
	note := &models.TaskNote{
		Id:          id,
		Description: noteModel.Description,
	}
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("notes")
	repo := &data.NoteRepository{C: col}
	if err := repo.Update(note); err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("notes")
	repo := &data.NoteRepository{C: col}
	err := repo.Delete(id)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
