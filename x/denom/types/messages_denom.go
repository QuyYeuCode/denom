package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateDenom = "create_denom"
	TypeMsgUpdateDenom = "update_denom"
	TypeMsgDeleteDenom = "delete_denom"
)

var _ sdk.Msg = &MsgCreateDenom{}

func NewMsgCreateDenom(
	owner string,
	denom string,
	description string,
	ticker string,
	precision int32,
	url string,
	maxSupply int32,
	//supply int32,
	canChangeMaxSupply bool,

) *MsgCreateDenom {
	return &MsgCreateDenom{
		Owner:       owner,
		Denom:       denom,
		Description: description,
		Ticker:      ticker,
		Precision:   precision,
		Url:         url,
		MaxSupply:   maxSupply,
		//	Supply:             supply,
		CanChangeMaxSupply: canChangeMaxSupply,
	}
}

func (msg *MsgCreateDenom) Route() string {
	return RouterKey
}

func (msg *MsgCreateDenom) Type() string {
	return TypeMsgCreateDenom
}

func (msg *MsgCreateDenom) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgCreateDenom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateDenom) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}

	tickerLength := len(msg.Ticker)
	if tickerLength < 3 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Ticker length must be at least 3 chars long")
	}
	if tickerLength > 10 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Ticker length must be 10 chars long maximum")
	}
	if msg.MaxSupply == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Max Supply must be greater than 0")
	}

	return nil
}

var _ sdk.Msg = &MsgUpdateDenom{}

func NewMsgUpdateDenom(
	owner string,
	denom string,
	description string,

	url string,
	maxSupply int32,

	canChangeMaxSupply bool,

) *MsgUpdateDenom {
	return &MsgUpdateDenom{
		Owner:       owner,
		Denom:       denom,
		Description: description,
		//	Ticker:             ticker,
		//	Precision:          precision,
		Url:       url,
		MaxSupply: maxSupply,
		//	Supply:             supply,
		CanChangeMaxSupply: canChangeMaxSupply,
	}
}

func (msg *MsgUpdateDenom) Route() string {
	return RouterKey
}

func (msg *MsgUpdateDenom) Type() string {
	return TypeMsgUpdateDenom
}

func (msg *MsgUpdateDenom) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgUpdateDenom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateDenom) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	if msg.MaxSupply == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Max Supply must be greater than 0")
	}
	return nil
}

// var _ sdk.Msg = &MsgDeleteDenom{}

// func NewMsgDeleteDenom(
// 	owner string,
// 	denom string,

// ) *MsgDeleteDenom {
// 	return &MsgDeleteDenom{
// 		Owner: owner,
// 		Denom: denom,
// 	}
// }
// func (msg *MsgDeleteDenom) Route() string {
// 	return RouterKey
// }

// func (msg *MsgDeleteDenom) Type() string {
// 	return TypeMsgDeleteDenom
// }

// func (msg *MsgDeleteDenom) GetSigners() []sdk.AccAddress {
// 	owner, err := sdk.AccAddressFromBech32(msg.Owner)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return []sdk.AccAddress{owner}
// }

// func (msg *MsgDeleteDenom) GetSignBytes() []byte {
// 	bz := ModuleCdc.MustMarshalJSON(msg)
// 	return sdk.MustSortJSON(bz)
// }

// func (msg *MsgDeleteDenom) ValidateBasic() error {
// 	_, err := sdk.AccAddressFromBech32(msg.Owner)
// 	if err != nil {
// 		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
// 	}
// 	return nil
// }
