package data

import(
	"github.com/Tyomnat/LygProgLab2/Lab2/automobile"
	"github.com/Tyomnat/LygProgLab2/Lab2/constants"
)

func StartData(dataInputChannel <-chan *automobile.Automobile, requestsChannel <-chan int, dataOutputChannel chan<- *automobile.Automobile) {
	var dataStorage []automobile.Automobile
	var inputFinished bool
	var outputFinished int

	for outputFinished != constants.WorkerThreadCount {
		if len(dataStorage) == 0 && !inputFinished {
			movie := <-dataInputChannel
			if movie != nil {
				dataStorage = append(dataStorage, *movie)
			} else {
				inputFinished = true
			}

			continue
		}

		if len(dataStorage) == constants.AllowedDataCount {
			<-requestsChannel
			if !inputFinished || len(dataStorage) != 0 {
				dataOutputChannel <- &dataStorage[0]
				dataStorage = dataStorage[1:]
			} else {
				dataOutputChannel <- nil
				outputFinished++
			}

			continue
		}

		select {
		case movie := <-dataInputChannel:
			if movie != nil {
				dataStorage = append(dataStorage, *movie)
			} else {
				inputFinished = true
			}

		case <-requestsChannel:
			if !inputFinished || len(dataStorage) != 0 {
				dataOutputChannel <- &dataStorage[0]
				dataStorage = dataStorage[1:]
			} else {
				dataOutputChannel <- nil
				outputFinished++
			}
		}
	}
}