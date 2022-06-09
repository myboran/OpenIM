package msg

//
//import (
//	"open-im/pkg/common/config"
//	"open-im/pkg/common/constant"
//	"open-im/pkg/common/log"
//	sdkws "open-im/pkg/proto/sdk_ws"
//	"open-im/pkg/utils"
//
//	"github.com/golang/protobuf/jsonpb"
//	"github.com/golang/protobuf/proto"
//)
//
//func SetConversationNotification(operationID, userID string) {
//	log.NewInfo(operationID, utils.GetSelfFuncName(), "userID: ", userID)
//	conversationUpdateTips := sdkws.ConversationUpdateTips{
//		UserID: userID,
//	}
//	conversationNotification(constant.ConversationOptChangeNotification, &conversationUpdateTips, operationID, userID)
//}
//
//func conversationNotification(contentType int32, m proto.Message, operationId, userID string) {
//	var err error
//	var tips sdkws.TipsComm
//	tips.Detail, err = proto.Marshal(m)
//	if err != nil {
//		log.Error(operationId, utils.GetSelfFuncName(), "Marshal failed ", err.Error(), m.String())
//		return
//	}
//	marshaler := jsonpb.Marshaler{
//		OrigName:     true,
//		EnumsAsInts:  false,
//		EmitDefaults: false,
//	}
//	tips.JsonDetail, _ = marshaler.MarshalToString(m)
//	cn := config.Config.Notification
//	switch contentType {
//	case constant.ConversationOptChangeNotification:
//		tips.DefaultTips = cn.ConversationOptUpdate.DefaultTips.Tips
//	}
//
//	var n NotificationMsg
//	n.SendID = userID
//	n.RecvID = userID
//	n.ContentType = contentType
//	n.
//}
