/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "rarifyprotocol.rarimocore.rarimocore";

export interface MsgCreateDeposit {
  creator: string;
  tx: string;
  fromChain: string;
  toChain: string;
  receiver: string;
  tokenAddress: string;
  tokenId: string;
}

export interface MsgCreateDepositResponse {}

export interface MsgUpdateDeposit {
  creator: string;
  tx: string;
  fromChain: string;
  toChain: string;
  receiver: string;
  tokenAddress: string;
  tokenId: string;
}

export interface MsgUpdateDepositResponse {}

export interface MsgDeleteDeposit {
  creator: string;
  tx: string;
}

export interface MsgDeleteDepositResponse {}

export interface MsgCreateConfirmation {
  creator: string;
  height: string;
  root: string;
  hashes: string[];
  sigECDSA: string;
  sigEdDSA: string;
}

export interface MsgCreateConfirmationResponse {}

export interface MsgUpdateConfirmation {
  creator: string;
  height: string;
  root: string;
  hashes: string[];
  sigECDSA: string;
  sigEdDSA: string;
}

export interface MsgUpdateConfirmationResponse {}

export interface MsgDeleteConfirmation {
  creator: string;
  height: string;
}

export interface MsgDeleteConfirmationResponse {}

export interface MsgCreateChangeKeyECDSA {
  creator: string;
  newKey: string;
  signature: string;
}

export interface MsgCreateChangeKeyECDSAResponse {}

export interface MsgUpdateChangeKeyECDSA {
  creator: string;
  newKey: string;
  signature: string;
}

export interface MsgUpdateChangeKeyECDSAResponse {}

export interface MsgDeleteChangeKeyECDSA {
  creator: string;
  newKey: string;
}

export interface MsgDeleteChangeKeyECDSAResponse {}

export interface MsgCreateChangeKeyEdDSA {
  creator: string;
  newKey: string;
  signature: string;
}

export interface MsgCreateChangeKeyEdDSAResponse {}

export interface MsgUpdateChangeKeyEdDSA {
  creator: string;
  newKey: string;
  signature: string;
}

export interface MsgUpdateChangeKeyEdDSAResponse {}

export interface MsgDeleteChangeKeyEdDSA {
  creator: string;
  newKey: string;
}

export interface MsgDeleteChangeKeyEdDSAResponse {}

const baseMsgCreateDeposit: object = {
  creator: "",
  tx: "",
  fromChain: "",
  toChain: "",
  receiver: "",
  tokenAddress: "",
  tokenId: "",
};

export const MsgCreateDeposit = {
  encode(message: MsgCreateDeposit, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.tx !== "") {
      writer.uint32(18).string(message.tx);
    }
    if (message.fromChain !== "") {
      writer.uint32(26).string(message.fromChain);
    }
    if (message.toChain !== "") {
      writer.uint32(34).string(message.toChain);
    }
    if (message.receiver !== "") {
      writer.uint32(42).string(message.receiver);
    }
    if (message.tokenAddress !== "") {
      writer.uint32(50).string(message.tokenAddress);
    }
    if (message.tokenId !== "") {
      writer.uint32(58).string(message.tokenId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateDeposit {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateDeposit } as MsgCreateDeposit;
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
          message.fromChain = reader.string();
          break;
        case 4:
          message.toChain = reader.string();
          break;
        case 5:
          message.receiver = reader.string();
          break;
        case 6:
          message.tokenAddress = reader.string();
          break;
        case 7:
          message.tokenId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateDeposit {
    const message = { ...baseMsgCreateDeposit } as MsgCreateDeposit;
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
    if (object.fromChain !== undefined && object.fromChain !== null) {
      message.fromChain = String(object.fromChain);
    } else {
      message.fromChain = "";
    }
    if (object.toChain !== undefined && object.toChain !== null) {
      message.toChain = String(object.toChain);
    } else {
      message.toChain = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = String(object.receiver);
    } else {
      message.receiver = "";
    }
    if (object.tokenAddress !== undefined && object.tokenAddress !== null) {
      message.tokenAddress = String(object.tokenAddress);
    } else {
      message.tokenAddress = "";
    }
    if (object.tokenId !== undefined && object.tokenId !== null) {
      message.tokenId = String(object.tokenId);
    } else {
      message.tokenId = "";
    }
    return message;
  },

  toJSON(message: MsgCreateDeposit): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.tx !== undefined && (obj.tx = message.tx);
    message.fromChain !== undefined && (obj.fromChain = message.fromChain);
    message.toChain !== undefined && (obj.toChain = message.toChain);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.tokenAddress !== undefined &&
      (obj.tokenAddress = message.tokenAddress);
    message.tokenId !== undefined && (obj.tokenId = message.tokenId);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateDeposit>): MsgCreateDeposit {
    const message = { ...baseMsgCreateDeposit } as MsgCreateDeposit;
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
    if (object.fromChain !== undefined && object.fromChain !== null) {
      message.fromChain = object.fromChain;
    } else {
      message.fromChain = "";
    }
    if (object.toChain !== undefined && object.toChain !== null) {
      message.toChain = object.toChain;
    } else {
      message.toChain = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = object.receiver;
    } else {
      message.receiver = "";
    }
    if (object.tokenAddress !== undefined && object.tokenAddress !== null) {
      message.tokenAddress = object.tokenAddress;
    } else {
      message.tokenAddress = "";
    }
    if (object.tokenId !== undefined && object.tokenId !== null) {
      message.tokenId = object.tokenId;
    } else {
      message.tokenId = "";
    }
    return message;
  },
};

const baseMsgCreateDepositResponse: object = {};

export const MsgCreateDepositResponse = {
  encode(
    _: MsgCreateDepositResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateDepositResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateDepositResponse,
    } as MsgCreateDepositResponse;
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

  fromJSON(_: any): MsgCreateDepositResponse {
    const message = {
      ...baseMsgCreateDepositResponse,
    } as MsgCreateDepositResponse;
    return message;
  },

  toJSON(_: MsgCreateDepositResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateDepositResponse>
  ): MsgCreateDepositResponse {
    const message = {
      ...baseMsgCreateDepositResponse,
    } as MsgCreateDepositResponse;
    return message;
  },
};

const baseMsgUpdateDeposit: object = {
  creator: "",
  tx: "",
  fromChain: "",
  toChain: "",
  receiver: "",
  tokenAddress: "",
  tokenId: "",
};

export const MsgUpdateDeposit = {
  encode(message: MsgUpdateDeposit, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.tx !== "") {
      writer.uint32(18).string(message.tx);
    }
    if (message.fromChain !== "") {
      writer.uint32(26).string(message.fromChain);
    }
    if (message.toChain !== "") {
      writer.uint32(34).string(message.toChain);
    }
    if (message.receiver !== "") {
      writer.uint32(42).string(message.receiver);
    }
    if (message.tokenAddress !== "") {
      writer.uint32(50).string(message.tokenAddress);
    }
    if (message.tokenId !== "") {
      writer.uint32(58).string(message.tokenId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateDeposit {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateDeposit } as MsgUpdateDeposit;
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
          message.fromChain = reader.string();
          break;
        case 4:
          message.toChain = reader.string();
          break;
        case 5:
          message.receiver = reader.string();
          break;
        case 6:
          message.tokenAddress = reader.string();
          break;
        case 7:
          message.tokenId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateDeposit {
    const message = { ...baseMsgUpdateDeposit } as MsgUpdateDeposit;
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
    if (object.fromChain !== undefined && object.fromChain !== null) {
      message.fromChain = String(object.fromChain);
    } else {
      message.fromChain = "";
    }
    if (object.toChain !== undefined && object.toChain !== null) {
      message.toChain = String(object.toChain);
    } else {
      message.toChain = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = String(object.receiver);
    } else {
      message.receiver = "";
    }
    if (object.tokenAddress !== undefined && object.tokenAddress !== null) {
      message.tokenAddress = String(object.tokenAddress);
    } else {
      message.tokenAddress = "";
    }
    if (object.tokenId !== undefined && object.tokenId !== null) {
      message.tokenId = String(object.tokenId);
    } else {
      message.tokenId = "";
    }
    return message;
  },

  toJSON(message: MsgUpdateDeposit): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.tx !== undefined && (obj.tx = message.tx);
    message.fromChain !== undefined && (obj.fromChain = message.fromChain);
    message.toChain !== undefined && (obj.toChain = message.toChain);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.tokenAddress !== undefined &&
      (obj.tokenAddress = message.tokenAddress);
    message.tokenId !== undefined && (obj.tokenId = message.tokenId);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateDeposit>): MsgUpdateDeposit {
    const message = { ...baseMsgUpdateDeposit } as MsgUpdateDeposit;
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
    if (object.fromChain !== undefined && object.fromChain !== null) {
      message.fromChain = object.fromChain;
    } else {
      message.fromChain = "";
    }
    if (object.toChain !== undefined && object.toChain !== null) {
      message.toChain = object.toChain;
    } else {
      message.toChain = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = object.receiver;
    } else {
      message.receiver = "";
    }
    if (object.tokenAddress !== undefined && object.tokenAddress !== null) {
      message.tokenAddress = object.tokenAddress;
    } else {
      message.tokenAddress = "";
    }
    if (object.tokenId !== undefined && object.tokenId !== null) {
      message.tokenId = object.tokenId;
    } else {
      message.tokenId = "";
    }
    return message;
  },
};

const baseMsgUpdateDepositResponse: object = {};

export const MsgUpdateDepositResponse = {
  encode(
    _: MsgUpdateDepositResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateDepositResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateDepositResponse,
    } as MsgUpdateDepositResponse;
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

  fromJSON(_: any): MsgUpdateDepositResponse {
    const message = {
      ...baseMsgUpdateDepositResponse,
    } as MsgUpdateDepositResponse;
    return message;
  },

  toJSON(_: MsgUpdateDepositResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateDepositResponse>
  ): MsgUpdateDepositResponse {
    const message = {
      ...baseMsgUpdateDepositResponse,
    } as MsgUpdateDepositResponse;
    return message;
  },
};

const baseMsgDeleteDeposit: object = { creator: "", tx: "" };

export const MsgDeleteDeposit = {
  encode(message: MsgDeleteDeposit, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.tx !== "") {
      writer.uint32(18).string(message.tx);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteDeposit {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteDeposit } as MsgDeleteDeposit;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.tx = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteDeposit {
    const message = { ...baseMsgDeleteDeposit } as MsgDeleteDeposit;
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
    return message;
  },

  toJSON(message: MsgDeleteDeposit): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.tx !== undefined && (obj.tx = message.tx);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeleteDeposit>): MsgDeleteDeposit {
    const message = { ...baseMsgDeleteDeposit } as MsgDeleteDeposit;
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
    return message;
  },
};

const baseMsgDeleteDepositResponse: object = {};

export const MsgDeleteDepositResponse = {
  encode(
    _: MsgDeleteDepositResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteDepositResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteDepositResponse,
    } as MsgDeleteDepositResponse;
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

  fromJSON(_: any): MsgDeleteDepositResponse {
    const message = {
      ...baseMsgDeleteDepositResponse,
    } as MsgDeleteDepositResponse;
    return message;
  },

  toJSON(_: MsgDeleteDepositResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteDepositResponse>
  ): MsgDeleteDepositResponse {
    const message = {
      ...baseMsgDeleteDepositResponse,
    } as MsgDeleteDepositResponse;
    return message;
  },
};

const baseMsgCreateConfirmation: object = {
  creator: "",
  height: "",
  root: "",
  hashes: "",
  sigECDSA: "",
  sigEdDSA: "",
};

export const MsgCreateConfirmation = {
  encode(
    message: MsgCreateConfirmation,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.height !== "") {
      writer.uint32(18).string(message.height);
    }
    if (message.root !== "") {
      writer.uint32(26).string(message.root);
    }
    for (const v of message.hashes) {
      writer.uint32(34).string(v!);
    }
    if (message.sigECDSA !== "") {
      writer.uint32(42).string(message.sigECDSA);
    }
    if (message.sigEdDSA !== "") {
      writer.uint32(50).string(message.sigEdDSA);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateConfirmation {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateConfirmation } as MsgCreateConfirmation;
    message.hashes = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.height = reader.string();
          break;
        case 3:
          message.root = reader.string();
          break;
        case 4:
          message.hashes.push(reader.string());
          break;
        case 5:
          message.sigECDSA = reader.string();
          break;
        case 6:
          message.sigEdDSA = reader.string();
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
    message.hashes = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.height !== undefined && object.height !== null) {
      message.height = String(object.height);
    } else {
      message.height = "";
    }
    if (object.root !== undefined && object.root !== null) {
      message.root = String(object.root);
    } else {
      message.root = "";
    }
    if (object.hashes !== undefined && object.hashes !== null) {
      for (const e of object.hashes) {
        message.hashes.push(String(e));
      }
    }
    if (object.sigECDSA !== undefined && object.sigECDSA !== null) {
      message.sigECDSA = String(object.sigECDSA);
    } else {
      message.sigECDSA = "";
    }
    if (object.sigEdDSA !== undefined && object.sigEdDSA !== null) {
      message.sigEdDSA = String(object.sigEdDSA);
    } else {
      message.sigEdDSA = "";
    }
    return message;
  },

  toJSON(message: MsgCreateConfirmation): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.height !== undefined && (obj.height = message.height);
    message.root !== undefined && (obj.root = message.root);
    if (message.hashes) {
      obj.hashes = message.hashes.map((e) => e);
    } else {
      obj.hashes = [];
    }
    message.sigECDSA !== undefined && (obj.sigECDSA = message.sigECDSA);
    message.sigEdDSA !== undefined && (obj.sigEdDSA = message.sigEdDSA);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateConfirmation>
  ): MsgCreateConfirmation {
    const message = { ...baseMsgCreateConfirmation } as MsgCreateConfirmation;
    message.hashes = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.height !== undefined && object.height !== null) {
      message.height = object.height;
    } else {
      message.height = "";
    }
    if (object.root !== undefined && object.root !== null) {
      message.root = object.root;
    } else {
      message.root = "";
    }
    if (object.hashes !== undefined && object.hashes !== null) {
      for (const e of object.hashes) {
        message.hashes.push(e);
      }
    }
    if (object.sigECDSA !== undefined && object.sigECDSA !== null) {
      message.sigECDSA = object.sigECDSA;
    } else {
      message.sigECDSA = "";
    }
    if (object.sigEdDSA !== undefined && object.sigEdDSA !== null) {
      message.sigEdDSA = object.sigEdDSA;
    } else {
      message.sigEdDSA = "";
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

const baseMsgUpdateConfirmation: object = {
  creator: "",
  height: "",
  root: "",
  hashes: "",
  sigECDSA: "",
  sigEdDSA: "",
};

export const MsgUpdateConfirmation = {
  encode(
    message: MsgUpdateConfirmation,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.height !== "") {
      writer.uint32(18).string(message.height);
    }
    if (message.root !== "") {
      writer.uint32(26).string(message.root);
    }
    for (const v of message.hashes) {
      writer.uint32(34).string(v!);
    }
    if (message.sigECDSA !== "") {
      writer.uint32(42).string(message.sigECDSA);
    }
    if (message.sigEdDSA !== "") {
      writer.uint32(50).string(message.sigEdDSA);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateConfirmation {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateConfirmation } as MsgUpdateConfirmation;
    message.hashes = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.height = reader.string();
          break;
        case 3:
          message.root = reader.string();
          break;
        case 4:
          message.hashes.push(reader.string());
          break;
        case 5:
          message.sigECDSA = reader.string();
          break;
        case 6:
          message.sigEdDSA = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateConfirmation {
    const message = { ...baseMsgUpdateConfirmation } as MsgUpdateConfirmation;
    message.hashes = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.height !== undefined && object.height !== null) {
      message.height = String(object.height);
    } else {
      message.height = "";
    }
    if (object.root !== undefined && object.root !== null) {
      message.root = String(object.root);
    } else {
      message.root = "";
    }
    if (object.hashes !== undefined && object.hashes !== null) {
      for (const e of object.hashes) {
        message.hashes.push(String(e));
      }
    }
    if (object.sigECDSA !== undefined && object.sigECDSA !== null) {
      message.sigECDSA = String(object.sigECDSA);
    } else {
      message.sigECDSA = "";
    }
    if (object.sigEdDSA !== undefined && object.sigEdDSA !== null) {
      message.sigEdDSA = String(object.sigEdDSA);
    } else {
      message.sigEdDSA = "";
    }
    return message;
  },

  toJSON(message: MsgUpdateConfirmation): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.height !== undefined && (obj.height = message.height);
    message.root !== undefined && (obj.root = message.root);
    if (message.hashes) {
      obj.hashes = message.hashes.map((e) => e);
    } else {
      obj.hashes = [];
    }
    message.sigECDSA !== undefined && (obj.sigECDSA = message.sigECDSA);
    message.sigEdDSA !== undefined && (obj.sigEdDSA = message.sigEdDSA);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgUpdateConfirmation>
  ): MsgUpdateConfirmation {
    const message = { ...baseMsgUpdateConfirmation } as MsgUpdateConfirmation;
    message.hashes = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.height !== undefined && object.height !== null) {
      message.height = object.height;
    } else {
      message.height = "";
    }
    if (object.root !== undefined && object.root !== null) {
      message.root = object.root;
    } else {
      message.root = "";
    }
    if (object.hashes !== undefined && object.hashes !== null) {
      for (const e of object.hashes) {
        message.hashes.push(e);
      }
    }
    if (object.sigECDSA !== undefined && object.sigECDSA !== null) {
      message.sigECDSA = object.sigECDSA;
    } else {
      message.sigECDSA = "";
    }
    if (object.sigEdDSA !== undefined && object.sigEdDSA !== null) {
      message.sigEdDSA = object.sigEdDSA;
    } else {
      message.sigEdDSA = "";
    }
    return message;
  },
};

const baseMsgUpdateConfirmationResponse: object = {};

export const MsgUpdateConfirmationResponse = {
  encode(
    _: MsgUpdateConfirmationResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateConfirmationResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateConfirmationResponse,
    } as MsgUpdateConfirmationResponse;
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

  fromJSON(_: any): MsgUpdateConfirmationResponse {
    const message = {
      ...baseMsgUpdateConfirmationResponse,
    } as MsgUpdateConfirmationResponse;
    return message;
  },

  toJSON(_: MsgUpdateConfirmationResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateConfirmationResponse>
  ): MsgUpdateConfirmationResponse {
    const message = {
      ...baseMsgUpdateConfirmationResponse,
    } as MsgUpdateConfirmationResponse;
    return message;
  },
};

const baseMsgDeleteConfirmation: object = { creator: "", height: "" };

export const MsgDeleteConfirmation = {
  encode(
    message: MsgDeleteConfirmation,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.height !== "") {
      writer.uint32(18).string(message.height);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteConfirmation {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteConfirmation } as MsgDeleteConfirmation;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.height = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteConfirmation {
    const message = { ...baseMsgDeleteConfirmation } as MsgDeleteConfirmation;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.height !== undefined && object.height !== null) {
      message.height = String(object.height);
    } else {
      message.height = "";
    }
    return message;
  },

  toJSON(message: MsgDeleteConfirmation): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.height !== undefined && (obj.height = message.height);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgDeleteConfirmation>
  ): MsgDeleteConfirmation {
    const message = { ...baseMsgDeleteConfirmation } as MsgDeleteConfirmation;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.height !== undefined && object.height !== null) {
      message.height = object.height;
    } else {
      message.height = "";
    }
    return message;
  },
};

const baseMsgDeleteConfirmationResponse: object = {};

export const MsgDeleteConfirmationResponse = {
  encode(
    _: MsgDeleteConfirmationResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteConfirmationResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteConfirmationResponse,
    } as MsgDeleteConfirmationResponse;
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

  fromJSON(_: any): MsgDeleteConfirmationResponse {
    const message = {
      ...baseMsgDeleteConfirmationResponse,
    } as MsgDeleteConfirmationResponse;
    return message;
  },

  toJSON(_: MsgDeleteConfirmationResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteConfirmationResponse>
  ): MsgDeleteConfirmationResponse {
    const message = {
      ...baseMsgDeleteConfirmationResponse,
    } as MsgDeleteConfirmationResponse;
    return message;
  },
};

const baseMsgCreateChangeKeyECDSA: object = {
  creator: "",
  newKey: "",
  signature: "",
};

export const MsgCreateChangeKeyECDSA = {
  encode(
    message: MsgCreateChangeKeyECDSA,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.newKey !== "") {
      writer.uint32(18).string(message.newKey);
    }
    if (message.signature !== "") {
      writer.uint32(26).string(message.signature);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateChangeKeyECDSA {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateChangeKeyECDSA,
    } as MsgCreateChangeKeyECDSA;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.newKey = reader.string();
          break;
        case 3:
          message.signature = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateChangeKeyECDSA {
    const message = {
      ...baseMsgCreateChangeKeyECDSA,
    } as MsgCreateChangeKeyECDSA;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.newKey !== undefined && object.newKey !== null) {
      message.newKey = String(object.newKey);
    } else {
      message.newKey = "";
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = String(object.signature);
    } else {
      message.signature = "";
    }
    return message;
  },

  toJSON(message: MsgCreateChangeKeyECDSA): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.newKey !== undefined && (obj.newKey = message.newKey);
    message.signature !== undefined && (obj.signature = message.signature);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateChangeKeyECDSA>
  ): MsgCreateChangeKeyECDSA {
    const message = {
      ...baseMsgCreateChangeKeyECDSA,
    } as MsgCreateChangeKeyECDSA;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.newKey !== undefined && object.newKey !== null) {
      message.newKey = object.newKey;
    } else {
      message.newKey = "";
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = object.signature;
    } else {
      message.signature = "";
    }
    return message;
  },
};

const baseMsgCreateChangeKeyECDSAResponse: object = {};

export const MsgCreateChangeKeyECDSAResponse = {
  encode(
    _: MsgCreateChangeKeyECDSAResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateChangeKeyECDSAResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateChangeKeyECDSAResponse,
    } as MsgCreateChangeKeyECDSAResponse;
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

  fromJSON(_: any): MsgCreateChangeKeyECDSAResponse {
    const message = {
      ...baseMsgCreateChangeKeyECDSAResponse,
    } as MsgCreateChangeKeyECDSAResponse;
    return message;
  },

  toJSON(_: MsgCreateChangeKeyECDSAResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateChangeKeyECDSAResponse>
  ): MsgCreateChangeKeyECDSAResponse {
    const message = {
      ...baseMsgCreateChangeKeyECDSAResponse,
    } as MsgCreateChangeKeyECDSAResponse;
    return message;
  },
};

const baseMsgUpdateChangeKeyECDSA: object = {
  creator: "",
  newKey: "",
  signature: "",
};

export const MsgUpdateChangeKeyECDSA = {
  encode(
    message: MsgUpdateChangeKeyECDSA,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.newKey !== "") {
      writer.uint32(18).string(message.newKey);
    }
    if (message.signature !== "") {
      writer.uint32(26).string(message.signature);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateChangeKeyECDSA {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateChangeKeyECDSA,
    } as MsgUpdateChangeKeyECDSA;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.newKey = reader.string();
          break;
        case 3:
          message.signature = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateChangeKeyECDSA {
    const message = {
      ...baseMsgUpdateChangeKeyECDSA,
    } as MsgUpdateChangeKeyECDSA;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.newKey !== undefined && object.newKey !== null) {
      message.newKey = String(object.newKey);
    } else {
      message.newKey = "";
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = String(object.signature);
    } else {
      message.signature = "";
    }
    return message;
  },

  toJSON(message: MsgUpdateChangeKeyECDSA): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.newKey !== undefined && (obj.newKey = message.newKey);
    message.signature !== undefined && (obj.signature = message.signature);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgUpdateChangeKeyECDSA>
  ): MsgUpdateChangeKeyECDSA {
    const message = {
      ...baseMsgUpdateChangeKeyECDSA,
    } as MsgUpdateChangeKeyECDSA;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.newKey !== undefined && object.newKey !== null) {
      message.newKey = object.newKey;
    } else {
      message.newKey = "";
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = object.signature;
    } else {
      message.signature = "";
    }
    return message;
  },
};

const baseMsgUpdateChangeKeyECDSAResponse: object = {};

export const MsgUpdateChangeKeyECDSAResponse = {
  encode(
    _: MsgUpdateChangeKeyECDSAResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateChangeKeyECDSAResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateChangeKeyECDSAResponse,
    } as MsgUpdateChangeKeyECDSAResponse;
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

  fromJSON(_: any): MsgUpdateChangeKeyECDSAResponse {
    const message = {
      ...baseMsgUpdateChangeKeyECDSAResponse,
    } as MsgUpdateChangeKeyECDSAResponse;
    return message;
  },

  toJSON(_: MsgUpdateChangeKeyECDSAResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateChangeKeyECDSAResponse>
  ): MsgUpdateChangeKeyECDSAResponse {
    const message = {
      ...baseMsgUpdateChangeKeyECDSAResponse,
    } as MsgUpdateChangeKeyECDSAResponse;
    return message;
  },
};

const baseMsgDeleteChangeKeyECDSA: object = { creator: "", newKey: "" };

export const MsgDeleteChangeKeyECDSA = {
  encode(
    message: MsgDeleteChangeKeyECDSA,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.newKey !== "") {
      writer.uint32(18).string(message.newKey);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteChangeKeyECDSA {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteChangeKeyECDSA,
    } as MsgDeleteChangeKeyECDSA;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.newKey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteChangeKeyECDSA {
    const message = {
      ...baseMsgDeleteChangeKeyECDSA,
    } as MsgDeleteChangeKeyECDSA;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.newKey !== undefined && object.newKey !== null) {
      message.newKey = String(object.newKey);
    } else {
      message.newKey = "";
    }
    return message;
  },

  toJSON(message: MsgDeleteChangeKeyECDSA): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.newKey !== undefined && (obj.newKey = message.newKey);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgDeleteChangeKeyECDSA>
  ): MsgDeleteChangeKeyECDSA {
    const message = {
      ...baseMsgDeleteChangeKeyECDSA,
    } as MsgDeleteChangeKeyECDSA;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.newKey !== undefined && object.newKey !== null) {
      message.newKey = object.newKey;
    } else {
      message.newKey = "";
    }
    return message;
  },
};

const baseMsgDeleteChangeKeyECDSAResponse: object = {};

export const MsgDeleteChangeKeyECDSAResponse = {
  encode(
    _: MsgDeleteChangeKeyECDSAResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteChangeKeyECDSAResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteChangeKeyECDSAResponse,
    } as MsgDeleteChangeKeyECDSAResponse;
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

  fromJSON(_: any): MsgDeleteChangeKeyECDSAResponse {
    const message = {
      ...baseMsgDeleteChangeKeyECDSAResponse,
    } as MsgDeleteChangeKeyECDSAResponse;
    return message;
  },

  toJSON(_: MsgDeleteChangeKeyECDSAResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteChangeKeyECDSAResponse>
  ): MsgDeleteChangeKeyECDSAResponse {
    const message = {
      ...baseMsgDeleteChangeKeyECDSAResponse,
    } as MsgDeleteChangeKeyECDSAResponse;
    return message;
  },
};

const baseMsgCreateChangeKeyEdDSA: object = {
  creator: "",
  newKey: "",
  signature: "",
};

export const MsgCreateChangeKeyEdDSA = {
  encode(
    message: MsgCreateChangeKeyEdDSA,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.newKey !== "") {
      writer.uint32(18).string(message.newKey);
    }
    if (message.signature !== "") {
      writer.uint32(26).string(message.signature);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateChangeKeyEdDSA {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateChangeKeyEdDSA,
    } as MsgCreateChangeKeyEdDSA;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.newKey = reader.string();
          break;
        case 3:
          message.signature = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateChangeKeyEdDSA {
    const message = {
      ...baseMsgCreateChangeKeyEdDSA,
    } as MsgCreateChangeKeyEdDSA;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.newKey !== undefined && object.newKey !== null) {
      message.newKey = String(object.newKey);
    } else {
      message.newKey = "";
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = String(object.signature);
    } else {
      message.signature = "";
    }
    return message;
  },

  toJSON(message: MsgCreateChangeKeyEdDSA): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.newKey !== undefined && (obj.newKey = message.newKey);
    message.signature !== undefined && (obj.signature = message.signature);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateChangeKeyEdDSA>
  ): MsgCreateChangeKeyEdDSA {
    const message = {
      ...baseMsgCreateChangeKeyEdDSA,
    } as MsgCreateChangeKeyEdDSA;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.newKey !== undefined && object.newKey !== null) {
      message.newKey = object.newKey;
    } else {
      message.newKey = "";
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = object.signature;
    } else {
      message.signature = "";
    }
    return message;
  },
};

const baseMsgCreateChangeKeyEdDSAResponse: object = {};

export const MsgCreateChangeKeyEdDSAResponse = {
  encode(
    _: MsgCreateChangeKeyEdDSAResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateChangeKeyEdDSAResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateChangeKeyEdDSAResponse,
    } as MsgCreateChangeKeyEdDSAResponse;
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

  fromJSON(_: any): MsgCreateChangeKeyEdDSAResponse {
    const message = {
      ...baseMsgCreateChangeKeyEdDSAResponse,
    } as MsgCreateChangeKeyEdDSAResponse;
    return message;
  },

  toJSON(_: MsgCreateChangeKeyEdDSAResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateChangeKeyEdDSAResponse>
  ): MsgCreateChangeKeyEdDSAResponse {
    const message = {
      ...baseMsgCreateChangeKeyEdDSAResponse,
    } as MsgCreateChangeKeyEdDSAResponse;
    return message;
  },
};

const baseMsgUpdateChangeKeyEdDSA: object = {
  creator: "",
  newKey: "",
  signature: "",
};

export const MsgUpdateChangeKeyEdDSA = {
  encode(
    message: MsgUpdateChangeKeyEdDSA,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.newKey !== "") {
      writer.uint32(18).string(message.newKey);
    }
    if (message.signature !== "") {
      writer.uint32(26).string(message.signature);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateChangeKeyEdDSA {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateChangeKeyEdDSA,
    } as MsgUpdateChangeKeyEdDSA;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.newKey = reader.string();
          break;
        case 3:
          message.signature = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateChangeKeyEdDSA {
    const message = {
      ...baseMsgUpdateChangeKeyEdDSA,
    } as MsgUpdateChangeKeyEdDSA;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.newKey !== undefined && object.newKey !== null) {
      message.newKey = String(object.newKey);
    } else {
      message.newKey = "";
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = String(object.signature);
    } else {
      message.signature = "";
    }
    return message;
  },

  toJSON(message: MsgUpdateChangeKeyEdDSA): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.newKey !== undefined && (obj.newKey = message.newKey);
    message.signature !== undefined && (obj.signature = message.signature);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgUpdateChangeKeyEdDSA>
  ): MsgUpdateChangeKeyEdDSA {
    const message = {
      ...baseMsgUpdateChangeKeyEdDSA,
    } as MsgUpdateChangeKeyEdDSA;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.newKey !== undefined && object.newKey !== null) {
      message.newKey = object.newKey;
    } else {
      message.newKey = "";
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = object.signature;
    } else {
      message.signature = "";
    }
    return message;
  },
};

const baseMsgUpdateChangeKeyEdDSAResponse: object = {};

export const MsgUpdateChangeKeyEdDSAResponse = {
  encode(
    _: MsgUpdateChangeKeyEdDSAResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateChangeKeyEdDSAResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateChangeKeyEdDSAResponse,
    } as MsgUpdateChangeKeyEdDSAResponse;
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

  fromJSON(_: any): MsgUpdateChangeKeyEdDSAResponse {
    const message = {
      ...baseMsgUpdateChangeKeyEdDSAResponse,
    } as MsgUpdateChangeKeyEdDSAResponse;
    return message;
  },

  toJSON(_: MsgUpdateChangeKeyEdDSAResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateChangeKeyEdDSAResponse>
  ): MsgUpdateChangeKeyEdDSAResponse {
    const message = {
      ...baseMsgUpdateChangeKeyEdDSAResponse,
    } as MsgUpdateChangeKeyEdDSAResponse;
    return message;
  },
};

const baseMsgDeleteChangeKeyEdDSA: object = { creator: "", newKey: "" };

export const MsgDeleteChangeKeyEdDSA = {
  encode(
    message: MsgDeleteChangeKeyEdDSA,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.newKey !== "") {
      writer.uint32(18).string(message.newKey);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteChangeKeyEdDSA {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteChangeKeyEdDSA,
    } as MsgDeleteChangeKeyEdDSA;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.newKey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteChangeKeyEdDSA {
    const message = {
      ...baseMsgDeleteChangeKeyEdDSA,
    } as MsgDeleteChangeKeyEdDSA;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.newKey !== undefined && object.newKey !== null) {
      message.newKey = String(object.newKey);
    } else {
      message.newKey = "";
    }
    return message;
  },

  toJSON(message: MsgDeleteChangeKeyEdDSA): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.newKey !== undefined && (obj.newKey = message.newKey);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgDeleteChangeKeyEdDSA>
  ): MsgDeleteChangeKeyEdDSA {
    const message = {
      ...baseMsgDeleteChangeKeyEdDSA,
    } as MsgDeleteChangeKeyEdDSA;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.newKey !== undefined && object.newKey !== null) {
      message.newKey = object.newKey;
    } else {
      message.newKey = "";
    }
    return message;
  },
};

const baseMsgDeleteChangeKeyEdDSAResponse: object = {};

export const MsgDeleteChangeKeyEdDSAResponse = {
  encode(
    _: MsgDeleteChangeKeyEdDSAResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteChangeKeyEdDSAResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteChangeKeyEdDSAResponse,
    } as MsgDeleteChangeKeyEdDSAResponse;
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

  fromJSON(_: any): MsgDeleteChangeKeyEdDSAResponse {
    const message = {
      ...baseMsgDeleteChangeKeyEdDSAResponse,
    } as MsgDeleteChangeKeyEdDSAResponse;
    return message;
  },

  toJSON(_: MsgDeleteChangeKeyEdDSAResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteChangeKeyEdDSAResponse>
  ): MsgDeleteChangeKeyEdDSAResponse {
    const message = {
      ...baseMsgDeleteChangeKeyEdDSAResponse,
    } as MsgDeleteChangeKeyEdDSAResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateDeposit(request: MsgCreateDeposit): Promise<MsgCreateDepositResponse>;
  UpdateDeposit(request: MsgUpdateDeposit): Promise<MsgUpdateDepositResponse>;
  DeleteDeposit(request: MsgDeleteDeposit): Promise<MsgDeleteDepositResponse>;
  CreateConfirmation(
    request: MsgCreateConfirmation
  ): Promise<MsgCreateConfirmationResponse>;
  UpdateConfirmation(
    request: MsgUpdateConfirmation
  ): Promise<MsgUpdateConfirmationResponse>;
  DeleteConfirmation(
    request: MsgDeleteConfirmation
  ): Promise<MsgDeleteConfirmationResponse>;
  CreateChangeKeyECDSA(
    request: MsgCreateChangeKeyECDSA
  ): Promise<MsgCreateChangeKeyECDSAResponse>;
  UpdateChangeKeyECDSA(
    request: MsgUpdateChangeKeyECDSA
  ): Promise<MsgUpdateChangeKeyECDSAResponse>;
  DeleteChangeKeyECDSA(
    request: MsgDeleteChangeKeyECDSA
  ): Promise<MsgDeleteChangeKeyECDSAResponse>;
  CreateChangeKeyEdDSA(
    request: MsgCreateChangeKeyEdDSA
  ): Promise<MsgCreateChangeKeyEdDSAResponse>;
  UpdateChangeKeyEdDSA(
    request: MsgUpdateChangeKeyEdDSA
  ): Promise<MsgUpdateChangeKeyEdDSAResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  DeleteChangeKeyEdDSA(
    request: MsgDeleteChangeKeyEdDSA
  ): Promise<MsgDeleteChangeKeyEdDSAResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateDeposit(request: MsgCreateDeposit): Promise<MsgCreateDepositResponse> {
    const data = MsgCreateDeposit.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.rarimocore.Msg",
      "CreateDeposit",
      data
    );
    return promise.then((data) =>
      MsgCreateDepositResponse.decode(new Reader(data))
    );
  }

  UpdateDeposit(request: MsgUpdateDeposit): Promise<MsgUpdateDepositResponse> {
    const data = MsgUpdateDeposit.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.rarimocore.Msg",
      "UpdateDeposit",
      data
    );
    return promise.then((data) =>
      MsgUpdateDepositResponse.decode(new Reader(data))
    );
  }

  DeleteDeposit(request: MsgDeleteDeposit): Promise<MsgDeleteDepositResponse> {
    const data = MsgDeleteDeposit.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.rarimocore.Msg",
      "DeleteDeposit",
      data
    );
    return promise.then((data) =>
      MsgDeleteDepositResponse.decode(new Reader(data))
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

  UpdateConfirmation(
    request: MsgUpdateConfirmation
  ): Promise<MsgUpdateConfirmationResponse> {
    const data = MsgUpdateConfirmation.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.rarimocore.Msg",
      "UpdateConfirmation",
      data
    );
    return promise.then((data) =>
      MsgUpdateConfirmationResponse.decode(new Reader(data))
    );
  }

  DeleteConfirmation(
    request: MsgDeleteConfirmation
  ): Promise<MsgDeleteConfirmationResponse> {
    const data = MsgDeleteConfirmation.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.rarimocore.Msg",
      "DeleteConfirmation",
      data
    );
    return promise.then((data) =>
      MsgDeleteConfirmationResponse.decode(new Reader(data))
    );
  }

  CreateChangeKeyECDSA(
    request: MsgCreateChangeKeyECDSA
  ): Promise<MsgCreateChangeKeyECDSAResponse> {
    const data = MsgCreateChangeKeyECDSA.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.rarimocore.Msg",
      "CreateChangeKeyECDSA",
      data
    );
    return promise.then((data) =>
      MsgCreateChangeKeyECDSAResponse.decode(new Reader(data))
    );
  }

  UpdateChangeKeyECDSA(
    request: MsgUpdateChangeKeyECDSA
  ): Promise<MsgUpdateChangeKeyECDSAResponse> {
    const data = MsgUpdateChangeKeyECDSA.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.rarimocore.Msg",
      "UpdateChangeKeyECDSA",
      data
    );
    return promise.then((data) =>
      MsgUpdateChangeKeyECDSAResponse.decode(new Reader(data))
    );
  }

  DeleteChangeKeyECDSA(
    request: MsgDeleteChangeKeyECDSA
  ): Promise<MsgDeleteChangeKeyECDSAResponse> {
    const data = MsgDeleteChangeKeyECDSA.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.rarimocore.Msg",
      "DeleteChangeKeyECDSA",
      data
    );
    return promise.then((data) =>
      MsgDeleteChangeKeyECDSAResponse.decode(new Reader(data))
    );
  }

  CreateChangeKeyEdDSA(
    request: MsgCreateChangeKeyEdDSA
  ): Promise<MsgCreateChangeKeyEdDSAResponse> {
    const data = MsgCreateChangeKeyEdDSA.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.rarimocore.Msg",
      "CreateChangeKeyEdDSA",
      data
    );
    return promise.then((data) =>
      MsgCreateChangeKeyEdDSAResponse.decode(new Reader(data))
    );
  }

  UpdateChangeKeyEdDSA(
    request: MsgUpdateChangeKeyEdDSA
  ): Promise<MsgUpdateChangeKeyEdDSAResponse> {
    const data = MsgUpdateChangeKeyEdDSA.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.rarimocore.Msg",
      "UpdateChangeKeyEdDSA",
      data
    );
    return promise.then((data) =>
      MsgUpdateChangeKeyEdDSAResponse.decode(new Reader(data))
    );
  }

  DeleteChangeKeyEdDSA(
    request: MsgDeleteChangeKeyEdDSA
  ): Promise<MsgDeleteChangeKeyEdDSAResponse> {
    const data = MsgDeleteChangeKeyEdDSA.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.rarimocore.Msg",
      "DeleteChangeKeyEdDSA",
      data
    );
    return promise.then((data) =>
      MsgDeleteChangeKeyEdDSAResponse.decode(new Reader(data))
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
