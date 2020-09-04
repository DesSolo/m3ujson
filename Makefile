PROJECTNAME="m3ujson"
BUILD_DIRECTORY="bin"
VERSION=${shell grep "const AppVersion" main.go | cut -d '"' -f 2}

clean:
	@echo "  >  Clean directory..."
	@rm -rf bin

build: clean
	@echo "  >  Building binary..."
	@go build -o ${BUILD_DIRECTORY}/${PROJECTNAME}-${VERSION}

run:
	@echo "  >  Run..."
	@go run ${FILE}

compile: clean
	@echo "  >  Build binary all arch..."
	@echo "+linux amd64"
	@GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ${BUILD_DIRECTORY}/${PROJECTNAME}-${VERSION}-linux-amd64
	@echo "+mipsle"
	@GOOS=linux GOARCH=mipsle go build -ldflags "-s -w" -o ${BUILD_DIRECTORY}/${PROJECTNAME}-${VERSION}-mipsle
