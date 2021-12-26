package service

import (
	"test/dao"
	"test/model"
)

func AddRecord(record model.Record) error {
	err := dao.AddRecord(record)
	return err
}
