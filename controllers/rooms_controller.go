package controllers

import (
	m "UTS/models"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM rooms")
	if err != nil {
		SendErrorResponse(w, 500, "Server failed to get rooms")
		return
	}
	defer rows.Close()

	var rooms []m.Rooms
	for rows.Next() {
		var room m.Rooms
		if err := rows.Scan(&room.ID, &room.Name, &room.GameID); err != nil {
			SendErrorResponse(w, 500, "Server failed to get rooms")
			return
		}
		rooms = append(rooms, room)
	}

	var response m.RoomResponses
	w.Header().Set("Content-Type", "application/json")
	response.Status = 200
	response.Message = "Success"
	response.Data = rooms
	json.NewEncoder(w).Encode(response)
}

func GetDetailRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		SendErrorResponse(w, 400, "Invalid request")
		return
	}

	roomid, _ := strconv.Atoi(r.URL.Query().Get("roomid"))
	if roomid == 0 {
		SendErrorResponse(w, 400, "Invalid request")
		return
	}
	query := "select r.id, r.room_name, a.id, a.username from rooms r join participants on id_room = r.id join accounts a on id_account = a.id where r.id = " + strconv.Itoa(roomid)
	rows, err := db.Query(query)
	if err != nil {
		SendErrorResponse(w, 500, "Server failed to get detail room")
		return
	}
	defer rows.Close()

	var room m.DetailRoom
	for rows.Next() {
		var account m.Accounts
		if err := rows.Scan(&room.ID, &room.Name, &account.ID, &account.Username); err != nil {
			SendErrorResponse(w, 500, "Server failed to get detail room")
			return
		}
	}

	var response m.DetailRoomResponse
	w.Header().Set("Content-Type", "application/json")
	response.Status = 200
	response.Message = "Success"
	response.Data = room
	json.NewEncoder(w).Encode(response)
}
