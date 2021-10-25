package module

import (
	"context"
	"fmt"
	"jenkinstools/utils"
	"strings"
)

func CopyJenkinsFolder(srcFolder, dstFolder, srcName, dstName string) {
	ctx := context.Background()

	c := utils.ConnectJenkins()
	parent, err := c.GetFolder(ctx, srcFolder)
	if err != nil {
		fmt.Println("Get folder failed ---> ", err)
	}
	jobs := parent.Raw.Jobs

	// 遍历源Folder中JobName
	for _, job := range jobs {
		j, _ := c.GetJob(ctx, job.Name, srcFolder)
		config, _ := j.GetConfig(ctx)
		newConfig := strings.Replace(config, srcName, dstName, -1)
		newJob := strings.Replace(job.Name, srcName, dstName, -1)
		// 在新的Folder中创建Job
		_, err = c.CreateJobInFolder(ctx, newConfig, newJob, dstFolder)
		if err != nil {
			fmt.Printf("copy job err: %v\n", err)
			_, _ = c.DeleteJob(ctx, newJob)
			_, _ = c.CreateJobInFolder(ctx, newConfig, newJob, dstFolder)
		}

		fmt.Printf("import %s success!\n", newJob)
	}
}