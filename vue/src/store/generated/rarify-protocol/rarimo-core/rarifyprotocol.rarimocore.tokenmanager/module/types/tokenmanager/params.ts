/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "rarifyprotocol.rarimocore.tokenmanager";

/** Params defines the parameters for the module. */
export interface Params {
  networks: { [key: string]: string };
}

export interface Params_NetworksEntry {
  key: string;
  value: string;
}

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
        message.networks[key] = String(value);
      });
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    obj.networks = {};
    if (message.networks) {
      Object.entries(message.networks).forEach(([k, v]) => {
        obj.networks[k] = v;
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
          message.networks[key] = String(value);
        }
      });
    }
    return message;
  },
};

const baseParams_NetworksEntry: object = { key: "", value: "" };

export const Params_NetworksEntry = {
  encode(
    message: Params_NetworksEntry,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
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
          message.value = reader.string();
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
      message.value = String(object.value);
    } else {
      message.value = "";
    }
    return message;
  },

  toJSON(message: Params_NetworksEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
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
      message.value = object.value;
    } else {
      message.value = "";
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
