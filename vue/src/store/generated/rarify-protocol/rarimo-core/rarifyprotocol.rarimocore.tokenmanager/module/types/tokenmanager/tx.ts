/* eslint-disable */
import { type, typeFromJSON, typeToJSON } from "../tokenmanager/info";
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "rarifyprotocol.rarimocore.tokenmanager";

export interface MsgCreateInfo {
  creator: string;
  index: string;
  /** hex-encoded */
  currentAddress: string;
  /** hex-encoded */
  currentId: string;
  currentChain: string;
  currentName: string;
  currentSymbol: string;
  currentURI: string;
  currentType: type;
}

export interface MsgCreateInfoResponse {}

export interface MsgDeleteInfo {
  creator: string;
  index: string;
}

export interface MsgDeleteInfoResponse {}

export interface MsgAddChain {
  creator: string;
  index: string;
  chainName: string;
  /** hex-encoded */
  tokenAddress: string;
  /** hex-encoded */
  tokenId: string;
  tokenType: type;
}

export interface MsgAddChainResponse {}

const baseMsgCreateInfo: object = {
  creator: "",
  index: "",
  currentAddress: "",
  currentId: "",
  currentChain: "",
  currentName: "",
  currentSymbol: "",
  currentURI: "",
  currentType: 0,
};

export const MsgCreateInfo = {
  encode(message: MsgCreateInfo, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    if (message.currentAddress !== "") {
      writer.uint32(26).string(message.currentAddress);
    }
    if (message.currentId !== "") {
      writer.uint32(34).string(message.currentId);
    }
    if (message.currentChain !== "") {
      writer.uint32(42).string(message.currentChain);
    }
    if (message.currentName !== "") {
      writer.uint32(50).string(message.currentName);
    }
    if (message.currentSymbol !== "") {
      writer.uint32(58).string(message.currentSymbol);
    }
    if (message.currentURI !== "") {
      writer.uint32(66).string(message.currentURI);
    }
    if (message.currentType !== 0) {
      writer.uint32(72).int32(message.currentType);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateInfo } as MsgCreateInfo;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        case 3:
          message.currentAddress = reader.string();
          break;
        case 4:
          message.currentId = reader.string();
          break;
        case 5:
          message.currentChain = reader.string();
          break;
        case 6:
          message.currentName = reader.string();
          break;
        case 7:
          message.currentSymbol = reader.string();
          break;
        case 8:
          message.currentURI = reader.string();
          break;
        case 9:
          message.currentType = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateInfo {
    const message = { ...baseMsgCreateInfo } as MsgCreateInfo;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.currentAddress !== undefined && object.currentAddress !== null) {
      message.currentAddress = String(object.currentAddress);
    } else {
      message.currentAddress = "";
    }
    if (object.currentId !== undefined && object.currentId !== null) {
      message.currentId = String(object.currentId);
    } else {
      message.currentId = "";
    }
    if (object.currentChain !== undefined && object.currentChain !== null) {
      message.currentChain = String(object.currentChain);
    } else {
      message.currentChain = "";
    }
    if (object.currentName !== undefined && object.currentName !== null) {
      message.currentName = String(object.currentName);
    } else {
      message.currentName = "";
    }
    if (object.currentSymbol !== undefined && object.currentSymbol !== null) {
      message.currentSymbol = String(object.currentSymbol);
    } else {
      message.currentSymbol = "";
    }
    if (object.currentURI !== undefined && object.currentURI !== null) {
      message.currentURI = String(object.currentURI);
    } else {
      message.currentURI = "";
    }
    if (object.currentType !== undefined && object.currentType !== null) {
      message.currentType = typeFromJSON(object.currentType);
    } else {
      message.currentType = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateInfo): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    message.currentAddress !== undefined &&
      (obj.currentAddress = message.currentAddress);
    message.currentId !== undefined && (obj.currentId = message.currentId);
    message.currentChain !== undefined &&
      (obj.currentChain = message.currentChain);
    message.currentName !== undefined &&
      (obj.currentName = message.currentName);
    message.currentSymbol !== undefined &&
      (obj.currentSymbol = message.currentSymbol);
    message.currentURI !== undefined && (obj.currentURI = message.currentURI);
    message.currentType !== undefined &&
      (obj.currentType = typeToJSON(message.currentType));
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateInfo>): MsgCreateInfo {
    const message = { ...baseMsgCreateInfo } as MsgCreateInfo;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.currentAddress !== undefined && object.currentAddress !== null) {
      message.currentAddress = object.currentAddress;
    } else {
      message.currentAddress = "";
    }
    if (object.currentId !== undefined && object.currentId !== null) {
      message.currentId = object.currentId;
    } else {
      message.currentId = "";
    }
    if (object.currentChain !== undefined && object.currentChain !== null) {
      message.currentChain = object.currentChain;
    } else {
      message.currentChain = "";
    }
    if (object.currentName !== undefined && object.currentName !== null) {
      message.currentName = object.currentName;
    } else {
      message.currentName = "";
    }
    if (object.currentSymbol !== undefined && object.currentSymbol !== null) {
      message.currentSymbol = object.currentSymbol;
    } else {
      message.currentSymbol = "";
    }
    if (object.currentURI !== undefined && object.currentURI !== null) {
      message.currentURI = object.currentURI;
    } else {
      message.currentURI = "";
    }
    if (object.currentType !== undefined && object.currentType !== null) {
      message.currentType = object.currentType;
    } else {
      message.currentType = 0;
    }
    return message;
  },
};

const baseMsgCreateInfoResponse: object = {};

export const MsgCreateInfoResponse = {
  encode(_: MsgCreateInfoResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateInfoResponse } as MsgCreateInfoResponse;
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

  fromJSON(_: any): MsgCreateInfoResponse {
    const message = { ...baseMsgCreateInfoResponse } as MsgCreateInfoResponse;
    return message;
  },

  toJSON(_: MsgCreateInfoResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgCreateInfoResponse>): MsgCreateInfoResponse {
    const message = { ...baseMsgCreateInfoResponse } as MsgCreateInfoResponse;
    return message;
  },
};

const baseMsgDeleteInfo: object = { creator: "", index: "" };

export const MsgDeleteInfo = {
  encode(message: MsgDeleteInfo, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteInfo } as MsgDeleteInfo;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteInfo {
    const message = { ...baseMsgDeleteInfo } as MsgDeleteInfo;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: MsgDeleteInfo): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeleteInfo>): MsgDeleteInfo {
    const message = { ...baseMsgDeleteInfo } as MsgDeleteInfo;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseMsgDeleteInfoResponse: object = {};

export const MsgDeleteInfoResponse = {
  encode(_: MsgDeleteInfoResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteInfoResponse } as MsgDeleteInfoResponse;
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

  fromJSON(_: any): MsgDeleteInfoResponse {
    const message = { ...baseMsgDeleteInfoResponse } as MsgDeleteInfoResponse;
    return message;
  },

  toJSON(_: MsgDeleteInfoResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgDeleteInfoResponse>): MsgDeleteInfoResponse {
    const message = { ...baseMsgDeleteInfoResponse } as MsgDeleteInfoResponse;
    return message;
  },
};

const baseMsgAddChain: object = {
  creator: "",
  index: "",
  chainName: "",
  tokenAddress: "",
  tokenId: "",
  tokenType: 0,
};

export const MsgAddChain = {
  encode(message: MsgAddChain, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    if (message.chainName !== "") {
      writer.uint32(26).string(message.chainName);
    }
    if (message.tokenAddress !== "") {
      writer.uint32(34).string(message.tokenAddress);
    }
    if (message.tokenId !== "") {
      writer.uint32(42).string(message.tokenId);
    }
    if (message.tokenType !== 0) {
      writer.uint32(48).int32(message.tokenType);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddChain {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddChain } as MsgAddChain;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        case 3:
          message.chainName = reader.string();
          break;
        case 4:
          message.tokenAddress = reader.string();
          break;
        case 5:
          message.tokenId = reader.string();
          break;
        case 6:
          message.tokenType = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddChain {
    const message = { ...baseMsgAddChain } as MsgAddChain;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.chainName !== undefined && object.chainName !== null) {
      message.chainName = String(object.chainName);
    } else {
      message.chainName = "";
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
    if (object.tokenType !== undefined && object.tokenType !== null) {
      message.tokenType = typeFromJSON(object.tokenType);
    } else {
      message.tokenType = 0;
    }
    return message;
  },

  toJSON(message: MsgAddChain): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    message.chainName !== undefined && (obj.chainName = message.chainName);
    message.tokenAddress !== undefined &&
      (obj.tokenAddress = message.tokenAddress);
    message.tokenId !== undefined && (obj.tokenId = message.tokenId);
    message.tokenType !== undefined &&
      (obj.tokenType = typeToJSON(message.tokenType));
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAddChain>): MsgAddChain {
    const message = { ...baseMsgAddChain } as MsgAddChain;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.chainName !== undefined && object.chainName !== null) {
      message.chainName = object.chainName;
    } else {
      message.chainName = "";
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
    if (object.tokenType !== undefined && object.tokenType !== null) {
      message.tokenType = object.tokenType;
    } else {
      message.tokenType = 0;
    }
    return message;
  },
};

const baseMsgAddChainResponse: object = {};

export const MsgAddChainResponse = {
  encode(_: MsgAddChainResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddChainResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddChainResponse } as MsgAddChainResponse;
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

  fromJSON(_: any): MsgAddChainResponse {
    const message = { ...baseMsgAddChainResponse } as MsgAddChainResponse;
    return message;
  },

  toJSON(_: MsgAddChainResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgAddChainResponse>): MsgAddChainResponse {
    const message = { ...baseMsgAddChainResponse } as MsgAddChainResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateInfo(request: MsgCreateInfo): Promise<MsgCreateInfoResponse>;
  DeleteInfo(request: MsgDeleteInfo): Promise<MsgDeleteInfoResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  AddChain(request: MsgAddChain): Promise<MsgAddChainResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateInfo(request: MsgCreateInfo): Promise<MsgCreateInfoResponse> {
    const data = MsgCreateInfo.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.tokenmanager.Msg",
      "CreateInfo",
      data
    );
    return promise.then((data) =>
      MsgCreateInfoResponse.decode(new Reader(data))
    );
  }

  DeleteInfo(request: MsgDeleteInfo): Promise<MsgDeleteInfoResponse> {
    const data = MsgDeleteInfo.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.tokenmanager.Msg",
      "DeleteInfo",
      data
    );
    return promise.then((data) =>
      MsgDeleteInfoResponse.decode(new Reader(data))
    );
  }

  AddChain(request: MsgAddChain): Promise<MsgAddChainResponse> {
    const data = MsgAddChain.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.tokenmanager.Msg",
      "AddChain",
      data
    );
    return promise.then((data) => MsgAddChainResponse.decode(new Reader(data)));
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
