package im_mysql_model

import "open-im/pkg/common/db"

func GetUserByUserID(userID string) (*db.User, error) {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	if err != nil {
		return nil, err
	}
	var user db.User
	err = dbConn.Table("users").Where("user_id=?", userID).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
