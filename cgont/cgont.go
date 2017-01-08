package cgont

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"encoding/json"
)

type Invalidation struct {
	Status     string
	CreateTime string
	Id         string
}

type InvalidationList struct {
	Items []Invalidation
}

type Response struct {
	InvalidationList InvalidationList
}

type WatchResponse struct {
	Invalidation Invalidation
}

func ListInvalidations(distId string) {
	out, err := exec.Command("aws", "cloudfront", "list-invalidations", "--distribution-id", distId).Output()
	if err != nil {
		fmt.Println(err)
	}
	res := new(Response)
	err = json.Unmarshal(out, res)
	if err != nil {
		fmt.Println(err)
	}
	items := res.InvalidationList.Items

	for _, item := range items {
		fmt.Printf("CreatedTime    : %s \n", item.CreateTime)
		fmt.Printf("Invalidation ID: %s \n", item.Id)
		fmt.Printf("Status         : %s \n", item.Status)
		fmt.Printf("-----------------------------------------\n")
	}
}

func WatchInvalidation(distId, watchId string) {
	status := getInvalidationStatus(distId, watchId)
	if status == "Completed" {
		fmt.Printf("Invalidation ID: %s is alread completed", watchId)
		os.Exit(0)
	}
	fmt.Println("Watching invalidation...")
	for {
		status = getInvalidationStatus(distId, watchId)
		if status == "Completed" {
			fmt.Printf("Invalidation ID: %s is completed!", watchId)
			os.Exit(0)
		}
		time.Sleep(30 * time.Second)
	}

}

func getInvalidationStatus(distId, watchId string) string {
	out, err := exec.Command("aws", "cloudfront", "get-invalidation", "--distribution-id", distId, "--id", watchId).Output()
	if err != nil {
		fmt.Println("Getting invalidation error is occured. Make sure your distribution and invalidation is existing")
		os.Exit(1)
	}
	res := new(WatchResponse)
	err = json.Unmarshal(out, res)
	if err != nil {
		panic(err)
	}

	return res.Invalidation.Status
}
