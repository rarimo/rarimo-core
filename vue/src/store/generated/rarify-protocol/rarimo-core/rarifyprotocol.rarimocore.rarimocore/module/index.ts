// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgUpdateDeposit } from "./types/rarimocore/tx";
import { MsgDeleteDeposit } from "./types/rarimocore/tx";
import { MsgCreateDeposit } from "./types/rarimocore/tx";


const types = [
  ["/rarifyprotocol.rarimocore.rarimocore.MsgUpdateDeposit", MsgUpdateDeposit],
  ["/rarifyprotocol.rarimocore.rarimocore.MsgDeleteDeposit", MsgDeleteDeposit],
  ["/rarifyprotocol.rarimocore.rarimocore.MsgCreateDeposit", MsgCreateDeposit],
  
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
    msgUpdateDeposit: (data: MsgUpdateDeposit): EncodeObject => ({ typeUrl: "/rarifyprotocol.rarimocore.rarimocore.MsgUpdateDeposit", value: MsgUpdateDeposit.fromPartial( data ) }),
    msgDeleteDeposit: (data: MsgDeleteDeposit): EncodeObject => ({ typeUrl: "/rarifyprotocol.rarimocore.rarimocore.MsgDeleteDeposit", value: MsgDeleteDeposit.fromPartial( data ) }),
    msgCreateDeposit: (data: MsgCreateDeposit): EncodeObject => ({ typeUrl: "/rarifyprotocol.rarimocore.rarimocore.MsgCreateDeposit", value: MsgCreateDeposit.fromPartial( data ) }),
    
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
