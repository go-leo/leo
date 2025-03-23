package lbx

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-leo/gox/cryptox/knuthx"
	"hash"
)

var (
	ErrNoExtractor = errors.New("lbx: Extractor is nil")

	ErrNoHasher = errors.New("lbx: Hash is nil")

	ErrExtractionFailed = errors.New("lbx: failed to extract identity")
)

// HashFactory create a hash balancer
type HashFactory struct {
	Extractor func(ctx context.Context) (string, bool)
	Hasher    hash.Hash64
}

func (f HashFactory) New(ctx context.Context, endpointer sd.Endpointer) lb.Balancer {
	return &hashBalancer{
		Extractor:  f.Extractor,
		Hasher:     f.Hasher,
		Ctx:        ctx,
		Endpointer: endpointer,
	}
}

type hashBalancer struct {
	Extractor  func(ctx context.Context) (string, bool)
	Hasher     hash.Hash64
	Ctx        context.Context
	Endpointer sd.Endpointer
}

func (b *hashBalancer) Endpoint() (endpoint.Endpoint, error) {
	if b.Extractor == nil {
		return nil, ErrNoExtractor
	}
	if b.Hasher == nil {
		return nil, ErrNoHasher
	}
	identity, ok := b.Extractor(b.Ctx)
	if !ok {
		return nil, ErrExtractionFailed
	}
	if _, err := b.Hasher.Write([]byte(identity)); err != nil {
		return nil, err
	}
	endpoints, err := b.Endpointer.Endpoints()
	if err != nil {
		return nil, err
	}
	length := uint(len(endpoints))
	if length <= 0 {
		return nil, lb.ErrNoEndpoints
	}
	if length == 1 {
		return endpoints[0], nil
	}
	index := knuthx.Knuth(uint(b.Hasher.Sum64())) % length
	return endpoints[index], nil
}

// 一致性hash
