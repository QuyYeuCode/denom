package keeper_test

import (
	"testing"

	testkeeper "github.com/QuyYeuCode/denom/testutil/keeper"
	"github.com/QuyYeuCode/denom/x/denom/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DenomKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
