package model



type Task struct {
	
	Id      uint   `gorm:"primaryKey"`
	Title   string `json:"title" gorm:"size:20"`
	Content string `json:"content" gorm:"size:100"`
}
