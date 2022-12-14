/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Any } from "../google/protobuf/any";

export const protobufPackage = "rarifyprotocol.rarimocore.rarimocore";

export enum opType {
  TRANSFER = 0,
  CHANGE_KEY = 1,
  UNRECOGNIZED = -1,
}

export function opTypeFromJSON(object: any): opType {
  switch (object) {
    case 0:
    case "TRANSFER":
      return opType.TRANSFER;
    case 1:
    case "CHANGE_KEY":
      return opType.CHANGE_KEY;
    case -1:
    case "UNRECOGNIZED":
    default:
      return opType.UNRECOGNIZED;
  }
}

export function opTypeToJSON(object: opType): string {
  switch (object) {
    case opType.TRANSFER:
      return "TRANSFER";
    case opType.CHANGE_KEY:
      return "CHANGE_KEY";
    case opType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface Operation {
  /** Should be in a hex format 0x... */
  index: string;
  operationType: opType;
  /** Corresponding to type details */
  details: Any | undefined;
  signed: boolean;
  creator: string;
  timestamp: number;
}

function createBaseOperation(): Operation {
  return { index: "", operationType: 0, details: undefined, signed: false, creator: "", timestamp: 0 };
}

export const Operation = {
  encode(message: Operation, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.operationType !== 0) {
      writer.uint32(16).int32(message.operationType);
    }
    if (message.details !== undefined) {
      Any.encode(message.details, writer.uint32(26).fork()).ldelim();
    }
    if (message.signed === true) {
      writer.uint32(32).bool(message.signed);
    }
    if (message.creator !== "") {
      writer.uint32(42).string(message.creator);
    }
    if (message.timestamp !== 0) {
      writer.uint32(48).int64(message.timestamp);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Operation {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOperation();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.operationType = reader.int32() as any;
          break;
        case 3:
          message.details = Any.decode(reader, reader.uint32());
          break;
        case 4:
          message.signed = reader.bool();
          break;
        case 5:
          message.creator = reader.string();
          break;
        case 6:
          message.timestamp = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Operation {
    return {
      index: isSet(object.index) ? String(object.index) : "",
      operationType: isSet(object.operationType) ? opTypeFromJSON(object.operationType) : 0,
      details: isSet(object.details) ? Any.fromJSON(object.details) : undefined,
      signed: isSet(object.signed) ? Boolean(object.signed) : false,
      creator: isSet(object.creator) ? String(object.creator) : "",
      timestamp: isSet(object.timestamp) ? Number(object.timestamp) : 0,
    };
  },

  toJSON(message: Operation): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.operationType !== undefined && (obj.operationType = opTypeToJSON(message.operationType));
    message.details !== undefined && (obj.details = message.details ? Any.toJSON(message.details) : undefined);
    message.signed !== undefined && (obj.signed = message.signed);
    message.creator !== undefined && (obj.creator = message.creator);
    message.timestamp !== undefined && (obj.timestamp = Math.round(message.timestamp));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Operation>, I>>(object: I): Operation {
    const message = createBaseOperation();
    message.index = object.index ?? "";
    message.operationType = object.operationType ?? 0;
    message.details = (object.details !== undefined && object.details !== null)
      ? Any.fromPartial(object.details)
      : undefined;
    message.signed = object.signed ?? false;
    message.creator = object.creator ?? "";
    message.timestamp = object.timestamp ?? 0;
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
