package results

import (
	"sort"

	"github.com/Tyomnat/LygProgLab2/Lab2/automobile"
	"github.com/Tyomnat/LygProgLab2/Lab2/constants"
)


func StartResults(workerChannel <-chan *automobile.Automobile, mainChannel chan<- []automobile.Automobile) {
	var resultsStorage []automobile.Automobile
	var workersFinished int

	for workersFinished != constants.WorkerThreadCount {
		result := <-workerChannel
		if result != nil {
			resultsStorage = insertSorted(resultsStorage, *result)
		} else {
			workersFinished++
		}
	}

	mainChannel <- resultsStorage
}

func insertSorted(slice []automobile.Automobile, element automobile.Automobile) []automobile.Automobile {
	i := sort.Search(len(slice), func(i int) bool { return slice[i].Price < element.Price })
	slice = append(slice, automobile.Automobile{Price: 0})
	copy(slice[i+1:], slice[i:])
	slice[i] = element
	return slice
}