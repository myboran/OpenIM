package msg

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"open-im/pkg/common/config"
	"open-im/pkg/common/constant"
	imdb "open-im/pkg/common/db/mysql_model/im_mysql_model"
	"open-im/pkg/common/log"
	pbFriend "open-im/pkg/proto/friend"
	sdkws "open-im/pkg/proto/sdk_ws"
	"open-im/pkg/utils"
)

func getFromToUserNickname(fromUserID, toUserID string) (string, string, error) {
	from, err := imdb.GetUserByUserID(fromUserID)
	if err != nil {
		return "", "", utils.Wrap(err, "")
	}
	to, err := imdb.GetUserByUserID(toUserID)
	if err != nil {
		return "", "", utils.Wrap(err, "")
	}
	return from.Nickname, to.Nickname, nil
}

func friendNotification(commID *pbFriend.CommID, contentType int32, m proto.Message) {
	log.Info(commID.OperationID, utils.GetSelfFuncName(), "args: ", commID, contentType)
	var err error
	var tips sdkws.TipsComm
	tips.Detail, err = proto.Marshal(m)
	if err != nil {
		log.Error(commID.OperationID, "Marshal failed ", err.Error(), m.String())
		return
	}

	marshaler := jsonpb.Marshaler{
		OrigName:     true,
		EmitDefaults: false,
		EnumsAsInts:  false,
	}

	tips.JsonDetail, _ = marshaler.MarshalToString(m)

	fromUserNickname, toUserNickname, err := getFromToUserNickname(commID.FromUserID, commID.ToUserID)
	if err != nil {
		log.Error(commID.OperationID, "getFromToUserNickname failed ", err.Error(), commID.FromUserID, commID.ToUserID)
		return
	}

	cn := config.Config.Notification

	switch contentType {
	case constant.FriendApplicationNotification:
		tips.DefaultTips = fromUserNickname + cn.FriendApplication.DefaultTips.Tips
	case constant.FriendApplicationApprovedNotification:
		tips.DefaultTips = fromUserNickname + cn.FriendApplicationApproved.DefaultTips.Tips
	case constant.FriendApplicationRejectedNotification:
		tips.DefaultTips = fromUserNickname + cn.FriendApplicationRejected.DefaultTips.Tips
	case constant.FriendAddedNotification:
		tips.DefaultTips = cn.FriendAdded.DefaultTips.Tips
	case constant.FriendDeletedNotification:
		tips.DefaultTips = cn.FriendDeleted.DefaultTips.Tips + toUserNickname
	case constant.FriendRemarkSetNotification:
		tips.DefaultTips = fromUserNickname + cn.FriendRemarkSet.DefaultTips.Tips
	case constant.BlackAddedNotification:
		tips.DefaultTips = cn.BlackAdded.DefaultTips.Tips
	case constant.BlackDeletedNotification:
		tips.DefaultTips = cn.BlackDeleted.DefaultTips.Tips + toUserNickname
	case constant.UserInfoUpdatedNotification:
		tips.DefaultTips = cn.UserInfoUpdated.DefaultTips.Tips
	default:
		log.Error(commID.OperationID, "contentType failed ", contentType)
		return
	}

	var n NotificationMsg
	n.SendID = commID.FromUserID
	n.RecvID = commID.ToUserID
	n.ContentType = contentType
	n.SessionType = constant.SingleChatType
	n.MsgFrom = constant.SysMsgType
	n.OperationID = commID.OperationID
	n.Content, err = proto.Marshal(&tips)
	if err != nil {
		log.Error(commID.OperationID, "Marshal failed ", err.Error(), tips.String())
		return
	}
	Notification(&n)
}

func UserInfoUpdatedNotification(operationID, userID string, needNotifiedUserID string) {
	selfInfoUpdatedTips := sdkws.UserInfoUpdatedTips{UserID: userID}
	commID := pbFriend.CommID{FromUserID: userID, ToUserID: needNotifiedUserID, OpUserID: userID, OperationID: operationID}
	friendNotification(&commID, constant.UserInfoUpdatedNotification, &selfInfoUpdatedTips)
}

func FriendApplicationNotification(req *pbFriend.AddFriendReq) {
	FriendApplicationTips := sdkws.FriendApplicationTips{FromToUserID: &sdkws.FromToUserID{}}
	FriendApplicationTips.FromToUserID.FromUserID = req.CommID.FromUserID
	FriendApplicationTips.FromToUserID.ToUserID = req.CommID.ToUserID
	friendNotification(req.CommID, constant.FriendApplicationNotification, &FriendApplicationTips)
}

func FriendDeletedNotification(req *pbFriend.DeleteFriendReq) {
	friendDeletedTips := sdkws.FriendDeletedTips{FromToUserID: &sdkws.FromToUserID{}}
	friendDeletedTips.FromToUserID.FromUserID = req.CommID.FromUserID
	friendDeletedTips.FromToUserID.ToUserID = req.CommID.ToUserID
	friendNotification(req.CommID, constant.FriendDeletedNotification, &friendDeletedTips)
}
