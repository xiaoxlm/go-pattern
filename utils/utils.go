package utils

import (
	"encoding/json"
	"fmt"
)

func LogJSON(v interface{}) {
	data, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(data))
}
