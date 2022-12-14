/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "rarifyprotocol.rarimocore.rarimocore";

export interface Transfer {
  /** hex-encoded keccak256 hash for tx||event||chain strings' bytes */
  origin: string;
  /** Deposit transaction data */
  tx: string;
  eventId: string;
  fromChain: string;
  toChain: string;
  /** hex-encoded */
  receiver: string;
  /** dec-encoded */
  amount: string;
  /** hex-encoded */
  bundleData: string;
  /** hex-encoded */
  bundleSalt: string;
  tokenIndex: string;
}

function createBaseTransfer(): Transfer {
  return {
    origin: "",
    tx: "",
    eventId: "",
    fromChain: "",
    toChain: "",
    receiver: "",
    amount: "",
    bundleData: "",
    bundleSalt: "",
    tokenIndex: "",
  };
}

export const Transfer = {
  encode(message: Transfer, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.origin !== "") {
      writer.uint32(10).string(message.origin);
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
    if (message.bundleData !== "") {
      writer.uint32(66).string(message.bundleData);
    }
    if (message.bundleSalt !== "") {
      writer.uint32(74).string(message.bundleSalt);
    }
    if (message.tokenIndex !== "") {
      writer.uint32(98).string(message.tokenIndex);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Transfer {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTransfer();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.origin = reader.string();
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
          message.bundleData = reader.string();
          break;
        case 9:
          message.bundleSalt = reader.string();
          break;
        case 12:
          message.tokenIndex = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Transfer {
    return {
      origin: isSet(object.origin) ? String(object.origin) : "",
      tx: isSet(object.tx) ? String(object.tx) : "",
      eventId: isSet(object.eventId) ? String(object.eventId) : "",
      fromChain: isSet(object.fromChain) ? String(object.fromChain) : "",
      toChain: isSet(object.toChain) ? String(object.toChain) : "",
      receiver: isSet(object.receiver) ? String(object.receiver) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
      bundleData: isSet(object.bundleData) ? String(object.bundleData) : "",
      bundleSalt: isSet(object.bundleSalt) ? String(object.bundleSalt) : "",
      tokenIndex: isSet(object.tokenIndex) ? String(object.tokenIndex) : "",
    };
  },

  toJSON(message: Transfer): unknown {
    const obj: any = {};
    message.origin !== undefined && (obj.origin = message.origin);
    message.tx !== undefined && (obj.tx = message.tx);
    message.eventId !== undefined && (obj.eventId = message.eventId);
    message.fromChain !== undefined && (obj.fromChain = message.fromChain);
    message.toChain !== undefined && (obj.toChain = message.toChain);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.amount !== undefined && (obj.amount = message.amount);
    message.bundleData !== undefined && (obj.bundleData = message.bundleData);
    message.bundleSalt !== undefined && (obj.bundleSalt = message.bundleSalt);
    message.tokenIndex !== undefined && (obj.tokenIndex = message.tokenIndex);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Transfer>, I>>(object: I): Transfer {
    const message = createBaseTransfer();
    message.origin = object.origin ?? "";
    message.tx = object.tx ?? "";
    message.eventId = object.eventId ?? "";
    message.fromChain = object.fromChain ?? "";
    message.toChain = object.toChain ?? "";
    message.receiver = object.receiver ?? "";
    message.amount = object.amount ?? "";
    message.bundleData = object.bundleData ?? "";
    message.bundleSalt = object.bundleSalt ?? "";
    message.tokenIndex = object.tokenIndex ?? "";
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
