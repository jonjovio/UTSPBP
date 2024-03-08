package models

type Participant struct {
	ID        int `json:"id"`
	IdAccount int `json:"id account"`
	IdRoom    int `json:"id room"`
}

type ParticipantResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    Participant `json:"data"`
}

type ParticipantsResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []Participant `json:"data"`
}
