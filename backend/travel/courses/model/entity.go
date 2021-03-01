package model

type Journey struct{
	ID			int	`json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Source		string	`json:"source" gorm:"type:varchar(255);not null"`
	Distance		string	`json:"distance"`
	Destination	string	`json:"destination"`
	Price		int		`json:"price" `
}


type Book struct {
	ID uint
	UserID string
	InfoID string
	Image string `gorm:"type:varchar(255)"`
}

type Info struct {

}

type Role struct {
	ID uint
	Name string `gorm:"type:varchar(255)"`
}

type User struct {
	ID       uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	FullName string `json:"name" gorm:"type:varchar(255);not null"`
	Email    string `json:"email" gorm:"type:varchar(255);not null;unique"`
	Password string `json:"pass" gorm:"type:varchar(255)"`
	RoleID   uint
	Amount   uint `json:"amount" gorm:"DEFAULT:300"`
}

type Schedule struct {
	ID           uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	TravelID      uint   `json:"travelid" gorm:"not null"`
	StartingTime string `json:"startingtime" gorm:"type:varchar(255);not null"`
	Class    string `json:"class" gorm:"type:varchar(255);not null"`
	TrainID     uint   `json:"trainid" gorm:"not null;"`
	Day          string `json:"day" gorm:"type:varchar(255);not null"`
	Booked       uint   `json:"booked" gorm:"DEFAULT:0"`
}

type Travel struct{
	ID	uint  `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name string `json:"name" gorm:"type:varchar(255);not null"`
	Price uint `json:"price"`
	Schedules []Schedule `json:"schedules"`
	VIPPrice uint `json:"vipprice"`
	Capacity uint `json:"capacity"`
	VIPCapacity uint `json:"vipcapacity"`
}