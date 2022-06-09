package user

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	jsonData "open-im/internal/utils"
	api "open-im/pkg/base_info"
	"open-im/pkg/common/config"
	"open-im/pkg/common/log"
	"open-im/pkg/common/token_verify"
	"open-im/pkg/grpc-etcdv3/getcdv3"
	sdkws "open-im/pkg/proto/sdk_ws"
	rpc "open-im/pkg/proto/user"
	"open-im/pkg/utils"

	"github.com/gin-gonic/gin"
)

func UpdateUserInfo(c *gin.Context) {
	fmt.Println("user.UpdateUserInfo")
	params := api.GetUsersInfoReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": http.StatusBadRequest, "errMsg": err.Error()})
		return
	}

	req := &rpc.GetUserInfoReq{}
	utils.CopyStructFields(req, &params)
	var ok bool
	ok, req.OpUserID = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), req.OperationID)
	if !ok {
		log.NewError(req.OperationID, "GetUserIDFromToken false ", c.Request.Header.Get("token"))
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "GetUserIDFromToken failed"})
		return
	}
	log.NewInfo(params.OperationID, "GetUserInfo args ", req.String())

	etcdConn := getcdv3.GetConn(
		config.Config.Etcd.EtcdSchema,
		strings.Join(config.Config.Etcd.EtcdAddr, ","),
		config.Config.RpcRegisterName.OpenImUserName,
	)
	client := rpc.NewUserClient(etcdConn)
	RpcResp, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		log.NewError(req.OperationID, "GetUserInfo failed ", err.Error(), req.String())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "call  rpc server failed"})
		return
	}
	var publicUserInfoList []*sdkws.PublicUserInfo
	for _, v := range RpcResp.UserInfoList {
		publicUserInfoList = append(publicUserInfoList,
			&sdkws.PublicUserInfo{UserID: v.UserID, Nickname: v.Nickname, FaceURL: v.FaceURL, Gender: v.Gender, Ex: v.Ex})
	}

	resp := api.GetUsersInfoResp{
		CommResp:     api.CommResp{ErrCode: RpcResp.CommonResp.ErrCode, ErrMsg: RpcResp.CommonResp.ErrMsg},
		UserInfoList: publicUserInfoList,
	}

	resp.Data = jsonData.JsonDataList(resp.UserInfoList)
	log.NewInfo(req.OperationID, "GetUserInfo api return ", resp)
	c.JSON(http.StatusOK, resp)
}
