package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateDenom{}, "denom/CreateDenom", nil)
	cdc.RegisterConcrete(&MsgUpdateDenom{}, "denom/UpdateDenom", nil)
	//cdc.RegisterConcrete(&MsgDeleteDenom{}, "denom/DeleteDenom", nil)
	cdc.RegisterConcrete(&MsgMintAndSendTokens{}, "denom/MintAndSendTokens", nil)
	cdc.RegisterConcrete(&MsgUpdateOwner{}, "denom/UpdateOwner", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateDenom{},
		&MsgUpdateDenom{},
		// &MsgDeleteDenom{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMintAndSendTokens{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateOwner{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
