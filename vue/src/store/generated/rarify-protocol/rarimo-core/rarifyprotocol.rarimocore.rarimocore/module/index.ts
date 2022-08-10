// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgDeleteChangeKeyECDSA } from "./types/rarimocore/tx";
import { MsgCreateDeposit } from "./types/rarimocore/tx";
import { MsgCreateChangeKeyECDSA } from "./types/rarimocore/tx";
import { MsgUpdateDeposit } from "./types/rarimocore/tx";
import { MsgUpdateChangeKeyECDSA } from "./types/rarimocore/tx";
import { MsgDeleteConfirmation } from "./types/rarimocore/tx";
import { MsgCreateConfirmation } from "./types/rarimocore/tx";
import { MsgDeleteDeposit } from "./types/rarimocore/tx";
import { MsgUpdateConfirmation } from "./types/rarimocore/tx";


const types = [
  ["/rarifyprotocol.rarimocore.rarimocore.MsgDeleteChangeKeyECDSA", MsgDeleteChangeKeyECDSA],
  ["/rarifyprotocol.rarimocore.rarimocore.MsgCreateDeposit", MsgCreateDeposit],
  ["/rarifyprotocol.rarimocore.rarimocore.MsgCreateChangeKeyECDSA", MsgCreateChangeKeyECDSA],
  ["/rarifyprotocol.rarimocore.rarimocore.MsgUpdateDeposit", MsgUpdateDeposit],
  ["/rarifyprotocol.rarimocore.rarimocore.MsgUpdateChangeKeyECDSA", MsgUpdateChangeKeyECDSA],
  ["/rarifyprotocol.rarimocore.rarimocore.MsgDeleteConfirmation", MsgDeleteConfirmation],
  ["/rarifyprotocol.rarimocore.rarimocore.MsgCreateConfirmation", MsgCreateConfirmation],
  ["/rarifyprotocol.rarimocore.rarimocore.MsgDeleteDeposit", MsgDeleteDeposit],
  ["/rarifyprotocol.rarimocore.rarimocore.MsgUpdateConfirmation", MsgUpdateConfirmation],
  
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
    msgDeleteChangeKeyECDSA: (data: MsgDeleteChangeKeyECDSA): EncodeObject => ({ typeUrl: "/rarifyprotocol.rarimocore.rarimocore.MsgDeleteChangeKeyECDSA", value: MsgDeleteChangeKeyECDSA.fromPartial( data ) }),
    msgCreateDeposit: (data: MsgCreateDeposit): EncodeObject => ({ typeUrl: "/rarifyprotocol.rarimocore.rarimocore.MsgCreateDeposit", value: MsgCreateDeposit.fromPartial( data ) }),
    msgCreateChangeKeyECDSA: (data: MsgCreateChangeKeyECDSA): EncodeObject => ({ typeUrl: "/rarifyprotocol.rarimocore.rarimocore.MsgCreateChangeKeyECDSA", value: MsgCreateChangeKeyECDSA.fromPartial( data ) }),
    msgUpdateDeposit: (data: MsgUpdateDeposit): EncodeObject => ({ typeUrl: "/rarifyprotocol.rarimocore.rarimocore.MsgUpdateDeposit", value: MsgUpdateDeposit.fromPartial( data ) }),
    msgUpdateChangeKeyECDSA: (data: MsgUpdateChangeKeyECDSA): EncodeObject => ({ typeUrl: "/rarifyprotocol.rarimocore.rarimocore.MsgUpdateChangeKeyECDSA", value: MsgUpdateChangeKeyECDSA.fromPartial( data ) }),
    msgDeleteConfirmation: (data: MsgDeleteConfirmation): EncodeObject => ({ typeUrl: "/rarifyprotocol.rarimocore.rarimocore.MsgDeleteConfirmation", value: MsgDeleteConfirmation.fromPartial( data ) }),
    msgCreateConfirmation: (data: MsgCreateConfirmation): EncodeObject => ({ typeUrl: "/rarifyprotocol.rarimocore.rarimocore.MsgCreateConfirmation", value: MsgCreateConfirmation.fromPartial( data ) }),
    msgDeleteDeposit: (data: MsgDeleteDeposit): EncodeObject => ({ typeUrl: "/rarifyprotocol.rarimocore.rarimocore.MsgDeleteDeposit", value: MsgDeleteDeposit.fromPartial( data ) }),
    msgUpdateConfirmation: (data: MsgUpdateConfirmation): EncodeObject => ({ typeUrl: "/rarifyprotocol.rarimocore.rarimocore.MsgUpdateConfirmation", value: MsgUpdateConfirmation.fromPartial( data ) }),
    
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
