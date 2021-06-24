package cmd

import (
	"fmt"
	"net/http"
	"os"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
    "github.com/spf13/viper"
    "github.com/hasura/go-graphql-client"
)

var cfgFile string
var listFlag bool
var apiEndpoint string
var adminSecret string
var client *graphql.Client

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A cli tool to manage your todos",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
//	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
    rootCmd.PersistentFlags().StringVar(&apiEndpoint, "api-endpoint",
        "", "API endpoint for graphql engine",
    )
	rootCmd.PersistentFlags().BoolVar(&listFlag, "list", false, "List todos after query")
	cobra.OnInitialize(initConfig)

}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		cobra.CheckErr(err)

        viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.SetConfigName(".todo")
	}

    viper.SetEnvPrefix("TODO")
    viper.BindEnv("api_endpoint")
	viper.BindEnv("admin_secret")
	viper.AutomaticEnv()

    viper.ReadInConfig()
//    if err := viper.ReadInConfig(); err == nil {
//		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
//	} else {
//        fmt.Fprintln(os.Stderr, err)
//    }

    if apiEndpoint == "" && viper.Get("api_endpoint") == nil {
        fmt.Fprintln(os.Stderr,
            "Using default address(http://localhost:8080/v1/graphql)",
        )
        apiEndpoint = "http://localhost:8080/v1/graphql"
    } else if apiEndpoint == "" {
        apiEndpoint = viper.Get("api_endpoint").(string)
    }
	adminSecret = fmt.Sprintf("%v", viper.Get("admin_secret"))

	httpClient := &http.Client{Transport: &transport{underlyingTransport:http.DefaultTransport}}
	client = graphql.NewClient(apiEndpoint, httpClient)
}
