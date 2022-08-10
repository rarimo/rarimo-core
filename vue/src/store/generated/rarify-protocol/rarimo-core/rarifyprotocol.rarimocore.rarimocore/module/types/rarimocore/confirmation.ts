/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "rarifyprotocol.rarimocore.rarimocore";

export interface Confirmation {
  height: string;
  root: string;
  hashes: string[];
  signature: string;
  creator: string;
}

const baseConfirmation: object = {
  height: "",
  root: "",
  hashes: "",
  signature: "",
  creator: "",
};

export const Confirmation = {
  encode(message: Confirmation, writer: Writer = Writer.create()): Writer {
    if (message.height !== "") {
      writer.uint32(10).string(message.height);
    }
    if (message.root !== "") {
      writer.uint32(18).string(message.root);
    }
    for (const v of message.hashes) {
      writer.uint32(26).string(v!);
    }
    if (message.signature !== "") {
      writer.uint32(34).string(message.signature);
    }
    if (message.creator !== "") {
      writer.uint32(42).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Confirmation {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseConfirmation } as Confirmation;
    message.hashes = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.height = reader.string();
          break;
        case 2:
          message.root = reader.string();
          break;
        case 3:
          message.hashes.push(reader.string());
          break;
        case 4:
          message.signature = reader.string();
          break;
        case 5:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Confirmation {
    const message = { ...baseConfirmation } as Confirmation;
    message.hashes = [];
    if (object.height !== undefined && object.height !== null) {
      message.height = String(object.height);
    } else {
      message.height = "";
    }
    if (object.root !== undefined && object.root !== null) {
      message.root = String(object.root);
    } else {
      message.root = "";
    }
    if (object.hashes !== undefined && object.hashes !== null) {
      for (const e of object.hashes) {
        message.hashes.push(String(e));
      }
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = String(object.signature);
    } else {
      message.signature = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: Confirmation): unknown {
    const obj: any = {};
    message.height !== undefined && (obj.height = message.height);
    message.root !== undefined && (obj.root = message.root);
    if (message.hashes) {
      obj.hashes = message.hashes.map((e) => e);
    } else {
      obj.hashes = [];
    }
    message.signature !== undefined && (obj.signature = message.signature);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<Confirmation>): Confirmation {
    const message = { ...baseConfirmation } as Confirmation;
    message.hashes = [];
    if (object.height !== undefined && object.height !== null) {
      message.height = object.height;
    } else {
      message.height = "";
    }
    if (object.root !== undefined && object.root !== null) {
      message.root = object.root;
    } else {
      message.root = "";
    }
    if (object.hashes !== undefined && object.hashes !== null) {
      for (const e of object.hashes) {
        message.hashes.push(e);
      }
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = object.signature;
    } else {
      message.signature = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
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
