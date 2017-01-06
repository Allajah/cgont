package cgont

import (
	"fmt"
	"os/exec"

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
	inv, err := getInvalidation(distId, watchId)
	if err != nil {
		panic(err)
	}
	res := new(WatchResponse)
	err = json.Unmarshal(inv, res)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Invalidation)

}

func getInvalidation(distId, watchId string) ([]byte, error) {
	out, err := exec.Command("aws", "cloudfront", "get-invalidation", "--distribution-id", distId, "--id", watchId).Output()
	if err != nil {
		return nil, err
	}
	return out, nil
}
