package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/QuyYeuCode/denom/testutil/keeper"
	"github.com/QuyYeuCode/denom/x/denom/keeper"
	"github.com/QuyYeuCode/denom/x/denom/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DenomKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
