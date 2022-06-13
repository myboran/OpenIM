package im_mysql_model

import "open-im/pkg/common/db"

func DeleteSingleFriendInfo(OwnerUserID, FriendUserID string) error {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	if err != nil {
		return err
	}
	err = dbConn.Table("friends").Where("owner_user_id=? and friend_user_id=?", OwnerUserID, FriendUserID).Delete(db.Friend{}).Error
	return err
}

func GetFriendRelationshipFromFriend(OwnerUserID, FriendUserID string) (*db.Friend, error) {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	if err != nil {
		return nil, err
	}
	var friend db.Friend
	err = dbConn.Table("friends").Where("owner_user_id=? and friend_user_id=?", OwnerUserID, FriendUserID).Take(&friend).Error
	if err != nil {
		return nil, err
	}
	return &friend, err
}
