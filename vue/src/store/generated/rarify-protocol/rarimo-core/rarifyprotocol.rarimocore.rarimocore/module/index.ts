// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateChangeKeyECDSA } from "./types/rarimocore/tx";
import { MsgCreateDeposit } from "./types/rarimocore/tx";
import { MsgCreateConfirmation } from "./types/rarimocore/tx";


const types = [
  ["/rarifyprotocol.rarimocore.rarimocore.MsgCreateChangeKeyECDSA", MsgCreateChangeKeyECDSA],
  ["/rarifyprotocol.rarimocore.rarimocore.MsgCreateDeposit", MsgCreateDeposit],
  ["/rarifyprotocol.rarimocore.rarimocore.MsgCreateConfirmation", MsgCreateConfirmation],
  
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
    msgCreateChangeKeyECDSA: (data: MsgCreateChangeKeyECDSA): EncodeObject => ({ typeUrl: "/rarifyprotocol.rarimocore.rarimocore.MsgCreateChangeKeyECDSA", value: MsgCreateChangeKeyECDSA.fromPartial( data ) }),
    msgCreateDeposit: (data: MsgCreateDeposit): EncodeObject => ({ typeUrl: "/rarifyprotocol.rarimocore.rarimocore.MsgCreateDeposit", value: MsgCreateDeposit.fromPartial( data ) }),
    msgCreateConfirmation: (data: MsgCreateConfirmation): EncodeObject => ({ typeUrl: "/rarifyprotocol.rarimocore.rarimocore.MsgCreateConfirmation", value: MsgCreateConfirmation.fromPartial( data ) }),
    
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
