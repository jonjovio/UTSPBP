package models

type Room struct {
	ID       int    `json:"id"`
	RoomName string `json:"room name"`
	IdGame   int    `json:"id game"`
}

type RoomResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Room   `json:"data"`
}

type RoomResponseDetail struct {
	Status       int           `json:"status"`
	Message      string        `json:"message"`
	Data         Room          `json:"data"`
	Participants []Participant `json:"participants"`
}

type RoomsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Room `json:"data"`
}
