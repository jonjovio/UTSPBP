package models

type Game struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Max_player int    `json:"max player"`
}

type GameResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Game   `json:"data"`
}

type GamesResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Game `json:"data"`
}
