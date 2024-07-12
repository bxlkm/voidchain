package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/bxlkm/voidchain/testutil/keeper"
	"github.com/bxlkm/voidchain/x/voidchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.VoidchainKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
