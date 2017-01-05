package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Allajah/cgont/cgont"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		cgont.Help()
		os.Exit(1)
	}

	listCommand := flag.NewFlagSet("list", flag.ExitOnError)
	listDistId := listCommand.String("dist-id", "DISTRIBUTIONID", "CloudFront Distribution ID")

	watchCommand := flag.NewFlagSet("watch", flag.ContinueOnError)
	watchDistId := watchCommand.String("dist-id", "DISTRIBUTIONID", "CloudFront Distribution ID")
	watchInvalidationId := watchCommand.String("invalidation-id", "INVALIDATIONID", "CloudFront Invalidation ID")

	switch os.Args[1] {
	case "list":
		if len(os.Args[2:]) < 1 {
			msg := "This command needs Distribution ID.\nPlease specify Distribution ID with --dist-id option."
			fmt.Println(msg)
			os.Exit(1)
		}
		listCommand.Parse(os.Args[2:])
		cgont.ListInvalidations(*listDistId)
	case "watch":
		watchCommand.Parse(os.Args[2:])
		cgont.WatchInvalidation(*watchDistId, *watchInvalidationId)
	default:
		cgont.Help()
		os.Exit(1)
	}

}
