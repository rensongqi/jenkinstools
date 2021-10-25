package module

import (
	"context"
	"fmt"
	"jenkinstools/utils"
)

func CreateFolder(folderName, viewName, parentName string) error {
	ctx := context.Background()

	c := utils.ConnectJenkins()
	if parentName == "" {
		if viewName == "" {
			_, err := c.CreateFolder(ctx, folderName)
			if err != nil {
				fmt.Println("Create folder failed ---> ", err)
				return err
			}
		} else {
			//_, err := c.CreateFolderInView(ctx, folderName, viewName)
			//fmt.Println("-------------------->CreateFolderInView")
			//if err != nil {
			//	fmt.Println("Create folder failed ---> ", err)
			//	return err
			//}
			fmt.Println("Functional is improving...")
		}
	} else if parentName != "" {
		if viewName == "" {
			_, err := c.CreateFolder(ctx, folderName, parentName)
			if err != nil {
				fmt.Println("Create folder failed ---> ", err)
				return err
			}
		} else {
			_, err := c.CreateFolderInView(ctx, folderName, viewName, parentName)
			if err != nil {
				fmt.Println("Create folder failed ---> ", err)
				return err
			}
		}
	}

	return nil
}
