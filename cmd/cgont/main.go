package main

import (
	//"github.com/Allajah/cgont/cgont"
	"flag"
	"os"
	"github.com/Allajah/cgont/cgont"
)

func main() {

	listCommand := flag.NewFlagSet("list", flag.ExitOnError)

	distId := listCommand.String("dist-id", "DISTRIBUTIONID", "CloudFront Distribution ID")

	switch os.Args[1] {
	case "list":
		listCommand.Parse(os.Args[2:])
		cgont.ListInvalidations(*distId)
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

}
