package models

type Guest struct {
	Name                string `json:"name"`
	Table               int64  `json:"table"`
	Accompanying_guests int64  `json:"accompanying_guests"`
}

type PartyGuest struct {
	Name                string `json:"name"`
	Time_arrived        string `json:"time_arrived"`
	Accompanying_guests int64  `json:"accompanying_guests"`
}

type GuestResponse struct {
	Name string `json:"name"`
}