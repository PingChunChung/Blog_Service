package model

import (
	"blog-service/pkg/setting"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
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
		// DisableForeignKeyConstraintWhenMigrating: false,
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Tag{})
	db.AutoMigrate(&Article{}, &Auth{})
	Initialize(db)
	return db, nil
	// return nil, nil

}

func Initialize(db *gorm.DB) error {
	db.Callback().Create().Replace("gorm:before_create", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:before_update", updateTimeStampForBeforeUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	return nil
}

func updateTimeStampForCreateCallback(db *gorm.DB) {
	db.Statement.SetColumn("CreatedOn", time.Now())
}

func updateTimeStampForBeforeUpdateCallback(db *gorm.DB) {
	db.Statement.SetColumn("ModifiedOn", time.Now().Unix())
}

func deleteCallback(db *gorm.DB) {
	if db.Error == nil {
		if db.Statement.Schema != nil {
			db.Statement.SQL.Grow(100)

			deleteField := db.Statement.Schema.LookUpField("DeletedOn")
			isDelField := db.Statement.Schema.LookUpField("IsDel")
			if !db.Statement.Unscoped && deleteField != nil && isDelField != nil {
				if db.Statement.SQL.String() == "" {
					now := time.Now().Unix()
					db.Statement.AddClause(
						clause.Set{{
							Column: clause.Column{Name: deleteField.DBName},
							Value:  now,
						},
							{Column: clause.Column{Name: isDelField.DBName},
								Value: 1}},
					)
					db.Statement.AddClauseIfNotExists(clause.Update{})
					db.Statement.Build("UPDATE", "SET", "WHERE")
				} else {
					if db.Statement.SQL.String() == "" {
						db.Statement.AddClauseIfNotExists(clause.Delete{})
						db.Statement.AddClauseIfNotExists(clause.From{})
						db.Statement.Build("DELETE", "FROM", "WHERE")
					}
				}
				fmt.Println(db.Statement.SQL.String())
				fmt.Println(db.Statement.Vars)

				if _, ok := db.Statement.Clauses["WHERE"]; !db.AllowGlobalUpdate && !ok {
					db.AddError(gorm.ErrMissingWhereClause)
					return
				}

				if !db.DryRun && db.Error == nil {
					result, err := db.Statement.ConnPool.ExecContext(db.Statement.Context, db.Statement.SQL.String(), db.Statement.Vars...)
					if err == nil {
						db.RowsAffected, _ = result.RowsAffected()
					} else {
						db.AddError(err)
					}
				}
			}
		}
	}
}
