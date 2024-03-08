package models

type Accounts struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type AccountResponses struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    []Accounts `json:"data"`
}

type Games struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	MaxPlayer int    `json:"maxplayer"`
}

type GameResponses struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Games `json:"data"`
}

type Rooms struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	GameID int    `json:"gameid"`
}

type RoomResponses struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Rooms `json:"data"`
}

type Participants struct {
	ID        int `json:"id"`
	RoomID    int `json:"roomid"`
	AccountID int `json:"accountid"`
}

type ParticipantResponse struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    Participants `json:"data"`
}

type DetailRoom struct {
	ID           int        `json:"id"`
	Name         string     `json:"room_name"`
	Participants []Accounts `json:"participants"`
}

type DetailRoomResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    DetailRoom `json:"rooms"`
}

type GlobalResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
