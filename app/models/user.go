package models

import (
	"database/sql"
	"errors"
	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
	"html"
	"strings"
)

type User struct {
	ID       uint64 `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
//хеширует пароль
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
// сравниевает пароли в хешированном виде
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.Nickname = html.EscapeString(strings.TrimSpace(u.Nickname))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Nickname == "" {
			return errors.New("Required Nickname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if u.Nickname == "" {
			return errors.New("Required Nickname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

func (u *User) CreateUser(db *sql.DB) (*User, error) {


	err := db.QueryRow("insert into  users (nickname,email,password)values ($1,$2,$3) Returning id", u.Nickname, u.Email, u.Password).Scan(&u.ID)

	if err != nil {

		return nil, err
	}
	return u, nil
}

func (u *User) FindAllUsers(db *sql.DB) (*[]User, error) {

	users := []User{}
	rows, err := db.Query("Select * from users order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&u.ID, &u.Nickname, &u.Email, &u.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, *u)
	}

	return &users, nil
}

func (u *User) FindUserByID(db *sql.DB) (*User ,error) {


	err:=db.QueryRow("select nickname,email,password from users where id=$1", u.ID).Scan(&u.Nickname, &u.Email, &u.Password)
	if err!=nil {
		return nil,err
	}
	return u,nil
}

func (u *User) UpdateAUser(db *sql.DB) (*User,error) {

	_, err := db.Exec("update users set nickname=$1, email=$2,password=$3 where id=$4", u.Nickname, u.Email, u.Password, u.ID)
	if err != nil {
		return nil,err
	}
	user,err:=u.FindUserByID(db)
	if err !=nil{
		return nil,err
	}
	return user,nil
}

func (u *User) DeleteAUser(db *sql.DB) (int64,error) {

	res, err := db.Exec("delete from users where id=$1", u.ID)

	if err != nil {
	return 0, err
	}
	r,err:=res.RowsAffected()
	if err!=nil {
		return 0,err
	}
	return r ,nil
}
