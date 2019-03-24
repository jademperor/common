package gormic

import (
	"github.com/jinzhu/gorm"

	"log"
)

// CreateTable ...
func CreateTable(scope *gorm.Scope) {
	tblName := scope.TableName()
	if !scope.DB().HasTable(tblName) {
		if err := scope.DB().CreateTable(scope.GetModelStruct()).Error; err != nil {
			log.Printf("error: could not create table for %s\n", tblName)
		}
	}
}

// RegisterPlugin ...
func RegisterPlugin(db *gorm.DB) {
	db.Callback().Create().Before("gorm:create").Register("plugin:create:create_table", CreateTable)
}
