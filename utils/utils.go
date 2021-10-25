package utils

import (
	"context"
	"flag"
	"fmt"
	"github.com/bndr/gojenkins"
	"github.com/go-ini/ini"
	"os"
)

// Connect jenkins
func ConnectJenkins() *gojenkins.Jenkins {
	ctx := context.Background()

	cfg, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("Reading config failed, please check config file ---> ", err)
	}
	jenkinsUrl := cfg.Section("jenkins").Key("JenkinsUrl").String()
	userName := cfg.Section("jenkins").Key("UserName").String()
	password := cfg.Section("jenkins").Key("Password").String()

	jenkins := gojenkins.CreateJenkins(nil, jenkinsUrl, userName, password)

	conn, err := jenkins.Init(ctx)
	if err != nil {
		fmt.Println("jenkins init Went Wrong ---> ", err)
	}

	return conn
}

func Usage() {
	fmt.Fprintf(os.Stderr,
		`
Jenkins Tools, Version: 1.0.0
Usage: jenkinsTools [-c|-f|-h|-v]
	-c,		copy jobs from different folder
	-f,		create folder
	-h,		for help
	-v,		create view

Options:
`)
	flag.PrintDefaults()
	fmt.Println()
	fmt.Fprintf(os.Stderr, `Use "jenkinstools <command>" for more information about a command.`)
	fmt.Println()
}