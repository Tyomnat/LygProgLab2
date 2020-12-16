package data

func StartData(dataInputChannel <-chan *Automobile, requestsChannel <-chan int, dataOutputChannel chan<- *Automobile) {
	var dataStorage []Automobile
	var inputFinished bool
	var outputFinished int

	for outputFinished != WorkerThreadCount {
		if len(dataStorage) == 0 && !inputFinished {
			movie := <-dataInputChannel
			if movie != nil {
				dataStorage = append(dataStorage, *movie)
			} else {
				inputFinished = true
			}

			continue
		}

		if len(dataStorage) == AllowedDataCount {
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