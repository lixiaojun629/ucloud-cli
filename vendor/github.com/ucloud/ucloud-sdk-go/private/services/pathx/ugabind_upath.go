//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api PathX UGABindUPath

package pathx

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// UGABindUPathRequest is request schema for UGABindUPath action
type UGABindUPathRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// UGA ID
	UGAId *string `required:"true"`

	// 加速线路id
	UPathId *string `required:"true"`

	// 代金券
	CouponId *string `required:"false"`
}

// UGABindUPathResponse is response schema for UGABindUPath action
type UGABindUPathResponse struct {
	response.CommonBase
}

// NewUGABindUPathRequest will create request of UGABindUPath action.
func (c *PathXClient) NewUGABindUPathRequest() *UGABindUPathRequest {
	req := &UGABindUPathRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// UGABindUPath - UGA绑定UPath
func (c *PathXClient) UGABindUPath(req *UGABindUPathRequest) (*UGABindUPathResponse, error) {
	var err error
	var res UGABindUPathResponse

	err = c.Client.InvokeAction("UGABindUPath", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
