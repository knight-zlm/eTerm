package model

import (
	"log"
	"os"
	"path"

	"github.com/jinzhu/gorm"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
)

var db *gorm.DB
var dbPath string

func init() {
	dir, err := homedir.Dir()
	if err != nil {
		log.Printf("get home dir err:%v", err)
	}

	dbPath = path.Join(dir, ".eTerm/sqlite.db")
	// 如果不存在是不是要新建文件
	if !IsExistByPath(dbPath) {
		file, err := os.Create(dbPath)
		if err != nil {
			log.Printf("create sqlite db path err:%v", err)
		}
		file.Close()
	}
}

func CreateSQLiteDb() {
	log.Printf("start create SQLite path:%v", dbPath)
	sqlite, err := gorm.Open("sqlite", dbPath)
	if err != nil {
		logrus.WithError(err).Fatalf("master fail to open its sqlite db in %s. please install master first.", dbPath)
		return
	}

	db = sqlite
	// 保持数据库字段和mode的字段一致
	db.AutoMigrate(Machine{})
	// 查看详细的sql日志
	//db.LogMode(true)
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
