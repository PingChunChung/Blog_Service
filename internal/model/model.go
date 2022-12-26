package model

import (
	"blog-service/pkg/setting"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Register struct{}

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      bool   `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		databaseSetting.Host,
		databaseSetting.Username,
		databaseSetting.Password,
		databaseSetting.DBName,
		databaseSetting.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Tag{})
	db.AutoMigrate(&Article{})
	return db, nil
	// return nil, nil

}

func (r *Register) Name() string {
	return "Register"
}

func (r *Register) Initialize(db *gorm.DB) error {
	db.Callback().Create().Before("gorm:create").Register("updateTimeStampForCreateCallback", updateTimeStampForCreateCallback)

	return nil
}

func updateTimeStampForCreateCallback(tx *gorm.DB) {
	ctx := tx.Statement.Context
	timeFieldsToInit := []string{"CreatedOn", "ModifiedOn"}
	for _, field := range timeFieldsToInit {
		if timeField := tx.Statement.Schema.LookUpField(field); timeField != nil {
			fmt.Println(field, timeField)
			timeField.Set(ctx, tx.Statement.ReflectValue, time.Now())
		}

	}
}
