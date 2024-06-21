package types

import "cosmossdk.io/collections"

const (
	// ModuleName defines the module name
	ModuleName = "testmodule"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_testmodule"
)

var (
	ParamsKey     = collections.NewPrefix(0)
	TestStructKey = collections.NewPrefix(1)
)
