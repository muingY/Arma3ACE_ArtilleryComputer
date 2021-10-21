package artilleryCalculator

import (
	"fmt"
	"strconv"
	"strings"
)

var M109A6Database M109A6RangeTable

func InitializeM109A6RangeTable() {
	var csvData [][]string
	var newTableCell RangeTableCell
	var err error

	var highFileList [5]string = [5]string{
		"./rangeTables/M109A6_rangeTable_high_0.csv",
		"./rangeTables/M109A6_rangeTable_high_1.csv",
		"./rangeTables/M109A6_rangeTable_high_2.csv",
		"./rangeTables/M109A6_rangeTable_high_3.csv",
		"./rangeTables/M109A6_rangeTable_high_4.csv",
	}
	for ch, filename := range highFileList {
		csvData, err = GetCSVFileData(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		for i := 1; i < len(csvData); i++ {
			newTableCell.Distance, _ = strconv.ParseFloat(strings.Trim(csvData[i][0], " "), 64)
			newTableCell.AimMill, _ = strconv.ParseFloat(strings.Trim(csvData[i][1], " "), 64)
			newTableCell.DeltaMillPer100m, _ = strconv.ParseFloat(strings.Trim(csvData[i][2], " "), 64)
			newTableCell.Eta, _ = strconv.ParseFloat(strings.Trim(csvData[i][3], " "), 64)

			M109A6Database.RangeTableHigh[ch] = append(M109A6Database.RangeTableHigh[ch], newTableCell)
		}
	}

	var lowFileList [5]string = [5]string{
		"./rangeTables/M109A6_rangeTable_low_0.csv",
		"./rangeTables/M109A6_rangeTable_low_1.csv",
		"./rangeTables/M109A6_rangeTable_low_2.csv",
		"./rangeTables/M109A6_rangeTable_low_3.csv",
		"./rangeTables/M109A6_rangeTable_low_4.csv",
	}
	for ch, filename := range lowFileList {
		csvData, err = GetCSVFileData(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		for i := 1; i < len(csvData); i++ {
			newTableCell.Distance, _ = strconv.ParseFloat(strings.Trim(csvData[i][0], " "), 64)
			newTableCell.AimMill, _ = strconv.ParseFloat(strings.Trim(csvData[i][1], " "), 64)
			newTableCell.DeltaMillPer100m, _ = strconv.ParseFloat(strings.Trim(csvData[i][2], " "), 64)
			newTableCell.Eta, _ = strconv.ParseFloat(strings.Trim(csvData[i][3], " "), 64)

			M109A6Database.RangeTableLow[ch] = append(M109A6Database.RangeTableLow[ch], newTableCell)
		}
	}
}

// M109A6GetShotSolution mode(1) = high, (0) = low
func M109A6GetShotSolution(distance float64, diffAlt float64, mode int) []ShotResultData {
	var result []ShotResultData
	var temp ShotResultData

	if mode == 1 {
		for ch, dataTable := range M109A6Database.RangeTableHigh {
			for i := 0; i < len(dataTable) - 1; i++ {
				if (distance >= dataTable[i].Distance) && (distance < dataTable[i + 1].Distance) {
					millPer1m := (dataTable[i + 1].AimMill - dataTable[i].AimMill) / (dataTable[i + 1].Distance - dataTable[i].Distance)
					temp.Charge = ch
					temp.AimMill = dataTable[i].AimMill + ((distance - dataTable[i].Distance) * millPer1m) + (dataTable[i].DeltaMillPer100m * (-diffAlt / 100))
					temp.Eta = dataTable[i].Eta
					result = append(result, temp)
					break
				}
			}
			if distance > dataTable[len(dataTable) - 1].Distance && distance < (dataTable[len(dataTable) - 1].Distance + 100) {
				millPer1m := -((dataTable[len(dataTable) - 2].AimMill - dataTable[len(dataTable) - 1].AimMill) / (dataTable[len(dataTable) - 2].Distance - dataTable[len(dataTable) - 1].Distance))
				temp.Charge = ch
				temp.AimMill = (dataTable[len(dataTable) - 1].AimMill + ((distance - dataTable[len(dataTable) - 1].Distance) * millPer1m)) + (dataTable[len(dataTable) - 1].DeltaMillPer100m * (-diffAlt / 100))
				temp.Eta = dataTable[len(dataTable) - 1].Eta
				result = append(result, temp)
			}
		}
	} else if mode == 0 {
		for ch, dataTable := range M109A6Database.RangeTableLow {
			for i := 0; i < len(dataTable) - 1; i++ {
				if (distance >= dataTable[i].Distance) && (distance < dataTable[i + 1].Distance) {
					millPer1m := (dataTable[i + 1].AimMill - dataTable[i].AimMill) / (dataTable[i + 1].Distance - dataTable[i].Distance)
					temp.Charge = ch
					temp.AimMill = dataTable[i].AimMill + ((distance - dataTable[i].Distance) * millPer1m) + (dataTable[i].DeltaMillPer100m * (-diffAlt / 100))
					temp.Eta = dataTable[i].Eta
					result = append(result, temp)
					break
				}
			}
			if distance > dataTable[len(dataTable) - 1].Distance && distance < (dataTable[len(dataTable) - 1].Distance + 100) {
				millPer1m := -((dataTable[len(dataTable) - 2].AimMill - dataTable[len(dataTable) - 1].AimMill) / (dataTable[len(dataTable) - 2].Distance - dataTable[len(dataTable) - 1].Distance))
				temp.Charge = ch
				temp.AimMill = (dataTable[len(dataTable) - 1].AimMill + ((distance - dataTable[len(dataTable) - 1].Distance) * millPer1m)) + (dataTable[len(dataTable) - 1].DeltaMillPer100m * (-diffAlt / 100))
				temp.Eta = dataTable[len(dataTable) - 1].Eta
				result = append(result, temp)
			}
		}
	}
	return result
}