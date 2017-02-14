GOCMD=go
GOINSTALL=$(GOCMD) install
GPATH=$$GOPATH

run:
	$(GOINSTALL) github.com/crazyfacka/pinky-02
	$(GPATH)/bin/pinky-02