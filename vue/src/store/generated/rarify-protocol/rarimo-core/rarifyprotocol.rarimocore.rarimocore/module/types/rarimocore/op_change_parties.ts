/* eslint-disable */
import { Party } from "../rarimocore/params";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "rarifyprotocol.rarimocore.rarimocore";

export enum ChangeType {
  REMOVE = 0,
  ADD = 1,
  UNRECOGNIZED = -1,
}

export function changeTypeFromJSON(object: any): ChangeType {
  switch (object) {
    case 0:
    case "REMOVE":
      return ChangeType.REMOVE;
    case 1:
    case "ADD":
      return ChangeType.ADD;
    case -1:
    case "UNRECOGNIZED":
    default:
      return ChangeType.UNRECOGNIZED;
  }
}

export function changeTypeToJSON(object: ChangeType): string {
  switch (object) {
    case ChangeType.REMOVE:
      return "REMOVE";
    case ChangeType.ADD:
      return "ADD";
    default:
      return "UNKNOWN";
  }
}

export interface ChangeParties {
  currentSet: Party[];
  newSet: Party[];
  type: ChangeType;
}

const baseChangeParties: object = { type: 0 };

export const ChangeParties = {
  encode(message: ChangeParties, writer: Writer = Writer.create()): Writer {
    for (const v of message.currentSet) {
      Party.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.newSet) {
      Party.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.type !== 0) {
      writer.uint32(24).int32(message.type);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ChangeParties {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseChangeParties } as ChangeParties;
    message.currentSet = [];
    message.newSet = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.currentSet.push(Party.decode(reader, reader.uint32()));
          break;
        case 2:
          message.newSet.push(Party.decode(reader, reader.uint32()));
          break;
        case 3:
          message.type = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ChangeParties {
    const message = { ...baseChangeParties } as ChangeParties;
    message.currentSet = [];
    message.newSet = [];
    if (object.currentSet !== undefined && object.currentSet !== null) {
      for (const e of object.currentSet) {
        message.currentSet.push(Party.fromJSON(e));
      }
    }
    if (object.newSet !== undefined && object.newSet !== null) {
      for (const e of object.newSet) {
        message.newSet.push(Party.fromJSON(e));
      }
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = changeTypeFromJSON(object.type);
    } else {
      message.type = 0;
    }
    return message;
  },

  toJSON(message: ChangeParties): unknown {
    const obj: any = {};
    if (message.currentSet) {
      obj.currentSet = message.currentSet.map((e) =>
        e ? Party.toJSON(e) : undefined
      );
    } else {
      obj.currentSet = [];
    }
    if (message.newSet) {
      obj.newSet = message.newSet.map((e) => (e ? Party.toJSON(e) : undefined));
    } else {
      obj.newSet = [];
    }
    message.type !== undefined && (obj.type = changeTypeToJSON(message.type));
    return obj;
  },

  fromPartial(object: DeepPartial<ChangeParties>): ChangeParties {
    const message = { ...baseChangeParties } as ChangeParties;
    message.currentSet = [];
    message.newSet = [];
    if (object.currentSet !== undefined && object.currentSet !== null) {
      for (const e of object.currentSet) {
        message.currentSet.push(Party.fromPartial(e));
      }
    }
    if (object.newSet !== undefined && object.newSet !== null) {
      for (const e of object.newSet) {
        message.newSet.push(Party.fromPartial(e));
      }
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = object.type;
    } else {
      message.type = 0;
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
