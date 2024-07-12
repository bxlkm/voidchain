package voidchain_test

import (
	"testing"

	keepertest "github.com/bxlkm/voidchain/testutil/keeper"
	"github.com/bxlkm/voidchain/testutil/nullify"
	voidchain "github.com/bxlkm/voidchain/x/voidchain/module"
	"github.com/bxlkm/voidchain/x/voidchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VoidchainKeeper(t)
	voidchain.InitGenesis(ctx, k, genesisState)
	got := voidchain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
