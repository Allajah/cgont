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

