package model

import (
	"log"
	"os"
	"path"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
)

var db *gorm.DB
var dbPath string

func init() {
	dir, err := homedir.Dir()
	if err != nil {
		log.Fatalf("get home dir err:%v", err)
	}

	dbPath = path.Join(dir, ".eTerm/sqlite.db")
	// 如果不存在是不是要新建文件
	if !IsExistByPath(dbPath) {
		// 判断目录是否存在
		fPath := path.Dir(dbPath)
		if !IsExistByPath(fPath) {
			if err := os.MkdirAll(fPath, os.ModePerm); err != nil {
				log.Fatalf("create sqlite db path err:%v", err)
			}
		}

		// 新建文件
		file, err := os.Create(dbPath)
		if err != nil {
			log.Fatalf("create sqlite db path err:%v", err)
		}
		file.Close()
	}
}

func CreateSQLiteDb(isDebug bool) {
	log.Printf("SQLite path:%v", dbPath)
	sqlite, err := gorm.Open("sqlite3", dbPath)
	if err != nil {
		logrus.WithError(err).Fatalf("master fail to open its sqlite db in %s. please install master first.", dbPath)
		return
	}

	db = sqlite
	// 保持数据库字段和mode的字段一致
	db.AutoMigrate(Machine{})
	// 查看详细的sql日志
	db.LogMode(isDebug)
}

// 删除数据库
func FlushSqliteDb() error {
	db.Close()
	return os.RemoveAll(dbPath)
}

func IsExistByPath(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}

	return true
}
