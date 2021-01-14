package models

import (
	"log"
	"time"
)

// User ユーザ定義
type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

// CreateUser ユーザ作成関数
func (u *User) CreateUser() (err error) {
	cmd := `insert into users(
		uuid,
		name,
		email,
		password,
		created_at) values (?,?,?,?,?)`

	_, err = Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.Password),
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// GetUser 指定したIDのユーザ取得関数
func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at
	from users where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	return user, err
}

// UpdateUser ユーザ情報更新関数
func (u *User) UpdateUser() (err error) {
	cmd := `update users set name = ?, email = ? where id = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// DeleteUser ユーザ情報削除関数
func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id = ?`
	_, err = Db.Exec(cmd, u.ID)

	if err != nil {
		log.Fatalln(err)
	}
	return err
}
