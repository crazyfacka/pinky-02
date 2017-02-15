GOCMD=go
GOINSTALL=$(GOCMD) install
GOBUILD=$(GOCMD) build
GPATH=$$GOPATH

install: 
	glide install

run:
	$(GOINSTALL) github.com/crazyfacka/pinky-02
	$(GPATH)/bin/pinky-02

build:
	CGO_ENABLED=0 $(GOBUILD) -a -ldflags '-s' -installsuffix cgo -o app .