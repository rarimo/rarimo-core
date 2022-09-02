/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "rarifyprotocol.rarimocore.tokenmanager";

export interface MsgCreateInfo {
  creator: string;
  index: string;
  name: string;
  symbol: string;
  image: string;
  currentAddress: Uint8Array;
  currentId: Uint8Array;
  currentChain: string;
}

export interface MsgCreateInfoResponse {}

export interface MsgUpdateInfo {
  creator: string;
  index: string;
  name: string;
  symbol: string;
  image: string;
}

export interface MsgUpdateInfoResponse {}

export interface MsgDeleteInfo {
  creator: string;
  index: string;
}

export interface MsgDeleteInfoResponse {}

export interface MsgAddChain {
  creator: string;
  index: string;
  chainName: string;
  tokenAddress: Uint8Array;
  tokenId: Uint8Array;
}

export interface MsgAddChainResponse {}

const baseMsgCreateInfo: object = {
  creator: "",
  index: "",
  name: "",
  symbol: "",
  image: "",
  currentChain: "",
};

export const MsgCreateInfo = {
  encode(message: MsgCreateInfo, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    if (message.symbol !== "") {
      writer.uint32(34).string(message.symbol);
    }
    if (message.image !== "") {
      writer.uint32(42).string(message.image);
    }
    if (message.currentAddress.length !== 0) {
      writer.uint32(50).bytes(message.currentAddress);
    }
    if (message.currentId.length !== 0) {
      writer.uint32(58).bytes(message.currentId);
    }
    if (message.currentChain !== "") {
      writer.uint32(66).string(message.currentChain);
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
          message.name = reader.string();
          break;
        case 4:
          message.symbol = reader.string();
          break;
        case 5:
          message.image = reader.string();
          break;
        case 6:
          message.currentAddress = reader.bytes();
          break;
        case 7:
          message.currentId = reader.bytes();
          break;
        case 8:
          message.currentChain = reader.string();
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
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.symbol !== undefined && object.symbol !== null) {
      message.symbol = String(object.symbol);
    } else {
      message.symbol = "";
    }
    if (object.image !== undefined && object.image !== null) {
      message.image = String(object.image);
    } else {
      message.image = "";
    }
    if (object.currentAddress !== undefined && object.currentAddress !== null) {
      message.currentAddress = bytesFromBase64(object.currentAddress);
    }
    if (object.currentId !== undefined && object.currentId !== null) {
      message.currentId = bytesFromBase64(object.currentId);
    }
    if (object.currentChain !== undefined && object.currentChain !== null) {
      message.currentChain = String(object.currentChain);
    } else {
      message.currentChain = "";
    }
    return message;
  },

  toJSON(message: MsgCreateInfo): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    message.name !== undefined && (obj.name = message.name);
    message.symbol !== undefined && (obj.symbol = message.symbol);
    message.image !== undefined && (obj.image = message.image);
    message.currentAddress !== undefined &&
      (obj.currentAddress = base64FromBytes(
        message.currentAddress !== undefined
          ? message.currentAddress
          : new Uint8Array()
      ));
    message.currentId !== undefined &&
      (obj.currentId = base64FromBytes(
        message.currentId !== undefined ? message.currentId : new Uint8Array()
      ));
    message.currentChain !== undefined &&
      (obj.currentChain = message.currentChain);
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
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.symbol !== undefined && object.symbol !== null) {
      message.symbol = object.symbol;
    } else {
      message.symbol = "";
    }
    if (object.image !== undefined && object.image !== null) {
      message.image = object.image;
    } else {
      message.image = "";
    }
    if (object.currentAddress !== undefined && object.currentAddress !== null) {
      message.currentAddress = object.currentAddress;
    } else {
      message.currentAddress = new Uint8Array();
    }
    if (object.currentId !== undefined && object.currentId !== null) {
      message.currentId = object.currentId;
    } else {
      message.currentId = new Uint8Array();
    }
    if (object.currentChain !== undefined && object.currentChain !== null) {
      message.currentChain = object.currentChain;
    } else {
      message.currentChain = "";
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

const baseMsgUpdateInfo: object = {
  creator: "",
  index: "",
  name: "",
  symbol: "",
  image: "",
};

export const MsgUpdateInfo = {
  encode(message: MsgUpdateInfo, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    if (message.symbol !== "") {
      writer.uint32(34).string(message.symbol);
    }
    if (message.image !== "") {
      writer.uint32(42).string(message.image);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateInfo } as MsgUpdateInfo;
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
          message.name = reader.string();
          break;
        case 4:
          message.symbol = reader.string();
          break;
        case 5:
          message.image = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateInfo {
    const message = { ...baseMsgUpdateInfo } as MsgUpdateInfo;
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
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.symbol !== undefined && object.symbol !== null) {
      message.symbol = String(object.symbol);
    } else {
      message.symbol = "";
    }
    if (object.image !== undefined && object.image !== null) {
      message.image = String(object.image);
    } else {
      message.image = "";
    }
    return message;
  },

  toJSON(message: MsgUpdateInfo): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    message.name !== undefined && (obj.name = message.name);
    message.symbol !== undefined && (obj.symbol = message.symbol);
    message.image !== undefined && (obj.image = message.image);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateInfo>): MsgUpdateInfo {
    const message = { ...baseMsgUpdateInfo } as MsgUpdateInfo;
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
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.symbol !== undefined && object.symbol !== null) {
      message.symbol = object.symbol;
    } else {
      message.symbol = "";
    }
    if (object.image !== undefined && object.image !== null) {
      message.image = object.image;
    } else {
      message.image = "";
    }
    return message;
  },
};

const baseMsgUpdateInfoResponse: object = {};

export const MsgUpdateInfoResponse = {
  encode(_: MsgUpdateInfoResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateInfoResponse } as MsgUpdateInfoResponse;
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

  fromJSON(_: any): MsgUpdateInfoResponse {
    const message = { ...baseMsgUpdateInfoResponse } as MsgUpdateInfoResponse;
    return message;
  },

  toJSON(_: MsgUpdateInfoResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgUpdateInfoResponse>): MsgUpdateInfoResponse {
    const message = { ...baseMsgUpdateInfoResponse } as MsgUpdateInfoResponse;
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

const baseMsgAddChain: object = { creator: "", index: "", chainName: "" };

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
    if (message.tokenAddress.length !== 0) {
      writer.uint32(34).bytes(message.tokenAddress);
    }
    if (message.tokenId.length !== 0) {
      writer.uint32(42).bytes(message.tokenId);
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
          message.tokenAddress = reader.bytes();
          break;
        case 5:
          message.tokenId = reader.bytes();
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
      message.tokenAddress = bytesFromBase64(object.tokenAddress);
    }
    if (object.tokenId !== undefined && object.tokenId !== null) {
      message.tokenId = bytesFromBase64(object.tokenId);
    }
    return message;
  },

  toJSON(message: MsgAddChain): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    message.chainName !== undefined && (obj.chainName = message.chainName);
    message.tokenAddress !== undefined &&
      (obj.tokenAddress = base64FromBytes(
        message.tokenAddress !== undefined
          ? message.tokenAddress
          : new Uint8Array()
      ));
    message.tokenId !== undefined &&
      (obj.tokenId = base64FromBytes(
        message.tokenId !== undefined ? message.tokenId : new Uint8Array()
      ));
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
      message.tokenAddress = new Uint8Array();
    }
    if (object.tokenId !== undefined && object.tokenId !== null) {
      message.tokenId = object.tokenId;
    } else {
      message.tokenId = new Uint8Array();
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
  UpdateInfo(request: MsgUpdateInfo): Promise<MsgUpdateInfoResponse>;
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

  UpdateInfo(request: MsgUpdateInfo): Promise<MsgUpdateInfoResponse> {
    const data = MsgUpdateInfo.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.tokenmanager.Msg",
      "UpdateInfo",
      data
    );
    return promise.then((data) =>
      MsgUpdateInfoResponse.decode(new Reader(data))
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

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

const atob: (b64: string) => string =
  globalThis.atob ||
  ((b64) => globalThis.Buffer.from(b64, "base64").toString("binary"));
function bytesFromBase64(b64: string): Uint8Array {
  const bin = atob(b64);
  const arr = new Uint8Array(bin.length);
  for (let i = 0; i < bin.length; ++i) {
    arr[i] = bin.charCodeAt(i);
  }
  return arr;
}

const btoa: (bin: string) => string =
  globalThis.btoa ||
  ((bin) => globalThis.Buffer.from(bin, "binary").toString("base64"));
function base64FromBytes(arr: Uint8Array): string {
  const bin: string[] = [];
  for (let i = 0; i < arr.byteLength; ++i) {
    bin.push(String.fromCharCode(arr[i]));
  }
  return btoa(bin.join(""));
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
