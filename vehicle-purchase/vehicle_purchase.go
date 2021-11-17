package purchase

import (
	"fmt"
)

// NeedsLicense determines whether a license is need to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return (kind == "car" || kind == "truck")
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	var option string
	if option1 < option2 {
		option = option1
	} else {
		option = option2
	}

	return fmt.Sprintf("%s is clearly the better choice.", option)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	return originalPrice * AgeDepreciation(age)
}

func AgeDepreciation(age float64) (result float64) {
	if age < 3 {
		result = 0.8
	} else if age < 9 {
		result = 0.7
	} else {
		result = 0.5
	}
	return result
}
