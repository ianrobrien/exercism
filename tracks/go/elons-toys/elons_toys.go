package elon

import "fmt"

func (car *Car) Drive() {
	if car.battery-car.batteryDrain < 0 {
		return
	}
	car.distance += car.speed
	car.battery -= car.batteryDrain
}

func (car *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

func (car *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

func (car *Car) CanFinish(trackDistance int) bool {
	availableDistance := (car.battery / car.batteryDrain) * car.speed
	return availableDistance >= trackDistance
}
