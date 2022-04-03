.PHONY: all clean

ifeq ($(OS),Windows_NT)
    MAKE = "C:/MinGW/msys/1.0/bin/make.exe"
endif

all: 
	"${MAKE}" -C cmd
clean:
	"${MAKE}" -C cmd clean