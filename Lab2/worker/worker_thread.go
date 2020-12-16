package worker

func StartWorker(dataRequestChannel chan<- int, dateReceiveChannel <-chan *Automobile, resultsChannel chan<- *Automobile, number int) {
	finished := false
	for !finished {
		dataRequestChannel <- number
		response := <-dateReceiveChannel

		if response != nil {
			response.CalculateAutomobileHash()

			if response.Price <= MaxAutoPrice {
				resultsChannel <- response
			}
		} else {
			resultsChannel <- response
			finished = true
		}
	}
}