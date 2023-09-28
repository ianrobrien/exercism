package purchase

import (
	"fmt"
	"strings"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return strings.ToLower(kind) == "car" || strings.ToLower(kind) == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var chosenVehicle string
	if option1 <= option2 {
		chosenVehicle = option1
	} else {
		chosenVehicle = option2
	}
	return fmt.Sprintf("%s is clearly the better choice.", chosenVehicle)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return originalPrice * .8
	}
	if age < 10 {
		return originalPrice * .7
	}
	return originalPrice * .5
}
