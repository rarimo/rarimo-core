/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { type, typeFromJSON, typeToJSON } from "../tokenmanager/item";

export const protobufPackage = "rarifyprotocol.rarimocore.rarimocore";

export interface MsgCreateTransferOp {
  creator: string;
  tx: string;
  eventId: string;
  fromChain: string;
  tokenType: type;
}

export interface MsgCreateTransferOpResponse {
}

export interface MsgCreateConfirmation {
  creator: string;
  /** hex-encoded */
  root: string;
  indexes: string[];
  /** hex-encoded */
  signatureECDSA: string;
}

export interface MsgCreateConfirmationResponse {
}

export interface MsgCreateChangeKeyECDSA {
  creator: string;
  /** hex-encoded */
  newKey: string;
}

export interface MsgCreateChangeKeyECDSAResponse {
}

function createBaseMsgCreateTransferOp(): MsgCreateTransferOp {
  return { creator: "", tx: "", eventId: "", fromChain: "", tokenType: 0 };
}

export const MsgCreateTransferOp = {
  encode(message: MsgCreateTransferOp, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
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
    if (message.tokenType !== 0) {
      writer.uint32(40).int32(message.tokenType);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateTransferOp {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateTransferOp();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
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
          message.tokenType = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateTransferOp {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      tx: isSet(object.tx) ? String(object.tx) : "",
      eventId: isSet(object.eventId) ? String(object.eventId) : "",
      fromChain: isSet(object.fromChain) ? String(object.fromChain) : "",
      tokenType: isSet(object.tokenType) ? typeFromJSON(object.tokenType) : 0,
    };
  },

  toJSON(message: MsgCreateTransferOp): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.tx !== undefined && (obj.tx = message.tx);
    message.eventId !== undefined && (obj.eventId = message.eventId);
    message.fromChain !== undefined && (obj.fromChain = message.fromChain);
    message.tokenType !== undefined && (obj.tokenType = typeToJSON(message.tokenType));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateTransferOp>, I>>(object: I): MsgCreateTransferOp {
    const message = createBaseMsgCreateTransferOp();
    message.creator = object.creator ?? "";
    message.tx = object.tx ?? "";
    message.eventId = object.eventId ?? "";
    message.fromChain = object.fromChain ?? "";
    message.tokenType = object.tokenType ?? 0;
    return message;
  },
};

function createBaseMsgCreateTransferOpResponse(): MsgCreateTransferOpResponse {
  return {};
}

export const MsgCreateTransferOpResponse = {
  encode(_: MsgCreateTransferOpResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateTransferOpResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateTransferOpResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateTransferOpResponse {
    return {};
  },

  toJSON(_: MsgCreateTransferOpResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateTransferOpResponse>, I>>(_: I): MsgCreateTransferOpResponse {
    const message = createBaseMsgCreateTransferOpResponse();
    return message;
  },
};

function createBaseMsgCreateConfirmation(): MsgCreateConfirmation {
  return { creator: "", root: "", indexes: [], signatureECDSA: "" };
}

export const MsgCreateConfirmation = {
  encode(message: MsgCreateConfirmation, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.root !== "") {
      writer.uint32(18).string(message.root);
    }
    for (const v of message.indexes) {
      writer.uint32(26).string(v!);
    }
    if (message.signatureECDSA !== "") {
      writer.uint32(34).string(message.signatureECDSA);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateConfirmation {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateConfirmation();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.root = reader.string();
          break;
        case 3:
          message.indexes.push(reader.string());
          break;
        case 4:
          message.signatureECDSA = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateConfirmation {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      root: isSet(object.root) ? String(object.root) : "",
      indexes: Array.isArray(object?.indexes) ? object.indexes.map((e: any) => String(e)) : [],
      signatureECDSA: isSet(object.signatureECDSA) ? String(object.signatureECDSA) : "",
    };
  },

  toJSON(message: MsgCreateConfirmation): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.root !== undefined && (obj.root = message.root);
    if (message.indexes) {
      obj.indexes = message.indexes.map((e) => e);
    } else {
      obj.indexes = [];
    }
    message.signatureECDSA !== undefined && (obj.signatureECDSA = message.signatureECDSA);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateConfirmation>, I>>(object: I): MsgCreateConfirmation {
    const message = createBaseMsgCreateConfirmation();
    message.creator = object.creator ?? "";
    message.root = object.root ?? "";
    message.indexes = object.indexes?.map((e) => e) || [];
    message.signatureECDSA = object.signatureECDSA ?? "";
    return message;
  },
};

function createBaseMsgCreateConfirmationResponse(): MsgCreateConfirmationResponse {
  return {};
}

export const MsgCreateConfirmationResponse = {
  encode(_: MsgCreateConfirmationResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateConfirmationResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateConfirmationResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateConfirmationResponse {
    return {};
  },

  toJSON(_: MsgCreateConfirmationResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateConfirmationResponse>, I>>(_: I): MsgCreateConfirmationResponse {
    const message = createBaseMsgCreateConfirmationResponse();
    return message;
  },
};

function createBaseMsgCreateChangeKeyECDSA(): MsgCreateChangeKeyECDSA {
  return { creator: "", newKey: "" };
}

export const MsgCreateChangeKeyECDSA = {
  encode(message: MsgCreateChangeKeyECDSA, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.newKey !== "") {
      writer.uint32(18).string(message.newKey);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateChangeKeyECDSA {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateChangeKeyECDSA();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.newKey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateChangeKeyECDSA {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      newKey: isSet(object.newKey) ? String(object.newKey) : "",
    };
  },

  toJSON(message: MsgCreateChangeKeyECDSA): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.newKey !== undefined && (obj.newKey = message.newKey);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateChangeKeyECDSA>, I>>(object: I): MsgCreateChangeKeyECDSA {
    const message = createBaseMsgCreateChangeKeyECDSA();
    message.creator = object.creator ?? "";
    message.newKey = object.newKey ?? "";
    return message;
  },
};

function createBaseMsgCreateChangeKeyECDSAResponse(): MsgCreateChangeKeyECDSAResponse {
  return {};
}

export const MsgCreateChangeKeyECDSAResponse = {
  encode(_: MsgCreateChangeKeyECDSAResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateChangeKeyECDSAResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateChangeKeyECDSAResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateChangeKeyECDSAResponse {
    return {};
  },

  toJSON(_: MsgCreateChangeKeyECDSAResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateChangeKeyECDSAResponse>, I>>(_: I): MsgCreateChangeKeyECDSAResponse {
    const message = createBaseMsgCreateChangeKeyECDSAResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateTransferOperation(request: MsgCreateTransferOp): Promise<MsgCreateTransferOpResponse>;
  CreateConfirmation(request: MsgCreateConfirmation): Promise<MsgCreateConfirmationResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateChangeKeyECDSA(request: MsgCreateChangeKeyECDSA): Promise<MsgCreateChangeKeyECDSAResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateTransferOperation = this.CreateTransferOperation.bind(this);
    this.CreateConfirmation = this.CreateConfirmation.bind(this);
    this.CreateChangeKeyECDSA = this.CreateChangeKeyECDSA.bind(this);
  }
  CreateTransferOperation(request: MsgCreateTransferOp): Promise<MsgCreateTransferOpResponse> {
    const data = MsgCreateTransferOp.encode(request).finish();
    const promise = this.rpc.request("rarifyprotocol.rarimocore.rarimocore.Msg", "CreateTransferOperation", data);
    return promise.then((data) => MsgCreateTransferOpResponse.decode(new _m0.Reader(data)));
  }

  CreateConfirmation(request: MsgCreateConfirmation): Promise<MsgCreateConfirmationResponse> {
    const data = MsgCreateConfirmation.encode(request).finish();
    const promise = this.rpc.request("rarifyprotocol.rarimocore.rarimocore.Msg", "CreateConfirmation", data);
    return promise.then((data) => MsgCreateConfirmationResponse.decode(new _m0.Reader(data)));
  }

  CreateChangeKeyECDSA(request: MsgCreateChangeKeyECDSA): Promise<MsgCreateChangeKeyECDSAResponse> {
    const data = MsgCreateChangeKeyECDSA.encode(request).finish();
    const promise = this.rpc.request("rarifyprotocol.rarimocore.rarimocore.Msg", "CreateChangeKeyECDSA", data);
    return promise.then((data) => MsgCreateChangeKeyECDSAResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

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
