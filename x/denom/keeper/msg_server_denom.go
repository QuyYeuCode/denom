package keeper

import (
	"context"

	"github.com/QuyYeuCode/denom/x/denom/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateDenom(goCtx context.Context, msg *types.MsgCreateDenom) (*types.MsgCreateDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetDenom(
		ctx,
		msg.Denom,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var denom = types.Denom{
		Owner:              msg.Owner,
		Denom:              msg.Denom,
		Description:        msg.Description,
		Ticker:             msg.Ticker,
		Precision:          msg.Precision,
		Url:                msg.Url,
		MaxSupply:          msg.MaxSupply,
	//	Supply:             msg.Supply,
		CanChangeMaxSupply: msg.CanChangeMaxSupply,
	}

	k.SetDenom(
		ctx,
		denom,
	)
	return &types.MsgCreateDenomResponse{}, nil
}

func (k msgServer) MintAndSendTokens(goCtx context.Context, msg *types.MsgMintAndSendTokens) (*types.MsgMintAndSendTokensResponse, error) {
    ctx := sdk.UnwrapSDKContext(goCtx)

    // Check if the value exists
    valFound, isFound := k.GetDenom(
        ctx,
        msg.Denom,
    )
    if !isFound {
        return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "denom does not exist")
    }

    // Checks if the the msg owner is the same as the current owner
    if msg.Owner != valFound.Owner {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
    }

    if valFound.Supply+msg.Amount > valFound.MaxSupply {
        return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Cannot mint more than Max Supply")
    }
    moduleAcct := k.accountKeeper.GetModuleAddress(types.ModuleName)

    recipientAddress, err := sdk.AccAddressFromBech32(msg.Recipient)
    if err != nil {
        return nil, err
    }

    var mintCoins sdk.Coins

    mintCoins = mintCoins.Add(sdk.NewCoin(msg.Denom, sdk.NewInt(int64(msg.Amount))))
    if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, mintCoins); err != nil {
        return nil, err
    }
    if err := k.bankKeeper.SendCoins(ctx, moduleAcct, recipientAddress, mintCoins); err != nil {
        return nil, err
    }

    var denom = types.Denom{
        Owner:              valFound.Owner,
        Denom:              valFound.Denom,
        Description:        valFound.Description,
        MaxSupply:          valFound.MaxSupply,
        Supply:             valFound.Supply + msg.Amount,
        Precision:          valFound.Precision,
        Ticker:             valFound.Ticker,
        Url:                valFound.Url,
        CanChangeMaxSupply: valFound.CanChangeMaxSupply,
    }

    k.SetDenom(
        ctx,
        denom,
    )
    return &types.MsgMintAndSendTokensResponse{}, nil
}

// func (k msgServer) DeleteDenom(goCtx context.Context, msg *types.MsgDeleteDenom) (*types.MsgDeleteDenomResponse, error) {
// 	ctx := sdk.UnwrapSDKContext(goCtx)

// 	// Check if the value exists
// 	valFound, isFound := k.GetDenom(
// 		ctx,
// 		msg.Denom,
// 	)
// 	if !isFound {
// 		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
// 	}

// 	// Checks if the the msg owner is the same as the current owner
// 	if msg.Owner != valFound.Owner {
// 		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
// 	}

// 	k.RemoveDenom(
// 		ctx,
// 		msg.Denom,
// 	)

// 	return &types.MsgDeleteDenomResponse{}, nil
// }
