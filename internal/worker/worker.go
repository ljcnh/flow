package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/camunda/zeebe/clients/go/v8/pkg/entities"
	"github.com/camunda/zeebe/clients/go/v8/pkg/worker"
	"github.com/ljcnh/flow/internal/infra/db"
	"github.com/ljcnh/flow/internal/infra/redis"
	"github.com/ljcnh/flow/internal/infra/zb"
)

func StartWorker() {
	fmt.Println("worker start")
	// mysql
	if err := db.InitMySQLClient(); err != nil {
		panic(err)
	}
	// redis
	if err := redis.InitRedis(); err != nil {
		panic(err)
	}

	// TODO: 初始化 producer
	// rmq.InitProducer()

	// TODO: 注册 feel
	// el.RegisterCustomFunction()
	// zeebe
	if err := zb.InitZeebeClient("localhost:26500"); err != nil {
		panic(err)
	}
	zbClient := zb.GetClient()

	// wm := lifecycle.NewJobWorkerManager(zbClient)
	// ticketStateTransitionBehavior := state.

	worker := zbClient.NewJobWorker().
		JobType("http-request").
		Handler(handleHTTPRequest).
		Name("http-worker").
		MaxJobsActive(10).
		Open()

	defer worker.Close()
	fmt.Println("HTTP Worker 启动成功，等待处理任务...")

	// 启动工作线程
	// worker := NewWorker()
	// worker.Start()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-sigChan

	// wm.Stop(5 * time.Second)
}

// 任务处理函数：发送 HTTP 请求并返回结果
func handleHTTPRequest(client worker.JobClient, job entities.Job) {
	// 1. 解析任务输入变量（从 Inputs 中获取的 HTTP 参数）
	var input struct {
		URL      string `json:"url"`    // 对应 Inputs 中的 url
		Method   string `json:"method"` // 对应 Inputs 中的 method
		Username string `json:"username"`
		Body     string `json:"body"` // 对应 Inputs 中的 body
	}

	// 从任务中获取变量（Inputs 配置的变量会被注入）
	if err := json.Unmarshal([]byte(job.Variables), &input); err != nil {
		failJob(client, job, err.Error())
		return
	}
	var (
		resp = make(map[string]interface{})
	)
	// 简单模拟一下 demo
	switch input.URL {
	case "get_auto_send":
		fmt.Println("get_auto_send")
		resp["is_auto_send"] = true
	case "send":
		fmt.Println("send")
		resp["send"] = "lijin"
	default:
		failJob(client, job, fmt.Sprintf("不支持的 URL: %s", input.URL))
		return
	}
	fmt.Println(input.Username)
	completeCmd, err := client.NewCompleteJobCommand().
		JobKey(job.Key).
		VariablesFromMap(resp)
	if err != nil {
		failJob(client, job, err.Error())
		return
	}
	_, err = completeCmd.Send(context.Background())
	if err != nil {
		fmt.Printf("任务 %d 完成通知: %v\n", job.Key, err)
	}
}

// 标记任务失败的工具函数
func failJob(client worker.JobClient, job entities.Job, message string) {
	_, err := client.NewFailJobCommand().
		JobKey(job.Key).
		Retries(job.Retries - 1). // 减少重试次数
		ErrorMessage(message).    // 错误信息
		Send(context.Background())
	if err != nil {
		fmt.Printf("任务 %d 失败通知失败: %v\n", job.Key, err)
	} else {
		fmt.Printf("任务 %d 失败: %s\n", job.Key, message)
	}
}
