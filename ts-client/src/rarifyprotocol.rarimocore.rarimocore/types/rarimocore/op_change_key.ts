/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "rarifyprotocol.rarimocore.rarimocore";

export interface ChangeKey {
  newKey: string;
}

function createBaseChangeKey(): ChangeKey {
  return { newKey: "" };
}

export const ChangeKey = {
  encode(message: ChangeKey, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.newKey !== "") {
      writer.uint32(10).string(message.newKey);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ChangeKey {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseChangeKey();
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
    return { newKey: isSet(object.newKey) ? String(object.newKey) : "" };
  },

  toJSON(message: ChangeKey): unknown {
    const obj: any = {};
    message.newKey !== undefined && (obj.newKey = message.newKey);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ChangeKey>, I>>(object: I): ChangeKey {
    const message = createBaseChangeKey();
    message.newKey = object.newKey ?? "";
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
