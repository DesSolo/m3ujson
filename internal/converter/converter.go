package converter

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

var sources = []sourceFunc{
	localFS(),
	remoteHTTP(),
}

type Converter struct {
	filePath string
}

func New(filePath string) *Converter {
	return &Converter{
		filePath: filePath,
	}
}

func (c *Converter) Convert(ctx context.Context, out io.Writer) error {
	source, err := c.source(ctx)
	if err != nil {
		return fmt.Errorf("source: %w", err)
	}
	defer source.Close()

	channels, err := extractByReader(source)
	if err != nil {
		return fmt.Errorf("extractByReader: %w", err)
	}

	if len(channels) == 0 {
		return errors.New("no channels found")
	}

	encoder := json.NewEncoder(out)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(channels); err != nil {
		return fmt.Errorf("encode: %w", err)
	}

	return nil
}

func (c *Converter) source(ctx context.Context) (io.ReadCloser, error) {
	for _, source := range sources {
		data, err := source(ctx, c.filePath)
		if err != nil {
			if errors.Is(err, errNoContent) {
				continue
			}
		}

		return data, nil
	}

	return nil, errNoContent
}
