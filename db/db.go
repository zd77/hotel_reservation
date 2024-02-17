package db

const MongoDbNameEnvName = "MONGO_DB_NAME"

type Store struct {
	Hotel   HotelStore
	User    UserStore
	Room    RoomStore
	Booking BookingStore
}
