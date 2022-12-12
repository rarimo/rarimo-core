/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "rarimo.rarimocore.rarimocore";

export interface Confirmation {
  /** hex-encoded */
  root: string;
  /** hex-encoded */
  indexes: string[];
  /** hex-encoded */
  signatureECDSA: string;
  creator: string;
}

const baseConfirmation: object = {
  root: "",
  indexes: "",
  signatureECDSA: "",
  creator: "",
};

export const Confirmation = {
  encode(message: Confirmation, writer: Writer = Writer.create()): Writer {
    if (message.root !== "") {
      writer.uint32(10).string(message.root);
    }
    for (const v of message.indexes) {
      writer.uint32(18).string(v!);
    }
    if (message.signatureECDSA !== "") {
      writer.uint32(26).string(message.signatureECDSA);
    }
    if (message.creator !== "") {
      writer.uint32(34).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Confirmation {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseConfirmation } as Confirmation;
    message.indexes = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.root = reader.string();
          break;
        case 2:
          message.indexes.push(reader.string());
          break;
        case 3:
          message.signatureECDSA = reader.string();
          break;
        case 4:
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
    message.indexes = [];
    if (object.root !== undefined && object.root !== null) {
      message.root = String(object.root);
    } else {
      message.root = "";
    }
    if (object.indexes !== undefined && object.indexes !== null) {
      for (const e of object.indexes) {
        message.indexes.push(String(e));
      }
    }
    if (object.signatureECDSA !== undefined && object.signatureECDSA !== null) {
      message.signatureECDSA = String(object.signatureECDSA);
    } else {
      message.signatureECDSA = "";
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
    message.root !== undefined && (obj.root = message.root);
    if (message.indexes) {
      obj.indexes = message.indexes.map((e) => e);
    } else {
      obj.indexes = [];
    }
    message.signatureECDSA !== undefined &&
      (obj.signatureECDSA = message.signatureECDSA);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<Confirmation>): Confirmation {
    const message = { ...baseConfirmation } as Confirmation;
    message.indexes = [];
    if (object.root !== undefined && object.root !== null) {
      message.root = object.root;
    } else {
      message.root = "";
    }
    if (object.indexes !== undefined && object.indexes !== null) {
      for (const e of object.indexes) {
        message.indexes.push(e);
      }
    }
    if (object.signatureECDSA !== undefined && object.signatureECDSA !== null) {
      message.signatureECDSA = object.signatureECDSA;
    } else {
      message.signatureECDSA = "";
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
