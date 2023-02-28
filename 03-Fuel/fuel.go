package main

import "fmt"

// Create the function fuelGauge() here
func fuelGauge(fuel int) {
	fmt.Printf("You have %d gallons of fuel left!\n", fuel)
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {
	var fuel int
	switch planet {
	case "Venus":
		fuel = 300000
	case "Mercury":
		fuel = 500000
	case "Mars":
		fuel = 700000
	case "Earth":
		fuel = 400000
	default:
		fuel = 0
	}
	fmt.Printf("Fuel required for the planet %v is %d \n", planet, fuel)
	return fuel
}

// Create the function greetPlanet() here
func greetPlanet(planet string) {
	fmt.Println("Welcome to planet", planet)
}

// Create the function cantFly() here
func cantFly() {
	fmt.Println("We do not have the available fuel to fly there.")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int {
	var fuelRemaining, fuelCost int
	fuelRemaining = fuel
	fuelCost = calculateFuel(planet)
	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining -= fuelCost
	} else {
		cantFly()
	}
	return fuelRemaining
}

func flyToHome(planet string, fuel int) int {
	var returnFuel, fuelCost int
	fuelCost = calculateFuel(planet)
	if fuel >= fuelCost {
		greetPlanet(planet)
		returnFuel = fuel - fuelCost
	} else {
		cantFly()
		returnFuel = fuel
	}
	return returnFuel
}

func main() {
	fuel := 1000000
	planetChoice := "Venus"
	homePlanet := "Earth"
	// fuelToHome := 500000

	fuel = flyToPlanet(planetChoice, fuel)
	fuelGauge(fuel)
	fuel = flyToHome(homePlanet, fuel)
	fuelGauge(fuel)
}
