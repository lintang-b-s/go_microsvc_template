package domain


type CarType string
const (
	MPV CarType = "MPV"
	SUV CarType = "SUV"
	Hatchback CarType = "Hatchback"
	Sedan CarType = "Sedan"
	Compact CarType = "Compact"
	Van CarType = "Van"
	Minibus CarType = "Minibus"
	PickUP CarType = "Pick-Up"
	Truk CarType = "Truk"
	DoubleCabin CarType = "Double Cabin"
	Wagon CarType = "Wagon"
	Coupe CarType = "Coupe"
	Jeep CarType = "Jeep"
	Convertible CarType = "Convertible"
	Offroad CarType = "Offroad"
	Sports CarType = "Sports"
	Classic CarType = "Classic"
	Bus CarType = "Bus"
)

type Car struct {
	ID uint64 `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	CarType CarType `json:"car_type"`
}
