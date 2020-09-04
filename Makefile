PROJECTNAME="m3ujson"
BUILD_DIRECTORY="bin"

clean:
	@echo "  >  Clean directory..."
	@rm -rf bin

build: clean
	@echo "  >  Building binary..."
	@go build -o ${BUILD_DIRECTORY}/${PROJECTNAME}

run:
	@echo "  >  Run..."
	@go run ${FILE}

compile: clean
	@echo "  >  Build binary all arch..."
	@echo "+linux amd64"
	@GOOS=linux GOARCH=amd64 go build -o ${BUILD_DIRECTORY}/${PROJECTNAME}-linux-amd64
	@echo "+mipsle"
	@GOOS=linux GOARCH=mipsle go build -o ${BUILD_DIRECTORY}/${PROJECTNAME}-mipsle
