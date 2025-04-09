package models

type Jurusan struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Jurusan string `json:"jurusan"`
}
