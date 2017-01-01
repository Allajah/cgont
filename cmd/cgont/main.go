package main

import (
	 "github.com/Allajah/cgont/cgont"
	"flag"
)

func main() {
	 var (
		distIdFlag string
	)

	flag.StringVar(&distIdFlag, "dist-id", "", "CloudFront Distribution ID")
	flag.Parse()

	cgont.Run(distIdFlag)
}
