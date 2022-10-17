/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "rarifyprotocol.rarimocore.rarimocore";

export interface ChangeKey {
  newKey: string;
}

const baseChangeKey: object = { newKey: "" };

export const ChangeKey = {
  encode(message: ChangeKey, writer: Writer = Writer.create()): Writer {
    if (message.newKey !== "") {
      writer.uint32(10).string(message.newKey);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ChangeKey {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseChangeKey } as ChangeKey;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.newKey = reader.string();
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
    if (object.newKey !== undefined && object.newKey !== null) {
      message.newKey = String(object.newKey);
    } else {
      message.newKey = "";
    }
    return message;
  },

  toJSON(message: ChangeKey): unknown {
    const obj: any = {};
    message.newKey !== undefined && (obj.newKey = message.newKey);
    return obj;
  },

  fromPartial(object: DeepPartial<ChangeKey>): ChangeKey {
    const message = { ...baseChangeKey } as ChangeKey;
    if (object.newKey !== undefined && object.newKey !== null) {
      message.newKey = object.newKey;
    } else {
      message.newKey = "";
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
