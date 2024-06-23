package domain

import "time"

// Typing Rating for age partners
type Rating string

const (
	RatingLivre Rating = "L"
	Raiting10   Rating = "L10"
	Raiting12   Rating = "L12"
	Raiting14   Rating = "L14"
	Raiting16   Rating = "L16"
	Raiting18   Rating = "L18"
)

type Event struct {
	ID           string
	Name         string
	Location     string
	Organization string
	Rating       Rating
	Date         time.Time
	ImageURL     string
	Capacity     int
	Price        float64
	PartnerID    int
	Spots        []Spot
	Tickets      []Ticket
}
