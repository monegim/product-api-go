package models

import (
	"log"
	"time"

	"github.com/monegim/product-api-go/pkg/setting"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}

func Setup() {
	var err error

	_, err = gorm.Open(sqlite.Open(setting.DatabaseSetting.Path), &gorm.Config{})

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	
	// db.Config.DryRun()
	// db.Table().Si
	// db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	// db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	// db.Callback().Delete().Replace("gorm:delete", deleteCallback)
}

// func CloseDB()  {
// 	defer db.
// }

func updateTimeStampForCreateCallback(scope *gorm.DB) {
	// if  scope.Error  != nil {
	// 	nowTime := time.Now().Unix()
	// 	if createTimeField , ok := scope.Statement.Schema.FieldsByName("CreatedOn"); ok {

	// 	}
	// if createTimeField, ok := scope.Fiel
	// }
}
func updateTimeStampForUpdateCallback(scope *gorm.DB) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.Statement.SetColumn("ModifiedOn", time.Now().Unix())
	}
}
func deleteCallback(scope *gorm.DB) {}
