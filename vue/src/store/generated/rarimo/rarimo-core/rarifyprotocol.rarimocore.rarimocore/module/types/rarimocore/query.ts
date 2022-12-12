/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../rarimocore/params";
import { Operation } from "../rarimocore/operation";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { Confirmation } from "../rarimocore/confirmation";

export const protobufPackage = "rarifyprotocol.rarimocore.rarimocore";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

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

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
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
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
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
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetOperationRequest: object = { index: "" };

export const QueryGetOperationRequest = {
  encode(
    message: QueryGetOperationRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetOperationRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetOperationRequest,
    } as QueryGetOperationRequest;
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
    const message = {
      ...baseQueryGetOperationRequest,
    } as QueryGetOperationRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetOperationRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetOperationRequest>
  ): QueryGetOperationRequest {
    const message = {
      ...baseQueryGetOperationRequest,
    } as QueryGetOperationRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetOperationResponse: object = {};

export const QueryGetOperationResponse = {
  encode(
    message: QueryGetOperationResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.operation !== undefined) {
      Operation.encode(message.operation, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetOperationResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetOperationResponse,
    } as QueryGetOperationResponse;
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
    const message = {
      ...baseQueryGetOperationResponse,
    } as QueryGetOperationResponse;
    if (object.operation !== undefined && object.operation !== null) {
      message.operation = Operation.fromJSON(object.operation);
    } else {
      message.operation = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetOperationResponse): unknown {
    const obj: any = {};
    message.operation !== undefined &&
      (obj.operation = message.operation
        ? Operation.toJSON(message.operation)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetOperationResponse>
  ): QueryGetOperationResponse {
    const message = {
      ...baseQueryGetOperationResponse,
    } as QueryGetOperationResponse;
    if (object.operation !== undefined && object.operation !== null) {
      message.operation = Operation.fromPartial(object.operation);
    } else {
      message.operation = undefined;
    }
    return message;
  },
};

const baseQueryAllOperationRequest: object = {};

export const QueryAllOperationRequest = {
  encode(
    message: QueryAllOperationRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllOperationRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllOperationRequest,
    } as QueryAllOperationRequest;
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
    const message = {
      ...baseQueryAllOperationRequest,
    } as QueryAllOperationRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllOperationRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllOperationRequest>
  ): QueryAllOperationRequest {
    const message = {
      ...baseQueryAllOperationRequest,
    } as QueryAllOperationRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllOperationResponse: object = {};

export const QueryAllOperationResponse = {
  encode(
    message: QueryAllOperationResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.operation) {
      Operation.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllOperationResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllOperationResponse,
    } as QueryAllOperationResponse;
    message.operation = [];
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
    const message = {
      ...baseQueryAllOperationResponse,
    } as QueryAllOperationResponse;
    message.operation = [];
    if (object.operation !== undefined && object.operation !== null) {
      for (const e of object.operation) {
        message.operation.push(Operation.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllOperationResponse): unknown {
    const obj: any = {};
    if (message.operation) {
      obj.operation = message.operation.map((e) =>
        e ? Operation.toJSON(e) : undefined
      );
    } else {
      obj.operation = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllOperationResponse>
  ): QueryAllOperationResponse {
    const message = {
      ...baseQueryAllOperationResponse,
    } as QueryAllOperationResponse;
    message.operation = [];
    if (object.operation !== undefined && object.operation !== null) {
      for (const e of object.operation) {
        message.operation.push(Operation.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetConfirmationRequest: object = { root: "" };

export const QueryGetConfirmationRequest = {
  encode(
    message: QueryGetConfirmationRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.root !== "") {
      writer.uint32(10).string(message.root);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetConfirmationRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetConfirmationRequest,
    } as QueryGetConfirmationRequest;
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
    const message = {
      ...baseQueryGetConfirmationRequest,
    } as QueryGetConfirmationRequest;
    if (object.root !== undefined && object.root !== null) {
      message.root = String(object.root);
    } else {
      message.root = "";
    }
    return message;
  },

  toJSON(message: QueryGetConfirmationRequest): unknown {
    const obj: any = {};
    message.root !== undefined && (obj.root = message.root);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetConfirmationRequest>
  ): QueryGetConfirmationRequest {
    const message = {
      ...baseQueryGetConfirmationRequest,
    } as QueryGetConfirmationRequest;
    if (object.root !== undefined && object.root !== null) {
      message.root = object.root;
    } else {
      message.root = "";
    }
    return message;
  },
};

const baseQueryGetConfirmationResponse: object = {};

export const QueryGetConfirmationResponse = {
  encode(
    message: QueryGetConfirmationResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.confirmation !== undefined) {
      Confirmation.encode(
        message.confirmation,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetConfirmationResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetConfirmationResponse,
    } as QueryGetConfirmationResponse;
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
    const message = {
      ...baseQueryGetConfirmationResponse,
    } as QueryGetConfirmationResponse;
    if (object.confirmation !== undefined && object.confirmation !== null) {
      message.confirmation = Confirmation.fromJSON(object.confirmation);
    } else {
      message.confirmation = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetConfirmationResponse): unknown {
    const obj: any = {};
    message.confirmation !== undefined &&
      (obj.confirmation = message.confirmation
        ? Confirmation.toJSON(message.confirmation)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetConfirmationResponse>
  ): QueryGetConfirmationResponse {
    const message = {
      ...baseQueryGetConfirmationResponse,
    } as QueryGetConfirmationResponse;
    if (object.confirmation !== undefined && object.confirmation !== null) {
      message.confirmation = Confirmation.fromPartial(object.confirmation);
    } else {
      message.confirmation = undefined;
    }
    return message;
  },
};

const baseQueryAllConfirmationRequest: object = {};

export const QueryAllConfirmationRequest = {
  encode(
    message: QueryAllConfirmationRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllConfirmationRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllConfirmationRequest,
    } as QueryAllConfirmationRequest;
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
    const message = {
      ...baseQueryAllConfirmationRequest,
    } as QueryAllConfirmationRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllConfirmationRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllConfirmationRequest>
  ): QueryAllConfirmationRequest {
    const message = {
      ...baseQueryAllConfirmationRequest,
    } as QueryAllConfirmationRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllConfirmationResponse: object = {};

export const QueryAllConfirmationResponse = {
  encode(
    message: QueryAllConfirmationResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.confirmation) {
      Confirmation.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllConfirmationResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllConfirmationResponse,
    } as QueryAllConfirmationResponse;
    message.confirmation = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.confirmation.push(
            Confirmation.decode(reader, reader.uint32())
          );
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
    const message = {
      ...baseQueryAllConfirmationResponse,
    } as QueryAllConfirmationResponse;
    message.confirmation = [];
    if (object.confirmation !== undefined && object.confirmation !== null) {
      for (const e of object.confirmation) {
        message.confirmation.push(Confirmation.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllConfirmationResponse): unknown {
    const obj: any = {};
    if (message.confirmation) {
      obj.confirmation = message.confirmation.map((e) =>
        e ? Confirmation.toJSON(e) : undefined
      );
    } else {
      obj.confirmation = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllConfirmationResponse>
  ): QueryAllConfirmationResponse {
    const message = {
      ...baseQueryAllConfirmationResponse,
    } as QueryAllConfirmationResponse;
    message.confirmation = [];
    if (object.confirmation !== undefined && object.confirmation !== null) {
      for (const e of object.confirmation) {
        message.confirmation.push(Confirmation.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a Deposit by index. */
  Operation(
    request: QueryGetOperationRequest
  ): Promise<QueryGetOperationResponse>;
  /** Queries a list of Deposit items. */
  OperationAll(
    request: QueryAllOperationRequest
  ): Promise<QueryAllOperationResponse>;
  /** Queries a Confirmation by index. */
  Confirmation(
    request: QueryGetConfirmationRequest
  ): Promise<QueryGetConfirmationResponse>;
  /** Queries a list of Confirmation items. */
  ConfirmationAll(
    request: QueryAllConfirmationRequest
  ): Promise<QueryAllConfirmationResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.rarimocore.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  Operation(
    request: QueryGetOperationRequest
  ): Promise<QueryGetOperationResponse> {
    const data = QueryGetOperationRequest.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.rarimocore.Query",
      "Operation",
      data
    );
    return promise.then((data) =>
      QueryGetOperationResponse.decode(new Reader(data))
    );
  }

  OperationAll(
    request: QueryAllOperationRequest
  ): Promise<QueryAllOperationResponse> {
    const data = QueryAllOperationRequest.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.rarimocore.Query",
      "OperationAll",
      data
    );
    return promise.then((data) =>
      QueryAllOperationResponse.decode(new Reader(data))
    );
  }

  Confirmation(
    request: QueryGetConfirmationRequest
  ): Promise<QueryGetConfirmationResponse> {
    const data = QueryGetConfirmationRequest.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.rarimocore.Query",
      "Confirmation",
      data
    );
    return promise.then((data) =>
      QueryGetConfirmationResponse.decode(new Reader(data))
    );
  }

  ConfirmationAll(
    request: QueryAllConfirmationRequest
  ): Promise<QueryAllConfirmationResponse> {
    const data = QueryAllConfirmationRequest.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.rarimocore.Query",
      "ConfirmationAll",
      data
    );
    return promise.then((data) =>
      QueryAllConfirmationResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
