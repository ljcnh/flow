package cmd

import (
	"fmt"

	"github.com/ljcnh/flow/internal/pkg/env"
	"github.com/ljcnh/flow/internal/worker"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(workerCmd)
}

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "worker 工作节点",
	Long:  `worker 工作节点`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		//  启动服务器进行初始化
		workerCfg := cfgFile
		if workerCfg == "" {
			workerCfg = fmt.Sprintf("conf/worker.%s.yaml", env.GetEnv())
		}
		viper.SetConfigType("yaml")
		viper.SetConfigFile(workerCfg)
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
		return viper.BindPFlags(cmd.Flags())
	},
	Run: func(cmd *cobra.Command, args []string) {
		worker.StartWorker()
	},
}
