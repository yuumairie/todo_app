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
	Todos     []Todo
}

// Session セッション定義
type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    int
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

// GetUserByEmail 指定したEmailのユーザ情報取得関数
func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at
		from users where email = ?`

	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt)

	return user, err
}

// CreateSession セッション作成関数
func (u *User) CreateSession() (session Session, err error) {
	session = Session{}
	cmd1 := `insert into sessions (
		uuid,
		email,
		user_id,
		created_at) values(?,?,?,?)`
	_, err = Db.Exec(cmd1, createUUID(), u.Email, u.ID, time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	cmd2 := `select id, uuid, email, user_id, created_at
	from sessions where user_id = ? and email = ?`

	err = Db.QueryRow(cmd2, u.ID, u.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt)

	return session, err
}

// CheckSession セッション確認関数
func (sess *Session) CheckSession() (valid bool, err error) {
	cmd := `select id, uuid, email, user_id, created_at
	from sessions where uuid = ?`

	err = Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.ID,
		&sess.UUID,
		&sess.Email,
		&sess.UserID,
		&sess.CreatedAt)

	if err != nil {
		valid = false
		return
	}
	if sess.ID != 0 {
		valid = true
	}
	return valid, err
}

// DeleteSessionByUUID セッション削除関数
func (sess *Session) DeleteSessionByUUID() (err error) {
	cmd := `delete from sessions where uuid = ?`
	_, err = Db.Exec(cmd, sess.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// GetUserBySession セッションからユーザ取得
func (sess *Session) GetUserBySession() (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, created_at FROM users
	where id = ?`
	err = Db.QueryRow(cmd, sess.UserID).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.CreatedAt)
	return user, err
}
