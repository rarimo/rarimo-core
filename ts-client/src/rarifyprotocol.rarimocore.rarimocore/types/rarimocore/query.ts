/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { Confirmation } from "./confirmation";
import { Operation } from "./operation";
import { Params } from "./params";

export const protobufPackage = "rarifyprotocol.rarimocore.rarimocore";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetOperationRequest {
  index: string;
}

export interface QueryGetOperationResponse {
  operation: Operation | undefined;
}

export interface QueryAllOperationRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllOperationResponse {
  operation: Operation[];
  pagination: PageResponse | undefined;
}

export interface QueryGetConfirmationRequest {
  root: string;
}

export interface QueryGetConfirmationResponse {
  confirmation: Confirmation | undefined;
}

export interface QueryAllConfirmationRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllConfirmationResponse {
  confirmation: Confirmation[];
  pagination: PageResponse | undefined;
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
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

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(_: I): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQueryGetOperationRequest(): QueryGetOperationRequest {
  return { index: "" };
}

export const QueryGetOperationRequest = {
  encode(message: QueryGetOperationRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetOperationRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetOperationRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetOperationRequest {
    return { index: isSet(object.index) ? String(object.index) : "" };
  },

  toJSON(message: QueryGetOperationRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetOperationRequest>, I>>(object: I): QueryGetOperationRequest {
    const message = createBaseQueryGetOperationRequest();
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseQueryGetOperationResponse(): QueryGetOperationResponse {
  return { operation: undefined };
}

export const QueryGetOperationResponse = {
  encode(message: QueryGetOperationResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.operation !== undefined) {
      Operation.encode(message.operation, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetOperationResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetOperationResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.operation = Operation.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetOperationResponse {
    return { operation: isSet(object.operation) ? Operation.fromJSON(object.operation) : undefined };
  },

  toJSON(message: QueryGetOperationResponse): unknown {
    const obj: any = {};
    message.operation !== undefined
      && (obj.operation = message.operation ? Operation.toJSON(message.operation) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetOperationResponse>, I>>(object: I): QueryGetOperationResponse {
    const message = createBaseQueryGetOperationResponse();
    message.operation = (object.operation !== undefined && object.operation !== null)
      ? Operation.fromPartial(object.operation)
      : undefined;
    return message;
  },
};

function createBaseQueryAllOperationRequest(): QueryAllOperationRequest {
  return { pagination: undefined };
}

export const QueryAllOperationRequest = {
  encode(message: QueryAllOperationRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllOperationRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllOperationRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllOperationRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllOperationRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllOperationRequest>, I>>(object: I): QueryAllOperationRequest {
    const message = createBaseQueryAllOperationRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllOperationResponse(): QueryAllOperationResponse {
  return { operation: [], pagination: undefined };
}

export const QueryAllOperationResponse = {
  encode(message: QueryAllOperationResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.operation) {
      Operation.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllOperationResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllOperationResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.operation.push(Operation.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllOperationResponse {
    return {
      operation: Array.isArray(object?.operation) ? object.operation.map((e: any) => Operation.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllOperationResponse): unknown {
    const obj: any = {};
    if (message.operation) {
      obj.operation = message.operation.map((e) => e ? Operation.toJSON(e) : undefined);
    } else {
      obj.operation = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllOperationResponse>, I>>(object: I): QueryAllOperationResponse {
    const message = createBaseQueryAllOperationResponse();
    message.operation = object.operation?.map((e) => Operation.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetConfirmationRequest(): QueryGetConfirmationRequest {
  return { root: "" };
}

export const QueryGetConfirmationRequest = {
  encode(message: QueryGetConfirmationRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.root !== "") {
      writer.uint32(10).string(message.root);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetConfirmationRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetConfirmationRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.root = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetConfirmationRequest {
    return { root: isSet(object.root) ? String(object.root) : "" };
  },

  toJSON(message: QueryGetConfirmationRequest): unknown {
    const obj: any = {};
    message.root !== undefined && (obj.root = message.root);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetConfirmationRequest>, I>>(object: I): QueryGetConfirmationRequest {
    const message = createBaseQueryGetConfirmationRequest();
    message.root = object.root ?? "";
    return message;
  },
};

function createBaseQueryGetConfirmationResponse(): QueryGetConfirmationResponse {
  return { confirmation: undefined };
}

export const QueryGetConfirmationResponse = {
  encode(message: QueryGetConfirmationResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.confirmation !== undefined) {
      Confirmation.encode(message.confirmation, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetConfirmationResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetConfirmationResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.confirmation = Confirmation.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetConfirmationResponse {
    return { confirmation: isSet(object.confirmation) ? Confirmation.fromJSON(object.confirmation) : undefined };
  },

  toJSON(message: QueryGetConfirmationResponse): unknown {
    const obj: any = {};
    message.confirmation !== undefined
      && (obj.confirmation = message.confirmation ? Confirmation.toJSON(message.confirmation) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetConfirmationResponse>, I>>(object: I): QueryGetConfirmationResponse {
    const message = createBaseQueryGetConfirmationResponse();
    message.confirmation = (object.confirmation !== undefined && object.confirmation !== null)
      ? Confirmation.fromPartial(object.confirmation)
      : undefined;
    return message;
  },
};

function createBaseQueryAllConfirmationRequest(): QueryAllConfirmationRequest {
  return { pagination: undefined };
}

export const QueryAllConfirmationRequest = {
  encode(message: QueryAllConfirmationRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllConfirmationRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllConfirmationRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllConfirmationRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllConfirmationRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllConfirmationRequest>, I>>(object: I): QueryAllConfirmationRequest {
    const message = createBaseQueryAllConfirmationRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllConfirmationResponse(): QueryAllConfirmationResponse {
  return { confirmation: [], pagination: undefined };
}

export const QueryAllConfirmationResponse = {
  encode(message: QueryAllConfirmationResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.confirmation) {
      Confirmation.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllConfirmationResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllConfirmationResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.confirmation.push(Confirmation.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllConfirmationResponse {
    return {
      confirmation: Array.isArray(object?.confirmation)
        ? object.confirmation.map((e: any) => Confirmation.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllConfirmationResponse): unknown {
    const obj: any = {};
    if (message.confirmation) {
      obj.confirmation = message.confirmation.map((e) => e ? Confirmation.toJSON(e) : undefined);
    } else {
      obj.confirmation = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllConfirmationResponse>, I>>(object: I): QueryAllConfirmationResponse {
    const message = createBaseQueryAllConfirmationResponse();
    message.confirmation = object.confirmation?.map((e) => Confirmation.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a Deposit by index. */
  Operation(request: QueryGetOperationRequest): Promise<QueryGetOperationResponse>;
  /** Queries a list of Deposit items. */
  OperationAll(request: QueryAllOperationRequest): Promise<QueryAllOperationResponse>;
  /** Queries a Confirmation by index. */
  Confirmation(request: QueryGetConfirmationRequest): Promise<QueryGetConfirmationResponse>;
  /** Queries a list of Confirmation items. */
  ConfirmationAll(request: QueryAllConfirmationRequest): Promise<QueryAllConfirmationResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.Operation = this.Operation.bind(this);
    this.OperationAll = this.OperationAll.bind(this);
    this.Confirmation = this.Confirmation.bind(this);
    this.ConfirmationAll = this.ConfirmationAll.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("rarifyprotocol.rarimocore.rarimocore.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  Operation(request: QueryGetOperationRequest): Promise<QueryGetOperationResponse> {
    const data = QueryGetOperationRequest.encode(request).finish();
    const promise = this.rpc.request("rarifyprotocol.rarimocore.rarimocore.Query", "Operation", data);
    return promise.then((data) => QueryGetOperationResponse.decode(new _m0.Reader(data)));
  }

  OperationAll(request: QueryAllOperationRequest): Promise<QueryAllOperationResponse> {
    const data = QueryAllOperationRequest.encode(request).finish();
    const promise = this.rpc.request("rarifyprotocol.rarimocore.rarimocore.Query", "OperationAll", data);
    return promise.then((data) => QueryAllOperationResponse.decode(new _m0.Reader(data)));
  }

  Confirmation(request: QueryGetConfirmationRequest): Promise<QueryGetConfirmationResponse> {
    const data = QueryGetConfirmationRequest.encode(request).finish();
    const promise = this.rpc.request("rarifyprotocol.rarimocore.rarimocore.Query", "Confirmation", data);
    return promise.then((data) => QueryGetConfirmationResponse.decode(new _m0.Reader(data)));
  }

  ConfirmationAll(request: QueryAllConfirmationRequest): Promise<QueryAllConfirmationResponse> {
    const data = QueryAllConfirmationRequest.encode(request).finish();
    const promise = this.rpc.request("rarifyprotocol.rarimocore.rarimocore.Query", "ConfirmationAll", data);
    return promise.then((data) => QueryAllConfirmationResponse.decode(new _m0.Reader(data)));
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
