/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../tokenmanager/params";
import { Item } from "../tokenmanager/item";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { Info } from "../tokenmanager/info";

export const protobufPackage = "rarifyprotocol.rarimocore.tokenmanager";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetItemRequest {
  tokenAddress: string;
  tokenId: string;
  chain: string;
}

export interface QueryGetItemResponse {
  item: Item | undefined;
}

export interface QueryAllItemRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllItemResponse {
  item: Item[];
  pagination: PageResponse | undefined;
}

export interface QueryGetInfoRequest {
  index: string;
}

export interface QueryGetInfoResponse {
  info: Info | undefined;
}

export interface QueryAllInfoRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllInfoResponse {
  info: Info[];
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

const baseQueryGetItemRequest: object = {
  tokenAddress: "",
  tokenId: "",
  chain: "",
};

export const QueryGetItemRequest = {
  encode(
    message: QueryGetItemRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.tokenAddress !== "") {
      writer.uint32(10).string(message.tokenAddress);
    }
    if (message.tokenId !== "") {
      writer.uint32(18).string(message.tokenId);
    }
    if (message.chain !== "") {
      writer.uint32(26).string(message.chain);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetItemRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetItemRequest } as QueryGetItemRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.tokenAddress = reader.string();
          break;
        case 2:
          message.tokenId = reader.string();
          break;
        case 3:
          message.chain = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetItemRequest {
    const message = { ...baseQueryGetItemRequest } as QueryGetItemRequest;
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
    if (object.chain !== undefined && object.chain !== null) {
      message.chain = String(object.chain);
    } else {
      message.chain = "";
    }
    return message;
  },

  toJSON(message: QueryGetItemRequest): unknown {
    const obj: any = {};
    message.tokenAddress !== undefined &&
      (obj.tokenAddress = message.tokenAddress);
    message.tokenId !== undefined && (obj.tokenId = message.tokenId);
    message.chain !== undefined && (obj.chain = message.chain);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetItemRequest>): QueryGetItemRequest {
    const message = { ...baseQueryGetItemRequest } as QueryGetItemRequest;
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
    if (object.chain !== undefined && object.chain !== null) {
      message.chain = object.chain;
    } else {
      message.chain = "";
    }
    return message;
  },
};

const baseQueryGetItemResponse: object = {};

export const QueryGetItemResponse = {
  encode(
    message: QueryGetItemResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.item !== undefined) {
      Item.encode(message.item, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetItemResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetItemResponse } as QueryGetItemResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.item = Item.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetItemResponse {
    const message = { ...baseQueryGetItemResponse } as QueryGetItemResponse;
    if (object.item !== undefined && object.item !== null) {
      message.item = Item.fromJSON(object.item);
    } else {
      message.item = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetItemResponse): unknown {
    const obj: any = {};
    message.item !== undefined &&
      (obj.item = message.item ? Item.toJSON(message.item) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetItemResponse>): QueryGetItemResponse {
    const message = { ...baseQueryGetItemResponse } as QueryGetItemResponse;
    if (object.item !== undefined && object.item !== null) {
      message.item = Item.fromPartial(object.item);
    } else {
      message.item = undefined;
    }
    return message;
  },
};

const baseQueryAllItemRequest: object = {};

export const QueryAllItemRequest = {
  encode(
    message: QueryAllItemRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllItemRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllItemRequest } as QueryAllItemRequest;
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

  fromJSON(object: any): QueryAllItemRequest {
    const message = { ...baseQueryAllItemRequest } as QueryAllItemRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllItemRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllItemRequest>): QueryAllItemRequest {
    const message = { ...baseQueryAllItemRequest } as QueryAllItemRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllItemResponse: object = {};

export const QueryAllItemResponse = {
  encode(
    message: QueryAllItemResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.item) {
      Item.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllItemResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllItemResponse } as QueryAllItemResponse;
    message.item = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.item.push(Item.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllItemResponse {
    const message = { ...baseQueryAllItemResponse } as QueryAllItemResponse;
    message.item = [];
    if (object.item !== undefined && object.item !== null) {
      for (const e of object.item) {
        message.item.push(Item.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllItemResponse): unknown {
    const obj: any = {};
    if (message.item) {
      obj.item = message.item.map((e) => (e ? Item.toJSON(e) : undefined));
    } else {
      obj.item = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllItemResponse>): QueryAllItemResponse {
    const message = { ...baseQueryAllItemResponse } as QueryAllItemResponse;
    message.item = [];
    if (object.item !== undefined && object.item !== null) {
      for (const e of object.item) {
        message.item.push(Item.fromPartial(e));
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

const baseQueryGetInfoRequest: object = { index: "" };

export const QueryGetInfoRequest = {
  encode(
    message: QueryGetInfoRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetInfoRequest } as QueryGetInfoRequest;
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

  fromJSON(object: any): QueryGetInfoRequest {
    const message = { ...baseQueryGetInfoRequest } as QueryGetInfoRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetInfoRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetInfoRequest>): QueryGetInfoRequest {
    const message = { ...baseQueryGetInfoRequest } as QueryGetInfoRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetInfoResponse: object = {};

export const QueryGetInfoResponse = {
  encode(
    message: QueryGetInfoResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.info !== undefined) {
      Info.encode(message.info, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetInfoResponse } as QueryGetInfoResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.info = Info.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetInfoResponse {
    const message = { ...baseQueryGetInfoResponse } as QueryGetInfoResponse;
    if (object.info !== undefined && object.info !== null) {
      message.info = Info.fromJSON(object.info);
    } else {
      message.info = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetInfoResponse): unknown {
    const obj: any = {};
    message.info !== undefined &&
      (obj.info = message.info ? Info.toJSON(message.info) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetInfoResponse>): QueryGetInfoResponse {
    const message = { ...baseQueryGetInfoResponse } as QueryGetInfoResponse;
    if (object.info !== undefined && object.info !== null) {
      message.info = Info.fromPartial(object.info);
    } else {
      message.info = undefined;
    }
    return message;
  },
};

const baseQueryAllInfoRequest: object = {};

export const QueryAllInfoRequest = {
  encode(
    message: QueryAllInfoRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllInfoRequest } as QueryAllInfoRequest;
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

  fromJSON(object: any): QueryAllInfoRequest {
    const message = { ...baseQueryAllInfoRequest } as QueryAllInfoRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllInfoRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllInfoRequest>): QueryAllInfoRequest {
    const message = { ...baseQueryAllInfoRequest } as QueryAllInfoRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllInfoResponse: object = {};

export const QueryAllInfoResponse = {
  encode(
    message: QueryAllInfoResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.info) {
      Info.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllInfoResponse } as QueryAllInfoResponse;
    message.info = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.info.push(Info.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllInfoResponse {
    const message = { ...baseQueryAllInfoResponse } as QueryAllInfoResponse;
    message.info = [];
    if (object.info !== undefined && object.info !== null) {
      for (const e of object.info) {
        message.info.push(Info.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllInfoResponse): unknown {
    const obj: any = {};
    if (message.info) {
      obj.info = message.info.map((e) => (e ? Info.toJSON(e) : undefined));
    } else {
      obj.info = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllInfoResponse>): QueryAllInfoResponse {
    const message = { ...baseQueryAllInfoResponse } as QueryAllInfoResponse;
    message.info = [];
    if (object.info !== undefined && object.info !== null) {
      for (const e of object.info) {
        message.info.push(Info.fromPartial(e));
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
  /** Queries a Item by index. */
  Item(request: QueryGetItemRequest): Promise<QueryGetItemResponse>;
  /** Queries a list of Item items. */
  ItemAll(request: QueryAllItemRequest): Promise<QueryAllItemResponse>;
  /** Queries a Info by index. */
  Info(request: QueryGetInfoRequest): Promise<QueryGetInfoResponse>;
  /** Queries a list of Info items. */
  InfoAll(request: QueryAllInfoRequest): Promise<QueryAllInfoResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.tokenmanager.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  Item(request: QueryGetItemRequest): Promise<QueryGetItemResponse> {
    const data = QueryGetItemRequest.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.tokenmanager.Query",
      "Item",
      data
    );
    return promise.then((data) =>
      QueryGetItemResponse.decode(new Reader(data))
    );
  }

  ItemAll(request: QueryAllItemRequest): Promise<QueryAllItemResponse> {
    const data = QueryAllItemRequest.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.tokenmanager.Query",
      "ItemAll",
      data
    );
    return promise.then((data) =>
      QueryAllItemResponse.decode(new Reader(data))
    );
  }

  Info(request: QueryGetInfoRequest): Promise<QueryGetInfoResponse> {
    const data = QueryGetInfoRequest.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.tokenmanager.Query",
      "Info",
      data
    );
    return promise.then((data) =>
      QueryGetInfoResponse.decode(new Reader(data))
    );
  }

  InfoAll(request: QueryAllInfoRequest): Promise<QueryAllInfoResponse> {
    const data = QueryAllInfoRequest.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.tokenmanager.Query",
      "InfoAll",
      data
    );
    return promise.then((data) =>
      QueryAllInfoResponse.decode(new Reader(data))
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
