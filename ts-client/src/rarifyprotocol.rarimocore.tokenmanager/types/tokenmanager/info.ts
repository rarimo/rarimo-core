/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "rarifyprotocol.rarimocore.tokenmanager";

export interface ChainInfo {
  /** hex-encoded */
  tokenAddress: string;
  /** hex-encoded */
  tokenId: string;
}

export interface Info {
  index: string;
  chains: { [key: string]: ChainInfo };
  creator: string;
}

export interface Info_ChainsEntry {
  key: string;
  value: ChainInfo | undefined;
}

function createBaseChainInfo(): ChainInfo {
  return { tokenAddress: "", tokenId: "" };
}

export const ChainInfo = {
  encode(message: ChainInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.tokenAddress !== "") {
      writer.uint32(10).string(message.tokenAddress);
    }
    if (message.tokenId !== "") {
      writer.uint32(18).string(message.tokenId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ChainInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseChainInfo();
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
    return {
      tokenAddress: isSet(object.tokenAddress) ? String(object.tokenAddress) : "",
      tokenId: isSet(object.tokenId) ? String(object.tokenId) : "",
    };
  },

  toJSON(message: ChainInfo): unknown {
    const obj: any = {};
    message.tokenAddress !== undefined && (obj.tokenAddress = message.tokenAddress);
    message.tokenId !== undefined && (obj.tokenId = message.tokenId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ChainInfo>, I>>(object: I): ChainInfo {
    const message = createBaseChainInfo();
    message.tokenAddress = object.tokenAddress ?? "";
    message.tokenId = object.tokenId ?? "";
    return message;
  },
};

function createBaseInfo(): Info {
  return { index: "", chains: {}, creator: "" };
}

export const Info = {
  encode(message: Info, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    Object.entries(message.chains).forEach(([key, value]) => {
      Info_ChainsEntry.encode({ key: key as any, value }, writer.uint32(18).fork()).ldelim();
    });
    if (message.creator !== "") {
      writer.uint32(26).string(message.creator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Info {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          const entry2 = Info_ChainsEntry.decode(reader, reader.uint32());
          if (entry2.value !== undefined) {
            message.chains[entry2.key] = entry2.value;
          }
          break;
        case 3:
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
    return {
      index: isSet(object.index) ? String(object.index) : "",
      chains: isObject(object.chains)
        ? Object.entries(object.chains).reduce<{ [key: string]: ChainInfo }>((acc, [key, value]) => {
          acc[key] = ChainInfo.fromJSON(value);
          return acc;
        }, {})
        : {},
      creator: isSet(object.creator) ? String(object.creator) : "",
    };
  },

  toJSON(message: Info): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    obj.chains = {};
    if (message.chains) {
      Object.entries(message.chains).forEach(([k, v]) => {
        obj.chains[k] = ChainInfo.toJSON(v);
      });
    }
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Info>, I>>(object: I): Info {
    const message = createBaseInfo();
    message.index = object.index ?? "";
    message.chains = Object.entries(object.chains ?? {}).reduce<{ [key: string]: ChainInfo }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = ChainInfo.fromPartial(value);
      }
      return acc;
    }, {});
    message.creator = object.creator ?? "";
    return message;
  },
};

function createBaseInfo_ChainsEntry(): Info_ChainsEntry {
  return { key: "", value: undefined };
}

export const Info_ChainsEntry = {
  encode(message: Info_ChainsEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== undefined) {
      ChainInfo.encode(message.value, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Info_ChainsEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInfo_ChainsEntry();
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
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? ChainInfo.fromJSON(object.value) : undefined,
    };
  },

  toJSON(message: Info_ChainsEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value ? ChainInfo.toJSON(message.value) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Info_ChainsEntry>, I>>(object: I): Info_ChainsEntry {
    const message = createBaseInfo_ChainsEntry();
    message.key = object.key ?? "";
    message.value = (object.value !== undefined && object.value !== null)
      ? ChainInfo.fromPartial(object.value)
      : undefined;
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isObject(value: any): boolean {
  return typeof value === "object" && value !== null;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
