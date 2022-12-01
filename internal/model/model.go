package model

import (
	"blog-service/pkg/setting"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"Modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      bool   `json:"is_del"`
}

func NewDBEngine(dataseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		dataseSetting.Host,
		dataseSetting.Username,
		dataseSetting.Password,
		dataseSetting.DBName,
		dataseSetting.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil

}