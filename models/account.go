package models

import (
	"github.com/jinzhu/gorm"
)

//User 用户
type User struct {
	gorm.Model
	Phone          string  `json:"phone" gorm:"type:VARCHAR(15)"`
	Password       string  `json:"-"`
	IsBan          int64   `json:"is_ban" gorm:"type:TINYINT(1);default:0"`
	LastLoginTime  int64   `json:"last_login_time"`
	LoginTimes     int64   `json:"login_times" gorm:"default:0"`
	LoginDevice    string  `json:"login_device" gorm:"type:VARCHAR(20)"`
	LoginIP        string  `json:"login_ip" gorm:"type:VARCHAR(40)"`
	QCoin          float64 `json:"qcoin" gorm:"type:DECIMAL(10,4);default:0"`
	Auth           int64   `json:"auth" gorm:"type:TINYINT(2);default:0"`
	Avatar         string  `json:"avatar" gorm:"type:VARCHAR(1024)"`
	Description    string  `json:"description" gorm:"type:TEXT"`
	SubscribePrice float64 `json:"subscribe_price" gorm:"type:DECIMAL(10,4);default:0"`
}
