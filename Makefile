FILE := main.go
PROJECTNAME := m3ujson

_mdir:
	@mkdir bin

clean:
	@echo "  >  Clean directory..."
	@rm -rf bin

build: clean _mdir
	@echo "  >  Building binary..."
	@go build -o bin/${PROJECTNAME} ${FILE}

run:
	@echo "  >  Run..."
	@go run ${FILE}

compile: clean _mdir
	@echo "  >  Build binary all arch..."
	@echo "+linux"
	@GOOS=linux GOARCH=amd64 go build -o bin/${PROJECTNAME}-linux-amd64 ${FILE}
	@echo "+mipsle"
	@GOOS=linux GOARCH=mipsle go build -o bin/${PROJECTNAME}-mipsle ${FILE}
