package im_mysql_model

import "open-im/pkg/common/db"

func CheckBlack(ownerUserID, blockUserID string) error {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	if err != nil {
		return err
	}
	var black db.Black
	err = dbConn.Table("blacks").Where("owner_user_id=? and block_user_id=?", ownerUserID, blockUserID).Find(&black).Error
	return err
}
