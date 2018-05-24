package main

import (
	"time"
	"net"
	"os"
	"os/exec"
	"path"
	"fmt"
	"github.com/naoina/toml"
)

type tomlConfig struct {
	Title string
	Owner struct {
		Name string
		Dob  time.Time
	}
	Database struct {
		Server        string
		Ports         []int
		ConnectionMax uint
		Enabled       bool
	}
	Servers map[string]ServerInfo
	Clients struct {
		Data  [][]interface{}
		Hosts []string
	}
}

type ServerInfo struct {
	IP net.IP
	DC string
}

func main() {

	currentPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("current path:", currentPath)

	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		panic(err)
	}
	fmt.Println("exec path:", file)

	dir, _ := path.Split(file)
	fmt.Println("exec folder relative path:", dir)

	os.Chdir(dir)
	wd, _ := os.Getwd()
	fmt.Println("exec folder absolute path:", wd)

	f, err := os.Open(currentPath+"/example.toml")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var config tomlConfig
	if err := toml.NewDecoder(f).Decode(&config); err != nil {
		panic(err)
	}

	// then to use the unmarshaled config...
	fmt.Println("IP of server 'alpha':", config.Servers["alpha"].IP)
}
