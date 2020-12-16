package worker

import (
	"github.com/Tyomnat/LygProgLab2/Lab2/automobile"
	"github.com/Tyomnat/LygProgLab2/Lab2/constants"
)

func StartWorker(dataRequestChannel chan<- int, dateReceiveChannel <-chan *automobile.Automobile, resultsChannel chan<- *automobile.Automobile, number int) {
	finished := false
	for !finished {
		dataRequestChannel <- number
		response := <-dateReceiveChannel

		if response != nil {
			response.CalculateAutomobileHash()

			if response.Price <= constants.MaxAutoPrice {
				resultsChannel <- response
			}
		} else {
			resultsChannel <- response
			finished = true
		}
	}
}