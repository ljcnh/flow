package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile = ""
)

// rootCmd 表示根命令，当没有子命令被指定时执行
var rootCmd = &cobra.Command{
	Use:   "flow",
	Short: "flow 工作流引擎",
	Long:  `flow 工作流引擎`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Usage()
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().String("log-level", "info", "设置日志级别")
	rootCmd.PersistentFlags().String("log-format", "text", "设置日志格式")
	rootCmd.PersistentFlags().String("env", "local", "设置当前环境")
	if err := viper.BindPFlags(rootCmd.PersistentFlags()); err != nil {
		panic(err)
	}
}

// 初始化配置文件
func initConfig() {
	if cfgFile != "" {
		// 从指定路径读取配置文件
		viper.SetConfigFile(cfgFile)
	} else {
		// 查找主目录
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// 设置配置文件名为 .flow.yaml
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".flow.yaml")
	}

	// 从环境变量读取配置
	viper.AutomaticEnv()

	// 尝试读取配置文件
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "使用配置文件:", viper.ConfigFileUsed())
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
