package user

import (
	"context"
	"net"
	"strconv"
	"strings"

	"open-im/pkg/common/config"
	"open-im/pkg/common/constant"
	imdb "open-im/pkg/common/db/mysql_model/im_mysql_model"
	"open-im/pkg/common/log"
	"open-im/pkg/grpc-etcdv3/getcdv3"
	sdkws "open-im/pkg/proto/sdk_ws"
	pbUser "open-im/pkg/proto/user"
	"open-im/pkg/utils"

	"google.golang.org/grpc"
)

type userServer struct {
	rpcPort         int
	rpcRegisterName string
	etcdSchema      string
	etcdAddr        []string
}

func NewUserServer(port int) *userServer {
	log.NewPrivateLog("user")
	return &userServer{
		rpcPort:         port,
		rpcRegisterName: config.Config.RpcRegisterName.OpenImUserName,
		etcdSchema:      config.Config.Etcd.EtcdSchema,
		etcdAddr:        config.Config.Etcd.EtcdAddr,
	}
}

func (s *userServer) Run() {
	log.NewInfo("0", "", "rpc user start...")
	ip := utils.ServerIP
	registerAddress := ip + ":" + strconv.Itoa(s.rpcPort)
	listener, err := net.Listen("tcp", registerAddress)
	if err != nil {
		log.NewError("0", "listen network failed", err.Error(), registerAddress)
		return
	}
	log.NewInfo("0", "listen network success, address ", registerAddress, listener)
	defer listener.Close()
	//grpc server
	srv := grpc.NewServer()
	defer srv.GracefulStop()
	// Service registers with etcd
	pbUser.RegisterUserServer(srv, s)

	err = getcdv3.RegisterEtcd(s.etcdSchema, strings.Join(s.etcdAddr, ","), ip, s.rpcPort, s.rpcRegisterName, 10)
	if err != nil {
		log.NewError("0", "RegisterEtcd failed", err.Error(), s.etcdSchema, strings.Join(s.etcdAddr, ","), ip, s.rpcPort, s.rpcRegisterName)
		return
	}
	err = srv.Serve(listener)
	if err != nil {
		log.NewError("0", "Serve failed", err.Error())
		return
	}
	log.NewInfo("0", "rpc user success")
}

func (s *userServer) GetUserInfo(ctx context.Context, req *pbUser.GetUserInfoReq) (*pbUser.GetUserInfoResp, error) {
	log.NewInfo(req.OperationID, "GetUserInfo args", req.String())
	var userInfoList []*sdkws.UserInfo

	if len(req.UserIDList) > 0 {
		for _, userId := range req.UserIDList {
			var userInfo sdkws.UserInfo
			user, err := imdb.GetUserByUserID(userId)
			if err != nil {
				log.NewError(req.OperationID, "GetUserByUserID failed", err.Error(), userId)
				continue
			}
			utils.CopyStructFields(&userInfo, user)
			userInfo.Birth = uint32(user.Birth.Unix())
			userInfoList = append(userInfoList, &userInfo)
		}
	} else {
		return &pbUser.GetUserInfoResp{CommonResp: &pbUser.CommonResp{ErrCode: constant.ErrArgs.ErrCode, ErrMsg: constant.ErrArgs.ErrMsg}}, nil
	}
	log.NewInfo(req.OperationID, "GetUserInfo rpc return ", pbUser.GetUserInfoResp{CommonResp: &pbUser.CommonResp{}, UserInfoList: userInfoList})
	return &pbUser.GetUserInfoResp{CommonResp: &pbUser.CommonResp{}, UserInfoList: userInfoList}, nil
}
