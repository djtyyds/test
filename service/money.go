package service

import (
	"test/dao"
)

func AddMoney(username string, money int) error {
	user, _ := dao.SelectUserByUsername(username)
	err := dao.AddMoney(user, money)
	return err
}

func GiveMoney(username1 string, username2 string, money int) (bool, error) {
	user1, _ := dao.SelectUserByUsername(username1)
	ok := dao.GiveMoney(user1.Money, money)
	if !ok {
		return false, nil
	}
	_ = dao.Consume(user1, money)
	user2, _ := dao.SelectUserByUsername(username2)
	err := dao.AddMoney(user2, money)
	return true, err
}

func Consume(username string, money int) (bool, error) {
	user, err := dao.SelectUserByUsername(username)
	ok := dao.GiveMoney(user.Money, money)
	if !ok {
		return false, nil
	}
	err = dao.Consume(user, money)
	return true, err
}
