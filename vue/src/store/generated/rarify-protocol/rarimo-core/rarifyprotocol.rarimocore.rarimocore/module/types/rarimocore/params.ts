/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "rarifyprotocol.rarimocore.rarimocore";

export enum ParamsUpdateType {
  CHANGE_SET = 0,
  OTHER = 1,
  UNRECOGNIZED = -1,
}

export function paramsUpdateTypeFromJSON(object: any): ParamsUpdateType {
  switch (object) {
    case 0:
    case "CHANGE_SET":
      return ParamsUpdateType.CHANGE_SET;
    case 1:
    case "OTHER":
      return ParamsUpdateType.OTHER;
    case -1:
    case "UNRECOGNIZED":
    default:
      return ParamsUpdateType.UNRECOGNIZED;
  }
}

export function paramsUpdateTypeToJSON(object: ParamsUpdateType): string {
  switch (object) {
    case ParamsUpdateType.CHANGE_SET:
      return "CHANGE_SET";
    case ParamsUpdateType.OTHER:
      return "OTHER";
    default:
      return "UNKNOWN";
  }
}

export interface Party {
  /** PublicKey in hex format */
  pubKey: string;
  /** Server address connect to */
  address: string;
  /** Rarimo core account */
  account: string;
  verified: boolean;
}

/** Params defines the parameters for the module. */
export interface Params {
  keyECDSA: string;
  threshold: number;
  parties: Party[];
  isUpdateRequired: boolean;
}

const baseParty: object = {
  pubKey: "",
  address: "",
  account: "",
  verified: false,
};

export const Party = {
  encode(message: Party, writer: Writer = Writer.create()): Writer {
    if (message.pubKey !== "") {
      writer.uint32(10).string(message.pubKey);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.account !== "") {
      writer.uint32(26).string(message.account);
    }
    if (message.verified === true) {
      writer.uint32(32).bool(message.verified);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Party {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseParty } as Party;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pubKey = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.account = reader.string();
          break;
        case 4:
          message.verified = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Party {
    const message = { ...baseParty } as Party;
    if (object.pubKey !== undefined && object.pubKey !== null) {
      message.pubKey = String(object.pubKey);
    } else {
      message.pubKey = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.account !== undefined && object.account !== null) {
      message.account = String(object.account);
    } else {
      message.account = "";
    }
    if (object.verified !== undefined && object.verified !== null) {
      message.verified = Boolean(object.verified);
    } else {
      message.verified = false;
    }
    return message;
  },

  toJSON(message: Party): unknown {
    const obj: any = {};
    message.pubKey !== undefined && (obj.pubKey = message.pubKey);
    message.address !== undefined && (obj.address = message.address);
    message.account !== undefined && (obj.account = message.account);
    message.verified !== undefined && (obj.verified = message.verified);
    return obj;
  },

  fromPartial(object: DeepPartial<Party>): Party {
    const message = { ...baseParty } as Party;
    if (object.pubKey !== undefined && object.pubKey !== null) {
      message.pubKey = object.pubKey;
    } else {
      message.pubKey = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.account !== undefined && object.account !== null) {
      message.account = object.account;
    } else {
      message.account = "";
    }
    if (object.verified !== undefined && object.verified !== null) {
      message.verified = object.verified;
    } else {
      message.verified = false;
    }
    return message;
  },
};

const baseParams: object = {
  keyECDSA: "",
  threshold: 0,
  isUpdateRequired: false,
};

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    if (message.keyECDSA !== "") {
      writer.uint32(10).string(message.keyECDSA);
    }
    if (message.threshold !== 0) {
      writer.uint32(16).uint64(message.threshold);
    }
    for (const v of message.parties) {
      Party.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    if (message.isUpdateRequired === true) {
      writer.uint32(40).bool(message.isUpdateRequired);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseParams } as Params;
    message.parties = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.keyECDSA = reader.string();
          break;
        case 2:
          message.threshold = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.parties.push(Party.decode(reader, reader.uint32()));
          break;
        case 5:
          message.isUpdateRequired = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Params {
    const message = { ...baseParams } as Params;
    message.parties = [];
    if (object.keyECDSA !== undefined && object.keyECDSA !== null) {
      message.keyECDSA = String(object.keyECDSA);
    } else {
      message.keyECDSA = "";
    }
    if (object.threshold !== undefined && object.threshold !== null) {
      message.threshold = Number(object.threshold);
    } else {
      message.threshold = 0;
    }
    if (object.parties !== undefined && object.parties !== null) {
      for (const e of object.parties) {
        message.parties.push(Party.fromJSON(e));
      }
    }
    if (
      object.isUpdateRequired !== undefined &&
      object.isUpdateRequired !== null
    ) {
      message.isUpdateRequired = Boolean(object.isUpdateRequired);
    } else {
      message.isUpdateRequired = false;
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.keyECDSA !== undefined && (obj.keyECDSA = message.keyECDSA);
    message.threshold !== undefined && (obj.threshold = message.threshold);
    if (message.parties) {
      obj.parties = message.parties.map((e) =>
        e ? Party.toJSON(e) : undefined
      );
    } else {
      obj.parties = [];
    }
    message.isUpdateRequired !== undefined &&
      (obj.isUpdateRequired = message.isUpdateRequired);
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    message.parties = [];
    if (object.keyECDSA !== undefined && object.keyECDSA !== null) {
      message.keyECDSA = object.keyECDSA;
    } else {
      message.keyECDSA = "";
    }
    if (object.threshold !== undefined && object.threshold !== null) {
      message.threshold = object.threshold;
    } else {
      message.threshold = 0;
    }
    if (object.parties !== undefined && object.parties !== null) {
      for (const e of object.parties) {
        message.parties.push(Party.fromPartial(e));
      }
    }
    if (
      object.isUpdateRequired !== undefined &&
      object.isUpdateRequired !== null
    ) {
      message.isUpdateRequired = object.isUpdateRequired;
    } else {
      message.isUpdateRequired = false;
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
