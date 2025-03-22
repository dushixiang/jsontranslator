install:
	@echo "Installing..."
	go build -ldflags "-s -w " -o bin/jsont cmd/jsont/main.go
	sudo mv bin/jsont /usr/local/bin/jsont
	@echo "Installed!"