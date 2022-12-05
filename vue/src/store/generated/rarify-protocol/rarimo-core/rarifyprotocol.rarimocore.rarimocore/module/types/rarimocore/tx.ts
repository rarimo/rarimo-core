/* eslint-disable */
import { type, typeFromJSON, typeToJSON } from "../tokenmanager/item";
import { Reader, Writer } from "protobufjs/minimal";
import { Party } from "../rarimocore/params";
import { Coin } from "../cosmos/base/v1beta1/coin";

export const protobufPackage = "rarifyprotocol.rarimocore.rarimocore";

export interface MsgCreateTransferOp {
  creator: string;
  tx: string;
  eventId: string;
  fromChain: string;
  tokenType: type;
}

export interface MsgCreateTransferOpResponse {}

export interface MsgCreateConfirmation {
  creator: string;
  /** hex-encoded */
  root: string;
  indexes: string[];
  /** hex-encoded */
  signatureECDSA: string;
}

export interface MsgCreateConfirmationResponse {}

export interface MsgCreateChangePartiesOp {
  creator: string;
  newSet: Party[];
  newPublicKey: string;
  signature: string;
}

export interface MsgCreateChangePartiesOpResponse {}

export interface MsgSetupInitial {
  creator: string;
  partyPublicKey: string;
  newPublicKey: string;
}

export interface MsgSetupInitialResponse {}

export interface MsgCreateAddSignerPartyProposal {
  creator: string;
  title: string;
  description: string;
  account: string;
  address: string;
  trialPublicKey: string;
  deposit: Coin[];
}

export interface MsgCreateAddSignerPartyProposalResponse {}

const baseMsgCreateTransferOp: object = {
  creator: "",
  tx: "",
  eventId: "",
  fromChain: "",
  tokenType: 0,
};

export const MsgCreateTransferOp = {
  encode(
    message: MsgCreateTransferOp,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.tx !== "") {
      writer.uint32(18).string(message.tx);
    }
    if (message.eventId !== "") {
      writer.uint32(26).string(message.eventId);
    }
    if (message.fromChain !== "") {
      writer.uint32(34).string(message.fromChain);
    }
    if (message.tokenType !== 0) {
      writer.uint32(40).int32(message.tokenType);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateTransferOp {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateTransferOp } as MsgCreateTransferOp;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.tx = reader.string();
          break;
        case 3:
          message.eventId = reader.string();
          break;
        case 4:
          message.fromChain = reader.string();
          break;
        case 5:
          message.tokenType = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateTransferOp {
    const message = { ...baseMsgCreateTransferOp } as MsgCreateTransferOp;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.tx !== undefined && object.tx !== null) {
      message.tx = String(object.tx);
    } else {
      message.tx = "";
    }
    if (object.eventId !== undefined && object.eventId !== null) {
      message.eventId = String(object.eventId);
    } else {
      message.eventId = "";
    }
    if (object.fromChain !== undefined && object.fromChain !== null) {
      message.fromChain = String(object.fromChain);
    } else {
      message.fromChain = "";
    }
    if (object.tokenType !== undefined && object.tokenType !== null) {
      message.tokenType = typeFromJSON(object.tokenType);
    } else {
      message.tokenType = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateTransferOp): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.tx !== undefined && (obj.tx = message.tx);
    message.eventId !== undefined && (obj.eventId = message.eventId);
    message.fromChain !== undefined && (obj.fromChain = message.fromChain);
    message.tokenType !== undefined &&
      (obj.tokenType = typeToJSON(message.tokenType));
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateTransferOp>): MsgCreateTransferOp {
    const message = { ...baseMsgCreateTransferOp } as MsgCreateTransferOp;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.tx !== undefined && object.tx !== null) {
      message.tx = object.tx;
    } else {
      message.tx = "";
    }
    if (object.eventId !== undefined && object.eventId !== null) {
      message.eventId = object.eventId;
    } else {
      message.eventId = "";
    }
    if (object.fromChain !== undefined && object.fromChain !== null) {
      message.fromChain = object.fromChain;
    } else {
      message.fromChain = "";
    }
    if (object.tokenType !== undefined && object.tokenType !== null) {
      message.tokenType = object.tokenType;
    } else {
      message.tokenType = 0;
    }
    return message;
  },
};

const baseMsgCreateTransferOpResponse: object = {};

export const MsgCreateTransferOpResponse = {
  encode(
    _: MsgCreateTransferOpResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateTransferOpResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateTransferOpResponse,
    } as MsgCreateTransferOpResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateTransferOpResponse {
    const message = {
      ...baseMsgCreateTransferOpResponse,
    } as MsgCreateTransferOpResponse;
    return message;
  },

  toJSON(_: MsgCreateTransferOpResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateTransferOpResponse>
  ): MsgCreateTransferOpResponse {
    const message = {
      ...baseMsgCreateTransferOpResponse,
    } as MsgCreateTransferOpResponse;
    return message;
  },
};

const baseMsgCreateConfirmation: object = {
  creator: "",
  root: "",
  indexes: "",
  signatureECDSA: "",
};

export const MsgCreateConfirmation = {
  encode(
    message: MsgCreateConfirmation,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.root !== "") {
      writer.uint32(18).string(message.root);
    }
    for (const v of message.indexes) {
      writer.uint32(26).string(v!);
    }
    if (message.signatureECDSA !== "") {
      writer.uint32(34).string(message.signatureECDSA);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateConfirmation {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateConfirmation } as MsgCreateConfirmation;
    message.indexes = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.root = reader.string();
          break;
        case 3:
          message.indexes.push(reader.string());
          break;
        case 4:
          message.signatureECDSA = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateConfirmation {
    const message = { ...baseMsgCreateConfirmation } as MsgCreateConfirmation;
    message.indexes = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.root !== undefined && object.root !== null) {
      message.root = String(object.root);
    } else {
      message.root = "";
    }
    if (object.indexes !== undefined && object.indexes !== null) {
      for (const e of object.indexes) {
        message.indexes.push(String(e));
      }
    }
    if (object.signatureECDSA !== undefined && object.signatureECDSA !== null) {
      message.signatureECDSA = String(object.signatureECDSA);
    } else {
      message.signatureECDSA = "";
    }
    return message;
  },

  toJSON(message: MsgCreateConfirmation): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.root !== undefined && (obj.root = message.root);
    if (message.indexes) {
      obj.indexes = message.indexes.map((e) => e);
    } else {
      obj.indexes = [];
    }
    message.signatureECDSA !== undefined &&
      (obj.signatureECDSA = message.signatureECDSA);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateConfirmation>
  ): MsgCreateConfirmation {
    const message = { ...baseMsgCreateConfirmation } as MsgCreateConfirmation;
    message.indexes = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.root !== undefined && object.root !== null) {
      message.root = object.root;
    } else {
      message.root = "";
    }
    if (object.indexes !== undefined && object.indexes !== null) {
      for (const e of object.indexes) {
        message.indexes.push(e);
      }
    }
    if (object.signatureECDSA !== undefined && object.signatureECDSA !== null) {
      message.signatureECDSA = object.signatureECDSA;
    } else {
      message.signatureECDSA = "";
    }
    return message;
  },
};

const baseMsgCreateConfirmationResponse: object = {};

export const MsgCreateConfirmationResponse = {
  encode(
    _: MsgCreateConfirmationResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateConfirmationResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateConfirmationResponse,
    } as MsgCreateConfirmationResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateConfirmationResponse {
    const message = {
      ...baseMsgCreateConfirmationResponse,
    } as MsgCreateConfirmationResponse;
    return message;
  },

  toJSON(_: MsgCreateConfirmationResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateConfirmationResponse>
  ): MsgCreateConfirmationResponse {
    const message = {
      ...baseMsgCreateConfirmationResponse,
    } as MsgCreateConfirmationResponse;
    return message;
  },
};

const baseMsgCreateChangePartiesOp: object = {
  creator: "",
  newPublicKey: "",
  signature: "",
};

export const MsgCreateChangePartiesOp = {
  encode(
    message: MsgCreateChangePartiesOp,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    for (const v of message.newSet) {
      Party.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.newPublicKey !== "") {
      writer.uint32(26).string(message.newPublicKey);
    }
    if (message.signature !== "") {
      writer.uint32(34).string(message.signature);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateChangePartiesOp {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateChangePartiesOp,
    } as MsgCreateChangePartiesOp;
    message.newSet = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.newSet.push(Party.decode(reader, reader.uint32()));
          break;
        case 3:
          message.newPublicKey = reader.string();
          break;
        case 4:
          message.signature = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateChangePartiesOp {
    const message = {
      ...baseMsgCreateChangePartiesOp,
    } as MsgCreateChangePartiesOp;
    message.newSet = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.newSet !== undefined && object.newSet !== null) {
      for (const e of object.newSet) {
        message.newSet.push(Party.fromJSON(e));
      }
    }
    if (object.newPublicKey !== undefined && object.newPublicKey !== null) {
      message.newPublicKey = String(object.newPublicKey);
    } else {
      message.newPublicKey = "";
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = String(object.signature);
    } else {
      message.signature = "";
    }
    return message;
  },

  toJSON(message: MsgCreateChangePartiesOp): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    if (message.newSet) {
      obj.newSet = message.newSet.map((e) => (e ? Party.toJSON(e) : undefined));
    } else {
      obj.newSet = [];
    }
    message.newPublicKey !== undefined &&
      (obj.newPublicKey = message.newPublicKey);
    message.signature !== undefined && (obj.signature = message.signature);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateChangePartiesOp>
  ): MsgCreateChangePartiesOp {
    const message = {
      ...baseMsgCreateChangePartiesOp,
    } as MsgCreateChangePartiesOp;
    message.newSet = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.newSet !== undefined && object.newSet !== null) {
      for (const e of object.newSet) {
        message.newSet.push(Party.fromPartial(e));
      }
    }
    if (object.newPublicKey !== undefined && object.newPublicKey !== null) {
      message.newPublicKey = object.newPublicKey;
    } else {
      message.newPublicKey = "";
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = object.signature;
    } else {
      message.signature = "";
    }
    return message;
  },
};

const baseMsgCreateChangePartiesOpResponse: object = {};

export const MsgCreateChangePartiesOpResponse = {
  encode(
    _: MsgCreateChangePartiesOpResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateChangePartiesOpResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateChangePartiesOpResponse,
    } as MsgCreateChangePartiesOpResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateChangePartiesOpResponse {
    const message = {
      ...baseMsgCreateChangePartiesOpResponse,
    } as MsgCreateChangePartiesOpResponse;
    return message;
  },

  toJSON(_: MsgCreateChangePartiesOpResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateChangePartiesOpResponse>
  ): MsgCreateChangePartiesOpResponse {
    const message = {
      ...baseMsgCreateChangePartiesOpResponse,
    } as MsgCreateChangePartiesOpResponse;
    return message;
  },
};

const baseMsgSetupInitial: object = {
  creator: "",
  partyPublicKey: "",
  newPublicKey: "",
};

export const MsgSetupInitial = {
  encode(message: MsgSetupInitial, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.partyPublicKey !== "") {
      writer.uint32(18).string(message.partyPublicKey);
    }
    if (message.newPublicKey !== "") {
      writer.uint32(26).string(message.newPublicKey);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetupInitial {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSetupInitial } as MsgSetupInitial;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.partyPublicKey = reader.string();
          break;
        case 3:
          message.newPublicKey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSetupInitial {
    const message = { ...baseMsgSetupInitial } as MsgSetupInitial;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.partyPublicKey !== undefined && object.partyPublicKey !== null) {
      message.partyPublicKey = String(object.partyPublicKey);
    } else {
      message.partyPublicKey = "";
    }
    if (object.newPublicKey !== undefined && object.newPublicKey !== null) {
      message.newPublicKey = String(object.newPublicKey);
    } else {
      message.newPublicKey = "";
    }
    return message;
  },

  toJSON(message: MsgSetupInitial): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.partyPublicKey !== undefined &&
      (obj.partyPublicKey = message.partyPublicKey);
    message.newPublicKey !== undefined &&
      (obj.newPublicKey = message.newPublicKey);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSetupInitial>): MsgSetupInitial {
    const message = { ...baseMsgSetupInitial } as MsgSetupInitial;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.partyPublicKey !== undefined && object.partyPublicKey !== null) {
      message.partyPublicKey = object.partyPublicKey;
    } else {
      message.partyPublicKey = "";
    }
    if (object.newPublicKey !== undefined && object.newPublicKey !== null) {
      message.newPublicKey = object.newPublicKey;
    } else {
      message.newPublicKey = "";
    }
    return message;
  },
};

const baseMsgSetupInitialResponse: object = {};

export const MsgSetupInitialResponse = {
  encode(_: MsgSetupInitialResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetupInitialResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSetupInitialResponse,
    } as MsgSetupInitialResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgSetupInitialResponse {
    const message = {
      ...baseMsgSetupInitialResponse,
    } as MsgSetupInitialResponse;
    return message;
  },

  toJSON(_: MsgSetupInitialResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgSetupInitialResponse>
  ): MsgSetupInitialResponse {
    const message = {
      ...baseMsgSetupInitialResponse,
    } as MsgSetupInitialResponse;
    return message;
  },
};

const baseMsgCreateAddSignerPartyProposal: object = {
  creator: "",
  title: "",
  description: "",
  account: "",
  address: "",
  trialPublicKey: "",
};

export const MsgCreateAddSignerPartyProposal = {
  encode(
    message: MsgCreateAddSignerPartyProposal,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.title !== "") {
      writer.uint32(18).string(message.title);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    if (message.account !== "") {
      writer.uint32(34).string(message.account);
    }
    if (message.address !== "") {
      writer.uint32(42).string(message.address);
    }
    if (message.trialPublicKey !== "") {
      writer.uint32(50).string(message.trialPublicKey);
    }
    for (const v of message.deposit) {
      Coin.encode(v!, writer.uint32(58).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateAddSignerPartyProposal {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateAddSignerPartyProposal,
    } as MsgCreateAddSignerPartyProposal;
    message.deposit = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.title = reader.string();
          break;
        case 3:
          message.description = reader.string();
          break;
        case 4:
          message.account = reader.string();
          break;
        case 5:
          message.address = reader.string();
          break;
        case 6:
          message.trialPublicKey = reader.string();
          break;
        case 7:
          message.deposit.push(Coin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateAddSignerPartyProposal {
    const message = {
      ...baseMsgCreateAddSignerPartyProposal,
    } as MsgCreateAddSignerPartyProposal;
    message.deposit = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = String(object.title);
    } else {
      message.title = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.account !== undefined && object.account !== null) {
      message.account = String(object.account);
    } else {
      message.account = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.trialPublicKey !== undefined && object.trialPublicKey !== null) {
      message.trialPublicKey = String(object.trialPublicKey);
    } else {
      message.trialPublicKey = "";
    }
    if (object.deposit !== undefined && object.deposit !== null) {
      for (const e of object.deposit) {
        message.deposit.push(Coin.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: MsgCreateAddSignerPartyProposal): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.title !== undefined && (obj.title = message.title);
    message.description !== undefined &&
      (obj.description = message.description);
    message.account !== undefined && (obj.account = message.account);
    message.address !== undefined && (obj.address = message.address);
    message.trialPublicKey !== undefined &&
      (obj.trialPublicKey = message.trialPublicKey);
    if (message.deposit) {
      obj.deposit = message.deposit.map((e) =>
        e ? Coin.toJSON(e) : undefined
      );
    } else {
      obj.deposit = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateAddSignerPartyProposal>
  ): MsgCreateAddSignerPartyProposal {
    const message = {
      ...baseMsgCreateAddSignerPartyProposal,
    } as MsgCreateAddSignerPartyProposal;
    message.deposit = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = object.title;
    } else {
      message.title = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.account !== undefined && object.account !== null) {
      message.account = object.account;
    } else {
      message.account = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.trialPublicKey !== undefined && object.trialPublicKey !== null) {
      message.trialPublicKey = object.trialPublicKey;
    } else {
      message.trialPublicKey = "";
    }
    if (object.deposit !== undefined && object.deposit !== null) {
      for (const e of object.deposit) {
        message.deposit.push(Coin.fromPartial(e));
      }
    }
    return message;
  },
};

const baseMsgCreateAddSignerPartyProposalResponse: object = {};

export const MsgCreateAddSignerPartyProposalResponse = {
  encode(
    _: MsgCreateAddSignerPartyProposalResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateAddSignerPartyProposalResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateAddSignerPartyProposalResponse,
    } as MsgCreateAddSignerPartyProposalResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateAddSignerPartyProposalResponse {
    const message = {
      ...baseMsgCreateAddSignerPartyProposalResponse,
    } as MsgCreateAddSignerPartyProposalResponse;
    return message;
  },

  toJSON(_: MsgCreateAddSignerPartyProposalResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateAddSignerPartyProposalResponse>
  ): MsgCreateAddSignerPartyProposalResponse {
    const message = {
      ...baseMsgCreateAddSignerPartyProposalResponse,
    } as MsgCreateAddSignerPartyProposalResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateTransferOperation(
    request: MsgCreateTransferOp
  ): Promise<MsgCreateTransferOpResponse>;
  CreateChangePartiesOperation(
    request: MsgCreateChangePartiesOp
  ): Promise<MsgCreateChangePartiesOpResponse>;
  CreateConfirmation(
    request: MsgCreateConfirmation
  ): Promise<MsgCreateConfirmationResponse>;
  SetupInitial(request: MsgSetupInitial): Promise<MsgSetupInitialResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateAddSignerPartyProposal(
    request: MsgCreateAddSignerPartyProposal
  ): Promise<MsgCreateAddSignerPartyProposalResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateTransferOperation(
    request: MsgCreateTransferOp
  ): Promise<MsgCreateTransferOpResponse> {
    const data = MsgCreateTransferOp.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.rarimocore.Msg",
      "CreateTransferOperation",
      data
    );
    return promise.then((data) =>
      MsgCreateTransferOpResponse.decode(new Reader(data))
    );
  }

  CreateChangePartiesOperation(
    request: MsgCreateChangePartiesOp
  ): Promise<MsgCreateChangePartiesOpResponse> {
    const data = MsgCreateChangePartiesOp.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.rarimocore.Msg",
      "CreateChangePartiesOperation",
      data
    );
    return promise.then((data) =>
      MsgCreateChangePartiesOpResponse.decode(new Reader(data))
    );
  }

  CreateConfirmation(
    request: MsgCreateConfirmation
  ): Promise<MsgCreateConfirmationResponse> {
    const data = MsgCreateConfirmation.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.rarimocore.Msg",
      "CreateConfirmation",
      data
    );
    return promise.then((data) =>
      MsgCreateConfirmationResponse.decode(new Reader(data))
    );
  }

  SetupInitial(request: MsgSetupInitial): Promise<MsgSetupInitialResponse> {
    const data = MsgSetupInitial.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.rarimocore.Msg",
      "SetupInitial",
      data
    );
    return promise.then((data) =>
      MsgSetupInitialResponse.decode(new Reader(data))
    );
  }

  CreateAddSignerPartyProposal(
    request: MsgCreateAddSignerPartyProposal
  ): Promise<MsgCreateAddSignerPartyProposalResponse> {
    const data = MsgCreateAddSignerPartyProposal.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.rarimocore.Msg",
      "CreateAddSignerPartyProposal",
      data
    );
    return promise.then((data) =>
      MsgCreateAddSignerPartyProposalResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
