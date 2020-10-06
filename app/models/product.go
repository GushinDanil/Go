package models

import (
	"database/sql"
	"errors"
	"html"
	"strings"
	"time"
)

type Product struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Price     int64     `json:"price"`
	User      User      `json:"user"`
	UserID    int       ` json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Product) Prepare() {
	p.ID = 0
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.User = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Product) Validate() error {

	if p.Name == "" {
		return errors.New("Required Name")
	}
	if p.Price <= 0 {
		return errors.New("Required Price")
	}
	if p.UserID < 1 {
		return errors.New("Required User")
	}
	return nil
}

func (p *Product) CreateProduct(db *sql.DB)(*Product, error) {
	err := db.QueryRow("insert into products (name,price,user_id,createdat,updatedat)values ($1,$2,$3,$4,$5) Returning id", p.Name, p.Price, p.UserID, p.CreatedAt, p.UpdatedAt).Scan(&p.ID)
	if err != nil {
		return nil,err
	}
	if p.ID != 0 {
		err := db.QueryRow("select * from users where id=$1", p.UserID).Scan(&p.User.ID, &p.User.Nickname, &p.User.Email, &p.User.Password)
		if err != nil {
			return nil,err
		}
	}
	return p,nil
}
func (p *Product) UpdateProduct(db *sql.DB) (*Product,error) {
	p.UpdatedAt = time.Now()
	_, err := db.Exec("update products set name=$1, price=$2,user_id=$3,updatedat=$4  where id=$5", p.Name, p.Price, p.UserID, p.UpdatedAt, p.ID)
	if err != nil {
		return nil,err
	}
	_,err=p.GetProduct(db)
	if err!=nil {
		return nil, err
	}
	return p,nil

}
func (p *Product) DeleteProduct(db *sql.DB) (int64,error) {
	res, err := db.Exec("delete from products where id=$1", p.ID)
	if err != nil {
		return 0,err
	}
	num,err:=res.RowsAffected()
	if err!=nil {
		return 0,err
	}
	return num,nil
}
func (p *Product) GetProduct(db *sql.DB) (*Product,error) {

	err := db.QueryRow("SELECT name, price,user_id,createdAt,updatedAt FROM products WHERE id=$1 ", p.ID).Scan(&p.Name, &p.Price, &p.UserID, &p.CreatedAt, &p.UpdatedAt)

	if err != nil {
		return nil,err
	}

	err = db.QueryRow("select * from users where id=$1", p.UserID).Scan(&p.User.ID, &p.User.Nickname, &p.User.Email, &p.User.Password)
	if err != nil {
		return nil,err
	}

	return p,nil
}
func (p *Product) GetProducts(db *sql.DB) (*[]Product, error) {
	var arr []Product
	rows, err := db.Query("SELECT * from products order by id")

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&p.ID, &p.Name, &p.Price, &p.UserID, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, err
		}
		err = db.QueryRow("select * from users where id=$1", p.UserID).Scan(&p.User.ID, &p.User.Nickname, &p.User.Email, &p.User.Password)

		if err != nil {
			return nil, err
		}
		arr = append(arr, *p)

	}

	return &arr, nil

}
