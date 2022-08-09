/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "rarifyprotocol.rarimocore.rarimocore";

export interface Deposit {
  tx: string;
  fromChain: string;
  toChain: string;
  receiver: string;
  tokenAddress: string;
  tokenId: string;
  creator: string;
}

const baseDeposit: object = {
  tx: "",
  fromChain: "",
  toChain: "",
  receiver: "",
  tokenAddress: "",
  tokenId: "",
  creator: "",
};

export const Deposit = {
  encode(message: Deposit, writer: Writer = Writer.create()): Writer {
    if (message.tx !== "") {
      writer.uint32(10).string(message.tx);
    }
    if (message.fromChain !== "") {
      writer.uint32(18).string(message.fromChain);
    }
    if (message.toChain !== "") {
      writer.uint32(26).string(message.toChain);
    }
    if (message.receiver !== "") {
      writer.uint32(34).string(message.receiver);
    }
    if (message.tokenAddress !== "") {
      writer.uint32(42).string(message.tokenAddress);
    }
    if (message.tokenId !== "") {
      writer.uint32(50).string(message.tokenId);
    }
    if (message.creator !== "") {
      writer.uint32(58).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Deposit {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseDeposit } as Deposit;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.tx = reader.string();
          break;
        case 2:
          message.fromChain = reader.string();
          break;
        case 3:
          message.toChain = reader.string();
          break;
        case 4:
          message.receiver = reader.string();
          break;
        case 5:
          message.tokenAddress = reader.string();
          break;
        case 6:
          message.tokenId = reader.string();
          break;
        case 7:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Deposit {
    const message = { ...baseDeposit } as Deposit;
    if (object.tx !== undefined && object.tx !== null) {
      message.tx = String(object.tx);
    } else {
      message.tx = "";
    }
    if (object.fromChain !== undefined && object.fromChain !== null) {
      message.fromChain = String(object.fromChain);
    } else {
      message.fromChain = "";
    }
    if (object.toChain !== undefined && object.toChain !== null) {
      message.toChain = String(object.toChain);
    } else {
      message.toChain = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = String(object.receiver);
    } else {
      message.receiver = "";
    }
    if (object.tokenAddress !== undefined && object.tokenAddress !== null) {
      message.tokenAddress = String(object.tokenAddress);
    } else {
      message.tokenAddress = "";
    }
    if (object.tokenId !== undefined && object.tokenId !== null) {
      message.tokenId = String(object.tokenId);
    } else {
      message.tokenId = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: Deposit): unknown {
    const obj: any = {};
    message.tx !== undefined && (obj.tx = message.tx);
    message.fromChain !== undefined && (obj.fromChain = message.fromChain);
    message.toChain !== undefined && (obj.toChain = message.toChain);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.tokenAddress !== undefined &&
      (obj.tokenAddress = message.tokenAddress);
    message.tokenId !== undefined && (obj.tokenId = message.tokenId);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<Deposit>): Deposit {
    const message = { ...baseDeposit } as Deposit;
    if (object.tx !== undefined && object.tx !== null) {
      message.tx = object.tx;
    } else {
      message.tx = "";
    }
    if (object.fromChain !== undefined && object.fromChain !== null) {
      message.fromChain = object.fromChain;
    } else {
      message.fromChain = "";
    }
    if (object.toChain !== undefined && object.toChain !== null) {
      message.toChain = object.toChain;
    } else {
      message.toChain = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = object.receiver;
    } else {
      message.receiver = "";
    }
    if (object.tokenAddress !== undefined && object.tokenAddress !== null) {
      message.tokenAddress = object.tokenAddress;
    } else {
      message.tokenAddress = "";
    }
    if (object.tokenId !== undefined && object.tokenId !== null) {
      message.tokenId = object.tokenId;
    } else {
      message.tokenId = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
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
