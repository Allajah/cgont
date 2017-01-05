package main

import (
	"flag"
	"os"
	"github.com/Allajah/cgont/cgont"
)

func main() {

	listCommand := flag.NewFlagSet("list", flag.ExitOnError)
	listDistId := listCommand.String("dist-id", "DISTRIBUTIONID", "CloudFront Distribution ID")

	watchCommand := flag.NewFlagSet("watch", flag.ContinueOnError)
	watchDistId := watchCommand.String("dist-id", "DISTRIBUTIONID", "CloudFront Distribution ID")
	watchInvalidationId := watchCommand.String("invalidation-id", "INVALIDATIONID", "CloudFront Invalidation ID")

	switch os.Args[1] {
	case "list":
		listCommand.Parse(os.Args[2:])
		cgont.ListInvalidations(*listDistId)
	case "watch":
		watchCommand.Parse(os.Args[2:])
		cgont.WatchInvalidation(*watchDistId, *watchInvalidationId)
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

}
