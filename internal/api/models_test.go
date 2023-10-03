package api_test

import (
	"context"
	"math/rand"

	"github.com/brianvoe/gofakeit/v6"
	"go.infratographer.com/x/gidx"

	ent "go.infratographer.com/virtual-machine-api/internal/ent/generated"
)

const (
	locationPrefix     = "testloc"
	ownerPrefix        = "testtnt"
	nodePrefix         = "testnod"
	cpuConfigPrefix    = "testcpu"
	memoryConfigPrefix = "testmem"
)

type VirtualMachineCPUConfigBuilder struct {
	ID      gidx.PrefixedID
	cores   int64
	sockets int64
}

func (p *VirtualMachineCPUConfigBuilder) MustNew(ctx context.Context) *ent.VirtualMachineCPUConfig {
	if p.ID == "" {
		p.ID = gidx.MustNewID(cpuConfigPrefix)
	}

	if p.cores == 0 {
		p.cores = int64(rand.Intn(128) + 1)
	}

	if p.sockets == 0 {
		p.sockets = int64(rand.Intn(128) + 1)
	}

	return EntClient.VirtualMachineCPUConfig.Create().SetID(p.ID).SetCores(p.cores).SetSockets(p.sockets).SaveX(ctx)
}

type VirtualMachineMemoryConfigBuilder struct {
	ID   gidx.PrefixedID
	Size int
}

func (m *VirtualMachineMemoryConfigBuilder) MustNew(ctx context.Context) *ent.VirtualMachineMemoryConfig {
	if m.ID == "" {
		m.ID = gidx.MustNewID(memoryConfigPrefix)
	}

	if m.Size == 0 {
		m.Size = rand.Intn(128) + 1
	}

	return EntClient.VirtualMachineMemoryConfig.Create().
		SetID(m.ID).
		SetSize(m.Size).
		SaveX(ctx)
}

type VirtualMachineBuilder struct {
	Name         string
	OwnerID      gidx.PrefixedID
	LocationID   gidx.PrefixedID
	CPUConfig    *ent.VirtualMachineCPUConfig
	MemoryConfig *ent.VirtualMachineMemoryConfig
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

	if l.CPUConfig == nil {
		pb := &VirtualMachineCPUConfigBuilder{}
		l.CPUConfig = pb.MustNew(ctx)
	}

	if l.MemoryConfig == nil {
		mb := &VirtualMachineMemoryConfigBuilder{}
		l.MemoryConfig = mb.MustNew(ctx)
	}

	return EntClient.VirtualMachine.Create().
		SetName(l.Name).
		SetOwnerID(l.OwnerID).
		SetLocationID(l.LocationID).
		SetVirtualMachineCPUConfig(l.CPUConfig).
		SetVirtualMachineMemoryConfig(l.MemoryConfig).
		SaveX(ctx)
}
