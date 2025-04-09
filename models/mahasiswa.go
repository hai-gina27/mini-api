package models

type Mahasiswa struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	Nama      string  `json:"nama"`
	Nim       string  `json:"nim"`
	Alamat    string  `json:"alamat"`
	Createdat string  `json:"created_at"`
	IDJurusan uint    `json:"id_jurusan"`
	Jurusan   Jurusan `json:"jurusan" gorm:"foreignKey:IDJurusan;references:ID"`
}
