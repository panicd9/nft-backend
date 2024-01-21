package database

import (
	"fmt"
	"nft-backend/model"
	"os/exec"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectToDB() *gorm.DB {
	var err error

	// Host address is changing :(
	cmdString := "ip route list default | awk '{print $3}'"
	cmd := exec.Command("bash", "-c", cmdString)
	ip, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	fmt.Println(string(ip))

	dsn := fmt.Sprintf("host=%s user=postgres password=ceres dbname=nft_db", string(ip))
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Database connected")

	AutoMigrate()
	fmt.Println("Database migrated")

	return db
}

func AutoMigrate() {
	err := db.AutoMigrate(&model.BlacklistedEvent{}, &model.MintedEvent{}, &model.RegisteredEvent{})
	if err != nil {
		fmt.Println(err)
	}
}

func AddEvent(event interface{}) {
	db.Create(event)
}
