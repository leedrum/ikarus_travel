package services

import (
	"encoding/csv"
	"mime/multipart"

	"github.com/leedrum/ikarus_travel/model"
	"gorm.io/gorm"
)

const (
	// TourImportLimit is the number of tours to import at once
	TourImportLimit = 500
)

func ImportTours(db *gorm.DB, file *multipart.FileHeader) error {
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

	tours := []model.Tour{}
	for index, record := range records {
		// skip the header
		if index == 0 {
			continue
		}
		tour := model.Tour{
			Name:        record[0],
			Description: record[1],
			Status:      model.TourStatusActive,
		}
		tours = append(tours, tour)

		if (len(records) == index+1) || len(tours) == TourImportLimit {
			db.Create(&tours)
			tours = []model.Tour{}
		}
	}

	return nil
}
