package main

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Menu struct {
	MenuID    uint    `gorm:"primaryKey"`
	Name      string  `gorm:"unique;not null"`
	Price     float64 `gorm:"not null"`
	Stock     int     `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Order struct {
	OrderID      uint `gorm:"primaryKey"`
	CustomerName string
	OrderDetails []OrderDetail
	TotalPrice   float64 `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type OrderDetail struct {
	OrderDetailID uint `gorm:"primaryKey"`
	OrderID       uint `gorm:"not null"`
	MenuID        uint `gorm:"not null"`
	Quantity      int  `gorm:"not null"`
	Menu          Menu
	Order         Order
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type MonthlyReport struct {
	gorm.Model
	Month            int    `gorm:"not null"`
	TotalSales       int    `gorm:"not null"`
	MostPopularDish  string `gorm:"not null"`
	MostFrequentCust string `gorm:"not null"`
}

type WeeklyReport struct {
	gorm.Model
	Start            time.Time `gorm:"not null"`
	End              time.Time `gorm:"not null"`
	TotalSales       int       `gorm:"not null"`
	MostPopularDish  string    `gorm:"not null"`
	MostFrequentCust string    `gorm:"not null"`
}

func main() {
	dsn := "host=localhost user=postgres password=changeme dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(
		&OrderDetail{},
		&Order{},
		&Menu{},
		&WeeklyReport{},
		&MonthlyReport{},
	)

}
