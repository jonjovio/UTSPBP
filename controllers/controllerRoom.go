package controllers

import (
	m "UTS/models"
	"encoding/json"
	"net/http"
)

// GORM Select
func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "failed to parse form data", http.StatusBadRequest)
		return
	}
	var rooms []m.Room
	db := connectgorm()

	result := db.Find(&rooms)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var response = m.RoomResponse{}
	response.Status = http.StatusOK
	response.Message = "Success Get the users"
	json.NewEncoder(w).Encode(rooms)
}

func GetDetailRoom(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "failed to parse form data", http.StatusBadRequest)
		return
	}

	idRoom := r.Form.Get("idRoom")

	var room m.Room
	var participants []m.Participant

	db := connectgorm()
	if err := db.First(&room, idRoom).Error; err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	result := db.Where("id_room = ?", room.ID).Find(&participants)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var response = m.RoomResponseDetail{}
	response.Status = http.StatusOK
	response.Message = "Success Deleting the Data"
	response.Data = room
	response.Participants = participants
	json.NewEncoder(w).Encode(response)
}
