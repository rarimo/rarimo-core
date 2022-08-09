/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../rarimocore/params";
import { Deposit } from "../rarimocore/deposit";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

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

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a Deposit by index. */
  Deposit(request: QueryGetDepositRequest): Promise<QueryGetDepositResponse>;
  /** Queries a list of Deposit items. */
  DepositAll(request: QueryAllDepositRequest): Promise<QueryAllDepositResponse>;
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
