package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

type BaseData struct {
	Host  string
	User  string
	Pass  string
	Token string
	Group string
}

var rootCmd = &cobra.Command{
	Use:   "controller",
	Short: "Client comman",
	Long:  `This tool is controller command tool.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.controller.yaml)")
	rootCmd.Flags().Bool("toggle", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringP("host", "H", "127.0.0.1:50200", "host example: 127.0.0.1:50001")
	rootCmd.PersistentFlags().StringP("token", "t", "", "")
	rootCmd.PersistentFlags().StringP("authuser", "u", "", "username")
	rootCmd.PersistentFlags().StringP("authpass", "p", "", "password")
	rootCmd.PersistentFlags().StringP("group", "g", "", "group")
	rootCmd.PersistentFlags().BoolP("direct", "D", false, "direct connection to node")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".controller")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func Base(cmd *cobra.Command) BaseData {
	host, err := cmd.Flags().GetString("host")
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	authuser, err := cmd.Flags().GetString("authuser")
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	authpass, err := cmd.Flags().GetString("authpass")
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	token, err := cmd.Flags().GetString("token")
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	group, err := cmd.Flags().GetString("group")
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	return BaseData{
		Host:  host,
		User:  authuser,
		Pass:  authpass,
		Token: token,
		Group: group,
	}
}
