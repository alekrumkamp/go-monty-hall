package main

import (
	"math/rand"
	"fmt"
	"time"
)

var sourcePseudoRandom = rand.NewSource(time.Now().UnixNano())

func chooseParticipantsSecondDoor(doors int,participantsDoor int ,hostDoor int) int  {
	participantsSecondDoor := 0
	for {
		participantsSecondDoor = getPseudoRandomDoor(doors)
		if participantsSecondDoor != participantsDoor && participantsSecondDoor != hostDoor {
			break
		}
	}
	return participantsSecondDoor
}

func getPseudoRandomDoor(doors int) int {
	return rand.New(sourcePseudoRandom).Intn(doors)
}

func chooseHostDoor(doors int,participantsDoor int,prizeDoor int) int {
	hostDoor := 0
	for {
		hostDoor = getPseudoRandomDoor(doors)
		if hostDoor != participantsDoor && hostDoor != prizeDoor {
			break
		}
	}
	return hostDoor
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
		hostDoor = chooseHostDoor(doors,participantsDoor,prizeDoor)

		if(participantChangesDoors){
			participantsDoor = chooseParticipantsSecondDoor(doors,participantsDoor,hostDoor)
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
