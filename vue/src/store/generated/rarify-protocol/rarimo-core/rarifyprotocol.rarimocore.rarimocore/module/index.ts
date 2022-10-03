// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateConfirmation } from "./types/rarimocore/tx";
import { MsgCreateTransferOp } from "./types/rarimocore/tx";
import { MsgCreateChangeKeyECDSA } from "./types/rarimocore/tx";


const types = [
  ["/rarifyprotocol.rarimocore.rarimocore.MsgCreateConfirmation", MsgCreateConfirmation],
  ["/rarifyprotocol.rarimocore.rarimocore.MsgCreateTransferOp", MsgCreateTransferOp],
  ["/rarifyprotocol.rarimocore.rarimocore.MsgCreateChangeKeyECDSA", MsgCreateChangeKeyECDSA],
  
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
    msgCreateConfirmation: (data: MsgCreateConfirmation): EncodeObject => ({ typeUrl: "/rarifyprotocol.rarimocore.rarimocore.MsgCreateConfirmation", value: MsgCreateConfirmation.fromPartial( data ) }),
    msgCreateTransferOp: (data: MsgCreateTransferOp): EncodeObject => ({ typeUrl: "/rarifyprotocol.rarimocore.rarimocore.MsgCreateTransferOp", value: MsgCreateTransferOp.fromPartial( data ) }),
    msgCreateChangeKeyECDSA: (data: MsgCreateChangeKeyECDSA): EncodeObject => ({ typeUrl: "/rarifyprotocol.rarimocore.rarimocore.MsgCreateChangeKeyECDSA", value: MsgCreateChangeKeyECDSA.fromPartial( data ) }),
    
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
