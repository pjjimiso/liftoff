package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
)

const url = "https://ll.thespacedevs.com/2.3.0/launches/upcoming/?location__ids=11&limit=1"
const cacheTTL = 5 * time.Minute

func main() {
	body, err := fetchWithCache(url)
	if err != nil {
		fmt.Printf("failed to fetch from api: %v", err)
		os.Exit(1)
	}

	launch := Launch{}
	err = json.Unmarshal(body, &launch)
	if err != nil {
		fmt.Printf("failed to parse json: %v", err)
		os.Exit(1)
	}

	if len(launch.Results) < 1 {
		fmt.Println("🚀 --")
		return
	}

	delta := time.Until(launch.Results[0].Net)
	if delta < 0 {
		fmt.Println("🚀 LIVE?")
		return
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

	fmt.Printf("🚀 %s", countdown)
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

	log.Printf("fetching new api data")
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
