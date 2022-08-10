/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../rarimocore/params";
import { Deposit } from "../rarimocore/deposit";
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

export interface QueryGetDepositRequest {
  tx: string;
}

export interface QueryGetDepositResponse {
  deposit: Deposit | undefined;
}

export interface QueryAllDepositRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllDepositResponse {
  deposit: Deposit[];
  pagination: PageResponse | undefined;
}

export interface QueryGetConfirmationRequest {
  height: string;
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

const baseQueryGetDepositRequest: object = { tx: "" };

export const QueryGetDepositRequest = {
  encode(
    message: QueryGetDepositRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.tx !== "") {
      writer.uint32(10).string(message.tx);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetDepositRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetDepositRequest } as QueryGetDepositRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.tx = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetDepositRequest {
    const message = { ...baseQueryGetDepositRequest } as QueryGetDepositRequest;
    if (object.tx !== undefined && object.tx !== null) {
      message.tx = String(object.tx);
    } else {
      message.tx = "";
    }
    return message;
  },

  toJSON(message: QueryGetDepositRequest): unknown {
    const obj: any = {};
    message.tx !== undefined && (obj.tx = message.tx);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetDepositRequest>
  ): QueryGetDepositRequest {
    const message = { ...baseQueryGetDepositRequest } as QueryGetDepositRequest;
    if (object.tx !== undefined && object.tx !== null) {
      message.tx = object.tx;
    } else {
      message.tx = "";
    }
    return message;
  },
};

const baseQueryGetDepositResponse: object = {};

export const QueryGetDepositResponse = {
  encode(
    message: QueryGetDepositResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.deposit !== undefined) {
      Deposit.encode(message.deposit, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetDepositResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetDepositResponse,
    } as QueryGetDepositResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.deposit = Deposit.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetDepositResponse {
    const message = {
      ...baseQueryGetDepositResponse,
    } as QueryGetDepositResponse;
    if (object.deposit !== undefined && object.deposit !== null) {
      message.deposit = Deposit.fromJSON(object.deposit);
    } else {
      message.deposit = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetDepositResponse): unknown {
    const obj: any = {};
    message.deposit !== undefined &&
      (obj.deposit = message.deposit
        ? Deposit.toJSON(message.deposit)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetDepositResponse>
  ): QueryGetDepositResponse {
    const message = {
      ...baseQueryGetDepositResponse,
    } as QueryGetDepositResponse;
    if (object.deposit !== undefined && object.deposit !== null) {
      message.deposit = Deposit.fromPartial(object.deposit);
    } else {
      message.deposit = undefined;
    }
    return message;
  },
};

const baseQueryAllDepositRequest: object = {};

export const QueryAllDepositRequest = {
  encode(
    message: QueryAllDepositRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllDepositRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllDepositRequest } as QueryAllDepositRequest;
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

  fromJSON(object: any): QueryAllDepositRequest {
    const message = { ...baseQueryAllDepositRequest } as QueryAllDepositRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllDepositRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllDepositRequest>
  ): QueryAllDepositRequest {
    const message = { ...baseQueryAllDepositRequest } as QueryAllDepositRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllDepositResponse: object = {};

export const QueryAllDepositResponse = {
  encode(
    message: QueryAllDepositResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.deposit) {
      Deposit.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllDepositResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllDepositResponse,
    } as QueryAllDepositResponse;
    message.deposit = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.deposit.push(Deposit.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllDepositResponse {
    const message = {
      ...baseQueryAllDepositResponse,
    } as QueryAllDepositResponse;
    message.deposit = [];
    if (object.deposit !== undefined && object.deposit !== null) {
      for (const e of object.deposit) {
        message.deposit.push(Deposit.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllDepositResponse): unknown {
    const obj: any = {};
    if (message.deposit) {
      obj.deposit = message.deposit.map((e) =>
        e ? Deposit.toJSON(e) : undefined
      );
    } else {
      obj.deposit = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllDepositResponse>
  ): QueryAllDepositResponse {
    const message = {
      ...baseQueryAllDepositResponse,
    } as QueryAllDepositResponse;
    message.deposit = [];
    if (object.deposit !== undefined && object.deposit !== null) {
      for (const e of object.deposit) {
        message.deposit.push(Deposit.fromPartial(e));
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

const baseQueryGetConfirmationRequest: object = { height: "" };

export const QueryGetConfirmationRequest = {
  encode(
    message: QueryGetConfirmationRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.height !== "") {
      writer.uint32(10).string(message.height);
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
          message.height = reader.string();
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
    if (object.height !== undefined && object.height !== null) {
      message.height = String(object.height);
    } else {
      message.height = "";
    }
    return message;
  },

  toJSON(message: QueryGetConfirmationRequest): unknown {
    const obj: any = {};
    message.height !== undefined && (obj.height = message.height);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetConfirmationRequest>
  ): QueryGetConfirmationRequest {
    const message = {
      ...baseQueryGetConfirmationRequest,
    } as QueryGetConfirmationRequest;
    if (object.height !== undefined && object.height !== null) {
      message.height = object.height;
    } else {
      message.height = "";
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
  Deposit(request: QueryGetDepositRequest): Promise<QueryGetDepositResponse>;
  /** Queries a list of Deposit items. */
  DepositAll(request: QueryAllDepositRequest): Promise<QueryAllDepositResponse>;
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

  Deposit(request: QueryGetDepositRequest): Promise<QueryGetDepositResponse> {
    const data = QueryGetDepositRequest.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.rarimocore.Query",
      "Deposit",
      data
    );
    return promise.then((data) =>
      QueryGetDepositResponse.decode(new Reader(data))
    );
  }

  DepositAll(
    request: QueryAllDepositRequest
  ): Promise<QueryAllDepositResponse> {
    const data = QueryAllDepositRequest.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.rarimocore.Query",
      "DepositAll",
      data
    );
    return promise.then((data) =>
      QueryAllDepositResponse.decode(new Reader(data))
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
