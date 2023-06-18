package testhelpers

import (
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var TestDB *gorm.DB

func Get() *gorm.DB {
	return TestDB
}

// ConnectToTestDB connects to the test database
func ConnectToTestDB(t *testing.T) error {
	var err error
	TestDB, err = gorm.Open("postgres", "postgres://postgres:@localhost:5432/shopping_cart_go_test?sslmode=disable")
	if err != nil {
		return err
	}
	TestDB.DB()
	err = TestDB.DB().Ping()
	if err != nil {
		return err
	}

	return nil
}

func Close() {
	var tables []string
	if err := TestDB.Table("information_schema.tables").Where("table_schema = ?", "public").Pluck("table_name", &tables).Error; err != nil {
		panic(err)
	}

	for _, t := range tables {
		_ = TestDB.Exec(fmt.Sprintf("DELETE FROM %s;", t)).Error
	}

	TestDB.Close()
}
