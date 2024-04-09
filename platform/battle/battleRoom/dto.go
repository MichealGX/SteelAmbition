package battleRoom

type RoomDTO struct {
	UserID      int
	UserName    string
	VehicleID   int
	VehicleName string
	ReadyFlag   bool
}

type RoomListDTO struct {
	RoomID int
	MaxNum int `json:"maxNum"`
}
