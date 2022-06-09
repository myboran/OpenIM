package utils

import (
	"open-im/pkg/common/db"
	imdb "open-im/pkg/common/db/mysql_model/im_mysql_model"
	sdkws "open-im/pkg/proto/sdk_ws"
	"open-im/pkg/utils"
)

func FriendDBCopyOpenIM(dst *sdkws.FriendInfo, src *db.Friend) error {
	utils.CopyStructFields(dst, src)
	user, err := imdb.GetUserByUserID(src.FriendUserID)
	if err != nil {
		return utils.Wrap(err, "")
	}

	utils.CopyStructFields(dst.FriendUser, user)
	dst.CreateTime = uint32(src.CreateTime.Unix())

	if dst.FriendUser == nil {
		dst.FriendUser = &sdkws.UserInfo{}
	}
	dst.FriendUser.CreateTime = uint32(user.CreateTime.Unix())
	return nil
}
