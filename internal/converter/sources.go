package converter

import (
	"context"
	"errors"
	"io"
	"net/http"
	"os"
)

var (
	errNoContent = errors.New("no content")
)

type sourceFunc func(ctx context.Context, filePath string) (io.ReadCloser, error)

func localFS() sourceFunc {
	return func(ctx context.Context, filePath string) (io.ReadCloser, error) {
		file, err := os.Open(filePath)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				return nil, errNoContent
			}
		}

		return file, err
	}
}

func remoteHTTP() sourceFunc {
	return func(ctx context.Context, filePath string) (io.ReadCloser, error) {
		resp, err := http.Get(filePath)
		if err != nil {
			return nil, errNoContent
		}

		return resp.Body, nil
	}
}
