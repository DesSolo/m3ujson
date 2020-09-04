package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// provider is variable of provider name like Edem, etc
var provider string

// TVChannel structure tv channel
type TVChannel struct {
	Name     string `json:"name"`
	Provider string `json:"provider"`
	URL      string `json:"url"`
	Source   string `json:"source"`
}

// makeRequest download file from source url
// return readed bytes or error
func makeRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return nil, err
	}
	return data, nil
}

// convertToStruct convert downloaded data bytes to list of TVChannels
func convertToStruct(m3u []byte) ([]TVChannel, error) {
	var TvChannels []TVChannel

	re := regexp.MustCompile("#EXTINF.+,([^\\n]+)\\r\\n.*\\n(http://.+)\\r")
	for _, items := range re.FindAllSubmatch(m3u, -1) {
		source := string(items[2])
		url := strings.Split(source, "/")
		channel := TVChannel{
			Name:     string(items[1]),
			Provider: provider,
			URL:      "/" + url[len(url)-2],
			Source:   source,
		}
		TvChannels = append(TvChannels, channel)
	}
	if len(TvChannels) == 0 {
		return nil, errors.New("not found channels in source data")
	}
	return TvChannels, nil
}

func main() {
	flag.StringVar(&provider, "p", "Edem", "Provider name")
	flag.Parse()

	if len(flag.Args()) < 1 {
		flag.Usage()
		log.Fatalln("url should be specifically")
	}

	data, err := makeRequest(flag.Args()[len(flag.Args())-1])
	if err != nil {
		log.Fatalln(err)
	}

	channels, err := convertToStruct(data)
	if err != nil {
		log.Fatalln(err)
	}

	jsonData, err := json.MarshalIndent(channels, "", " ")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(jsonData))
}
