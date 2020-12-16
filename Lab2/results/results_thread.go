package results

import "sort"

func StartResults(workerChannel <-chan *Automobile, mainChannel chan<- []Automobile) {
	var resultsStorage []Automobile
	var workersFinished int

	for workersFinished != WorkerThreadCount {
		result := <-workerChannel
		if result != nil {
			resultsStorage = insertSorted(resultsStorage, *result)
		} else {
			workersFinished++
		}
	}

	mainChannel <- resultsStorage
}

func insertSorted(slice []Automobile, element Automobile) []Automobile {
	i := sort.Search(len(slice), func(i int) bool { return slice[i].Price < element.Price })
	slice = append(slice, Automobile{Price: 0})
	copy(slice[i+1:], slice[i:])
	slice[i] = element
	return slice
}