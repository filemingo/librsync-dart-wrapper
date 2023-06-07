
MAKEFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
DIR := $(dir $(MAKEFILE_PATH))

OUT?=librsync.aar

all: build

build:
	GOROOT=$(DIR)/go gomobile bind -target=android -o $(OUT) github.com/filemingo/app/bridge

