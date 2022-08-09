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

/** Msg defines the Msg service. */
export interface Msg {
  CreateDeposit(request: MsgCreateDeposit): Promise<MsgCreateDepositResponse>;
  UpdateDeposit(request: MsgUpdateDeposit): Promise<MsgUpdateDepositResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  DeleteDeposit(request: MsgDeleteDeposit): Promise<MsgDeleteDepositResponse>;
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
