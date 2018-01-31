BUILD_DATE := `date +%Y-%m-%d\ %H:%M`
VERSION_NUMBER := 0.1.0
VERSIONFILE := version.go

.PHONY: test

all: gensrc install

gensrc:
		rm -f $(VERSIONFILE)
		@echo "package main" > $(VERSIONFILE)
		@echo "const (" >> $(VERSIONFILE)
		@echo "  VERSION = \"$(VERSION_NUMBER)\"" >> $(VERSIONFILE)
		@echo "  BUILD_DATE = \"$(BUILD_DATE)\"" >> $(VERSIONFILE)
		@echo ")" >> $(VERSIONFILE)

install:
		go install

test:
		go test -v ./...
