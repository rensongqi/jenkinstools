package module

import (
	"context"
	"fmt"
	"github.com/bndr/gojenkins"
	"jenkinstools/utils"
)

// Possible Types:
// 		gojenkins.LIST_VIEW
// 		gojenkins.NESTED_VIEW
// 		gojenkins.MY_VIEW
// 		gojenkins.DASHBOARD_VIEW
// 		gojenkins.PIPELINE_VIEW
func CreateView(viewName, viewType string) error {
	ctx := context.Background()

	c := utils.ConnectJenkins()
	if viewType == "" {
		_, err := c.CreateView(ctx,viewName, gojenkins.LIST_VIEW)
		if err != nil {
			fmt.Println("Create view failed ---> ", err)
			return err
		}
	} else {
		_, err := c.CreateView(ctx,viewName,"hudson.model." + viewType)
		if err != nil {
			fmt.Println("Create view failed ---> ", err)
			return err
		}
	}
	return nil
}