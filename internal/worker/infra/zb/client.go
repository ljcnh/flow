package zb

import (
	"context"
	"time"

	"github.com/camunda/zeebe/clients/go/v8/pkg/zbc"
	"github.com/ljcnh/flow/internal/pkg/consts"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

var client zbc.Client

func GetClient() zbc.Client {
	return client
}

type TenantIdCredentialsProvider struct{}

func (t *TenantIdCredentialsProvider) ApplyCredentials(ctx context.Context, headers map[string]string) error {
	tenantID := viper.GetString("zeebe.tenantID")
	headers[consts.HeaderTenantID] = tenantID
	return nil
}

func (t *TenantIdCredentialsProvider) ShouldRetryRequest(ctx context.Context, _ error) bool {
	return false
}

func InitZeebeClient(address string) error {
	var (
		err error
	)
	retryPolicy := `{
		"retryPolicy" : {
			"maxAttempts": 30,
		}
	}`

	opts := []grpc.DialOption{
		grpc.WithKeepaliveParams(
			keepalive.ClientParameters{
				Time:                10 * time.Second,
				Timeout:             5 * time.Second,
				PermitWithoutStream: true,
			},
		),
		grpc.WithDefaultServiceConfig(retryPolicy),
	}

	config := &zbc.ClientConfig{
		GatewayAddress:         address,
		UsePlaintextConnection: true,
		DialOpts:               opts,
		CredentialsProvider:    &TenantIdCredentialsProvider{},
	}

	client, err = zbc.NewClient(config)
	if err != nil {
		panic(err)
	}

	return nil
}
