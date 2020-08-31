// Code is generated by ucloud-model, DO NOT EDIT IT.

package umem

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// UMemClient is the client of UMem
type UMemClient struct {
	*ucloud.Client
}

// NewClient will return a instance of UMemClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *UMemClient {
	meta := ucloud.ClientMeta{Product: "UMem"}
	client := ucloud.NewClientWithMeta(config, credential, meta)
	return &UMemClient{
		client,
	}
}