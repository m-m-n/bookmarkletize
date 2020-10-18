GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
EXECFILE=bookmarkletize

all: build
build:
	$(GOBUILD) -o $(EXECFILE) -v
clean:
	$(GOCLEAN)
	rm -f $(EXECFILE)
install:
	ln -sfv $(HOME)/go/src/$(EXECFILE)/$(EXECFILE) $(HOME)/bin/$(EXECFILE)
