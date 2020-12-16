package main

import (
	"fmt"
	"time"
	"automobile"
)

func main() {
	AutomobilesData, err := automobile.ReadAutomobilesJsonData(AllDataPassFile)
	if err != nil {
		panic(err)
	}

	mainToDataChannel := make(chan *Automobile)
	dataToWorkersChannel := make(chan *Automobile)
	workersToDataChannel := make(chan int)
	workersToResultsChannel := make(chan *Automobile)
	resultsToMainChannel := make(chan []Automobile)

	go StartData(mainToDataChannel, workersToDataChannel, dataToWorkersChannel)
	go StartResults(workersToResultsChannel, resultsToMainChannel)

	for i := 1; i <= WorkerThreadCount; i++ {
		go StartWorker(workersToDataChannel, dataToWorkersChannel, workersToResultsChannel, i)
	}

	startTime := time.Now()
	for _, Automobile := range AutomobilesData {
		AutomobileCopy := Automobile // https://stackoverflow.com/questions/49123133/sending-pointers-over-a-channel/49125045#49125045
		mainToDataChannel <- &AutomobileCopy
	}

	// signaling that we've finished
	mainToDataChannel <- nil
	fmt.Printf("Results received in %f seconds.\n", time.Since(startTime).Seconds())

	result := <-resultsToMainChannel
	WriteResultsToFile(result, ResultsFile)

	fmt.Println("The application has finished.")
}