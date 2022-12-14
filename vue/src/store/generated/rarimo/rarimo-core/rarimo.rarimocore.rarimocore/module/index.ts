// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateChangePartiesOp } from "./types/rarimocore/tx";
import { MsgCreateConfirmation } from "./types/rarimocore/tx";
import { MsgCreateTransferOp } from "./types/rarimocore/tx";
import { MsgSetupInitial } from "./types/rarimocore/tx";


const types = [
  ["/rarimo.rarimocore.rarimocore.MsgCreateChangePartiesOp", MsgCreateChangePartiesOp],
  ["/rarimo.rarimocore.rarimocore.MsgCreateConfirmation", MsgCreateConfirmation],
  ["/rarimo.rarimocore.rarimocore.MsgCreateTransferOp", MsgCreateTransferOp],
  ["/rarimo.rarimocore.rarimocore.MsgSetupInitial", MsgSetupInitial],
  
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
    msgCreateChangePartiesOp: (data: MsgCreateChangePartiesOp): EncodeObject => ({ typeUrl: "/rarimo.rarimocore.rarimocore.MsgCreateChangePartiesOp", value: MsgCreateChangePartiesOp.fromPartial( data ) }),
    msgCreateConfirmation: (data: MsgCreateConfirmation): EncodeObject => ({ typeUrl: "/rarimo.rarimocore.rarimocore.MsgCreateConfirmation", value: MsgCreateConfirmation.fromPartial( data ) }),
    msgCreateTransferOp: (data: MsgCreateTransferOp): EncodeObject => ({ typeUrl: "/rarimo.rarimocore.rarimocore.MsgCreateTransferOp", value: MsgCreateTransferOp.fromPartial( data ) }),
    msgSetupInitial: (data: MsgSetupInitial): EncodeObject => ({ typeUrl: "/rarimo.rarimocore.rarimocore.MsgSetupInitial", value: MsgSetupInitial.fromPartial( data ) }),
    
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
