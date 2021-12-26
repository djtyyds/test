package dao

import "test/model"

func AddRecord(record model.Record) error {
	_, err := DB.Exec("INSERT INTO record(name ,money, trade_time, txt) value(?, ?, ?, ?) ", record.Name, record.Money, record.RecordTime, record.Txt)
	return err
}
