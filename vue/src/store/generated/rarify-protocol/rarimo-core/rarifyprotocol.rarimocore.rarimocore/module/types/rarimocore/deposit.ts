/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "rarifyprotocol.rarimocore.rarimocore";

export interface Deposit {
  /** hex-encoded keccak256 hash for tx||event||chain strings' bytes */
  index: string;
  tx: string;
  eventId: string;
  fromChain: string;
  toChain: string;
  /** hex-encoded */
  receiver: string;
  /** dec-encoded */
  amount: string;
  creator: string;
  signed: boolean;
  tokenIndex: string;
  timestamp: number;
}

const baseDeposit: object = {
  index: "",
  tx: "",
  eventId: "",
  fromChain: "",
  toChain: "",
  receiver: "",
  amount: "",
  creator: "",
  signed: false,
  tokenIndex: "",
  timestamp: 0,
};

export const Deposit = {
  encode(message: Deposit, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
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
    if (message.creator !== "") {
      writer.uint32(66).string(message.creator);
    }
    if (message.signed === true) {
      writer.uint32(72).bool(message.signed);
    }
    if (message.tokenIndex !== "") {
      writer.uint32(82).string(message.tokenIndex);
    }
    if (message.timestamp !== 0) {
      writer.uint32(88).uint64(message.timestamp);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Deposit {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseDeposit } as Deposit;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
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
          message.creator = reader.string();
          break;
        case 9:
          message.signed = reader.bool();
          break;
        case 10:
          message.tokenIndex = reader.string();
          break;
        case 11:
          message.timestamp = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Deposit {
    const message = { ...baseDeposit } as Deposit;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
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
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.signed !== undefined && object.signed !== null) {
      message.signed = Boolean(object.signed);
    } else {
      message.signed = false;
    }
    if (object.tokenIndex !== undefined && object.tokenIndex !== null) {
      message.tokenIndex = String(object.tokenIndex);
    } else {
      message.tokenIndex = "";
    }
    if (object.timestamp !== undefined && object.timestamp !== null) {
      message.timestamp = Number(object.timestamp);
    } else {
      message.timestamp = 0;
    }
    return message;
  },

  toJSON(message: Deposit): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.tx !== undefined && (obj.tx = message.tx);
    message.eventId !== undefined && (obj.eventId = message.eventId);
    message.fromChain !== undefined && (obj.fromChain = message.fromChain);
    message.toChain !== undefined && (obj.toChain = message.toChain);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.amount !== undefined && (obj.amount = message.amount);
    message.creator !== undefined && (obj.creator = message.creator);
    message.signed !== undefined && (obj.signed = message.signed);
    message.tokenIndex !== undefined && (obj.tokenIndex = message.tokenIndex);
    message.timestamp !== undefined && (obj.timestamp = message.timestamp);
    return obj;
  },

  fromPartial(object: DeepPartial<Deposit>): Deposit {
    const message = { ...baseDeposit } as Deposit;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
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
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.signed !== undefined && object.signed !== null) {
      message.signed = object.signed;
    } else {
      message.signed = false;
    }
    if (object.tokenIndex !== undefined && object.tokenIndex !== null) {
      message.tokenIndex = object.tokenIndex;
    } else {
      message.tokenIndex = "";
    }
    if (object.timestamp !== undefined && object.timestamp !== null) {
      message.timestamp = object.timestamp;
    } else {
      message.timestamp = 0;
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
