package model

var Feed []Opinion

type Opinion struct {
	ID      int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Rating  int    `json:"rating" gorm:"column:rating"`
	Comment string `json:"comment"`
}

var Feedback Opinion
