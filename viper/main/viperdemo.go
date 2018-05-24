package main

import (
	"strings"
	"os"
	"fmt"
	"path/filepath"
	"github.com/spf13/viper"
)

const cmdRoot = "core"

func main() {
	viper.SetEnvPrefix(cmdRoot)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigName(cmdRoot)
	viper.AddConfigPath("./viper")
	gopath := os.Getenv("GOPATH")
	for _, p := range filepath.SplitList(gopath) {
		peerpath := filepath.Join(p, "src/vip")
		viper.AddConfigPath(peerpath)
	}
	fmt.Println("GOPATH = " + gopath)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		fmt.Println(fmt.Errorf("Fatal error when reading %s config file: %s\n", cmdRoot, err))
	}
	environment := viper.GetBool("security.enabled")
	fmt.Println("environment:", environment)

	fullstate := viper.GetString("statetransfer.timeout.fullstate")
	fmt.Println("fullstate:", fullstate)

	envValue := viper.GetString("viperconfig.env")
	fmt.Println("envValue is:", envValue)
}
