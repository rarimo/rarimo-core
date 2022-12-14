/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "rarifyprotocol.rarimocore.tokenmanager";

export enum NetworkType {
  EVM = 0,
  Solana = 1,
  Near = 2,
  Other = 3,
  UNRECOGNIZED = -1,
}

export function networkTypeFromJSON(object: any): NetworkType {
  switch (object) {
    case 0:
    case "EVM":
      return NetworkType.EVM;
    case 1:
    case "Solana":
      return NetworkType.Solana;
    case 2:
    case "Near":
      return NetworkType.Near;
    case 3:
    case "Other":
      return NetworkType.Other;
    case -1:
    case "UNRECOGNIZED":
    default:
      return NetworkType.UNRECOGNIZED;
  }
}

export function networkTypeToJSON(object: NetworkType): string {
  switch (object) {
    case NetworkType.EVM:
      return "EVM";
    case NetworkType.Solana:
      return "Solana";
    case NetworkType.Near:
      return "Near";
    case NetworkType.Other:
      return "Other";
    case NetworkType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface ChainParams {
  contract: string;
  types: { [key: string]: number };
  type: NetworkType;
}

export interface ChainParams_TypesEntry {
  key: string;
  value: number;
}

/** Params defines the parameters for the module. */
export interface Params {
  networks: { [key: string]: ChainParams };
}

export interface Params_NetworksEntry {
  key: string;
  value: ChainParams | undefined;
}

function createBaseChainParams(): ChainParams {
  return { contract: "", types: {}, type: 0 };
}

export const ChainParams = {
  encode(message: ChainParams, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.contract !== "") {
      writer.uint32(10).string(message.contract);
    }
    Object.entries(message.types).forEach(([key, value]) => {
      ChainParams_TypesEntry.encode({ key: key as any, value }, writer.uint32(18).fork()).ldelim();
    });
    if (message.type !== 0) {
      writer.uint32(24).int32(message.type);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ChainParams {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseChainParams();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.contract = reader.string();
          break;
        case 2:
          const entry2 = ChainParams_TypesEntry.decode(reader, reader.uint32());
          if (entry2.value !== undefined) {
            message.types[entry2.key] = entry2.value;
          }
          break;
        case 3:
          message.type = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ChainParams {
    return {
      contract: isSet(object.contract) ? String(object.contract) : "",
      types: isObject(object.types)
        ? Object.entries(object.types).reduce<{ [key: string]: number }>((acc, [key, value]) => {
          acc[key] = Number(value);
          return acc;
        }, {})
        : {},
      type: isSet(object.type) ? networkTypeFromJSON(object.type) : 0,
    };
  },

  toJSON(message: ChainParams): unknown {
    const obj: any = {};
    message.contract !== undefined && (obj.contract = message.contract);
    obj.types = {};
    if (message.types) {
      Object.entries(message.types).forEach(([k, v]) => {
        obj.types[k] = Math.round(v);
      });
    }
    message.type !== undefined && (obj.type = networkTypeToJSON(message.type));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ChainParams>, I>>(object: I): ChainParams {
    const message = createBaseChainParams();
    message.contract = object.contract ?? "";
    message.types = Object.entries(object.types ?? {}).reduce<{ [key: string]: number }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = Number(value);
      }
      return acc;
    }, {});
    message.type = object.type ?? 0;
    return message;
  },
};

function createBaseChainParams_TypesEntry(): ChainParams_TypesEntry {
  return { key: "", value: 0 };
}

export const ChainParams_TypesEntry = {
  encode(message: ChainParams_TypesEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== 0) {
      writer.uint32(16).uint32(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ChainParams_TypesEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseChainParams_TypesEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = reader.uint32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ChainParams_TypesEntry {
    return { key: isSet(object.key) ? String(object.key) : "", value: isSet(object.value) ? Number(object.value) : 0 };
  },

  toJSON(message: ChainParams_TypesEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = Math.round(message.value));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ChainParams_TypesEntry>, I>>(object: I): ChainParams_TypesEntry {
    const message = createBaseChainParams_TypesEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? 0;
    return message;
  },
};

function createBaseParams(): Params {
  return { networks: {} };
}

export const Params = {
  encode(message: Params, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    Object.entries(message.networks).forEach(([key, value]) => {
      Params_NetworksEntry.encode({ key: key as any, value }, writer.uint32(10).fork()).ldelim();
    });
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseParams();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          const entry1 = Params_NetworksEntry.decode(reader, reader.uint32());
          if (entry1.value !== undefined) {
            message.networks[entry1.key] = entry1.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Params {
    return {
      networks: isObject(object.networks)
        ? Object.entries(object.networks).reduce<{ [key: string]: ChainParams }>((acc, [key, value]) => {
          acc[key] = ChainParams.fromJSON(value);
          return acc;
        }, {})
        : {},
    };
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    obj.networks = {};
    if (message.networks) {
      Object.entries(message.networks).forEach(([k, v]) => {
        obj.networks[k] = ChainParams.toJSON(v);
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Params>, I>>(object: I): Params {
    const message = createBaseParams();
    message.networks = Object.entries(object.networks ?? {}).reduce<{ [key: string]: ChainParams }>(
      (acc, [key, value]) => {
        if (value !== undefined) {
          acc[key] = ChainParams.fromPartial(value);
        }
        return acc;
      },
      {},
    );
    return message;
  },
};

function createBaseParams_NetworksEntry(): Params_NetworksEntry {
  return { key: "", value: undefined };
}

export const Params_NetworksEntry = {
  encode(message: Params_NetworksEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== undefined) {
      ChainParams.encode(message.value, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Params_NetworksEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseParams_NetworksEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = ChainParams.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Params_NetworksEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? ChainParams.fromJSON(object.value) : undefined,
    };
  },

  toJSON(message: Params_NetworksEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value ? ChainParams.toJSON(message.value) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Params_NetworksEntry>, I>>(object: I): Params_NetworksEntry {
    const message = createBaseParams_NetworksEntry();
    message.key = object.key ?? "";
    message.value = (object.value !== undefined && object.value !== null)
      ? ChainParams.fromPartial(object.value)
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
