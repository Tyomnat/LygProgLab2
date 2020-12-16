package automobile

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Automobile struct {
	Name   string  `json:"pavadinimas"`
	Year   int     `json:"metai"`
	Price float64 `json:"kaina"`
}

func (a Automobile) CalculateAutomobileHash() string {
	return string(sha1.New().Sum([]byte(fmt.Sprintf("%s %d %f", a.Name, a.Year, a.Price))))
}

func ReadAutomobilesJsonData(fileName string) ([]Automobile, error) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	automobiles := make([]Automobile, 0)
	err = json.Unmarshal(jsonData, &automobiles)
	if err != nil {
		return nil, err
	}

	return automobiles, nil
}

func WriteResultsToFile(results []Automobile, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	file.WriteString(fmt.Sprintf("%3s %20s %5s %6s\n", "No.", "Name", "Year", "Price"))
	for i, r := range results {
		file.WriteString(fmt.Sprintf("%3d %20s %5d %6.2f\n", i+1, r.Name, r.Year, r.Price))
	}

	return nil
}