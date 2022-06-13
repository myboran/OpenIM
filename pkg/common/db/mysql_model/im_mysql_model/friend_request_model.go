package im_mysql_model

import (
	"open-im/pkg/common/db"
	"open-im/pkg/utils"
	"time"
)

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

func InsertFriendApplication(friendRequest *db.FriendRequest, args map[string]interface{}) error {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	if err != nil {
		return err
	}

	if err = dbConn.Table("friend_requests").Create(friendRequest).Error; err == nil {
		return nil
	}

	friendRequest.CreateTime = time.Now()
	args["create_time"] = friendRequest.CreateTime
	u := dbConn.Model(friendRequest).Updates(args)

	if u.RowsAffected != 0 {
		return nil
	}

	if friendRequest.CreateTime.Unix() < 0 {
		friendRequest.CreateTime = time.Now()
	}
	if friendRequest.HandleTime.Unix() < 0 {
		friendRequest.HandleTime = utils.UnixSecondToTime(0)
	}
	err = dbConn.Table("friend_requests").Create(friendRequest).Error
	if err != nil {
		return err
	}

	return nil
}
