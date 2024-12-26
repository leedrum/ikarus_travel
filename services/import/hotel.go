package services

import (
	"encoding/csv"
	"mime/multipart"

	"github.com/leedrum/ikarus_travel/model"
	"gorm.io/gorm"
)

const (
	// HotelImportLimit is the number of hotels to import at once
	HotelImportLimit = 500
)

func ImportHotels(db *gorm.DB, file *multipart.FileHeader) error {
	f, err := file.Open()
	if err != nil {
		return err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	// reader.Comma = ';'
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	hotels := []model.Hotel{}
	for index, record := range records {
		// skip the header
		if index == 0 {
			continue
		}
		hotel := model.Hotel{
			Name:        record[0],
			Address:     record[1],
			Description: record[2],
		}
		hotels = append(hotels, hotel)

		if (len(records) == index+1) || len(hotels) == HotelImportLimit {
			db.Create(&hotels)
			hotels = []model.Hotel{}
		}
	}

	return nil
}
