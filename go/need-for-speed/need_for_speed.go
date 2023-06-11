package speed

type Car struct {
    battery			int
    batterDrain		int
    speed			int
    distance		int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car {
        battery:100,
        batteryDrain: batteryDrain,
        speed: speed,
    }
}

// TODO: define the 'Track' type struct

// NewTrack creates a new track
func NewTrack(distance int) Track {
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
}
