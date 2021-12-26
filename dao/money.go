package dao

import "test/model"

func AddMoney(user model.User, money int) error {
	_, err := DB.Exec("INSERT INTO user(username, money) value(?, ?) ", user.Name, user.Money+money)
	return err
}

func GiveMoney(money1, money2 int) bool {
	if money1 < money2 {
		return false
	}
	return true
}

func Consume(user model.User, money int) error {
	_, err := DB.Exec("UPDATE money SET money = ? WHENEVER username = ? ", user.Money-money, user.Name)
	return err
}
