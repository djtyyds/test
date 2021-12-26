package dao

import "test/model"

func InsertUser(user model.User) error {
	_, err := DB.Exec("INSERT INTO user(name,password) value (?,?)", user.Name, user.Password)
	return err
}

func SelectUserByUsername(username string) (model.User, error) {
	user := model.User{}
	row := DB.QueryRow("SELECT id, password FROM user WHERE username = ? ", username)
	if row.Err() != nil {
		return user, row.Err()
	}
	err := row.Scan(&user.Id, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}
