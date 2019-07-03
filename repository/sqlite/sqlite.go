package sqlite

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"

	"github.com/charlesfan/go-api/repository/user"
	"github.com/charlesfan/go-api/utils/log"
	"github.com/charlesfan/go-api/utils/tmpfile"
)

func NewSqlite(path string, migrate bool) (db *gorm.DB, err error) {
	if path == "" {
		f, err := tempfile.TempFileWithSuffix(os.TempDir(), "gorm", ".db")
		if f == nil || err != nil {
			return nil, err
		}
		path = f.Name()
		migrate = true
	}

	log.Info("DB host: ", path)
	db, err = gorm.Open("sqlite3", path)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if migrate {
		runMigration(db)
	}

	db.DB().SetMaxIdleConns(10)

	return
}

func runMigration(db *gorm.DB) {
	values := []interface{}{
		&user.User{},
	}

	for _, value := range values {
		db.DropTableIfExists(value)
	}

	if err := db.AutoMigrate(values...).Error; err != nil {
		log.Error(err)
	}
}
