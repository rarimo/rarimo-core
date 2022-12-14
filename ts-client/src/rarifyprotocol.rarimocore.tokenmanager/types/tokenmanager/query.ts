/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { Info } from "./info";
import { Item } from "./item";
import { Params } from "./params";

export const protobufPackage = "rarifyprotocol.rarimocore.tokenmanager";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

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

export interface QueryGetItemByChainRequest {
  infoIndex: string;
  chain: string;
}

export interface QueryGetItemByChainResponse {
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

function createBaseQueryGetItemRequest(): QueryGetItemRequest {
  return { tokenAddress: "", tokenId: "", chain: "" };
}

export const QueryGetItemRequest = {
  encode(message: QueryGetItemRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetItemRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetItemRequest();
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
    return {
      tokenAddress: isSet(object.tokenAddress) ? String(object.tokenAddress) : "",
      tokenId: isSet(object.tokenId) ? String(object.tokenId) : "",
      chain: isSet(object.chain) ? String(object.chain) : "",
    };
  },

  toJSON(message: QueryGetItemRequest): unknown {
    const obj: any = {};
    message.tokenAddress !== undefined && (obj.tokenAddress = message.tokenAddress);
    message.tokenId !== undefined && (obj.tokenId = message.tokenId);
    message.chain !== undefined && (obj.chain = message.chain);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetItemRequest>, I>>(object: I): QueryGetItemRequest {
    const message = createBaseQueryGetItemRequest();
    message.tokenAddress = object.tokenAddress ?? "";
    message.tokenId = object.tokenId ?? "";
    message.chain = object.chain ?? "";
    return message;
  },
};

function createBaseQueryGetItemResponse(): QueryGetItemResponse {
  return { item: undefined };
}

export const QueryGetItemResponse = {
  encode(message: QueryGetItemResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.item !== undefined) {
      Item.encode(message.item, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetItemResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetItemResponse();
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
    return { item: isSet(object.item) ? Item.fromJSON(object.item) : undefined };
  },

  toJSON(message: QueryGetItemResponse): unknown {
    const obj: any = {};
    message.item !== undefined && (obj.item = message.item ? Item.toJSON(message.item) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetItemResponse>, I>>(object: I): QueryGetItemResponse {
    const message = createBaseQueryGetItemResponse();
    message.item = (object.item !== undefined && object.item !== null) ? Item.fromPartial(object.item) : undefined;
    return message;
  },
};

function createBaseQueryGetItemByChainRequest(): QueryGetItemByChainRequest {
  return { infoIndex: "", chain: "" };
}

export const QueryGetItemByChainRequest = {
  encode(message: QueryGetItemByChainRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.infoIndex !== "") {
      writer.uint32(10).string(message.infoIndex);
    }
    if (message.chain !== "") {
      writer.uint32(18).string(message.chain);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetItemByChainRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetItemByChainRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.infoIndex = reader.string();
          break;
        case 2:
          message.chain = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetItemByChainRequest {
    return {
      infoIndex: isSet(object.infoIndex) ? String(object.infoIndex) : "",
      chain: isSet(object.chain) ? String(object.chain) : "",
    };
  },

  toJSON(message: QueryGetItemByChainRequest): unknown {
    const obj: any = {};
    message.infoIndex !== undefined && (obj.infoIndex = message.infoIndex);
    message.chain !== undefined && (obj.chain = message.chain);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetItemByChainRequest>, I>>(object: I): QueryGetItemByChainRequest {
    const message = createBaseQueryGetItemByChainRequest();
    message.infoIndex = object.infoIndex ?? "";
    message.chain = object.chain ?? "";
    return message;
  },
};

function createBaseQueryGetItemByChainResponse(): QueryGetItemByChainResponse {
  return { item: undefined };
}

export const QueryGetItemByChainResponse = {
  encode(message: QueryGetItemByChainResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.item !== undefined) {
      Item.encode(message.item, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetItemByChainResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetItemByChainResponse();
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

  fromJSON(object: any): QueryGetItemByChainResponse {
    return { item: isSet(object.item) ? Item.fromJSON(object.item) : undefined };
  },

  toJSON(message: QueryGetItemByChainResponse): unknown {
    const obj: any = {};
    message.item !== undefined && (obj.item = message.item ? Item.toJSON(message.item) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetItemByChainResponse>, I>>(object: I): QueryGetItemByChainResponse {
    const message = createBaseQueryGetItemByChainResponse();
    message.item = (object.item !== undefined && object.item !== null) ? Item.fromPartial(object.item) : undefined;
    return message;
  },
};

function createBaseQueryAllItemRequest(): QueryAllItemRequest {
  return { pagination: undefined };
}

export const QueryAllItemRequest = {
  encode(message: QueryAllItemRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllItemRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllItemRequest();
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
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllItemRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllItemRequest>, I>>(object: I): QueryAllItemRequest {
    const message = createBaseQueryAllItemRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllItemResponse(): QueryAllItemResponse {
  return { item: [], pagination: undefined };
}

export const QueryAllItemResponse = {
  encode(message: QueryAllItemResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.item) {
      Item.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllItemResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllItemResponse();
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
    return {
      item: Array.isArray(object?.item) ? object.item.map((e: any) => Item.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllItemResponse): unknown {
    const obj: any = {};
    if (message.item) {
      obj.item = message.item.map((e) => e ? Item.toJSON(e) : undefined);
    } else {
      obj.item = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllItemResponse>, I>>(object: I): QueryAllItemResponse {
    const message = createBaseQueryAllItemResponse();
    message.item = object.item?.map((e) => Item.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetInfoRequest(): QueryGetInfoRequest {
  return { index: "" };
}

export const QueryGetInfoRequest = {
  encode(message: QueryGetInfoRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetInfoRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetInfoRequest();
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
    return { index: isSet(object.index) ? String(object.index) : "" };
  },

  toJSON(message: QueryGetInfoRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetInfoRequest>, I>>(object: I): QueryGetInfoRequest {
    const message = createBaseQueryGetInfoRequest();
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseQueryGetInfoResponse(): QueryGetInfoResponse {
  return { info: undefined };
}

export const QueryGetInfoResponse = {
  encode(message: QueryGetInfoResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.info !== undefined) {
      Info.encode(message.info, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetInfoResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetInfoResponse();
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
    return { info: isSet(object.info) ? Info.fromJSON(object.info) : undefined };
  },

  toJSON(message: QueryGetInfoResponse): unknown {
    const obj: any = {};
    message.info !== undefined && (obj.info = message.info ? Info.toJSON(message.info) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetInfoResponse>, I>>(object: I): QueryGetInfoResponse {
    const message = createBaseQueryGetInfoResponse();
    message.info = (object.info !== undefined && object.info !== null) ? Info.fromPartial(object.info) : undefined;
    return message;
  },
};

function createBaseQueryAllInfoRequest(): QueryAllInfoRequest {
  return { pagination: undefined };
}

export const QueryAllInfoRequest = {
  encode(message: QueryAllInfoRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllInfoRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllInfoRequest();
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
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllInfoRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllInfoRequest>, I>>(object: I): QueryAllInfoRequest {
    const message = createBaseQueryAllInfoRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllInfoResponse(): QueryAllInfoResponse {
  return { info: [], pagination: undefined };
}

export const QueryAllInfoResponse = {
  encode(message: QueryAllInfoResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.info) {
      Info.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllInfoResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllInfoResponse();
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
    return {
      info: Array.isArray(object?.info) ? object.info.map((e: any) => Info.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllInfoResponse): unknown {
    const obj: any = {};
    if (message.info) {
      obj.info = message.info.map((e) => e ? Info.toJSON(e) : undefined);
    } else {
      obj.info = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllInfoResponse>, I>>(object: I): QueryAllInfoResponse {
    const message = createBaseQueryAllInfoResponse();
    message.info = object.info?.map((e) => Info.fromPartial(e)) || [];
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
  /** Queries a Item by index. */
  Item(request: QueryGetItemRequest): Promise<QueryGetItemResponse>;
  /** Queries a Item by chain. */
  ItemByChain(request: QueryGetItemByChainRequest): Promise<QueryGetItemByChainResponse>;
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
    this.Params = this.Params.bind(this);
    this.Item = this.Item.bind(this);
    this.ItemByChain = this.ItemByChain.bind(this);
    this.ItemAll = this.ItemAll.bind(this);
    this.Info = this.Info.bind(this);
    this.InfoAll = this.InfoAll.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("rarifyprotocol.rarimocore.tokenmanager.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  Item(request: QueryGetItemRequest): Promise<QueryGetItemResponse> {
    const data = QueryGetItemRequest.encode(request).finish();
    const promise = this.rpc.request("rarifyprotocol.rarimocore.tokenmanager.Query", "Item", data);
    return promise.then((data) => QueryGetItemResponse.decode(new _m0.Reader(data)));
  }

  ItemByChain(request: QueryGetItemByChainRequest): Promise<QueryGetItemByChainResponse> {
    const data = QueryGetItemByChainRequest.encode(request).finish();
    const promise = this.rpc.request("rarifyprotocol.rarimocore.tokenmanager.Query", "ItemByChain", data);
    return promise.then((data) => QueryGetItemByChainResponse.decode(new _m0.Reader(data)));
  }

  ItemAll(request: QueryAllItemRequest): Promise<QueryAllItemResponse> {
    const data = QueryAllItemRequest.encode(request).finish();
    const promise = this.rpc.request("rarifyprotocol.rarimocore.tokenmanager.Query", "ItemAll", data);
    return promise.then((data) => QueryAllItemResponse.decode(new _m0.Reader(data)));
  }

  Info(request: QueryGetInfoRequest): Promise<QueryGetInfoResponse> {
    const data = QueryGetInfoRequest.encode(request).finish();
    const promise = this.rpc.request("rarifyprotocol.rarimocore.tokenmanager.Query", "Info", data);
    return promise.then((data) => QueryGetInfoResponse.decode(new _m0.Reader(data)));
  }

  InfoAll(request: QueryAllInfoRequest): Promise<QueryAllInfoResponse> {
    const data = QueryAllInfoRequest.encode(request).finish();
    const promise = this.rpc.request("rarifyprotocol.rarimocore.tokenmanager.Query", "InfoAll", data);
    return promise.then((data) => QueryAllInfoResponse.decode(new _m0.Reader(data)));
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
