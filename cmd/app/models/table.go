package models
 
type Table struct {
    Id   string `form:"id" json:"id"`
    Capacity int `form:"capacity" json:"name"`
}



type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Table
}