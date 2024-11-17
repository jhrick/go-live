GO = go
GOFLAGS = -ldflags="-s -w"

TARGET = go-live
SRC = $(wildcard cmd/**/main.go)
GO_BINS := $(shell $(GO) env GOBIN)

all: $(TARGET)

$(TARGET): $(SRC) 
	$(GO) build $(GOFLAGS) -o $(TARGET) $(SRC)

clean:
	rm -rf $(TARGET)

run: $(TARGET)
	./$(TARGET)

install: all
	cp -f $(TARGET) $(GO_BINS)
	rm -rf $(TARGET)
	echo done!

uninstall:
	rm -rf $(GO_BINS)/$(TARGET)
	echo uninstalled
