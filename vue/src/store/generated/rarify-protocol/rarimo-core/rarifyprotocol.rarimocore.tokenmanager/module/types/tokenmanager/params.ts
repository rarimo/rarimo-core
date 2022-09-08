/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "rarifyprotocol.rarimocore.tokenmanager";

export interface Network {
  contract: string;
  saver: string;
}

/** Params defines the parameters for the module. */
export interface Params {
  networks: { [key: string]: Network };
}

export interface Params_NetworksEntry {
  key: string;
  value: Network | undefined;
}

const baseNetwork: object = { contract: "", saver: "" };

export const Network = {
  encode(message: Network, writer: Writer = Writer.create()): Writer {
    if (message.contract !== "") {
      writer.uint32(10).string(message.contract);
    }
    if (message.saver !== "") {
      writer.uint32(18).string(message.saver);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Network {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseNetwork } as Network;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.contract = reader.string();
          break;
        case 2:
          message.saver = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Network {
    const message = { ...baseNetwork } as Network;
    if (object.contract !== undefined && object.contract !== null) {
      message.contract = String(object.contract);
    } else {
      message.contract = "";
    }
    if (object.saver !== undefined && object.saver !== null) {
      message.saver = String(object.saver);
    } else {
      message.saver = "";
    }
    return message;
  },

  toJSON(message: Network): unknown {
    const obj: any = {};
    message.contract !== undefined && (obj.contract = message.contract);
    message.saver !== undefined && (obj.saver = message.saver);
    return obj;
  },

  fromPartial(object: DeepPartial<Network>): Network {
    const message = { ...baseNetwork } as Network;
    if (object.contract !== undefined && object.contract !== null) {
      message.contract = object.contract;
    } else {
      message.contract = "";
    }
    if (object.saver !== undefined && object.saver !== null) {
      message.saver = object.saver;
    } else {
      message.saver = "";
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
        message.networks[key] = Network.fromJSON(value);
      });
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    obj.networks = {};
    if (message.networks) {
      Object.entries(message.networks).forEach(([k, v]) => {
        obj.networks[k] = Network.toJSON(v);
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
          message.networks[key] = Network.fromPartial(value);
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
      Network.encode(message.value, writer.uint32(18).fork()).ldelim();
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
          message.value = Network.decode(reader, reader.uint32());
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
      message.value = Network.fromJSON(object.value);
    } else {
      message.value = undefined;
    }
    return message;
  },

  toJSON(message: Params_NetworksEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined &&
      (obj.value = message.value ? Network.toJSON(message.value) : undefined);
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
      message.value = Network.fromPartial(object.value);
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
