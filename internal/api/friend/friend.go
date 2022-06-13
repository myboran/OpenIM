package friend

import (
	"context"
	"net/http"
	"strings"

	api "open-im/pkg/base_info"
	"open-im/pkg/common/config"
	"open-im/pkg/common/log"
	"open-im/pkg/common/token_verify"
	"open-im/pkg/grpc-etcdv3/getcdv3"
	rpc "open-im/pkg/proto/friend"
	"open-im/pkg/utils"

	"github.com/gin-gonic/gin"
)

func AddFriend(c *gin.Context) {
	params := api.AddFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	req := &rpc.AddFriendReq{}
	utils.CopyStructFields(req.CommID, &params.ParamsCommFriend)
	req.ReqMsg = params.ReqMsg
	var ok bool
	ok, req.CommID.OpUserID = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), req.CommID.OperationID)
	if !ok {
		log.NewError(req.CommID.OperationID, "GetUserIDFromToken false ", c.Request.Header.Get("token"))
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "GetUserIDFromToken failed"})
		return
	}
	log.NewInfo(req.CommID.OperationID, "AddFriend args ", req.String())

	etcdConn := getcdv3.GetConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImFriendName)
	client := rpc.NewFriendClient(etcdConn)
	RpcResp, err := client.AddFriend(context.Background(), req)
	if err != nil {
		log.NewError(req.CommID.OperationID, "AddFriend failed ", err.Error(), req.String())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "call AddFriend rpc server failed"})
		return
	}

	resp := api.AddFriendResp{CommResp: api.CommResp{ErrCode: RpcResp.CommonResp.ErrCode, ErrMsg: RpcResp.CommonResp.ErrMsg}}
	log.NewInfo(req.CommID.OperationID, "AddFriend api return ", resp)
	c.JSON(http.StatusOK, resp)
}

func DeleteFriend(c *gin.Context) {
	params := api.DeleteFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	req := &rpc.DeleteFriendReq{CommID: &rpc.CommID{}}
	utils.CopyStructFields(req.CommID, &params.ParamsCommFriend)
	var ok bool
	ok, req.CommID.OpUserID = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), req.CommID.OperationID)
	if !ok {
		log.NewError(req.CommID.OperationID, "GetUserIDFromToken false ", c.Request.Header.Get("token"))
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "GetUserIDFromToken failed"})
		return
	}
	log.NewInfo(req.CommID.OperationID, "DeleteFriend args ", req.String())

	etcdConn := getcdv3.GetConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImFriendName)
	client := rpc.NewFriendClient(etcdConn)
	RpcResp, err := client.DeleteFriend(context.Background(), req)
	if err != nil {
		log.NewError(req.CommID.OperationID, "DeleteFriend failed ", err, req.String())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "call delete_friend rpc server failed"})
		return
	}

	resp := api.DeleteFriendResp{CommResp: api.CommResp{ErrCode: RpcResp.CommonResp.ErrCode, ErrMsg: RpcResp.CommonResp.ErrMsg}}
	log.NewInfo(req.CommID.OperationID, "DeleteFriend api return ", resp)
	c.JSON(http.StatusOK, resp)
}
