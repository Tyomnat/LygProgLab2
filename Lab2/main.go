package main

import (
	"fmt"
	"time"

	"github.com/Tyomnat/LygProgLab2/Lab2/automobile"
	"github.com/Tyomnat/LygProgLab2/Lab2/constants"
	"github.com/Tyomnat/LygProgLab2/Lab2/data"
	"github.com/Tyomnat/LygProgLab2/Lab2/results"
	"github.com/Tyomnat/LygProgLab2/Lab2/worker"
)

func main() {
	AutomobilesData, err := automobile.ReadAutomobilesJsonData(constants.HalfDataPassFile)
	if err != nil {
		panic(err)
	}

	mainToDataChannel := make(chan *automobile.Automobile)
	dataToWorkersChannel := make(chan *automobile.Automobile)
	workersToDataChannel := make(chan int)
	workersToResultsChannel := make(chan *automobile.Automobile)
	resultsToMainChannel := make(chan []automobile.Automobile)

	go data.StartData(mainToDataChannel, workersToDataChannel, dataToWorkersChannel)
	go results.StartResults(workersToResultsChannel, resultsToMainChannel)

	for i := 1; i <= constants.WorkerThreadCount; i++ {
		go worker.StartWorker(workersToDataChannel, dataToWorkersChannel, workersToResultsChannel, i)
	}

	startTime := time.Now()
	for _, automobile := range AutomobilesData {
		automobileCopy := automobile // https://stackoverflow.com/questions/49123133/sending-pointers-over-a-channel/49125045#49125045
		mainToDataChannel <- &automobileCopy
	}

	// signaling that we've finished
	mainToDataChannel <- nil
	fmt.Printf("Results received in %f seconds.\n", time.Since(startTime).Seconds())

	result := <-resultsToMainChannel
	automobile.WriteResultsToFile(result, constants.ResultsFile)

	fmt.Println("The application has finished.")
}