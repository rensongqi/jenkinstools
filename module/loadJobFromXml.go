package module

import (
	"context"
	"fmt"
	_ "github.com/bndr/gojenkins"
	"github.com/go-ini/ini"
	"io/ioutil"
	"jenkinstools/utils"
	"log"
	"net/http"
)

func CreateJenkinsJob(jobName string, configXmlUrl string) (string, error) {
	ctx := context.Background()
	c := utils.ConnectJenkins()
	client := &http.Client{}

	cfg, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("Reading config failed, please check config file ---> ", err)
	}
	userName := cfg.Section("jenkins").Key("UserName").String()
	password := cfg.Section("jenkins").Key("Password").String()

	req, err := http.NewRequest("GET", configXmlUrl, nil)
	if err != nil {
		log.Println("request failed", err)
	}
	req.SetBasicAuth(userName, password)
	response, err := client.Do(req)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	createJob, err := c.CreateJob(ctx, string(body), jobName)
	if err != nil {
		return "", err
	}

	return createJob.GetDetails().URL, nil
}