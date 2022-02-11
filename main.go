package main

import (
	"flag"
	"fmt"
	"jenkinstools/module"
	"jenkinstools/utils"
	"os"
)

var (
	c string
	f string
	h bool
	j string
	v string
)

func init() {
	flag.BoolVar(&h, "h", false, "This Is Help.")
	flag.StringVar(&c, "c", "", `usage: -r <srcFolder> <dstFolder> <srcName> <dstName>`)
	flag.StringVar(&v, "v", "", `usage: -c <viewName> <viewType>
	Only two view types are supported ---> [ListView|MyView]
	Default view is ListView
	If viewType is nil then this param must be "". example: -f <viewName> ""`)
	flag.StringVar(&f, "f", "", `usage: -f <folderName> <viewName> <parentName>
	If viewName or parentName is nil then this param must be "". example: -f <folderName> "" ""`)
	flag.StringVar(&j, "j", "", `usage: -j <jobName> <jobConfigXmlUrl>
	Default parentName is "". example: -f "SimOne_Test" "xxx"`)
	flag.Usage = utils.Usage
}

func main() {
	flag.Parse()
	switch {
	case h:
		flag.Usage()
	case c != "":
		if len(os.Args) > 5 {
			module.CopyJenkinsFolder(os.Args[2], os.Args[3], os.Args[4], os.Args[5])
		} else {
			fmt.Println("Please Input 4 Parameters.")
			flag.Usage()
		}
	case v != "":
		if len(os.Args) > 3 {
			if err := module.CreateView(os.Args[2], os.Args[3]); err == nil {
				fmt.Printf("Create view %v success.", os.Args[2])
			} else {
				fmt.Printf("Create view %v failed, error ---> %v", os.Args[2], err)
			}
		} else {
			fmt.Println("Please Input 3 Parameters.")
		}
	case f != "":
		if len(os.Args) > 4 {
			if err := module.CreateFolder(os.Args[2], os.Args[3], os.Args[4]); err == nil {
				fmt.Printf("Create folder %v success.", os.Args[2])
			} else {
				fmt.Printf("Create folder %v failed, error ---> %v", os.Args[2], err)
			}
		} else {
			fmt.Println("Please Input 4 Parameters.")
		}
	case j != "":
		if len(os.Args) > 3 {
			if url, err := module.CreateJenkinsJob(os.Args[2], os.Args[3]); err == nil {
				fmt.Printf("Create jenkins job %v success.\nJenkins job url ---> %s", os.Args[2], url)
			} else {
				fmt.Printf("Create Jenkins Job %v failed, error ---> %v", os.Args[2], err)
			}
		} else {
			fmt.Println("Please Input 4 Parameters.")
		}
	default:
		flag.Usage()
		os.Exit(1)
	}
}
