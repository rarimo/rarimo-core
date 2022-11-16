/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "rarifyprotocol.rarimocore.rarimocore";

export enum StepType {
  Proposing = 0,
  Accepting = 1,
  Signing = 2,
  UNRECOGNIZED = -1,
}

export function stepTypeFromJSON(object: any): StepType {
  switch (object) {
    case 0:
    case "Proposing":
      return StepType.Proposing;
    case 1:
    case "Accepting":
      return StepType.Accepting;
    case 2:
    case "Signing":
      return StepType.Signing;
    case -1:
    case "UNRECOGNIZED":
    default:
      return StepType.UNRECOGNIZED;
  }
}

export function stepTypeToJSON(object: StepType): string {
  switch (object) {
    case StepType.Proposing:
      return "Proposing";
    case StepType.Accepting:
      return "Accepting";
    case StepType.Signing:
      return "Signing";
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
  active: boolean;
}

export interface Step {
  /** Duration in blocks */
  duration: number;
  type: StepType;
}

/** Params defines the parameters for the module. */
export interface Params {
  keyECDSA: string;
  threshold: number;
  parties: Party[];
  steps: Step[];
}

const baseParty: object = {
  pubKey: "",
  address: "",
  account: "",
  active: false,
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
    if (message.active === true) {
      writer.uint32(32).bool(message.active);
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
          message.active = reader.bool();
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
    if (object.active !== undefined && object.active !== null) {
      message.active = Boolean(object.active);
    } else {
      message.active = false;
    }
    return message;
  },

  toJSON(message: Party): unknown {
    const obj: any = {};
    message.pubKey !== undefined && (obj.pubKey = message.pubKey);
    message.address !== undefined && (obj.address = message.address);
    message.account !== undefined && (obj.account = message.account);
    message.active !== undefined && (obj.active = message.active);
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
    if (object.active !== undefined && object.active !== null) {
      message.active = object.active;
    } else {
      message.active = false;
    }
    return message;
  },
};

const baseStep: object = { duration: 0, type: 0 };

export const Step = {
  encode(message: Step, writer: Writer = Writer.create()): Writer {
    if (message.duration !== 0) {
      writer.uint32(8).uint64(message.duration);
    }
    if (message.type !== 0) {
      writer.uint32(16).int32(message.type);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Step {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseStep } as Step;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.duration = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.type = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Step {
    const message = { ...baseStep } as Step;
    if (object.duration !== undefined && object.duration !== null) {
      message.duration = Number(object.duration);
    } else {
      message.duration = 0;
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = stepTypeFromJSON(object.type);
    } else {
      message.type = 0;
    }
    return message;
  },

  toJSON(message: Step): unknown {
    const obj: any = {};
    message.duration !== undefined && (obj.duration = message.duration);
    message.type !== undefined && (obj.type = stepTypeToJSON(message.type));
    return obj;
  },

  fromPartial(object: DeepPartial<Step>): Step {
    const message = { ...baseStep } as Step;
    if (object.duration !== undefined && object.duration !== null) {
      message.duration = object.duration;
    } else {
      message.duration = 0;
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = object.type;
    } else {
      message.type = 0;
    }
    return message;
  },
};

const baseParams: object = { keyECDSA: "", threshold: 0 };

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
    for (const v of message.steps) {
      Step.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseParams } as Params;
    message.parties = [];
    message.steps = [];
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
        case 4:
          message.steps.push(Step.decode(reader, reader.uint32()));
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
    message.steps = [];
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
    if (object.steps !== undefined && object.steps !== null) {
      for (const e of object.steps) {
        message.steps.push(Step.fromJSON(e));
      }
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
    if (message.steps) {
      obj.steps = message.steps.map((e) => (e ? Step.toJSON(e) : undefined));
    } else {
      obj.steps = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    message.parties = [];
    message.steps = [];
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
    if (object.steps !== undefined && object.steps !== null) {
      for (const e of object.steps) {
        message.steps.push(Step.fromPartial(e));
      }
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
