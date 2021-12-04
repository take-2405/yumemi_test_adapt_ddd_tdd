package repository

import (
	"encoding/csv"
	"yumemi-coding-test-practice/src/domain"
)

type PlayDataRepository interface {
	ParsePlayData(csv *csv.Reader) (domain.PlayDatas, error)
	ReadPlayData(filePath string) (*csv.Reader, error)
}
