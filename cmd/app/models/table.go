package models

type Table struct {
	Id       int64 `json:"id"`
	Capacity int64 `json:"capacity"`
}

type SeatsResponse struct {
	Empty_seats int64 `json:"empty_seats"`
}