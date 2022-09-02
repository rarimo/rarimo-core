// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgAddChain } from "./types/tokenmanager/tx";
import { MsgCreateInfo } from "./types/tokenmanager/tx";
import { MsgUpdateInfo } from "./types/tokenmanager/tx";
import { MsgDeleteInfo } from "./types/tokenmanager/tx";


const types = [
  ["/rarifyprotocol.rarimocore.tokenmanager.MsgAddChain", MsgAddChain],
  ["/rarifyprotocol.rarimocore.tokenmanager.MsgCreateInfo", MsgCreateInfo],
  ["/rarifyprotocol.rarimocore.tokenmanager.MsgUpdateInfo", MsgUpdateInfo],
  ["/rarifyprotocol.rarimocore.tokenmanager.MsgDeleteInfo", MsgDeleteInfo],
  
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
    msgAddChain: (data: MsgAddChain): EncodeObject => ({ typeUrl: "/rarifyprotocol.rarimocore.tokenmanager.MsgAddChain", value: MsgAddChain.fromPartial( data ) }),
    msgCreateInfo: (data: MsgCreateInfo): EncodeObject => ({ typeUrl: "/rarifyprotocol.rarimocore.tokenmanager.MsgCreateInfo", value: MsgCreateInfo.fromPartial( data ) }),
    msgUpdateInfo: (data: MsgUpdateInfo): EncodeObject => ({ typeUrl: "/rarifyprotocol.rarimocore.tokenmanager.MsgUpdateInfo", value: MsgUpdateInfo.fromPartial( data ) }),
    msgDeleteInfo: (data: MsgDeleteInfo): EncodeObject => ({ typeUrl: "/rarifyprotocol.rarimocore.tokenmanager.MsgDeleteInfo", value: MsgDeleteInfo.fromPartial( data ) }),
    
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
