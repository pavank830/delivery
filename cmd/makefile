.PHONY: all clean

all:
	go mod tidy
ifeq ($(OS),Windows_NT)
	CGO_ENABLED=0 go build  -o ../bin/delivery.exe
else
	CGO_ENABLED=0 go build  -o ../bin/delivery
endif

clean:
ifeq ($(OS),Windows_NT)
	-powershell.exe -ExecutionPolicy Bypass -Command 'Remove-Item -Path ../bin/delivery.exe -ErrorAction Continue'
else
	-rm -f ../bin/delivery
endif