/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "rarimo.rarimocore.rarimocore";

export interface Transfer {
  /** hex-encoded keccak256 hash for tx||event||chain strings' bytes */
  origin: string;
  /** Deposit transaction data */
  tx: string;
  eventId: string;
  fromChain: string;
  toChain: string;
  /** hex-encoded */
  receiver: string;
  /** dec-encoded */
  amount: string;
  /** hex-encoded */
  bundleData: string;
  /** hex-encoded */
  bundleSalt: string;
  tokenIndex: string;
}

const baseTransfer: object = {
  origin: "",
  tx: "",
  eventId: "",
  fromChain: "",
  toChain: "",
  receiver: "",
  amount: "",
  bundleData: "",
  bundleSalt: "",
  tokenIndex: "",
};

export const Transfer = {
  encode(message: Transfer, writer: Writer = Writer.create()): Writer {
    if (message.origin !== "") {
      writer.uint32(10).string(message.origin);
    }
    if (message.tx !== "") {
      writer.uint32(18).string(message.tx);
    }
    if (message.eventId !== "") {
      writer.uint32(26).string(message.eventId);
    }
    if (message.fromChain !== "") {
      writer.uint32(34).string(message.fromChain);
    }
    if (message.toChain !== "") {
      writer.uint32(42).string(message.toChain);
    }
    if (message.receiver !== "") {
      writer.uint32(50).string(message.receiver);
    }
    if (message.amount !== "") {
      writer.uint32(58).string(message.amount);
    }
    if (message.bundleData !== "") {
      writer.uint32(66).string(message.bundleData);
    }
    if (message.bundleSalt !== "") {
      writer.uint32(74).string(message.bundleSalt);
    }
    if (message.tokenIndex !== "") {
      writer.uint32(98).string(message.tokenIndex);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Transfer {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseTransfer } as Transfer;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.origin = reader.string();
          break;
        case 2:
          message.tx = reader.string();
          break;
        case 3:
          message.eventId = reader.string();
          break;
        case 4:
          message.fromChain = reader.string();
          break;
        case 5:
          message.toChain = reader.string();
          break;
        case 6:
          message.receiver = reader.string();
          break;
        case 7:
          message.amount = reader.string();
          break;
        case 8:
          message.bundleData = reader.string();
          break;
        case 9:
          message.bundleSalt = reader.string();
          break;
        case 12:
          message.tokenIndex = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Transfer {
    const message = { ...baseTransfer } as Transfer;
    if (object.origin !== undefined && object.origin !== null) {
      message.origin = String(object.origin);
    } else {
      message.origin = "";
    }
    if (object.tx !== undefined && object.tx !== null) {
      message.tx = String(object.tx);
    } else {
      message.tx = "";
    }
    if (object.eventId !== undefined && object.eventId !== null) {
      message.eventId = String(object.eventId);
    } else {
      message.eventId = "";
    }
    if (object.fromChain !== undefined && object.fromChain !== null) {
      message.fromChain = String(object.fromChain);
    } else {
      message.fromChain = "";
    }
    if (object.toChain !== undefined && object.toChain !== null) {
      message.toChain = String(object.toChain);
    } else {
      message.toChain = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = String(object.receiver);
    } else {
      message.receiver = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    if (object.bundleData !== undefined && object.bundleData !== null) {
      message.bundleData = String(object.bundleData);
    } else {
      message.bundleData = "";
    }
    if (object.bundleSalt !== undefined && object.bundleSalt !== null) {
      message.bundleSalt = String(object.bundleSalt);
    } else {
      message.bundleSalt = "";
    }
    if (object.tokenIndex !== undefined && object.tokenIndex !== null) {
      message.tokenIndex = String(object.tokenIndex);
    } else {
      message.tokenIndex = "";
    }
    return message;
  },

  toJSON(message: Transfer): unknown {
    const obj: any = {};
    message.origin !== undefined && (obj.origin = message.origin);
    message.tx !== undefined && (obj.tx = message.tx);
    message.eventId !== undefined && (obj.eventId = message.eventId);
    message.fromChain !== undefined && (obj.fromChain = message.fromChain);
    message.toChain !== undefined && (obj.toChain = message.toChain);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.amount !== undefined && (obj.amount = message.amount);
    message.bundleData !== undefined && (obj.bundleData = message.bundleData);
    message.bundleSalt !== undefined && (obj.bundleSalt = message.bundleSalt);
    message.tokenIndex !== undefined && (obj.tokenIndex = message.tokenIndex);
    return obj;
  },

  fromPartial(object: DeepPartial<Transfer>): Transfer {
    const message = { ...baseTransfer } as Transfer;
    if (object.origin !== undefined && object.origin !== null) {
      message.origin = object.origin;
    } else {
      message.origin = "";
    }
    if (object.tx !== undefined && object.tx !== null) {
      message.tx = object.tx;
    } else {
      message.tx = "";
    }
    if (object.eventId !== undefined && object.eventId !== null) {
      message.eventId = object.eventId;
    } else {
      message.eventId = "";
    }
    if (object.fromChain !== undefined && object.fromChain !== null) {
      message.fromChain = object.fromChain;
    } else {
      message.fromChain = "";
    }
    if (object.toChain !== undefined && object.toChain !== null) {
      message.toChain = object.toChain;
    } else {
      message.toChain = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = object.receiver;
    } else {
      message.receiver = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    if (object.bundleData !== undefined && object.bundleData !== null) {
      message.bundleData = object.bundleData;
    } else {
      message.bundleData = "";
    }
    if (object.bundleSalt !== undefined && object.bundleSalt !== null) {
      message.bundleSalt = object.bundleSalt;
    } else {
      message.bundleSalt = "";
    }
    if (object.tokenIndex !== undefined && object.tokenIndex !== null) {
      message.tokenIndex = object.tokenIndex;
    } else {
      message.tokenIndex = "";
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
