/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "rarifyprotocol.rarimocore.rarimocore";

export interface Confirmation {
  /** hex-encoded */
  root: string;
  /** hex-encoded */
  indexes: string[];
  /** hex-encoded */
  signatureECDSA: string;
  creator: string;
}

function createBaseConfirmation(): Confirmation {
  return { root: "", indexes: [], signatureECDSA: "", creator: "" };
}

export const Confirmation = {
  encode(message: Confirmation, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.root !== "") {
      writer.uint32(10).string(message.root);
    }
    for (const v of message.indexes) {
      writer.uint32(18).string(v!);
    }
    if (message.signatureECDSA !== "") {
      writer.uint32(26).string(message.signatureECDSA);
    }
    if (message.creator !== "") {
      writer.uint32(34).string(message.creator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Confirmation {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseConfirmation();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.root = reader.string();
          break;
        case 2:
          message.indexes.push(reader.string());
          break;
        case 3:
          message.signatureECDSA = reader.string();
          break;
        case 4:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Confirmation {
    return {
      root: isSet(object.root) ? String(object.root) : "",
      indexes: Array.isArray(object?.indexes) ? object.indexes.map((e: any) => String(e)) : [],
      signatureECDSA: isSet(object.signatureECDSA) ? String(object.signatureECDSA) : "",
      creator: isSet(object.creator) ? String(object.creator) : "",
    };
  },

  toJSON(message: Confirmation): unknown {
    const obj: any = {};
    message.root !== undefined && (obj.root = message.root);
    if (message.indexes) {
      obj.indexes = message.indexes.map((e) => e);
    } else {
      obj.indexes = [];
    }
    message.signatureECDSA !== undefined && (obj.signatureECDSA = message.signatureECDSA);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Confirmation>, I>>(object: I): Confirmation {
    const message = createBaseConfirmation();
    message.root = object.root ?? "";
    message.indexes = object.indexes?.map((e) => e) || [];
    message.signatureECDSA = object.signatureECDSA ?? "";
    message.creator = object.creator ?? "";
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
