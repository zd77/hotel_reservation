package db

const (
	DBURI      = "mongodb://localhost:27017"
	DBNAME     = "hotel_reservation"
	DBNAMETEST = "hotel_reservation_test"
)

type Store struct {
	Hotel HotelStore
	User  UserStore
	Room  RoomStore
}
