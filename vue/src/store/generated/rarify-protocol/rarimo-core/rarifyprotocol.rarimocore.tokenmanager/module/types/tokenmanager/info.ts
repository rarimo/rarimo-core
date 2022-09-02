/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "rarifyprotocol.rarimocore.tokenmanager";

export interface ChainInfo {
  /** hex-encoded */
  tokenAddress: string;
  /** hex-encoded */
  tokenId: string;
}

export interface Info {
  index: string;
  name: string;
  symbol: string;
  image: string;
  currentChain: string;
  chains: { [key: string]: ChainInfo };
  creator: string;
}

export interface Info_ChainsEntry {
  key: string;
  value: ChainInfo | undefined;
}

const baseChainInfo: object = { tokenAddress: "", tokenId: "" };

export const ChainInfo = {
  encode(message: ChainInfo, writer: Writer = Writer.create()): Writer {
    if (message.tokenAddress !== "") {
      writer.uint32(10).string(message.tokenAddress);
    }
    if (message.tokenId !== "") {
      writer.uint32(18).string(message.tokenId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ChainInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseChainInfo } as ChainInfo;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.tokenAddress = reader.string();
          break;
        case 2:
          message.tokenId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ChainInfo {
    const message = { ...baseChainInfo } as ChainInfo;
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

  toJSON(message: ChainInfo): unknown {
    const obj: any = {};
    message.tokenAddress !== undefined &&
      (obj.tokenAddress = message.tokenAddress);
    message.tokenId !== undefined && (obj.tokenId = message.tokenId);
    return obj;
  },

  fromPartial(object: DeepPartial<ChainInfo>): ChainInfo {
    const message = { ...baseChainInfo } as ChainInfo;
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

const baseInfo: object = {
  index: "",
  name: "",
  symbol: "",
  image: "",
  currentChain: "",
  creator: "",
};

export const Info = {
  encode(message: Info, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.symbol !== "") {
      writer.uint32(26).string(message.symbol);
    }
    if (message.image !== "") {
      writer.uint32(34).string(message.image);
    }
    if (message.currentChain !== "") {
      writer.uint32(74).string(message.currentChain);
    }
    Object.entries(message.chains).forEach(([key, value]) => {
      Info_ChainsEntry.encode(
        { key: key as any, value },
        writer.uint32(82).fork()
      ).ldelim();
    });
    if (message.creator !== "") {
      writer.uint32(90).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Info {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseInfo } as Info;
    message.chains = {};
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.symbol = reader.string();
          break;
        case 4:
          message.image = reader.string();
          break;
        case 9:
          message.currentChain = reader.string();
          break;
        case 10:
          const entry10 = Info_ChainsEntry.decode(reader, reader.uint32());
          if (entry10.value !== undefined) {
            message.chains[entry10.key] = entry10.value;
          }
          break;
        case 11:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Info {
    const message = { ...baseInfo } as Info;
    message.chains = {};
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
    if (object.currentChain !== undefined && object.currentChain !== null) {
      message.currentChain = String(object.currentChain);
    } else {
      message.currentChain = "";
    }
    if (object.chains !== undefined && object.chains !== null) {
      Object.entries(object.chains).forEach(([key, value]) => {
        message.chains[key] = ChainInfo.fromJSON(value);
      });
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: Info): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.name !== undefined && (obj.name = message.name);
    message.symbol !== undefined && (obj.symbol = message.symbol);
    message.image !== undefined && (obj.image = message.image);
    message.currentChain !== undefined &&
      (obj.currentChain = message.currentChain);
    obj.chains = {};
    if (message.chains) {
      Object.entries(message.chains).forEach(([k, v]) => {
        obj.chains[k] = ChainInfo.toJSON(v);
      });
    }
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<Info>): Info {
    const message = { ...baseInfo } as Info;
    message.chains = {};
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
    if (object.currentChain !== undefined && object.currentChain !== null) {
      message.currentChain = object.currentChain;
    } else {
      message.currentChain = "";
    }
    if (object.chains !== undefined && object.chains !== null) {
      Object.entries(object.chains).forEach(([key, value]) => {
        if (value !== undefined) {
          message.chains[key] = ChainInfo.fromPartial(value);
        }
      });
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseInfo_ChainsEntry: object = { key: "" };

export const Info_ChainsEntry = {
  encode(message: Info_ChainsEntry, writer: Writer = Writer.create()): Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== undefined) {
      ChainInfo.encode(message.value, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Info_ChainsEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseInfo_ChainsEntry } as Info_ChainsEntry;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = ChainInfo.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Info_ChainsEntry {
    const message = { ...baseInfo_ChainsEntry } as Info_ChainsEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key);
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = ChainInfo.fromJSON(object.value);
    } else {
      message.value = undefined;
    }
    return message;
  },

  toJSON(message: Info_ChainsEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined &&
      (obj.value = message.value ? ChainInfo.toJSON(message.value) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<Info_ChainsEntry>): Info_ChainsEntry {
    const message = { ...baseInfo_ChainsEntry } as Info_ChainsEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = ChainInfo.fromPartial(object.value);
    } else {
      message.value = undefined;
    }
    return message;
  },
};

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
