package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/mattn/go-sqlite3"
)

var (
	o                orm.Ormer
	TabPrefix        = "siot_"
	secondsEastOfUTC = int((8 * time.Hour).Seconds())
	beijing          = time.FixedZone("Beijing Time", secondsEastOfUTC)
)

const (
	FormatTime     = "15:04:05"
	FormatDate     = "2006-01-02"
	FormatDateTime = "2006-01-02 15:04:05"
)

// "data.db?_loc=Asia%2FShanghai"

type DBConfig struct {
	Name   string
	Driver string
	Url    string
}

func InitDB(db *DBConfig) (err error) {
	orm.DefaultTimeLoc = beijing
	err = orm.RegisterDriver("sqlite3", orm.DRSqlite)
	if err != nil {
		return
	}
	err = orm.RegisterDataBase(db.Name, db.Driver, db.Url)
	if err != nil {
		return
	}
	o = orm.NewOrmUsingDB(db.Name)
	return
}

func InitData() (err error) {
	// default user
	u := &User{
		LoginName: "llz",
		LoginPwd:  "1",
		UserName:  "林雷洲",
	}
	err, _ = u.ReadOrAdd()
	if err != nil {
		return
	}

	// default appuser
	appuser := new(AppUser)
	err = appuser.Add(u)
	if err != nil {
		return
	}

	return
}
