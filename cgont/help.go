package cgont

import "fmt"

func Help() {
	help := `
Usage: cgont <command> [options]
Comands:
	list    Show CloudFront Invalidation list
	watch   Watch specified Invalidation
	`
	fmt.Println(help)
}
