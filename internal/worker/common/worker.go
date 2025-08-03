package common

import (
	"github.com/camunda/zeebe/clients/go/v8/pkg/zbc"
)

type WrapperWorker struct {
	Client zbc.Client
}
