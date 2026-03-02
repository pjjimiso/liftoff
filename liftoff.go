package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	url := "https://ll.thespacedevs.com/2.3.0/launches/upcoming/?location__ids=11&limit=1"

	//TODO implement caching

	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("http request failed: %v", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("failed to read response body: %v", err)
		os.Exit(1)
	}

	launch := Launch{}
	err = json.Unmarshal(body, &launch)
	if err != nil {
		fmt.Printf("failed to parse json: %v", err)
		os.Exit(1)
	}

	delta := time.Until(launch.Results[0].Net)
	days := int(delta.Hours()) / 24
	hours := int(delta.Hours()) % 24
	minutes := int(delta.Minutes()) % 60

	fmt.Printf("🚀 %dd %dh %dm\n", days, hours, minutes)
}
