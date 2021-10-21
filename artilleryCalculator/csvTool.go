package artilleryCalculator

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func GetCSVFileData(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()
	rdr := csv.NewReader(bufio.NewReader(file))
	rows, err := rdr.ReadAll()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return rows, nil
}