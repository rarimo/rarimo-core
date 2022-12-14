/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface ProtobufAny {
  "@type"?: string;
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
}

export interface TokenmanagerChainInfo {
  /** hex-encoded */
  tokenAddress?: string;

  /** hex-encoded */
  tokenId?: string;
}

export interface TokenmanagerChainParams {
  contract?: string;
  types?: Record<string, number>;
  type?: TokenmanagerNetworkType;
}

export interface TokenmanagerInfo {
  index?: string;
  chains?: Record<string, TokenmanagerChainInfo>;
  creator?: string;
}

export interface TokenmanagerItem {
  /** hex-encoded */
  tokenAddress?: string;

  /** hex-encoded */
  tokenId?: string;
  chain?: string;
  index?: string;
  name?: string;
  symbol?: string;
  uri?: string;

  /** @format int64 */
  decimals?: number;

  /** Seed for deriving address for Solana networks. Encoded into hex string. (optional) */
  seed?: string;
  imageUri?: string;

  /** Hash of the token image. Encoded into hex string. (optional) */
  imageHash?: string;
  tokenType?: Tokenmanagertype;
  wrapped?: boolean;
}

export type TokenmanagerMsgAddChainResponse = object;

export type TokenmanagerMsgCreateInfoResponse = object;

export type TokenmanagerMsgDeleteInfoResponse = object;

export enum TokenmanagerNetworkType {
  EVM = "EVM",
  Solana = "Solana",
  Near = "Near",
  Other = "Other",
}

/**
 * Params defines the parameters for the module.
 */
export interface TokenmanagerParams {
  networks?: Record<string, TokenmanagerChainParams>;
}

export interface TokenmanagerQueryAllInfoResponse {
  info?: TokenmanagerInfo[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface TokenmanagerQueryAllItemResponse {
  item?: TokenmanagerItem[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface TokenmanagerQueryGetInfoResponse {
  info?: TokenmanagerInfo;
}

export interface TokenmanagerQueryGetItemByChainResponse {
  item?: TokenmanagerItem;
}

export interface TokenmanagerQueryGetItemResponse {
  item?: TokenmanagerItem;
}

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface TokenmanagerQueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: TokenmanagerParams;
}

export enum Tokenmanagertype {
  NATIVE = "NATIVE",
  ERC20 = "ERC20",
  ERC721 = "ERC721",
  ERC1155 = "ERC1155",
  METAPLEX_NFT = "METAPLEX_NFT",
  METAPLEX_FT = "METAPLEX_FT",
  NEAR_FT = "NEAR_FT",
  NEAR_NFT = "NEAR_NFT",
}

/**
* message SomeRequest {
         Foo some_parameter = 1;
         PageRequest pagination = 2;
 }
*/
export interface V1Beta1PageRequest {
  /**
   * key is a value returned in PageResponse.next_key to begin
   * querying the next page most efficiently. Only one of offset or key
   * should be set.
   * @format byte
   */
  key?: string;

  /**
   * offset is a numeric offset that can be used when key is unavailable.
   * It is less efficient than using key. Only one of offset or key should
   * be set.
   * @format uint64
   */
  offset?: string;

  /**
   * limit is the total number of results to be returned in the result page.
   * If left empty it will default to a value to be set by each app.
   * @format uint64
   */
  limit?: string;

  /**
   * count_total is set to true  to indicate that the result set should include
   * a count of the total number of items available for pagination in UIs.
   * count_total is only respected when offset is used. It is ignored when key
   * is set.
   */
  count_total?: boolean;
}

/**
* PageResponse is to be embedded in gRPC response messages where the
corresponding request message has used PageRequest.

 message SomeResponse {
         repeated Bar results = 1;
         PageResponse page = 2;
 }
*/
export interface V1Beta1PageResponse {
  /**
   * next_key is the key to be passed to PageRequest.key to
   * query the next page most efficiently
   * @format byte
   */
  next_key?: string;

  /**
   * total is total number of results available if PageRequest.count_total
   * was set, its value is undefined otherwise
   * @format uint64
   */
  total?: string;
}

import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse, ResponseType } from "axios";

export type QueryParamsType = Record<string | number, any>;

export interface FullRequestParams extends Omit<AxiosRequestConfig, "data" | "params" | "url" | "responseType"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: ResponseType;
  /** request body */
  body?: unknown;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> extends Omit<AxiosRequestConfig, "data" | "cancelToken"> {
  securityWorker?: (
    securityData: SecurityDataType | null,
  ) => Promise<AxiosRequestConfig | void> | AxiosRequestConfig | void;
  secure?: boolean;
  format?: ResponseType;
}

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public instance: AxiosInstance;
  private securityData: SecurityDataType | null = null;
  private securityWorker?: ApiConfig<SecurityDataType>["securityWorker"];
  private secure?: boolean;
  private format?: ResponseType;

  constructor({ securityWorker, secure, format, ...axiosConfig }: ApiConfig<SecurityDataType> = {}) {
    this.instance = axios.create({ ...axiosConfig, baseURL: axiosConfig.baseURL || "" });
    this.secure = secure;
    this.format = format;
    this.securityWorker = securityWorker;
  }

  public setSecurityData = (data: SecurityDataType | null) => {
    this.securityData = data;
  };

  private mergeRequestParams(params1: AxiosRequestConfig, params2?: AxiosRequestConfig): AxiosRequestConfig {
    return {
      ...this.instance.defaults,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.instance.defaults.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createFormData(input: Record<string, unknown>): FormData {
    return Object.keys(input || {}).reduce((formData, key) => {
      const property = input[key];
      formData.append(
        key,
        property instanceof Blob
          ? property
          : typeof property === "object" && property !== null
          ? JSON.stringify(property)
          : `${property}`,
      );
      return formData;
    }, new FormData());
  }

  public request = async <T = any, _E = any>({
    secure,
    path,
    type,
    query,
    format,
    body,
    ...params
  }: FullRequestParams): Promise<AxiosResponse<T>> => {
    const secureParams =
      ((typeof secure === "boolean" ? secure : this.secure) &&
        this.securityWorker &&
        (await this.securityWorker(this.securityData))) ||
      {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const responseFormat = (format && this.format) || void 0;

    if (type === ContentType.FormData && body && body !== null && typeof body === "object") {
      requestParams.headers.common = { Accept: "*/*" };
      requestParams.headers.post = {};
      requestParams.headers.put = {};

      body = this.createFormData(body as Record<string, unknown>);
    }

    return this.instance.request({
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      params: query,
      responseType: responseFormat,
      data: body,
      url: path,
    });
  };
}

/**
 * @title tokenmanager/genesis.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryInfoAll
   * @summary Queries a list of Info items.
   * @request GET:/rarimo/rarimo-core/tokenmanager/info
   */
  queryInfoAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<TokenmanagerQueryAllInfoResponse, RpcStatus>({
      path: `/rarimo/rarimo-core/tokenmanager/info`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryInfo
   * @summary Queries a Info by index.
   * @request GET:/rarimo/rarimo-core/tokenmanager/info/{index}
   */
  queryInfo = (index: string, params: RequestParams = {}) =>
    this.request<TokenmanagerQueryGetInfoResponse, RpcStatus>({
      path: `/rarimo/rarimo-core/tokenmanager/info/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryItemAll
   * @summary Queries a list of Item items.
   * @request GET:/rarimo/rarimo-core/tokenmanager/item
   */
  queryItemAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<TokenmanagerQueryAllItemResponse, RpcStatus>({
      path: `/rarimo/rarimo-core/tokenmanager/item`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryItem
   * @summary Queries a Item by index.
   * @request GET:/rarimo/rarimo-core/tokenmanager/item/{chain}
   */
  queryItem = (chain: string, query?: { tokenAddress?: string; tokenId?: string }, params: RequestParams = {}) =>
    this.request<TokenmanagerQueryGetItemResponse, RpcStatus>({
      path: `/rarimo/rarimo-core/tokenmanager/item/${chain}`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryItemByChain
   * @summary Queries a Item by chain.
   * @request GET:/rarimo/rarimo-core/tokenmanager/item/{infoIndex}/{chain}
   */
  queryItemByChain = (infoIndex: string, chain: string, params: RequestParams = {}) =>
    this.request<TokenmanagerQueryGetItemByChainResponse, RpcStatus>({
      path: `/rarimo/rarimo-core/tokenmanager/item/${infoIndex}/${chain}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryParams
   * @summary Parameters queries the parameters of the module.
   * @request GET:/rarimo/rarimo-core/tokenmanager/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<TokenmanagerQueryParamsResponse, RpcStatus>({
      path: `/rarimo/rarimo-core/tokenmanager/params`,
      method: "GET",
      format: "json",
      ...params,
    });
}
