package api_test

import (
	"context"

	"github.com/brianvoe/gofakeit/v6"
	"go.infratographer.com/x/gidx"

	ent "go.infratographer.com/virtual-machine-api/internal/ent/generated"
)

const (
	locationPrefix  = "testloc"
	ownerPrefix     = "testtnt"
	nodePrefix      = "testnod"
	cpuConfigPrefix = "testcpu"
)

type VirtualMachineCPUConfigBuilder struct {
	ID      gidx.PrefixedID
	OwnerID gidx.PrefixedID
	cores   int
	sockets int
}

func (p *VirtualMachineCPUConfigBuilder) MustNew(ctx context.Context) *ent.VirtualMachineCPUConfig {
	if p.ID == "" {
		p.ID = gidx.MustNewID(cpuConfigPrefix)
	}

	if p.OwnerID == "" {
		p.OwnerID = gidx.MustNewID(ownerPrefix)
	}

	if p.cores == 0 {
		p.cores = 2
	}

	if p.sockets == 0 {
		p.sockets = 2
	}

	return EntClient.VirtualMachineCPUConfig.Create().SetID(p.ID).SetOwnerID(p.OwnerID).SetCores(p.cores).SetSockets(p.sockets).SaveX(ctx)
}

type VirtualMachineBuilder struct {
	Name       string
	OwnerID    gidx.PrefixedID
	LocationID gidx.PrefixedID
	CPUConfig  *ent.VirtualMachineCPUConfig
}

func (l *VirtualMachineBuilder) MustNew(ctx context.Context) *ent.VirtualMachine {
	if l.CPUConfig == nil {
		pb := &VirtualMachineCPUConfigBuilder{OwnerID: l.OwnerID}
		l.CPUConfig = pb.MustNew(ctx)
	}

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
		SetVirtualMachineCPUConfig(l.CPUConfig).
		SaveX(ctx)
}
