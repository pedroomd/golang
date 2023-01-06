package models
 
type Guest struct {
    Name  string `json:"name"`
    Table Table `json:"table"`
	Accompanying_guests int `json:"accompanying_guests"`
}