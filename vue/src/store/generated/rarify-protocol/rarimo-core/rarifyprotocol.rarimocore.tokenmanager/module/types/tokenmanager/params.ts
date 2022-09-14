/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "rarifyprotocol.rarimocore.tokenmanager";

export interface ChainParams {
  contract: string;
  types: { [key: number]: number };
}

export interface ChainParams_TypesEntry {
  key: number;
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

const baseChainParams: object = { contract: "" };

export const ChainParams = {
  encode(message: ChainParams, writer: Writer = Writer.create()): Writer {
    if (message.contract !== "") {
      writer.uint32(10).string(message.contract);
    }
    Object.entries(message.types).forEach(([key, value]) => {
      ChainParams_TypesEntry.encode(
        { key: key as any, value },
        writer.uint32(18).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ChainParams {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseChainParams } as ChainParams;
    message.types = {};
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
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ChainParams {
    const message = { ...baseChainParams } as ChainParams;
    message.types = {};
    if (object.contract !== undefined && object.contract !== null) {
      message.contract = String(object.contract);
    } else {
      message.contract = "";
    }
    if (object.types !== undefined && object.types !== null) {
      Object.entries(object.types).forEach(([key, value]) => {
        message.types[Number(key)] = Number(value);
      });
    }
    return message;
  },

  toJSON(message: ChainParams): unknown {
    const obj: any = {};
    message.contract !== undefined && (obj.contract = message.contract);
    obj.types = {};
    if (message.types) {
      Object.entries(message.types).forEach(([k, v]) => {
        obj.types[k] = v;
      });
    }
    return obj;
  },

  fromPartial(object: DeepPartial<ChainParams>): ChainParams {
    const message = { ...baseChainParams } as ChainParams;
    message.types = {};
    if (object.contract !== undefined && object.contract !== null) {
      message.contract = object.contract;
    } else {
      message.contract = "";
    }
    if (object.types !== undefined && object.types !== null) {
      Object.entries(object.types).forEach(([key, value]) => {
        if (value !== undefined) {
          message.types[Number(key)] = Number(value);
        }
      });
    }
    return message;
  },
};

const baseChainParams_TypesEntry: object = { key: 0, value: 0 };

export const ChainParams_TypesEntry = {
  encode(
    message: ChainParams_TypesEntry,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.key !== 0) {
      writer.uint32(8).uint32(message.key);
    }
    if (message.value !== 0) {
      writer.uint32(16).uint32(message.value);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ChainParams_TypesEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseChainParams_TypesEntry } as ChainParams_TypesEntry;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.uint32();
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
    const message = { ...baseChainParams_TypesEntry } as ChainParams_TypesEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = Number(object.key);
    } else {
      message.key = 0;
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = Number(object.value);
    } else {
      message.value = 0;
    }
    return message;
  },

  toJSON(message: ChainParams_TypesEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial(
    object: DeepPartial<ChainParams_TypesEntry>
  ): ChainParams_TypesEntry {
    const message = { ...baseChainParams_TypesEntry } as ChainParams_TypesEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = 0;
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = 0;
    }
    return message;
  },
};

const baseParams: object = {};

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    Object.entries(message.networks).forEach(([key, value]) => {
      Params_NetworksEntry.encode(
        { key: key as any, value },
        writer.uint32(10).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseParams } as Params;
    message.networks = {};
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
    const message = { ...baseParams } as Params;
    message.networks = {};
    if (object.networks !== undefined && object.networks !== null) {
      Object.entries(object.networks).forEach(([key, value]) => {
        message.networks[key] = ChainParams.fromJSON(value);
      });
    }
    return message;
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

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    message.networks = {};
    if (object.networks !== undefined && object.networks !== null) {
      Object.entries(object.networks).forEach(([key, value]) => {
        if (value !== undefined) {
          message.networks[key] = ChainParams.fromPartial(value);
        }
      });
    }
    return message;
  },
};

const baseParams_NetworksEntry: object = { key: "" };

export const Params_NetworksEntry = {
  encode(
    message: Params_NetworksEntry,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== undefined) {
      ChainParams.encode(message.value, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Params_NetworksEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseParams_NetworksEntry } as Params_NetworksEntry;
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
    const message = { ...baseParams_NetworksEntry } as Params_NetworksEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key);
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = ChainParams.fromJSON(object.value);
    } else {
      message.value = undefined;
    }
    return message;
  },

  toJSON(message: Params_NetworksEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined &&
      (obj.value = message.value
        ? ChainParams.toJSON(message.value)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<Params_NetworksEntry>): Params_NetworksEntry {
    const message = { ...baseParams_NetworksEntry } as Params_NetworksEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = ChainParams.fromPartial(object.value);
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
