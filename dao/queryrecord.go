package dao

import "test/model"

func GetRecordByTxt(txt string) (model.Record, error) {
	var record model.Record
	row := DB.QueryRow("SELECT  username, money, trade_time, txt  FROM post WHERE txt = ? ", txt)
	if row.Err() != nil {
		return record, row.Err()
	}
	err := row.Scan(&record.Name, &record.Money, &record.RecordTime, &record.Txt)
	if err != nil {
		return record, err
	}
	return record, nil
}
