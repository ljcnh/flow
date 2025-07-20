package worker

import (
	"github.com/ljcnh/flow/internal/worker/infra/zb"
	"github.com/spf13/viper"
)

func StartWorker() {
	// zeebe
	if err := zb.InitZeebeClient(viper.GetString("zeebe.gateway.address")); err != nil {
		panic(err)
	}
	// client = zb.GetClient()

}
