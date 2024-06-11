package utils

import (
	"fmt"
	"os"

	"resto_go/types"

	"github.com/gocarina/gocsv"
)

func ReadFile(path string) ([]types.MerchantInfo, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return []types.MerchantInfo{}, fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	var merchants []types.MerchantInfo
	if err := gocsv.UnmarshalFile(file, &merchants); err != nil {
		return []types.MerchantInfo{}, fmt.Errorf("could not read file: %v", err)
	}
	return merchants, nil
}
