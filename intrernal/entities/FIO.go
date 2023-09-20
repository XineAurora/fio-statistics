package entities

type FIO struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name" gorm:"not null"`
	Surname     string `json:"surname" gorm:"not null"`
	Patronymic  string `json:"patronymic,omitempty"`
	Age         int    `json:"age" gorm:"not null"`
	Gender      string `json:"gender" gorm:"not null"`
	Nationality string `json:"nationality" gorm:"not null"`
}
