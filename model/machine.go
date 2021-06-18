package model

import (
	"time"
)

type Machine struct {
	Id        uint      `gorm:"primary_key" json:"id" form:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `gorm:"type:varchar(50);unique_index" json:"name"`
	Host      string    `gorm:"type:varchar(50)"`
	Ip        string    `gorm:"type:varchar(80)"`
	Port      uint      `json:"port" gorm:"type:int(6)"`
	User      string    `json:"user" gorm:"type:varchar(20)"`
	Password  string    `json:"password,omitempty"`
	Key       string    `json:"key,omitempty"`
	Type      string    `json:"type" gorm:"type:varchar(20)"`
}

//MachineAdd
func MachineAdd(name, addr, ip, user, password, key, auth string, port uint) error {
	ins := &Machine{Name: name, Ip: ip, Host: addr, User: user, Password: password, Key: key, Type: auth, Port: port}
	return db.Create(ins).Error
}

// MachineAll 查询所有的数据
func MachineAll(search string) ([]Machine, error) {
	//db.Order("")
	var resp []Machine
	query := db
	if search != "" {
		query = db.Where("name like ? or host like ?", "%"+search+"%", "%"+search+"%")
	}

	err := query.Find(&resp).Error
	return resp, err
}

// GetMachineByID 查询所有的数据
func GetMachineByID(id int) (*Machine, error) {
	var resp Machine
	err := db.Where("id = ?", id).First(&resp).Error

	return &resp, err
}

// DelMachineByID 通过Id删除信息
func DelMachineByID(id int) error {

	return db.Where("id = ?", id).Delete(&Machine{}).Error
}
