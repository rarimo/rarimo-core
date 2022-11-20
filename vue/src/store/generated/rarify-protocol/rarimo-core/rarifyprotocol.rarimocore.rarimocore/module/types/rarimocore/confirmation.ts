/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "rarifyprotocol.rarimocore.rarimocore";

export interface ConfirmationMeta {
  newKeyECDSA: string;
  partyKey: string[];
}

export interface Confirmation {
  /** hex-encoded */
  root: string;
  /** hex-encoded */
  indexes: string[];
  /** hex-encoded */
  signatureECDSA: string;
  creator: string;
  meta: ConfirmationMeta | undefined;
}

const baseConfirmationMeta: object = { newKeyECDSA: "", partyKey: "" };

export const ConfirmationMeta = {
  encode(message: ConfirmationMeta, writer: Writer = Writer.create()): Writer {
    if (message.newKeyECDSA !== "") {
      writer.uint32(10).string(message.newKeyECDSA);
    }
    for (const v of message.partyKey) {
      writer.uint32(18).string(v!);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ConfirmationMeta {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseConfirmationMeta } as ConfirmationMeta;
    message.partyKey = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.newKeyECDSA = reader.string();
          break;
        case 2:
          message.partyKey.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ConfirmationMeta {
    const message = { ...baseConfirmationMeta } as ConfirmationMeta;
    message.partyKey = [];
    if (object.newKeyECDSA !== undefined && object.newKeyECDSA !== null) {
      message.newKeyECDSA = String(object.newKeyECDSA);
    } else {
      message.newKeyECDSA = "";
    }
    if (object.partyKey !== undefined && object.partyKey !== null) {
      for (const e of object.partyKey) {
        message.partyKey.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: ConfirmationMeta): unknown {
    const obj: any = {};
    message.newKeyECDSA !== undefined &&
      (obj.newKeyECDSA = message.newKeyECDSA);
    if (message.partyKey) {
      obj.partyKey = message.partyKey.map((e) => e);
    } else {
      obj.partyKey = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<ConfirmationMeta>): ConfirmationMeta {
    const message = { ...baseConfirmationMeta } as ConfirmationMeta;
    message.partyKey = [];
    if (object.newKeyECDSA !== undefined && object.newKeyECDSA !== null) {
      message.newKeyECDSA = object.newKeyECDSA;
    } else {
      message.newKeyECDSA = "";
    }
    if (object.partyKey !== undefined && object.partyKey !== null) {
      for (const e of object.partyKey) {
        message.partyKey.push(e);
      }
    }
    return message;
  },
};

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
    if (message.meta !== undefined) {
      ConfirmationMeta.encode(message.meta, writer.uint32(42).fork()).ldelim();
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
        case 5:
          message.meta = ConfirmationMeta.decode(reader, reader.uint32());
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
    if (object.meta !== undefined && object.meta !== null) {
      message.meta = ConfirmationMeta.fromJSON(object.meta);
    } else {
      message.meta = undefined;
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
    message.meta !== undefined &&
      (obj.meta = message.meta
        ? ConfirmationMeta.toJSON(message.meta)
        : undefined);
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
    if (object.meta !== undefined && object.meta !== null) {
      message.meta = ConfirmationMeta.fromPartial(object.meta);
    } else {
      message.meta = undefined;
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
