/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { type, typeFromJSON, typeToJSON } from "./item";

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
  currentDecimals: number;
  /** Seed for deriving address for Solana networks. Encoded into hex string */
  currentSeed: string;
  currentImageUri: string;
  /** hex-encoded */
  currentImageHash: string;
  currentType: type;
}

export interface MsgCreateInfoResponse {
}

export interface MsgDeleteInfo {
  creator: string;
  index: string;
}

export interface MsgDeleteInfoResponse {
}

export interface MsgAddChain {
  creator: string;
  index: string;
  chainName: string;
  /** hex-encoded */
  tokenAddress: string;
  /** hex-encoded */
  tokenId: string;
  name: string;
  symbol: string;
  uri: string;
  decimals: number;
  /** Seed for deriving address for Solana networks. Encoded into hex string */
  seed: string;
  imageUri: string;
  /** hex-encoded */
  imageHash: string;
  tokenType: type;
}

export interface MsgAddChainResponse {
}

function createBaseMsgCreateInfo(): MsgCreateInfo {
  return {
    creator: "",
    index: "",
    currentAddress: "",
    currentId: "",
    currentChain: "",
    currentName: "",
    currentSymbol: "",
    currentURI: "",
    currentDecimals: 0,
    currentSeed: "",
    currentImageUri: "",
    currentImageHash: "",
    currentType: 0,
  };
}

export const MsgCreateInfo = {
  encode(message: MsgCreateInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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
    if (message.currentDecimals !== 0) {
      writer.uint32(72).uint32(message.currentDecimals);
    }
    if (message.currentSeed !== "") {
      writer.uint32(82).string(message.currentSeed);
    }
    if (message.currentImageUri !== "") {
      writer.uint32(90).string(message.currentImageUri);
    }
    if (message.currentImageHash !== "") {
      writer.uint32(98).string(message.currentImageHash);
    }
    if (message.currentType !== 0) {
      writer.uint32(104).int32(message.currentType);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateInfo();
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
          message.currentDecimals = reader.uint32();
          break;
        case 10:
          message.currentSeed = reader.string();
          break;
        case 11:
          message.currentImageUri = reader.string();
          break;
        case 12:
          message.currentImageHash = reader.string();
          break;
        case 13:
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      index: isSet(object.index) ? String(object.index) : "",
      currentAddress: isSet(object.currentAddress) ? String(object.currentAddress) : "",
      currentId: isSet(object.currentId) ? String(object.currentId) : "",
      currentChain: isSet(object.currentChain) ? String(object.currentChain) : "",
      currentName: isSet(object.currentName) ? String(object.currentName) : "",
      currentSymbol: isSet(object.currentSymbol) ? String(object.currentSymbol) : "",
      currentURI: isSet(object.currentURI) ? String(object.currentURI) : "",
      currentDecimals: isSet(object.currentDecimals) ? Number(object.currentDecimals) : 0,
      currentSeed: isSet(object.currentSeed) ? String(object.currentSeed) : "",
      currentImageUri: isSet(object.currentImageUri) ? String(object.currentImageUri) : "",
      currentImageHash: isSet(object.currentImageHash) ? String(object.currentImageHash) : "",
      currentType: isSet(object.currentType) ? typeFromJSON(object.currentType) : 0,
    };
  },

  toJSON(message: MsgCreateInfo): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    message.currentAddress !== undefined && (obj.currentAddress = message.currentAddress);
    message.currentId !== undefined && (obj.currentId = message.currentId);
    message.currentChain !== undefined && (obj.currentChain = message.currentChain);
    message.currentName !== undefined && (obj.currentName = message.currentName);
    message.currentSymbol !== undefined && (obj.currentSymbol = message.currentSymbol);
    message.currentURI !== undefined && (obj.currentURI = message.currentURI);
    message.currentDecimals !== undefined && (obj.currentDecimals = Math.round(message.currentDecimals));
    message.currentSeed !== undefined && (obj.currentSeed = message.currentSeed);
    message.currentImageUri !== undefined && (obj.currentImageUri = message.currentImageUri);
    message.currentImageHash !== undefined && (obj.currentImageHash = message.currentImageHash);
    message.currentType !== undefined && (obj.currentType = typeToJSON(message.currentType));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateInfo>, I>>(object: I): MsgCreateInfo {
    const message = createBaseMsgCreateInfo();
    message.creator = object.creator ?? "";
    message.index = object.index ?? "";
    message.currentAddress = object.currentAddress ?? "";
    message.currentId = object.currentId ?? "";
    message.currentChain = object.currentChain ?? "";
    message.currentName = object.currentName ?? "";
    message.currentSymbol = object.currentSymbol ?? "";
    message.currentURI = object.currentURI ?? "";
    message.currentDecimals = object.currentDecimals ?? 0;
    message.currentSeed = object.currentSeed ?? "";
    message.currentImageUri = object.currentImageUri ?? "";
    message.currentImageHash = object.currentImageHash ?? "";
    message.currentType = object.currentType ?? 0;
    return message;
  },
};

function createBaseMsgCreateInfoResponse(): MsgCreateInfoResponse {
  return {};
}

export const MsgCreateInfoResponse = {
  encode(_: MsgCreateInfoResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateInfoResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateInfoResponse();
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
    return {};
  },

  toJSON(_: MsgCreateInfoResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateInfoResponse>, I>>(_: I): MsgCreateInfoResponse {
    const message = createBaseMsgCreateInfoResponse();
    return message;
  },
};

function createBaseMsgDeleteInfo(): MsgDeleteInfo {
  return { creator: "", index: "" };
}

export const MsgDeleteInfo = {
  encode(message: MsgDeleteInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteInfo();
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      index: isSet(object.index) ? String(object.index) : "",
    };
  },

  toJSON(message: MsgDeleteInfo): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteInfo>, I>>(object: I): MsgDeleteInfo {
    const message = createBaseMsgDeleteInfo();
    message.creator = object.creator ?? "";
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseMsgDeleteInfoResponse(): MsgDeleteInfoResponse {
  return {};
}

export const MsgDeleteInfoResponse = {
  encode(_: MsgDeleteInfoResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteInfoResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteInfoResponse();
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
    return {};
  },

  toJSON(_: MsgDeleteInfoResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteInfoResponse>, I>>(_: I): MsgDeleteInfoResponse {
    const message = createBaseMsgDeleteInfoResponse();
    return message;
  },
};

function createBaseMsgAddChain(): MsgAddChain {
  return {
    creator: "",
    index: "",
    chainName: "",
    tokenAddress: "",
    tokenId: "",
    name: "",
    symbol: "",
    uri: "",
    decimals: 0,
    seed: "",
    imageUri: "",
    imageHash: "",
    tokenType: 0,
  };
}

export const MsgAddChain = {
  encode(message: MsgAddChain, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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
    if (message.name !== "") {
      writer.uint32(50).string(message.name);
    }
    if (message.symbol !== "") {
      writer.uint32(58).string(message.symbol);
    }
    if (message.uri !== "") {
      writer.uint32(66).string(message.uri);
    }
    if (message.decimals !== 0) {
      writer.uint32(72).uint32(message.decimals);
    }
    if (message.seed !== "") {
      writer.uint32(82).string(message.seed);
    }
    if (message.imageUri !== "") {
      writer.uint32(90).string(message.imageUri);
    }
    if (message.imageHash !== "") {
      writer.uint32(98).string(message.imageHash);
    }
    if (message.tokenType !== 0) {
      writer.uint32(104).int32(message.tokenType);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddChain {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddChain();
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
          message.name = reader.string();
          break;
        case 7:
          message.symbol = reader.string();
          break;
        case 8:
          message.uri = reader.string();
          break;
        case 9:
          message.decimals = reader.uint32();
          break;
        case 10:
          message.seed = reader.string();
          break;
        case 11:
          message.imageUri = reader.string();
          break;
        case 12:
          message.imageHash = reader.string();
          break;
        case 13:
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      index: isSet(object.index) ? String(object.index) : "",
      chainName: isSet(object.chainName) ? String(object.chainName) : "",
      tokenAddress: isSet(object.tokenAddress) ? String(object.tokenAddress) : "",
      tokenId: isSet(object.tokenId) ? String(object.tokenId) : "",
      name: isSet(object.name) ? String(object.name) : "",
      symbol: isSet(object.symbol) ? String(object.symbol) : "",
      uri: isSet(object.uri) ? String(object.uri) : "",
      decimals: isSet(object.decimals) ? Number(object.decimals) : 0,
      seed: isSet(object.seed) ? String(object.seed) : "",
      imageUri: isSet(object.imageUri) ? String(object.imageUri) : "",
      imageHash: isSet(object.imageHash) ? String(object.imageHash) : "",
      tokenType: isSet(object.tokenType) ? typeFromJSON(object.tokenType) : 0,
    };
  },

  toJSON(message: MsgAddChain): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    message.chainName !== undefined && (obj.chainName = message.chainName);
    message.tokenAddress !== undefined && (obj.tokenAddress = message.tokenAddress);
    message.tokenId !== undefined && (obj.tokenId = message.tokenId);
    message.name !== undefined && (obj.name = message.name);
    message.symbol !== undefined && (obj.symbol = message.symbol);
    message.uri !== undefined && (obj.uri = message.uri);
    message.decimals !== undefined && (obj.decimals = Math.round(message.decimals));
    message.seed !== undefined && (obj.seed = message.seed);
    message.imageUri !== undefined && (obj.imageUri = message.imageUri);
    message.imageHash !== undefined && (obj.imageHash = message.imageHash);
    message.tokenType !== undefined && (obj.tokenType = typeToJSON(message.tokenType));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddChain>, I>>(object: I): MsgAddChain {
    const message = createBaseMsgAddChain();
    message.creator = object.creator ?? "";
    message.index = object.index ?? "";
    message.chainName = object.chainName ?? "";
    message.tokenAddress = object.tokenAddress ?? "";
    message.tokenId = object.tokenId ?? "";
    message.name = object.name ?? "";
    message.symbol = object.symbol ?? "";
    message.uri = object.uri ?? "";
    message.decimals = object.decimals ?? 0;
    message.seed = object.seed ?? "";
    message.imageUri = object.imageUri ?? "";
    message.imageHash = object.imageHash ?? "";
    message.tokenType = object.tokenType ?? 0;
    return message;
  },
};

function createBaseMsgAddChainResponse(): MsgAddChainResponse {
  return {};
}

export const MsgAddChainResponse = {
  encode(_: MsgAddChainResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddChainResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddChainResponse();
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
    return {};
  },

  toJSON(_: MsgAddChainResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddChainResponse>, I>>(_: I): MsgAddChainResponse {
    const message = createBaseMsgAddChainResponse();
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
    this.CreateInfo = this.CreateInfo.bind(this);
    this.DeleteInfo = this.DeleteInfo.bind(this);
    this.AddChain = this.AddChain.bind(this);
  }
  CreateInfo(request: MsgCreateInfo): Promise<MsgCreateInfoResponse> {
    const data = MsgCreateInfo.encode(request).finish();
    const promise = this.rpc.request("rarifyprotocol.rarimocore.tokenmanager.Msg", "CreateInfo", data);
    return promise.then((data) => MsgCreateInfoResponse.decode(new _m0.Reader(data)));
  }

  DeleteInfo(request: MsgDeleteInfo): Promise<MsgDeleteInfoResponse> {
    const data = MsgDeleteInfo.encode(request).finish();
    const promise = this.rpc.request("rarifyprotocol.rarimocore.tokenmanager.Msg", "DeleteInfo", data);
    return promise.then((data) => MsgDeleteInfoResponse.decode(new _m0.Reader(data)));
  }

  AddChain(request: MsgAddChain): Promise<MsgAddChainResponse> {
    const data = MsgAddChain.encode(request).finish();
    const promise = this.rpc.request("rarifyprotocol.rarimocore.tokenmanager.Msg", "AddChain", data);
    return promise.then((data) => MsgAddChainResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
