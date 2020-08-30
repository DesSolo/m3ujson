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

// Provider ...
var Provider string

// TVChannel ...
type TVChannel struct {
	Name string `json:"name"`
	Provider string `json:"provider"`
	URL string `json:"url"`
	Source string `json:"source"`
}

func getM3U(url string) ([]byte, error) {
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

func convertToStruct(m3u []byte) ([]TVChannel, error){
	var TvChannels []TVChannel

	re := regexp.MustCompile("#EXTINF.+,([^\\n]+)\\r\\n.*\\n(http://.+)\\r")
	for _, items := range re.FindAllSubmatch(m3u, -1) {
		source := string(items[2])
		url := strings.Split(source, "/")
		channel := TVChannel{
			string(items[1]),
			Provider,
			"/" + url[len(url)-2],
			source,
		}
		TvChannels = append(TvChannels, channel)
	}
	if len(TvChannels) == 0 {
		return nil, errors.New("not found channels in target")
	}
	return TvChannels, nil
}


func main()  {
	flag.StringVar(&Provider, "p", "Edem", "Provider name")
	flag.Parse()

	if len(flag.Args()) < 1 {
		flag.Usage()
		log.Fatalln("Url should be specifically")
	}

	m3u, err := getM3U(flag.Args()[len(flag.Args())-1])
	if err != nil {
		log.Fatalln(err)
	}

	channels, err := convertToStruct(m3u)
	if err != nil {
		log.Fatalln(err)
	}
	
	data, err := json.MarshalIndent(channels, "", " ")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(data))
}
