package keeper

import (
	"context"

	"github.com/QuyYeuCode/denom/x/denom/types"
)

type msgServer struct {
	Keeper
}

// UpdateDenom implements types.MsgServer.
func (k *msgServer) UpdateDenom(context.Context, *types.MsgUpdateDenom) (*types.MsgUpdateDenomResponse, error) {
	panic("unimplemented")
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}
