package main

import "github.com/knative/pkg/signals"

func main() {
	// set up signals so we handle the first shutdown signal gracefully
	stopCh := signals.SetupSignalHandler()
	<-stopCh
}
