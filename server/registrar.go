package server

import "github.com/go-leo/leo/v3/runner"

type Server interface {
	runner.StartStopper
}
