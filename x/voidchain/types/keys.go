package types

const (
	// ModuleName defines the module name
	ModuleName = "voidchain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_voidchain"
)

var (
	ParamsKey = []byte("p_voidchain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
