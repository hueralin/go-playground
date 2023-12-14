package data

import (
	"errors"
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

func GetUserByEmail(email string) (User, error) {
	var user User
	row := Db.QueryRow("SELECT * FROM users WHERE email = $1;", email)
	err := row.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return user, err
}

// CreateSession create a new session for an existing user
func (user *User) CreateSession() (Session, error) {
	var session Session
	statement := "INSERT INTO sessions (uuid, email, user_id, created_at) VALUES ($1, $2, $3, $4);"
	// Prepare creates a prepared statement for later queries or executions.
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return session, err
	}
	defer stmt.Close()
	// use QueryRow to return a row and scan the returned id into the Session struct
	row := stmt.QueryRow(createUUID(), user.Email, user.Id, time.Now())
	err = row.Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return session, nil
}

func (sess *Session) Check() (bool, error) {
	row := Db.QueryRow("SELECT * FROM session WHERE uuid = $1;", sess.Uuid)
	err := row.Scan(&sess.Id, &sess.Uuid, &sess.UserId, &sess.Email, &sess.CreatedAt)
	if err != nil {
		return false, err
	}
	if sess.Id == 0 {
		return false, errors.New("session not found")
	}
	return true, nil
}
