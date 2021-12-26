package service

import (
	"test/dao"
	"test/model"
)

func GetRecordByTxt(txt string) (model.Record, error) {
	return dao.GetRecordByTxt(txt)
}
