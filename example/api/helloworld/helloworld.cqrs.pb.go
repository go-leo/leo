// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package helloworld

import (
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

// GreeterAssembler responsible for completing the transformation between domain model objects and DTOs
type GreeterAssembler interface {
}

// GreeterCQRSService implement the Greeter service with CQRS pattern
type GreeterCQRSService struct {
	bus       cqrs.Bus
	assembler GreeterAssembler
}

func NewGreeterCQRSService(bus cqrs.Bus, assembler GreeterAssembler) *GreeterCQRSService {
	return &GreeterCQRSService{bus: bus, assembler: assembler}
}

func NewGreeterBus() (cqrs.Bus, error) {
	bus := cqrs.NewBus()
	return bus, nil
}
