package api_test

import (
	"context"

	"github.com/brianvoe/gofakeit/v6"
	"go.infratographer.com/x/gidx"

	ent "go.infratographer.com/virtual-machine-api/internal/ent/generated"
)

const (
	locationPrefix = "testloc"
	ownerPrefix    = "testtnt"
	nodePrefix     = "testnod"
)

type VirtualMachineBuilder struct {
	Name       string
	OwnerID    gidx.PrefixedID
	LocationID gidx.PrefixedID
}

func (l *VirtualMachineBuilder) MustNew(ctx context.Context) *ent.VirtualMachine {
	if l.Name == "" {
		l.Name = gofakeit.AppName()
	}

	if l.OwnerID == "" {
		l.OwnerID = gidx.MustNewID(ownerPrefix)
	}

	if l.LocationID == "" {
		l.LocationID = gidx.MustNewID(locationPrefix)
	}

	return EntClient.VirtualMachine.Create().
		SetName(l.Name).
		SetOwnerID(l.OwnerID).
		SetLocationID(l.LocationID).
		SaveX(ctx)
}
