package lasagna

// TODO: define the 'OvenTime()' function
func OvenTime () int {
    return 40
}
// TODO: define the 'RemainingOvenTime()' function

func RemainingOvenTime(timeInOven int) int {
    return OvenTime() - timeInOven
}
// TODO: define the 'PreparationTime()' function
func PreparationTime (numberOfLayers int) int {
    return numberOfLayers * 2
}
// TODO: define the 'ElapsedTime()' function
func ElapsedTime (numberOfLayers int, MinutesInOven int) int {
	return PreparationTime(numberOfLayers) + MinutesInOven
}