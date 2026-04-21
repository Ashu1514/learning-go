package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetBalanceFromFile(filename string) (float64, error){
	data, readErr := os.ReadFile(filename);

	if readErr != nil {
		return 0, errors.New("Failed to find file!")
	}

	valueText := string(data);
	value, parseErr := strconv.ParseFloat(valueText, 64)

	if parseErr != nil {
		return 0, errors.New("Failed to parse stored value!")
	}

	return value, nil
}

func WriteDataToFile(filename string, balance float64) {
	valueText := fmt.Sprint(balance)
	os.WriteFile(filename, []byte(valueText), 0644)
}