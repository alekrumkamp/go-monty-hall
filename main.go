package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {
	sourcePseudoRandom := rand.NewSource(time.Now().UnixNano())

	// In this version of the code only three doors can be benchmarked for the sake of simplicity
	const doors = 3
	const numberOfExperiments = 1000000
	const participantChangesDoors = true
	hostDoor := 0
	losers := 0
	participantsDoor := 0
	participantsSecondDoor := 0
	prizeDoor := 0
	winners := 0

	i := 0
	for i < numberOfExperiments {

		// Assign random doors to prize and to participant
		prizeDoor = rand.New(sourcePseudoRandom).Intn(doors)
		participantsDoor = rand.New(sourcePseudoRandom).Intn(doors)

		// Choose Host door
		for {
			hostDoor = rand.New(sourcePseudoRandom).Intn(doors)
			if hostDoor != participantsDoor && hostDoor != prizeDoor {
				break
			}
		}

		if(participantChangesDoors){
			// Model change of participant's door choice
			for {
				participantsSecondDoor = rand.New(sourcePseudoRandom).Intn(doors)
				if participantsSecondDoor != participantsDoor && participantsSecondDoor != hostDoor {
					break
				}
			}
			participantsDoor = participantsSecondDoor
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
