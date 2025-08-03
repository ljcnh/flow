package lifecycle

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/camunda/zeebe/clients/go/v8/pkg/worker"
	"github.com/camunda/zeebe/clients/go/v8/pkg/zbc"
	"github.com/spf13/viper"
)

type JobWorkerManager struct {
	client     zbc.Client
	handlers   map[string]worker.JobHandler
	jobWorkers []worker.JobWorker
}

func (j *JobWorkerManager) DeployResource() {
	dirPath := "conf/bpmn"
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return
	}

	var bpmnFiles []string
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), "bpmn") {
			bpmnFiles = append(bpmnFiles, filepath.Join(dirPath, entry.Name()))
		}
	}

	if len(bpmnFiles) == 0 {
		return
	}

	deployRequst := j.client.NewDeployResourceCommand().TenantId(viper.GetString("zeebe.tenantId"))
	for _, filePath := range bpmnFiles {
		fileBytes, err := os.ReadFile(filePath)
		if err != nil {
			continue
		}
		deployRequst = deployRequst.AddResource(fileBytes, filepath.Base(filePath))
	}

	ctx := context.Background()
	_, err = deployRequst.Send(ctx)
	if err != nil {
		panic(err)
	}

	// for _, deployment := range deployRespons.GetDeployments() {
	// 	process := deployment.GetProcess()
	// 	// log
	// }
}

// func (m *JobWorkerManager) Open(handlers map[string]common.W) {
// 	m.handlers = handlers
// }
