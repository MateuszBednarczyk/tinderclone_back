package services

import (
	"context"
)

type ILocaliser interface {
	GetNearbyTowns(r int, ctx context.Context) []string
}

type localiser struct {
}

func NewLocaliser() *localiser {
	return &localiser{}
}

func (*localiser) GetNearbyTowns(r int, town string, ctx context.Context) []string {
	return []string{}
}
