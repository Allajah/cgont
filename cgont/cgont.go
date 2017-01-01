package cgont

import (
	"fmt"
	"os/exec"

	"encoding/json"
)
type Invalidation struct {
	Status string `json:"Status"`
	CreateTime string `json:"CreateTime"`
	Id string `json:"Id"`
}

type InvalidationList struct {
	Items []Invalidation `json:"Items"`
}


func Run(distId string)  {
	 out, err := exec.Command("aws", "cloudfront", "list-invalidations", "--distribution-id", distId).Output()
	inv := new(InvalidationList)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
	err = json.Unmarshal(out, inv)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(inv)
}
