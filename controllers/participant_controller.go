package controllers

import (
	m "UTS/models"
	"encoding/json"
	"net/http"
	"strconv"
)

func InsertRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		SendErrorResponse(w, 400, "Invalid request")
		return
	}

	account, _ := strconv.Atoi(r.Form.Get("accountid"))
	room, _ := strconv.Atoi(r.Form.Get("roomid"))
	var response m.ParticipantResponse
	maxed := false

	if account == 0 || room == 0 {
		SendErrorResponse(w, 400, "Invalid request")
		return
	}

	rows, err := db.Query("SELECT id_game FROM rooms WHERE id = ?", room)
	if err != nil {
		SendErrorResponse(w, 500, "Server failed to get room")
		return
	}
	defer rows.Close()
	for rows.Next() {
		var gameid int
		if err := rows.Scan(&gameid); err != nil {
			SendErrorResponse(w, 500, "Server failed to get room")
			return
		}
		rows, err = db.Query("SELECT max_player FROM games WHERE id = ?", gameid)
		if err != nil {
			SendErrorResponse(w, 500, "Server failed to get game")
			return
		}
		defer rows.Close()
		for rows.Next() {
			var maxplayer int
			if err := rows.Scan(&maxplayer); err != nil {
				SendErrorResponse(w, 500, "Server failed to get game")
				return
			}
			rows, err = db.Query("SELECT COUNT(*) FROM participants WHERE id_room = ?", room)
			if err != nil {
				SendErrorResponse(w, 500, "Server failed to get participants")
				return
			}
			defer rows.Close()
			for rows.Next() {
				var count int
				if err := rows.Scan(&count); err != nil {
					SendErrorResponse(w, 500, "Server failed to get participants")
					return
				}
				if count >= maxplayer {
					maxed = true
				}
			}
		}
	}

	if !maxed {
		_, err = db.Exec("INSERT INTO participants (id_room, id_account) VALUES (?, ?)", room, account)
		if err != nil {
			SendErrorResponse(w, 500, "Server failed to insert room")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		response.Status = 200
		response.Message = "Success, Participant has been added to room"
		response.Data.AccountID = account
		response.Data.RoomID = room
		json.NewEncoder(w).Encode(response)
	} else {
		SendErrorResponse(w, 500, "Room is full")
	}
}
