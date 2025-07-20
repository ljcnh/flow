package cmd

import (
	"fmt"

	"github.com/ljcnh/flow/internal/pkg/env"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(apiCmd)
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "api 接口服务",
	Long:  `api 接口服务`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// 启动api服务器进行初始化
		apiCfg := cfgFile
		if apiCfg == "" {
			apiCfg = fmt.Sprintf("conf/api.%s.yaml", env.GetEnv())
		}
		viper.SetConfigType("yaml")
		viper.SetConfigFile(apiCfg)
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
		return viper.BindPFlags(cmd.Flags())
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
