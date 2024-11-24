package models

type Book struct {
	Id        int64   `gorm:"primaryKey" db:"id" json:"id"`
	Title     string  `db:"title" json:"title"`
	Author    string  `db:"author" json:"author"`
	Price     float64 `db:"price" json:"price"`
	Purchased int     `db:"purchased" json:"purchased"`
}
