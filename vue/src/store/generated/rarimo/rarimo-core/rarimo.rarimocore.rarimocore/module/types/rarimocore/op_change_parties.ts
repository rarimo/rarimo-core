/* eslint-disable */
import { Party } from "../rarimocore/params";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "rarimo.rarimocore.rarimocore";

export interface ChangeParties {
  parties: Party[];
  newPublicKey: string;
  signature: string;
}

const baseChangeParties: object = { newPublicKey: "", signature: "" };

export const ChangeParties = {
  encode(message: ChangeParties, writer: Writer = Writer.create()): Writer {
    for (const v of message.parties) {
      Party.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.newPublicKey !== "") {
      writer.uint32(18).string(message.newPublicKey);
    }
    if (message.signature !== "") {
      writer.uint32(26).string(message.signature);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ChangeParties {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseChangeParties } as ChangeParties;
    message.parties = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.parties.push(Party.decode(reader, reader.uint32()));
          break;
        case 2:
          message.newPublicKey = reader.string();
          break;
        case 3:
          message.signature = reader.string();
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
    message.parties = [];
    if (object.parties !== undefined && object.parties !== null) {
      for (const e of object.parties) {
        message.parties.push(Party.fromJSON(e));
      }
    }
    if (object.newPublicKey !== undefined && object.newPublicKey !== null) {
      message.newPublicKey = String(object.newPublicKey);
    } else {
      message.newPublicKey = "";
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = String(object.signature);
    } else {
      message.signature = "";
    }
    return message;
  },

  toJSON(message: ChangeParties): unknown {
    const obj: any = {};
    if (message.parties) {
      obj.parties = message.parties.map((e) =>
        e ? Party.toJSON(e) : undefined
      );
    } else {
      obj.parties = [];
    }
    message.newPublicKey !== undefined &&
      (obj.newPublicKey = message.newPublicKey);
    message.signature !== undefined && (obj.signature = message.signature);
    return obj;
  },

  fromPartial(object: DeepPartial<ChangeParties>): ChangeParties {
    const message = { ...baseChangeParties } as ChangeParties;
    message.parties = [];
    if (object.parties !== undefined && object.parties !== null) {
      for (const e of object.parties) {
        message.parties.push(Party.fromPartial(e));
      }
    }
    if (object.newPublicKey !== undefined && object.newPublicKey !== null) {
      message.newPublicKey = object.newPublicKey;
    } else {
      message.newPublicKey = "";
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = object.signature;
    } else {
      message.signature = "";
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
