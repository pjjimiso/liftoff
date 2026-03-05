package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
)

const cacheTTL = 5 * time.Minute

func main() {
	argLocation := flag.Int("location", 11, "Location ID to query (default: 11, Vandenberg SFB)")
	argLimit := flag.Int("limit", 1, "number of upcoming launches to show")
	argFull := flag.Bool("full", false, "show full launch details")

	flag.Parse()

	url := fmt.Sprintf(
		"https://ll.thespacedevs.com/2.3.0/launches/upcoming/?location__ids=%d&limit=%d",
		*argLocation,
		*argLimit,
	)

	body, err := fetchWithCache(url)
	if err != nil {
		log.Printf("failed to fetch from api: %v", err)
		os.Exit(1)
	}

	launch := Launch{}
	err = json.Unmarshal(body, &launch)
	if err != nil {
		log.Printf("failed to parse json: %v", err)
		os.Exit(1)
	}

	if len(launch.Results) < 1 {
		fmt.Println("🚀 --")
		return
	}

	for _, result := range launch.Results {
		if *argFull {
			fmt.Println("Name:", result.Name)
			fmt.Println("Status:", result.Status.Name)
			fmt.Println("Launch Time (local):", result.Net.In(time.Local).Format("2006-01-02 15:04:05 MST"))
			fmt.Println("Service Provider:", result.LaunchServiceProvider.Name)
			fmt.Println("Rocket:", result.Rocket.Configuration.FullName)
			fmt.Printf("Mission: %s - %s\n", result.Mission.Name, result.Mission.Description)
			fmt.Printf("Launch Pad: %s - %s\n", result.Pad.Name, result.Pad.Location.Name)
			fmt.Println("Launch details last updated", result.LastUpdated)
		}

		delta := time.Until(result.Net)
		if delta < 0 {
			fmt.Println("🚀 LIVE?")
			continue
		}

		var countdown string
		days := int(delta.Hours()) / 24
		hours := int(delta.Hours()) % 24
		minutes := int(delta.Minutes()) % 60
		if days > 0 {
			countdown = fmt.Sprintf("%dd %dh %dm", days, hours, minutes)
		} else {
			countdown = fmt.Sprintf("%dh %dm", hours, minutes)
		}

		fmt.Println("🚀", countdown)
	}
}

func fetchWithCache(url string) ([]byte, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		fmt.Printf("failed to read cache dir %s: %v", cacheDir, err)
		fmt.Println("falling back to /tmp instead")
		cacheDir = os.TempDir()
	}

	cacheFile := filepath.Join(cacheDir, "liftoff_cache.json")
	fileInfo, err := os.Stat(cacheFile)
	if err == nil { // cache hit
		if time.Since(fileInfo.ModTime()) < cacheTTL {
			log.Printf("reading from cache file: %s", cacheFile)
			return os.ReadFile(cacheFile)
		}
		log.Printf("cache has expired")
	}

	log.Printf("fetching new api data from url %s", url)
	res, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "http request failed")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	// non-fatal if write fails - log any errors and return the data
	err = os.WriteFile(cacheFile, body, 0600)
	if err != nil {
		log.Printf("couldn't write to cache file: %v", err)
	}

	return body, nil
}
