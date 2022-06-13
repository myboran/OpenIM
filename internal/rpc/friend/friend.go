package friend

import (
	"context"
	"net"
	"open-im/pkg/common/db"
	"strconv"
	"strings"
	"time"

	chat "open-im/internal/rpc/msg"
	"open-im/pkg/common/config"
	"open-im/pkg/common/constant"
	imdb "open-im/pkg/common/db/mysql_model/im_mysql_model"
	"open-im/pkg/common/log"
	"open-im/pkg/common/token_verify"
	cp "open-im/pkg/common/utils"
	"open-im/pkg/grpc-etcdv3/getcdv3"
	pbFriend "open-im/pkg/proto/friend"
	sdkws "open-im/pkg/proto/sdk_ws"
	"open-im/pkg/utils"

	"google.golang.org/grpc"
)

type friendServer struct {
	rpcPort         int
	rpcRegisterName string
	etcdSchema      string
	etcdAddr        []string
}

func NewFriendServer(port int) *friendServer {
	log.NewPrivateLog("friend")
	return &friendServer{
		rpcPort:         port,
		rpcRegisterName: config.Config.RpcRegisterName.OpenImFriendName,
		etcdSchema:      config.Config.Etcd.EtcdSchema,
		etcdAddr:        config.Config.Etcd.EtcdAddr,
	}
}

func (s *friendServer) Run() {
	log.NewInfo("0", "friendServer run...")

	ip := utils.ServerIP
	registerAddress := ip + ":" + strconv.Itoa(s.rpcPort)
	//listener network
	listener, err := net.Listen("tcp", registerAddress)
	if err != nil {
		log.NewError("0", "Listen failed ", err.Error(), registerAddress)
		return
	}
	log.NewInfo("0", "listen ok ", registerAddress)
	defer listener.Close()
	//grpc server
	srv := grpc.NewServer()
	defer srv.GracefulStop()
	//User friend related services register to etcd

	pbFriend.RegisterFriendServer(srv, s)
	err = getcdv3.RegisterEtcd(s.etcdSchema, strings.Join(s.etcdAddr, ","), ip, s.rpcPort, s.rpcRegisterName, 10)
	if err != nil {
		log.NewError("0", "RegisterEtcd failed ", err.Error(), s.etcdSchema, strings.Join(s.etcdAddr, ","), ip, s.rpcPort, s.rpcRegisterName)
		return
	}
	err = srv.Serve(listener)
	if err != nil {
		log.NewError("0", "Serve failed ", err.Error(), listener)
		return
	}
}

func (s *friendServer) GetFriendList(ctx context.Context, req *pbFriend.GetFriendListReq) (*pbFriend.GetFriendListResp, error) {
	log.NewInfo("GetFriendList args ", req.String())
	// 操作用户是否在管理用户列表
	if !token_verify.CheckAccess(req.CommID.OpUserID, req.CommID.FromUserID) {
		log.NewError(req.CommID.OperationID, "CheckAccess false ", req.CommID.OpUserID, req.CommID.FromUserID)
		return &pbFriend.GetFriendListResp{ErrCode: constant.ErrAccess.ErrCode, ErrMsg: constant.ErrAccess.ErrMsg}, nil
	}

	friends, err := imdb.GetFriendListByUserID(req.CommID.FromUserID)
	if err != nil {
		log.NewError(req.CommID.OperationID, "FindUserInfoFromFriend failed ", err.Error(), req.CommID.FromUserID)
		return &pbFriend.GetFriendListResp{ErrCode: constant.ErrDB.ErrCode, ErrMsg: constant.ErrDB.ErrMsg}, nil
	}

	var userInfoList []*sdkws.FriendInfo
	for _, friendUser := range friends {
		friendUserInfo := sdkws.FriendInfo{FriendUser: &sdkws.UserInfo{}}
		cp.FriendDBCopyOpenIM(&friendUserInfo, &friendUser)
		log.NewDebug(req.CommID.OperationID, "friends : ", friendUser, "openim friends: ", friendUserInfo)
		userInfoList = append(userInfoList, &friendUserInfo)
	}
	log.NewInfo(req.CommID.OperationID, "rpc GetFriendList ok", pbFriend.GetFriendListResp{FriendInfoList: userInfoList})
	return &pbFriend.GetFriendListResp{FriendInfoList: userInfoList}, nil
}

func (s *friendServer) AddFriend(ctx context.Context, req *pbFriend.AddFriendReq) (*pbFriend.AddFriendResp, error) {
	log.NewInfo(req.CommID.OperationID, "AddFriend args ", req.String())
	ok := token_verify.CheckAccess(req.CommID.OpUserID, req.CommID.FromUserID)
	if !ok {
		log.NewError(req.CommID.OperationID, "CheckAccess false ", req.CommID.OpUserID, req.CommID.FromUserID)
		return &pbFriend.AddFriendResp{CommonResp: &pbFriend.CommonResp{ErrCode: constant.ErrAccess.ErrCode, ErrMsg: constant.ErrAccess.ErrMsg}}, nil
	}

	// 不能添加不存在的用户
	if _, err := imdb.GetUserByUserID(req.CommID.ToUserID); err != nil {
		log.NewError(req.CommID.OperationID, "GetUserByUserID failed ", err.Error(), req.CommID.ToUserID)
		return &pbFriend.AddFriendResp{CommonResp: &pbFriend.CommonResp{ErrCode: constant.ErrDB.ErrCode, ErrMsg: constant.ErrDB.ErrMsg}}, nil
	}

	friendRequest := db.FriendRequest{
		HandleResult: 0, ReqMsg: req.ReqMsg, CreateTime: time.Now()}
	utils.CopyStructFields(&friendRequest, req.CommID)
	// {openIM001 openIM002 0 test add friend 0001-01-01 00:00:00 +0000 UTC   0001-01-01 00:00:00 +0000 UTC }]
	log.NewDebug(req.CommID.OperationID, "UpdateFriendApplication args ", friendRequest)

	err := imdb.InsertFriendApplication(&friendRequest,
		map[string]interface{}{"handle_result": 0, "req_msg": friendRequest.ReqMsg, "create_time": friendRequest.CreateTime,
			"handler_user_id": "", "handle_msg": "", "handle_time": utils.UnixSecondToTime(0), "ex": ""})
	if err != nil {
		log.NewError(req.CommID.OperationID, "UpdateFriendApplication failed ", err.Error(), friendRequest)
		return &pbFriend.AddFriendResp{CommonResp: &pbFriend.CommonResp{ErrCode: constant.ErrDB.ErrCode, ErrMsg: constant.ErrDB.ErrMsg}}, nil
	}

	chat.FriendApplicationNotification(req)
	return &pbFriend.AddFriendResp{CommonResp: &pbFriend.CommonResp{}}, nil
}

func (s *friendServer) DeleteFriend(ctx context.Context, req *pbFriend.DeleteFriendReq) (*pbFriend.DeleteFriendResp, error) {
	log.NewInfo(req.CommID.OperationID, "DeleteFriend args ", req.String())
	//Parse token, to find current user information
	if !token_verify.CheckAccess(req.CommID.OpUserID, req.CommID.FromUserID) {
		log.NewError(req.CommID.OperationID, "CheckAccess false ", req.CommID.OpUserID, req.CommID.FromUserID)
		return &pbFriend.DeleteFriendResp{CommonResp: &pbFriend.CommonResp{ErrCode: constant.ErrAccess.ErrCode, ErrMsg: constant.ErrAccess.ErrMsg}}, nil
	}
	err := imdb.DeleteSingleFriendInfo(req.CommID.FromUserID, req.CommID.ToUserID)
	if err != nil {
		log.NewError(req.CommID.OperationID, "DeleteSingleFriendInfo failed", err.Error(), req.CommID.FromUserID, req.CommID.ToUserID)
		return &pbFriend.DeleteFriendResp{CommonResp: &pbFriend.CommonResp{ErrCode: constant.ErrAccess.ErrCode, ErrMsg: constant.ErrAccess.ErrMsg}}, nil
	}

	log.NewInfo(req.CommID.OperationID, "DeleteFriend rpc ok")
	chat.FriendDeletedNotification(req)
	return &pbFriend.DeleteFriendResp{CommonResp: &pbFriend.CommonResp{}}, nil
}

func (s *friendServer) IsFriend(ctx context.Context, req *pbFriend.IsFriendReq) (*pbFriend.IsFriendResp, error) {
	log.NewInfo(req.CommID.OperationID, req.String())
	var isFriend bool
	if !token_verify.CheckAccess(req.CommID.OpUserID, req.CommID.FromUserID) {
		log.NewError(req.CommID.OperationID, "CheckAccess false ", req.CommID.OpUserID, req.CommID.FromUserID)
		return &pbFriend.IsFriendResp{ErrCode: constant.ErrAccess.ErrCode, ErrMsg: constant.ErrAccess.ErrMsg}, nil
	}
	_, err := imdb.GetFriendRelationshipFromFriend(req.CommID.FromUserID, req.CommID.ToUserID)
	if err == nil {
		isFriend = true
	} else {
		isFriend = false
	}
	log.NewInfo(req.CommID.OperationID, pbFriend.IsFriendResp{Response: isFriend})
	return &pbFriend.IsFriendResp{Response: isFriend}, nil
}

func (s *friendServer) IsInBlackList(ctx context.Context, req *pbFriend.IsInBlackListReq) (*pbFriend.IsInBlackListResp, error) {
	log.NewInfo("IsInBlackList args ", req.String())
	if !token_verify.CheckAccess(req.CommID.OpUserID, req.CommID.FromUserID) {
		log.NewError(req.CommID.OperationID, "CheckAccess false ", req.CommID.OpUserID, req.CommID.FromUserID)
		return &pbFriend.IsInBlackListResp{ErrCode: constant.ErrAccess.ErrCode, ErrMsg: constant.ErrAccess.ErrMsg}, nil
	}

	var isInBlacklist = false
	err := imdb.CheckBlack(req.CommID.FromUserID, req.CommID.ToUserID)
	if err == nil {
		isInBlacklist = true
	}
	log.NewInfo(req.CommID.OperationID, "IsInBlackList rpc ok ", pbFriend.IsInBlackListResp{Response: isInBlacklist})
	return &pbFriend.IsInBlackListResp{Response: isInBlacklist}, nil
}
