package denom_test

import (
	"testing"

	keepertest "github.com/QuyYeuCode/denom/testutil/keeper"
	"github.com/QuyYeuCode/denom/testutil/nullify"
	"github.com/QuyYeuCode/denom/x/denom"
	"github.com/QuyYeuCode/denom/x/denom/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		DenomList: []types.Denom{
			{
				Denom: "0",
			},
			{
				Denom: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DenomKeeper(t)
	denom.InitGenesis(ctx, *k, genesisState)
	got := denom.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.DenomList, got.DenomList)
	// this line is used by starport scaffolding # genesis/test/assert
}
