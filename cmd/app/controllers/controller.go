package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	models 
)

// AllEmployee = Select Employee API
func AllEmployee(w http.ResponseWriter, r *http.Request) {
    var table Table
    var response model.Response
    var arrEmployee []model.Table
 
    db := config.Connect()
    defer db.Close()
 
    rows, err := db.Query("SELECT id, name, city FROM employee")
 
    if err != nil {  
        log.Print(err)
    }
 
    for rows.Next() {
        err = rows.Scan(&employee.Id, &employee.Name, &employee.City)
        if err != nil {
            log.Fatal(err.Error())
        } else {
            arrEmployee = append(arrEmployee, employee)
        }
    }
 
    response.Status = 200
    response.Message = "Success"
    response.Data = arrEmployee
 
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(response)
}
 
// InsertEmployee = Insert Employee API
func InsertTable(w http.ResponseWriter, r *http.Request) {
    var response model.AddTable
 
    db := config.Connect()
    defer db.Close()
 
    err := r.ParseMultipartForm(4096)
    if err != nil {
        panic(err)
    }
    capacity := r.FormValue("capacity")
 
    k, err = db.Exec("INSERT INTO table(capacity) VALUES(?)", capacity)
	fmt.Print(k)
 
    if err != nil {
        log.Print(err)
        return
    }

    //response.Status = 200
    //response.Message = "Table inserted successfully"
    fmt.Print("Inserting data to database")
 
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(response)
}