package msg

import (
	"context"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"open-im/pkg/common/config"
	"open-im/pkg/common/constant"
	"open-im/pkg/common/db"
	"open-im/pkg/common/log"
	"open-im/pkg/grpc-etcdv3/getcdv3"
	pbChat "open-im/pkg/proto/chat"
	rpc "open-im/pkg/proto/friend"
	pbGroup "open-im/pkg/proto/group"
	sdkws "open-im/pkg/proto/sdk_ws"
	"open-im/pkg/utils"

	"github.com/garyburd/redigo/redis"
	"github.com/golang/protobuf/proto"
)

type NotificationMsg struct {
	SendID      string
	RecvID      string
	Content     []byte //  open_im_sdk.TipsComm
	MsgFrom     int32
	ContentType int32
	SessionType int32
	OperationID string
}

func Notification(n *NotificationMsg) {
	var (
		req                     pbChat.SendMsgReq
		msg                     sdkws.MsgData
		offlineInfo             sdkws.OfflinePushInfo
		title, desc, ex         string
		pushSwitch, unReadCount bool
		reliabilityLevel        int
	)
	req.OperationID = n.OperationID
	msg.SendID = n.SendID
	msg.RecvID = n.RecvID
	msg.Content = n.Content
	msg.MsgFrom = n.MsgFrom
	msg.ContentType = n.ContentType
	msg.SessionType = n.SessionType
	msg.CreateTime = utils.GetCurrentTimestampByMill()
	msg.ClientMsgID = utils.GetMsgID(n.SendID)
	msg.Options = make(map[string]bool, 7)

	switch n.SessionType {
	case constant.GroupChatType:
		msg.RecvID = ""
		msg.GroupID = n.RecvID
	}
	offlineInfo.IOSBadgeCount = config.Config.IOSPush.BadgeCount
	offlineInfo.IOSPushSound = config.Config.IOSPush.PushSound
	switch msg.ContentType {
	case constant.UserInfoUpdatedNotification:
		pushSwitch = config.Config.Notification.UserInfoUpdated.OfflinePush.PushSwitch
		title = config.Config.Notification.UserInfoUpdated.OfflinePush.Title
		desc = config.Config.Notification.UserInfoUpdated.OfflinePush.Desc
		ex = config.Config.Notification.UserInfoUpdated.OfflinePush.Ext
		reliabilityLevel = config.Config.Notification.UserInfoUpdated.Conversation.ReliabilityLevel
		unReadCount = config.Config.Notification.UserInfoUpdated.Conversation.UnreadCount
	case constant.FriendApplicationNotification:
		pushSwitch = config.Config.Notification.FriendApplication.OfflinePush.PushSwitch
		title = config.Config.Notification.FriendApplication.OfflinePush.Title
		desc = config.Config.Notification.FriendApplication.OfflinePush.Desc
		ex = config.Config.Notification.FriendApplication.OfflinePush.Ext
		reliabilityLevel = config.Config.Notification.FriendApplication.Conversation.ReliabilityLevel
		unReadCount = config.Config.Notification.FriendApplication.Conversation.UnreadCount
	case constant.FriendDeletedNotification:
		pushSwitch = config.Config.Notification.FriendDeleted.OfflinePush.PushSwitch
		title = config.Config.Notification.FriendDeleted.OfflinePush.Title
		desc = config.Config.Notification.FriendDeleted.OfflinePush.Desc
		ex = config.Config.Notification.FriendDeleted.OfflinePush.Ext
		reliabilityLevel = config.Config.Notification.FriendDeleted.Conversation.ReliabilityLevel
		unReadCount = config.Config.Notification.FriendDeleted.Conversation.UnreadCount
	}

	switch reliabilityLevel {
	case constant.UnreliableNotification:
		utils.SetSwitchFromOptions(msg.Options, constant.IsHistory, false)
		utils.SetSwitchFromOptions(msg.Options, constant.IsPersistent, false)
		utils.SetSwitchFromOptions(msg.Options, constant.IsConversationUpdate, false)
		utils.SetSwitchFromOptions(msg.Options, constant.IsSenderConversationUpdate, false)
	case constant.ReliableNotificationNoMsg:
		utils.SetSwitchFromOptions(msg.Options, constant.IsConversationUpdate, false)
		utils.SetSwitchFromOptions(msg.Options, constant.IsSenderConversationUpdate, false)
	case constant.ReliableNotificationMsg:

	}
	utils.SetSwitchFromOptions(msg.Options, constant.IsUnreadCount, unReadCount)
	utils.SetSwitchFromOptions(msg.Options, constant.IsOfflinePush, pushSwitch)

	offlineInfo.Title = title
	offlineInfo.Desc = desc
	offlineInfo.Ex = ex
	msg.OfflinePushInfo = &offlineInfo
	req.MsgData = &msg

	etcdConn := getcdv3.GetConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImOfflineMessageName)
	client := pbChat.NewChatClient(etcdConn)
	reply, err := client.SendMsg(context.Background(), &req)
	if err != nil {
		log.NewError(req.OperationID, "SendMsg rpc failed, ", req.String(), err.Error())
	} else if reply.ErrCode != 0 {
		log.NewError(req.OperationID, "SendMsg rpc failed, ", req.String())
	}
}

type MsgCallBackReq struct {
	SendID       string `json:"sendID"`
	RecvID       string `json:"recvID"`
	Content      string `json:"content"`
	SendTime     int64  `json:"sendTime"`
	MsgFrom      int32  `json:"msgFrom"`
	ContentType  int32  `json:"contentType"`
	SessionType  int32  `json:"sessionType"`
	PlatformID   int32  `json:"senderPlatformID"`
	MsgID        string `json:"msgID"`
	IsOnlineOnly bool   `json:"isOnlineOnly"`
}
type MsgCallBackResp struct {
	ErrCode         int32  `json:"errCode"`
	ErrMsg          string `json:"errMsg"`
	ResponseErrCode int32  `json:"responseErrCode"`
	ResponseResult  struct {
		ModifiedMsg string `json:"modifiedMsg"`
		Ext         string `json:"ext"`
	}
}

func GetMsgID(sendID string) string {
	t := time.Now().Format("2006-01-02 15:04:05")
	return utils.Md5(t + "-" + sendID + "-" + strconv.Itoa(rand.Int()))
}
func (rpc *rpcChat) encapsulateMsgData(msg *sdkws.MsgData) {
	msg.ServerMsgID = GetMsgID(msg.SendID)
	msg.SendTime = utils.GetCurrentTimestampByMill()
	switch msg.ContentType {
	case constant.Text:
		fallthrough
	case constant.Picture:
		fallthrough
	case constant.Voice:
		fallthrough
	case constant.Video:
		fallthrough
	case constant.File:
		fallthrough
	case constant.AtText:
		fallthrough
	case constant.Merger:
		fallthrough
	case constant.Card:
		fallthrough
	case constant.Location:
		fallthrough
	case constant.Custom:
		fallthrough
	case constant.Quote:
		utils.SetSwitchFromOptions(msg.Options, constant.IsConversationUpdate, true)
		utils.SetSwitchFromOptions(msg.Options, constant.IsUnreadCount, true)
		utils.SetSwitchFromOptions(msg.Options, constant.IsSenderSync, true)
	case constant.Revoke:
		utils.SetSwitchFromOptions(msg.Options, constant.IsUnreadCount, false)
		utils.SetSwitchFromOptions(msg.Options, constant.IsOfflinePush, false)
	case constant.HasReadReceipt:
		log.Info("", "this is a test start", msg, msg.Options)
		utils.SetSwitchFromOptions(msg.Options, constant.IsConversationUpdate, false)
		utils.SetSwitchFromOptions(msg.Options, constant.IsSenderConversationUpdate, false)
		utils.SetSwitchFromOptions(msg.Options, constant.IsUnreadCount, false)
		utils.SetSwitchFromOptions(msg.Options, constant.IsOfflinePush, false)
		log.Info("", "this is a test end", msg, msg.Options)
	case constant.Typing:
		utils.SetSwitchFromOptions(msg.Options, constant.IsHistory, false)
		utils.SetSwitchFromOptions(msg.Options, constant.IsPersistent, false)
		utils.SetSwitchFromOptions(msg.Options, constant.IsSenderSync, false)
		utils.SetSwitchFromOptions(msg.Options, constant.IsConversationUpdate, false)
		utils.SetSwitchFromOptions(msg.Options, constant.IsSenderConversationUpdate, false)
		utils.SetSwitchFromOptions(msg.Options, constant.IsUnreadCount, false)
		utils.SetSwitchFromOptions(msg.Options, constant.IsOfflinePush, false)

	}
}

func (rpc *rpcChat) sendMsgToKafka(m *pbChat.MsgDataToMQ, key string) error {
	pid, offset, err := rpc.producer.SendMessage(m, key)
	if err != nil {
		log.ErrorByKv("kafka send failed", m.OperationID, "send data", m.String(), "pid", pid, "offset", offset, "err", err.Error(), "key", key)
	}
	return err
}

func (rpc *rpcChat) SendMsg(_ context.Context, pb *pbChat.SendMsgReq) (*pbChat.SendMsgResp, error) {
	replay := pbChat.SendMsgResp{}
	log.NewDebug(pb.OperationID, "rpc sendMsg come here", pb.String())
	flag, errCode, errMsg := userRelationshipVerification(pb)
	if !flag {
		return returnMsg(&replay, pb, errCode, errMsg, "", 0)
	}
	rpc.encapsulateMsgData(pb.MsgData)
	log.Info("", "this is a test MsgData ", pb.MsgData)
	msgToMQ := pbChat.MsgDataToMQ{Token: pb.Token, OperationID: pb.OperationID, MsgData: pb.MsgData}
	//options := utils.JsonStringToMap(pbData.Options)
	isHistory := utils.GetSwitchFromOptions(pb.MsgData.Options, constant.IsHistory)
	mReq := MsgCallBackReq{
		SendID:      pb.MsgData.SendID,
		RecvID:      pb.MsgData.RecvID,
		Content:     string(pb.MsgData.Content),
		SendTime:    pb.MsgData.SendTime,
		MsgFrom:     pb.MsgData.MsgFrom,
		ContentType: pb.MsgData.ContentType,
		SessionType: pb.MsgData.SessionType,
		PlatformID:  pb.MsgData.SenderPlatformID,
		MsgID:       pb.MsgData.ClientMsgID,
	}
	if !isHistory {
		mReq.IsOnlineOnly = true
	}
	// callback
	canSend, err := callbackWordFilter(pb)
	if err != nil {
		log.NewError(pb.OperationID, utils.GetSelfFuncName(), "callbackWordFilter failed", err.Error(), pb.MsgData)
	}
	if !canSend {
		log.NewDebug(pb.OperationID, utils.GetSelfFuncName(), "callbackWordFilter result", canSend, "end rpc and return", pb.MsgData)
		return returnMsg(&replay, pb, 201, "callbackWordFilter result stop rpc and return", "", 0)
	}
	switch pb.MsgData.SessionType {
	case constant.SingleChatType:
		// callback
		canSend, err := callbackBeforeSendSingleMsg(pb)
		if err != nil {
			log.NewError(pb.OperationID, utils.GetSelfFuncName(), "callbackBeforeSendSingleMsg failed", err.Error())
		}
		if !canSend {
			log.NewDebug(pb.OperationID, utils.GetSelfFuncName(), "callbackBeforeSendSingleMsg result", canSend, "end rpc and return")
			return returnMsg(&replay, pb, 201, "callbackBeforeSendSingleMsg result stop rpc and return", "", 0)
		}
		isSend := modifyMessageByUserMessageReceiveOpt(pb.MsgData.RecvID, pb.MsgData.SendID, constant.SingleChatType, pb)
		if isSend {
			msgToMQ.MsgData = pb.MsgData
			log.NewInfo(msgToMQ.OperationID, msgToMQ)
			err1 := rpc.sendMsgToKafka(&msgToMQ, msgToMQ.MsgData.RecvID)
			if err1 != nil {
				log.NewError(msgToMQ.OperationID, "kafka send msg err:RecvID", msgToMQ.MsgData.RecvID, msgToMQ.String())
				return returnMsg(&replay, pb, 201, "kafka send msg err", "", 0)
			}
		}
		if msgToMQ.MsgData.SendID != msgToMQ.MsgData.RecvID { //Filter messages sent to yourself
			err2 := rpc.sendMsgToKafka(&msgToMQ, msgToMQ.MsgData.SendID)
			if err2 != nil {
				log.NewError(msgToMQ.OperationID, "kafka send msg err:SendID", msgToMQ.MsgData.SendID, msgToMQ.String())
				return returnMsg(&replay, pb, 201, "kafka send msg err", "", 0)
			}
		}
		// callback
		if err := callbackAfterSendSingleMsg(pb); err != nil {
			log.NewError(pb.OperationID, utils.GetSelfFuncName(), "callbackAfterSendSingleMsg failed", err.Error())
		}
		return returnMsg(&replay, pb, 0, "", msgToMQ.MsgData.ServerMsgID, msgToMQ.MsgData.SendTime)
	case constant.GroupChatType:
		// callback
		canSend, err := callbackBeforeSendGroupMsg(pb)
		if err != nil {
			log.NewError(pb.OperationID, utils.GetSelfFuncName(), "callbackBeforeSendGroupMsg failed", err.Error())
		}
		if !canSend {
			log.NewDebug(pb.OperationID, utils.GetSelfFuncName(), "callbackBeforeSendGroupMsg result", canSend, "end rpc and return")
			return returnMsg(&replay, pb, 201, "callbackBeforeSendGroupMsg result stop rpc and return", "", 0)
		}
		etcdConn := getcdv3.GetConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImGroupName)
		client := pbGroup.NewGroupClient(etcdConn)
		req := &pbGroup.GetGroupAllMemberReq{
			GroupID:     pb.MsgData.GroupID,
			OperationID: pb.OperationID,
		}
		reply, err := client.GetGroupAllMember(context.Background(), req)
		if err != nil {
			log.Error(pb.Token, pb.OperationID, "rpc send_msg getGroupInfo failed, err = %s", err.Error())
			return returnMsg(&replay, pb, 201, err.Error(), "", 0)
		}
		if reply.ErrCode != 0 {
			log.Error(pb.Token, pb.OperationID, "rpc send_msg getGroupInfo failed, err = %s", reply.ErrMsg)
			return returnMsg(&replay, pb, reply.ErrCode, reply.ErrMsg, "", 0)
		}
		var addUidList []string
		switch pb.MsgData.ContentType {
		case constant.MemberKickedNotification:
			var tips sdkws.TipsComm
			var memberKickedTips sdkws.MemberKickedTips
			err := proto.Unmarshal(pb.MsgData.Content, &tips)
			if err != nil {
				log.Error(pb.OperationID, "Unmarshal err", err.Error())
			}
			err = proto.Unmarshal(tips.Detail, &memberKickedTips)
			if err != nil {
				log.Error(pb.OperationID, "Unmarshal err", err.Error())
			}
			log.Info(pb.OperationID, "data is ", memberKickedTips.String())
			for _, v := range memberKickedTips.KickedUserList {
				addUidList = append(addUidList, v.UserID)
			}
		case constant.MemberQuitNotification:
			addUidList = append(addUidList, pb.MsgData.SendID)
		default:
		}
		groupID := pb.MsgData.GroupID
		for _, v := range reply.MemberList {
			pb.MsgData.RecvID = v.UserID
			isSend := modifyMessageByUserMessageReceiveOpt(v.UserID, groupID, constant.GroupChatType, pb)
			if isSend {
				msgToMQ.MsgData = pb.MsgData
				err := rpc.sendMsgToKafka(&msgToMQ, v.UserID)
				if err != nil {
					log.NewError(msgToMQ.OperationID, "kafka send msg err:UserId", v.UserID, msgToMQ.String())
					return returnMsg(&replay, pb, 201, "kafka send msg err", "", 0)
				}
			}

		}
		log.Info(msgToMQ.OperationID, "addUidList", addUidList)
		for _, v := range addUidList {
			pb.MsgData.RecvID = v
			isSend := modifyMessageByUserMessageReceiveOpt(v, groupID, constant.GroupChatType, pb)
			log.Info(msgToMQ.OperationID, "isSend", isSend)
			if isSend {
				msgToMQ.MsgData = pb.MsgData
				err := rpc.sendMsgToKafka(&msgToMQ, v)
				if err != nil {
					log.NewError(msgToMQ.OperationID, "kafka send msg err:UserId", v, msgToMQ.String())
					return returnMsg(&replay, pb, 201, "kafka send msg err", "", 0)
				}
			}
		}
		// callback
		if err := callbackAfterSendGroupMsg(pb); err != nil {
			log.NewError(pb.OperationID, utils.GetSelfFuncName(), "callbackAfterSendGroupMsg failed", err.Error())
		}
		return returnMsg(&replay, pb, 0, "", msgToMQ.MsgData.ServerMsgID, msgToMQ.MsgData.SendTime)
	default:
		return returnMsg(&replay, pb, 203, "unkonwn sessionType", "", 0)
	}
}

func modifyMessageByUserMessageReceiveOpt(userID, sourceID string, sessionType int, pb *pbChat.SendMsgReq) bool {
	conversationID := utils.GetConversationIDBySessionType(sourceID, sessionType)
	opt, err := db.DB.GetSingleConversationRecvMsgOpt(userID, conversationID)
	if err != nil && err != redis.ErrNil {
		log.NewError(pb.OperationID, "GetSingleConversationMsgOpt from redis err", conversationID, pb.String(), err.Error())
		return true
	}
	switch opt {
	case constant.ReceiveMessage:
		return true
	case constant.NotReceiveMessage:
		return false
	case constant.ReceiveNotNotifyMessage:
		if pb.MsgData.Options == nil {
			pb.MsgData.Options = make(map[string]bool, 10)
		}
		utils.SetSwitchFromOptions(pb.MsgData.Options, constant.IsOfflinePush, false)
		return true
	}

	return true
}

func returnMsg(replay *pbChat.SendMsgResp, pb *pbChat.SendMsgReq, errCode int32, errMsg, serverMsgID string, sendTime int64) (*pbChat.SendMsgResp, error) {
	replay.ErrCode = errCode
	replay.ErrMsg = errMsg
	replay.ServerMsgID = serverMsgID
	replay.ClientMsgID = pb.MsgData.ClientMsgID
	replay.SendTime = sendTime
	return replay, nil
}

func userRelationshipVerification(data *pbChat.SendMsgReq) (bool, int32, string) {
	if data.MsgData.SessionType == constant.GroupChatType {
		return true, 0, ""
	}
	log.NewDebug(data.OperationID, config.Config.MessageVerify.FriendVerify)
	req := &rpc.IsInBlackListReq{}
	req.CommID.OperationID = data.OperationID
	req.CommID.OpUserID = data.MsgData.RecvID
	req.CommID.FromUserID = data.MsgData.RecvID
	req.CommID.ToUserID = data.MsgData.SendID

	etcdConn := getcdv3.GetConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImFriendName)
	client := rpc.NewFriendClient(etcdConn)
	reply, err := client.IsInBlackList(context.Background(), req)
	if err != nil {
		log.NewDebug(data.OperationID, "IsInBlackListReq rpc failed, ", req.String(), err.Error())
	} else if reply.Response == true {
		log.NewDebug(data.OperationID, "IsInBlackListReq  ", req.String())
		return false, 600, "in black list"
	}
	log.NewDebug(data.OperationID, config.Config.MessageVerify.FriendVerify)
	if config.Config.MessageVerify.FriendVerify {
		friendReq := &rpc.IsFriendReq{CommID: &rpc.CommID{}}
		friendReq.CommID.OperationID = data.OperationID
		friendReq.CommID.OpUserID = data.MsgData.RecvID
		friendReq.CommID.FromUserID = data.MsgData.RecvID
		friendReq.CommID.ToUserID = data.MsgData.SendID
		friendReply, err := client.IsFriend(context.Background(), friendReq)
		if err != nil {
			log.NewDebug(data.OperationID, "IsFriendReq rpc failed, ", req.String(), err.Error())
			return true, 0, ""
		} else if friendReply.Response == false {
			log.NewDebug(data.OperationID, "not friend  ", req.String())
			return friendReply.Response, 601, "not friend"
		}
		log.NewDebug(data.OperationID, config.Config.MessageVerify.FriendVerify, friendReply.Response)
		return true, 0, ""
	} else {
		return true, 0, ""
	}
}
