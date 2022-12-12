/* eslint-disable */
import { Party } from "../rarimocore/params";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "rarimo.rarimocore.rarimocore";

export interface RemoveParty {
  currentSet: Party[];
  partyIndex: number;
}

const baseRemoveParty: object = { partyIndex: 0 };

export const RemoveParty = {
  encode(message: RemoveParty, writer: Writer = Writer.create()): Writer {
    for (const v of message.currentSet) {
      Party.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.partyIndex !== 0) {
      writer.uint32(16).uint32(message.partyIndex);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): RemoveParty {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseRemoveParty } as RemoveParty;
    message.currentSet = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.currentSet.push(Party.decode(reader, reader.uint32()));
          break;
        case 2:
          message.partyIndex = reader.uint32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RemoveParty {
    const message = { ...baseRemoveParty } as RemoveParty;
    message.currentSet = [];
    if (object.currentSet !== undefined && object.currentSet !== null) {
      for (const e of object.currentSet) {
        message.currentSet.push(Party.fromJSON(e));
      }
    }
    if (object.partyIndex !== undefined && object.partyIndex !== null) {
      message.partyIndex = Number(object.partyIndex);
    } else {
      message.partyIndex = 0;
    }
    return message;
  },

  toJSON(message: RemoveParty): unknown {
    const obj: any = {};
    if (message.currentSet) {
      obj.currentSet = message.currentSet.map((e) =>
        e ? Party.toJSON(e) : undefined
      );
    } else {
      obj.currentSet = [];
    }
    message.partyIndex !== undefined && (obj.partyIndex = message.partyIndex);
    return obj;
  },

  fromPartial(object: DeepPartial<RemoveParty>): RemoveParty {
    const message = { ...baseRemoveParty } as RemoveParty;
    message.currentSet = [];
    if (object.currentSet !== undefined && object.currentSet !== null) {
      for (const e of object.currentSet) {
        message.currentSet.push(Party.fromPartial(e));
      }
    }
    if (object.partyIndex !== undefined && object.partyIndex !== null) {
      message.partyIndex = object.partyIndex;
    } else {
      message.partyIndex = 0;
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
