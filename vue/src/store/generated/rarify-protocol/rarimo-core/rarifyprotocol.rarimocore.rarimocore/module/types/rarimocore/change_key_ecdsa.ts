/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "rarimo.rarimocore.rarimocore";

export interface ChangeKeyECDSA {
  newKey: string;
  signature: string;
  creator: string;
}

const baseChangeKeyECDSA: object = { newKey: "", signature: "", creator: "" };

export const ChangeKeyECDSA = {
  encode(message: ChangeKeyECDSA, writer: Writer = Writer.create()): Writer {
    if (message.newKey !== "") {
      writer.uint32(10).string(message.newKey);
    }
    if (message.signature !== "") {
      writer.uint32(18).string(message.signature);
    }
    if (message.creator !== "") {
      writer.uint32(26).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ChangeKeyECDSA {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseChangeKeyECDSA } as ChangeKeyECDSA;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.newKey = reader.string();
          break;
        case 2:
          message.signature = reader.string();
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

  fromJSON(object: any): ChangeKeyECDSA {
    const message = { ...baseChangeKeyECDSA } as ChangeKeyECDSA;
    if (object.newKey !== undefined && object.newKey !== null) {
      message.newKey = String(object.newKey);
    } else {
      message.newKey = "";
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

  toJSON(message: ChangeKeyECDSA): unknown {
    const obj: any = {};
    message.newKey !== undefined && (obj.newKey = message.newKey);
    message.signature !== undefined && (obj.signature = message.signature);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<ChangeKeyECDSA>): ChangeKeyECDSA {
    const message = { ...baseChangeKeyECDSA } as ChangeKeyECDSA;
    if (object.newKey !== undefined && object.newKey !== null) {
      message.newKey = object.newKey;
    } else {
      message.newKey = "";
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
