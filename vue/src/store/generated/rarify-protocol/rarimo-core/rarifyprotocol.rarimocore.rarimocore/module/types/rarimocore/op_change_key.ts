/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Party } from "../rarimocore/params";

export const protobufPackage = "rarimo.rarimocore.rarimocore";

export interface ChangeKey {
  newKey: string;
  parties: Party[];
  threshold: number;
}

const baseChangeKey: object = { newKey: "", threshold: 0 };

export const ChangeKey = {
  encode(message: ChangeKey, writer: Writer = Writer.create()): Writer {
    if (message.newKey !== "") {
      writer.uint32(10).string(message.newKey);
    }
    for (const v of message.parties) {
      Party.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.threshold !== 0) {
      writer.uint32(24).uint64(message.threshold);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ChangeKey {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseChangeKey } as ChangeKey;
    message.parties = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.newKey = reader.string();
          break;
        case 2:
          message.parties.push(Party.decode(reader, reader.uint32()));
          break;
        case 3:
          message.threshold = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ChangeKey {
    const message = { ...baseChangeKey } as ChangeKey;
    message.parties = [];
    if (object.newKey !== undefined && object.newKey !== null) {
      message.newKey = String(object.newKey);
    } else {
      message.newKey = "";
    }
    if (object.parties !== undefined && object.parties !== null) {
      for (const e of object.parties) {
        message.parties.push(Party.fromJSON(e));
      }
    }
    if (object.threshold !== undefined && object.threshold !== null) {
      message.threshold = Number(object.threshold);
    } else {
      message.threshold = 0;
    }
    return message;
  },

  toJSON(message: ChangeKey): unknown {
    const obj: any = {};
    message.newKey !== undefined && (obj.newKey = message.newKey);
    if (message.parties) {
      obj.parties = message.parties.map((e) =>
        e ? Party.toJSON(e) : undefined
      );
    } else {
      obj.parties = [];
    }
    message.threshold !== undefined && (obj.threshold = message.threshold);
    return obj;
  },

  fromPartial(object: DeepPartial<ChangeKey>): ChangeKey {
    const message = { ...baseChangeKey } as ChangeKey;
    message.parties = [];
    if (object.newKey !== undefined && object.newKey !== null) {
      message.newKey = object.newKey;
    } else {
      message.newKey = "";
    }
    if (object.parties !== undefined && object.parties !== null) {
      for (const e of object.parties) {
        message.parties.push(Party.fromPartial(e));
      }
    }
    if (object.threshold !== undefined && object.threshold !== null) {
      message.threshold = object.threshold;
    } else {
      message.threshold = 0;
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
