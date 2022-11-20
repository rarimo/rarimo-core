/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Any } from "../google/protobuf/any";

export const protobufPackage = "rarifyprotocol.rarimocore.rarimocore";

export enum op_type {
  TRANSFER = 0,
  CHANGE_PARTIES = 1,
  UNRECOGNIZED = -1,
}

export function op_typeFromJSON(object: any): op_type {
  switch (object) {
    case 0:
    case "TRANSFER":
      return op_type.TRANSFER;
    case 1:
    case "CHANGE_PARTIES":
      return op_type.CHANGE_PARTIES;
    case -1:
    case "UNRECOGNIZED":
    default:
      return op_type.UNRECOGNIZED;
  }
}

export function op_typeToJSON(object: op_type): string {
  switch (object) {
    case op_type.TRANSFER:
      return "TRANSFER";
    case op_type.CHANGE_PARTIES:
      return "CHANGE_PARTIES";
    default:
      return "UNKNOWN";
  }
}

export interface Operation {
  /** Should be in a hex format 0x... */
  index: string;
  operationType: op_type;
  /** Corresponding to type details */
  details: Any | undefined;
  signed: boolean;
  creator: string;
  timestamp: number;
}

const baseOperation: object = {
  index: "",
  operationType: 0,
  signed: false,
  creator: "",
  timestamp: 0,
};

export const Operation = {
  encode(message: Operation, writer: Writer = Writer.create()): Writer {
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

  decode(input: Reader | Uint8Array, length?: number): Operation {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseOperation } as Operation;
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
    const message = { ...baseOperation } as Operation;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.operationType !== undefined && object.operationType !== null) {
      message.operationType = op_typeFromJSON(object.operationType);
    } else {
      message.operationType = 0;
    }
    if (object.details !== undefined && object.details !== null) {
      message.details = Any.fromJSON(object.details);
    } else {
      message.details = undefined;
    }
    if (object.signed !== undefined && object.signed !== null) {
      message.signed = Boolean(object.signed);
    } else {
      message.signed = false;
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.timestamp !== undefined && object.timestamp !== null) {
      message.timestamp = Number(object.timestamp);
    } else {
      message.timestamp = 0;
    }
    return message;
  },

  toJSON(message: Operation): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.operationType !== undefined &&
      (obj.operationType = op_typeToJSON(message.operationType));
    message.details !== undefined &&
      (obj.details = message.details ? Any.toJSON(message.details) : undefined);
    message.signed !== undefined && (obj.signed = message.signed);
    message.creator !== undefined && (obj.creator = message.creator);
    message.timestamp !== undefined && (obj.timestamp = message.timestamp);
    return obj;
  },

  fromPartial(object: DeepPartial<Operation>): Operation {
    const message = { ...baseOperation } as Operation;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.operationType !== undefined && object.operationType !== null) {
      message.operationType = object.operationType;
    } else {
      message.operationType = 0;
    }
    if (object.details !== undefined && object.details !== null) {
      message.details = Any.fromPartial(object.details);
    } else {
      message.details = undefined;
    }
    if (object.signed !== undefined && object.signed !== null) {
      message.signed = object.signed;
    } else {
      message.signed = false;
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
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
