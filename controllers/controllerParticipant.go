package controllers

import (
	m "UTS/models"
	"encoding/json"
	"net/http"
	"strconv"
)

func InsertRoom(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "failed to parse form data", http.StatusBadRequest)
		return
	}
	idAccount, _ := strconv.Atoi(r.Form.Get("idAccount"))
	idRoom, _ := strconv.Atoi(r.Form.Get("idRoom"))

	db := connectgorm()

	user := m.Participant{IdAccount: idAccount, IdRoom: idRoom}

	var room m.Room
	var game m.Game
	var participants []m.Participant

	db.First(&game, room.IdGame)

	if err := db.First(&room, idRoom).Error; err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	result := db.Where("id_room = ?", room.ID).Find(&participants)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if len(participants) >= game.Max_player {

		return
	}

	w.Header().Set("Content-Type", "application/json")
	var response = m.GameResponse{}
	response.Status = http.StatusOK
	response.Message = "Success"
	json.NewEncoder(w).Encode(response)
}
