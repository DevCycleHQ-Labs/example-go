package main

import (
	"log"
	"os"
	"time"
	devcycle "github.com/devcyclehq/go-server-sdk/v2"
)

var devcycleClient *devcycle.Client

func initalizeDevCycle() *devcycle.Client {
	sdkKey := os.Getenv("DEVCYCLE_SERVER_SDK_KEY")

	if len(sdkKey) == 0 {
		log.Fatalf("Add your DEVCYCLE_SERVER_SDK_KEY to the .env file")
	}

	options := devcycle.Options{
		EnableEdgeDB:                 false,
		EnableCloudBucketing:         false,
		EventFlushIntervalMS:         30 * time.Second,
		ConfigPollingIntervalMS:      5 * time.Second,
		RequestTimeout:               30 * time.Second,
		DisableAutomaticEventLogging: false,
		DisableCustomEventLogging:    false,
	}

	devcycleClient, err := devcycle.NewClient(sdkKey, &options)
	if err != nil {
		log.Fatalf("Error initializing DevCycle client: %v", err)
	}

	return devcycleClient
}

func getDevCycleClient() *devcycle.Client {
	if devcycleClient == nil {
		devcycleClient = initalizeDevCycle()
	}

	return devcycleClient
}