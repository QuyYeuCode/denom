// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgUpdateDenom } from "./types/denom/tx";
import { MsgCreateDenom } from "./types/denom/tx";
import { MsgMintAndSendTokens } from "./types/denom/tx";
import { MsgUpdateOwner } from "./types/denom/tx";


const types = [
  ["/QuyYeuCode.denom.denom.MsgUpdateDenom", MsgUpdateDenom],
  ["/QuyYeuCode.denom.denom.MsgCreateDenom", MsgCreateDenom],
  ["/QuyYeuCode.denom.denom.MsgMintAndSendTokens", MsgMintAndSendTokens],
  ["/QuyYeuCode.denom.denom.MsgUpdateOwner", MsgUpdateOwner],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgUpdateDenom: (data: MsgUpdateDenom): EncodeObject => ({ typeUrl: "/QuyYeuCode.denom.denom.MsgUpdateDenom", value: MsgUpdateDenom.fromPartial( data ) }),
    msgCreateDenom: (data: MsgCreateDenom): EncodeObject => ({ typeUrl: "/QuyYeuCode.denom.denom.MsgCreateDenom", value: MsgCreateDenom.fromPartial( data ) }),
    msgMintAndSendTokens: (data: MsgMintAndSendTokens): EncodeObject => ({ typeUrl: "/QuyYeuCode.denom.denom.MsgMintAndSendTokens", value: MsgMintAndSendTokens.fromPartial( data ) }),
    msgUpdateOwner: (data: MsgUpdateOwner): EncodeObject => ({ typeUrl: "/QuyYeuCode.denom.denom.MsgUpdateOwner", value: MsgUpdateOwner.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
