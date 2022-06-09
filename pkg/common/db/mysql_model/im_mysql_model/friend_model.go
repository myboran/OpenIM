package im_mysql_model

import "open-im/pkg/common/db"

func GetFriendListByUserID(OwnerUserID string) ([]db.Friend, error) {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	if err != nil {
		return nil, err
	}

	var friends []db.Friend
	var x db.Friend
	x.OwnerUserID = OwnerUserID
	err = dbConn.Table("friends").Where("owner_user_id=?", OwnerUserID).Find(&friends).Error
	if err != nil {
		return nil, err
	}
	return friends, nil
}
