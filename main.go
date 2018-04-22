package main

import (
	"math/rand"
	"fmt"
	"time"
)

var sourcePseudoRandom = rand.NewSource(time.Now().UnixNano())

func containsCertainDoor(doors []int, aDoor int) bool {
	for _, currentDoor := range doors {
		if currentDoor == aDoor {
			return true
		}
	}
	return false
}

func chooseDoorDifferentThan(doors int, notEligibleDoors []int) int{
	chosenDoor := 0
	for {
		chosenDoor = getPseudoRandomDoor(doors)
		if !containsCertainDoor(notEligibleDoors,chosenDoor) {
			break
		}
	}
	return chosenDoor
}

func getPseudoRandomDoor(doors int) int {
	return rand.New(sourcePseudoRandom).Intn(doors)
}

func main() {

	// In this version of the code only three doors can be benchmarked for the sake of simplicity
	const doors = 3
	const numberOfExperiments = 1000000
	const participantChangesDoors = true
	hostDoor := 0
	losers := 0
	participantsDoor := 0
	prizeDoor := 0
	winners := 0

	i := 0
	for i < numberOfExperiments {

		// Assign random doors to prize and to participant
		prizeDoor = getPseudoRandomDoor(doors)
		participantsDoor = getPseudoRandomDoor(doors)

		// Choose Host door
		hostDoor = chooseDoorDifferentThan(doors,[]int {participantsDoor,prizeDoor})

		if(participantChangesDoors){
			participantsDoor = chooseDoorDifferentThan(doors,[]int {participantsDoor,hostDoor})
		}

		if participantsDoor == prizeDoor {
			winners++
		} else{
			losers++
		}

		i++
	}

	//Show results
	fmt.Println("Out of",numberOfExperiments,"participants, the results are:")
	fmt.Println("Winners:",float32(winners)/numberOfExperiments*100,"%")
	fmt.Println("Losers:",float32(losers)/numberOfExperiments*100,"%")

}
