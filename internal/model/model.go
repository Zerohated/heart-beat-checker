package model

import (
	"time"

	conf "github.com/Zerohated/heart-beat-checker/configs"
	"github.com/Zerohated/heart-beat-checker/pkg/dao"

	"github.com/Zerohated/tools/pkg/logger"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	pgConn *gorm.DB
	log    = logger.Logger
)

// Init connect to DB
func Init(dbConf *conf.DatabaseConfig) {
	err := dao.ConnectPG(dbConf.Host, dbConf.Port, dbConf.User, dbConf.DBName, dbConf.Password)
	if err != nil {
		log.Warnln(err.Error())
	}
	dao.PgConn.AutoMigrate(&User{})
	pgConn = dao.PgConn
}

type User struct {
	gorm.Model `json:"-"`
	UID        int        `json:"uid" gorm:"uniqueIndex"`
	Username   string     `json:"username"`
	Message    string     `json:"message"`
	Path       string     `json:"path"`
	LastSeen   *time.Time `json:"lastSeen"`
}

func CreateOrUpdateUser(user *User) error {
	err := pgConn.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "uid"}},
		DoUpdates: clause.AssignmentColumns([]string{"username", "message", "path", "last_seen"}),
	}).Create(user).Error
	return err
}

func GetUserList() (users []*User, err error) {
	users = []*User{}
	err = pgConn.Find(&users).Error
	return
}

func GetUser(uid int) (user *User) {
	user = &User{UID: uid}
	pgConn.Model(user).First(user)
	return
}
