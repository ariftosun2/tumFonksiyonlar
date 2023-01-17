package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const loadNerve = 200

type propertyFeatures struct {
	load         int
	price        int
	propertyname string
}
type propertyRespons struct {
	loadRespons         int
	priceRespons        int
	averageWeight       float64
	propertynameRespons string
}

func transportNerve(transport []propertyRespons) []propertyRespons {
	amountCarried := 0
	var goodsReceived []propertyRespons
	for _, v := range transport {
		if amountCarried != loadNerve {
			amountCarried += v.loadRespons
			goodsReceived = append(goodsReceived, v)
		}
	}
	return goodsReceived
}

func averageCalculation(calcultion []propertyRespons) []propertyRespons {

	goodsMap := make(map[propertyRespons]float64)
	for _, p := range calcultion {
		k := propertyRespons{loadRespons: p.loadRespons, priceRespons: p.priceRespons, propertynameRespons: p.propertynameRespons}
		goodsMap[k] = p.averageWeight
	}

	var averageSort []float64
	for _, v := range goodsMap {
		averageSort = append(averageSort, v)
		sort.Float64s(averageSort)
	}
	var calculationSort []propertyRespons
	for i, v := range goodsMap {
		for _, b := range averageSort {
			if v == b {
				calculationSort = append(calculationSort, i)
			}
		}
	}

	return calculationSort

}

func inverseProportion(property []propertyFeatures) []propertyRespons {
	var propertyWeight []propertyRespons

	for _, v := range property {
		var averageValue float64
		averageValue = float64(v.load / v.price)
		propertyWeight = append(propertyWeight, propertyRespons{
			loadRespons:         v.load,
			priceRespons:        v.price,
			averageWeight:       averageValue,
			propertynameRespons: v.propertyname})

	}
	return propertyWeight
}

func ReadCsvFile(fileName string) ([]propertyFeatures, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var personList []propertyFeatures
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		record := strings.Split(line, ",")
		loads, _ := strconv.Atoi(record[0])
		prices, _ := strconv.Atoi(record[1])
		personList = append(personList, propertyFeatures{load: loads, price: prices, propertyname: record[2]})
	}
	return personList, nil
}
func WriteFileCsv(employee []propertyRespons, resultFileName string) error {
	file, err := os.OpenFile(resultFileName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	dataWriter := bufio.NewWriter(file)
	for _, h := range employee {
		s := fmt.Sprintf("%s,\n", h.propertynameRespons)
		_, err = dataWriter.WriteString(s)
		if err != nil {
			return err
		}
	}

	dataWriter.Flush()
	return nil
}

func main() {
	file, err := ReadCsvFile("mallarlistesi.csv")
	if err != nil {
		log.Fatal(err)
	}
	inverse := inverseProportion(file)
	averega := averageCalculation(inverse)
	trans := transportNerve(averega)
	WriteFileCsv(trans, "alinacakMallar")
}
