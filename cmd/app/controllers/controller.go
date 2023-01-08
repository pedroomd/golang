package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	config "github.com/getground/tech-tasks/backend/cmd/app/configs"
	model "github.com/getground/tech-tasks/backend/cmd/app/models"
	"github.com/gorilla/mux"
)

func InsertTable(w http.ResponseWriter, r *http.Request) {

    db := config.Connect()
    defer db.Close()
    
    decoder := json.NewDecoder(r.Body)
    var tableData model.Table

    err := decoder.Decode(&tableData)
    if err != nil {
        panic(err)
    }
  
    result , err_insert := db.Exec("INSERT INTO tables(capacity) VALUES (?)", tableData.Capacity)
    
    if err_insert != nil {
        log.Print(err_insert)
        return
    }

    tableData.Id, _ = result.LastInsertId()
 
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(tableData)
}

func InsertGuest(w http.ResponseWriter, r *http.Request) {
    var guestData model.Guest
    var guestResponse model.GuestResponse

    db := config.Connect()
    defer db.Close()
    
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&guestData)

    if err != nil {
        panic(err)
    }


    vars := mux.Vars(r)
    name := vars["name"]

    row, err := db.Query("SELECT * FROM tables WHERE id = ? AND capacity >= ?", guestData.Table, guestData.Accompanying_guests)

    if err != nil {
        log.Print(err)
    }

    if !row.Next(){
        w.WriteHeader(http.StatusConflict) 
        return 
    }

    _ , err_insert := db.Exec("INSERT INTO guests(name, accompanying_guests, table_id) VALUES (?, ?, ?)", name, guestData.Accompanying_guests, guestData.Table)
    
    if err_insert != nil {
        log.Print(err_insert)
        return
    }
    
    guestResponse.Name = name
 
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(guestResponse)
}

func GetGuestlist(w http.ResponseWriter, r *http.Request) {
    var guestlist []model.Guest
    var guest model.Guest

    db := config.Connect()
    defer db.Close()

    rows, err := db.Query("SELECT name, table_id, accompanying_guests FROM guests")

    if err != nil {
        log.Print(err)
    }

    for rows.Next() {
        err = rows.Scan(&guest.Name, &guest.Table, &guest.Accompanying_guests)
        if err != nil {
            log.Fatal(err.Error())
        } else {
            guestlist = append(guestlist, guest)
        }
    }

    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(guestlist)
}

func CheckInGuest(w http.ResponseWriter, r *http.Request) {

    var partyGuest model.Guest
    var guestResponse model.GuestResponse
    var guest_id int64

    db := config.Connect()
    defer db.Close()

    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&partyGuest)

    if err != nil {
        panic(err)
    }

    vars := mux.Vars(r)
    name := vars["name"]
  

    row, err := db.Query("SELECT guests.id FROM tables, guests WHERE guests.name LIKE ? AND capacity >= ?", name, partyGuest.Accompanying_guests + 1)

    if err != nil {
        log.Print(err)
    }

    //should return only one value
    if row.Next() {
        err = row.Scan(&guest_id)
        if err != nil {
            log.Fatal(err.Error())
            return
        }
        _ , err_insert := db.Exec("INSERT INTO party_guests(guest_id, time_arrived, accompanying_guests) VALUES (?, NOW(), ?)", guest_id, partyGuest.Accompanying_guests) 
        if err_insert != nil {
            log.Print(err_insert)
            return
        }
        guestResponse.Name = name
    } else {
        w.WriteHeader(http.StatusConflict) 
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(guestResponse)
}


func DeleteGuest(w http.ResponseWriter, r *http.Request) {

    db := config.Connect()
    defer db.Close()

    vars := mux.Vars(r)
    name:= vars["name"]

    _ , err := db.Exec("DELETE FROM party_guests WHERE guest_id IN (SELECT id FROM guests WHERE name LIKE ?)", name)

    if err != nil {
        log.Print(err)
        return
    }

    
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusNoContent) 
}

func GetPartyGuests(w http.ResponseWriter, r *http.Request) {
    var partyGuests []model.PartyGuest
    var partyGuest model.PartyGuest

    db := config.Connect()
    defer db.Close()

    rows, err := db.Query("SELECT guests.name, DATE_FORMAT(party_guests.time_arrived,'%H:%i:%s'), party_guests.accompanying_guests FROM party_guests, guests")

    if err != nil {
        log.Print(err)
    }

    for rows.Next() {
        err = rows.Scan(&partyGuest.Name, &partyGuest.Time_arrived, &partyGuest.Accompanying_guests)
        if err != nil {
            log.Fatal(err.Error())
        } else {
            partyGuests = append(partyGuests, partyGuest)
        }
    }

    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(partyGuests)
}


func GetSeatsEmpty(w http.ResponseWriter, r *http.Request) {
    var seatsResponse model.SeatsResponse

    db := config.Connect()
    defer db.Close()

    row := db.QueryRow("SELECT (SELECT SUM(capacity) FROM tables) - (SELECT IFNULL(SUM(accompanying_guests) + 1, 0) FROM party_guests) AS empty_seats")

    if row.Err() != nil {
        log.Print(row.Err())
    }

    err := row.Scan(&seatsResponse.Empty_seats)

    if err != nil {
        log.Fatal(err.Error())
    }

    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(seatsResponse)
}