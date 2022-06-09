package base_info

import open_im_sdk "open-im/pkg/proto/sdk_ws"

type GetUsersInfoReq struct {
	OperationID string   `json:"operationID" binding:"required"`
	UserIDList  []string `json:"userIDList" binding:"required"`
}
type GetUsersInfoResp struct {
	CommResp
	UserInfoList []*open_im_sdk.PublicUserInfo `json:"-"`
	Data         []map[string]interface{}      `json:"data"`
}

type UpdateSelfUserInfoReq struct {
	ApiUserInfo
	OperationId string `json:"operationID" binding:"required"`
}

type UpdateUserInfoResp struct {
	CommResp
}
