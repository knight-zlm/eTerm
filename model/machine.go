package model

import "time"

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

// 
func MachineAll(search string) ([]Machine, error) {
	//db.Order("")
	var resp []Machine
	if search != "" {
		db.Where("name like ?", "%"+search+"%")
	}

	err := db.Find(&resp).Error
	return resp, err
}
