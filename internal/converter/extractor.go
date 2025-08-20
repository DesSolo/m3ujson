package converter

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func extractByReader(src io.Reader) ([]Channel, error) {
	scanner := bufio.NewScanner(src)

	var (
		channels []Channel
		channel  *Channel
		id       uint64
	)

	for scanner.Scan() {
		line := scanner.Text()

		if !strings.HasPrefix(line, "#") && channel != nil {
			channel.Source = line
			channels = append(channels, *channel)
			channel = nil
		}

		if strings.HasPrefix(line, "#EXTINF") {
			id++

			channel = &Channel{
				ID: id,
			}

			parts := strings.SplitN(line, ",", 2)
			if len(parts) == 2 {
				channel.Name = strings.TrimSpace(parts[1])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner.Err: %w", err)
	}

	return channels, nil
}
