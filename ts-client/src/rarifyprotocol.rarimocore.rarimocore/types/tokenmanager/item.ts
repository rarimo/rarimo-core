/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "rarifyprotocol.rarimocore.tokenmanager";

export enum type {
  NATIVE = 0,
  ERC20 = 1,
  ERC721 = 2,
  ERC1155 = 3,
  METAPLEX_NFT = 4,
  METAPLEX_FT = 5,
  NEAR_FT = 6,
  NEAR_NFT = 7,
  UNRECOGNIZED = -1,
}

export function typeFromJSON(object: any): type {
  switch (object) {
    case 0:
    case "NATIVE":
      return type.NATIVE;
    case 1:
    case "ERC20":
      return type.ERC20;
    case 2:
    case "ERC721":
      return type.ERC721;
    case 3:
    case "ERC1155":
      return type.ERC1155;
    case 4:
    case "METAPLEX_NFT":
      return type.METAPLEX_NFT;
    case 5:
    case "METAPLEX_FT":
      return type.METAPLEX_FT;
    case 6:
    case "NEAR_FT":
      return type.NEAR_FT;
    case 7:
    case "NEAR_NFT":
      return type.NEAR_NFT;
    case -1:
    case "UNRECOGNIZED":
    default:
      return type.UNRECOGNIZED;
  }
}

export function typeToJSON(object: type): string {
  switch (object) {
    case type.NATIVE:
      return "NATIVE";
    case type.ERC20:
      return "ERC20";
    case type.ERC721:
      return "ERC721";
    case type.ERC1155:
      return "ERC1155";
    case type.METAPLEX_NFT:
      return "METAPLEX_NFT";
    case type.METAPLEX_FT:
      return "METAPLEX_FT";
    case type.NEAR_FT:
      return "NEAR_FT";
    case type.NEAR_NFT:
      return "NEAR_NFT";
    case type.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface Item {
  /** hex-encoded */
  tokenAddress: string;
  /** hex-encoded */
  tokenId: string;
  chain: string;
  index: string;
  name: string;
  symbol: string;
  uri: string;
  decimals: number;
  /** Seed for deriving address for Solana networks. Encoded into hex string. (optional) */
  seed: string;
  imageUri: string;
  /** Hash of the token image. Encoded into hex string. (optional) */
  imageHash: string;
  tokenType: type;
  wrapped: boolean;
}

function createBaseItem(): Item {
  return {
    tokenAddress: "",
    tokenId: "",
    chain: "",
    index: "",
    name: "",
    symbol: "",
    uri: "",
    decimals: 0,
    seed: "",
    imageUri: "",
    imageHash: "",
    tokenType: 0,
    wrapped: false,
  };
}

export const Item = {
  encode(message: Item, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.tokenAddress !== "") {
      writer.uint32(10).string(message.tokenAddress);
    }
    if (message.tokenId !== "") {
      writer.uint32(18).string(message.tokenId);
    }
    if (message.chain !== "") {
      writer.uint32(26).string(message.chain);
    }
    if (message.index !== "") {
      writer.uint32(34).string(message.index);
    }
    if (message.name !== "") {
      writer.uint32(42).string(message.name);
    }
    if (message.symbol !== "") {
      writer.uint32(50).string(message.symbol);
    }
    if (message.uri !== "") {
      writer.uint32(58).string(message.uri);
    }
    if (message.decimals !== 0) {
      writer.uint32(64).uint32(message.decimals);
    }
    if (message.seed !== "") {
      writer.uint32(74).string(message.seed);
    }
    if (message.imageUri !== "") {
      writer.uint32(82).string(message.imageUri);
    }
    if (message.imageHash !== "") {
      writer.uint32(90).string(message.imageHash);
    }
    if (message.tokenType !== 0) {
      writer.uint32(96).int32(message.tokenType);
    }
    if (message.wrapped === true) {
      writer.uint32(104).bool(message.wrapped);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Item {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseItem();
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
        case 4:
          message.index = reader.string();
          break;
        case 5:
          message.name = reader.string();
          break;
        case 6:
          message.symbol = reader.string();
          break;
        case 7:
          message.uri = reader.string();
          break;
        case 8:
          message.decimals = reader.uint32();
          break;
        case 9:
          message.seed = reader.string();
          break;
        case 10:
          message.imageUri = reader.string();
          break;
        case 11:
          message.imageHash = reader.string();
          break;
        case 12:
          message.tokenType = reader.int32() as any;
          break;
        case 13:
          message.wrapped = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Item {
    return {
      tokenAddress: isSet(object.tokenAddress) ? String(object.tokenAddress) : "",
      tokenId: isSet(object.tokenId) ? String(object.tokenId) : "",
      chain: isSet(object.chain) ? String(object.chain) : "",
      index: isSet(object.index) ? String(object.index) : "",
      name: isSet(object.name) ? String(object.name) : "",
      symbol: isSet(object.symbol) ? String(object.symbol) : "",
      uri: isSet(object.uri) ? String(object.uri) : "",
      decimals: isSet(object.decimals) ? Number(object.decimals) : 0,
      seed: isSet(object.seed) ? String(object.seed) : "",
      imageUri: isSet(object.imageUri) ? String(object.imageUri) : "",
      imageHash: isSet(object.imageHash) ? String(object.imageHash) : "",
      tokenType: isSet(object.tokenType) ? typeFromJSON(object.tokenType) : 0,
      wrapped: isSet(object.wrapped) ? Boolean(object.wrapped) : false,
    };
  },

  toJSON(message: Item): unknown {
    const obj: any = {};
    message.tokenAddress !== undefined && (obj.tokenAddress = message.tokenAddress);
    message.tokenId !== undefined && (obj.tokenId = message.tokenId);
    message.chain !== undefined && (obj.chain = message.chain);
    message.index !== undefined && (obj.index = message.index);
    message.name !== undefined && (obj.name = message.name);
    message.symbol !== undefined && (obj.symbol = message.symbol);
    message.uri !== undefined && (obj.uri = message.uri);
    message.decimals !== undefined && (obj.decimals = Math.round(message.decimals));
    message.seed !== undefined && (obj.seed = message.seed);
    message.imageUri !== undefined && (obj.imageUri = message.imageUri);
    message.imageHash !== undefined && (obj.imageHash = message.imageHash);
    message.tokenType !== undefined && (obj.tokenType = typeToJSON(message.tokenType));
    message.wrapped !== undefined && (obj.wrapped = message.wrapped);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Item>, I>>(object: I): Item {
    const message = createBaseItem();
    message.tokenAddress = object.tokenAddress ?? "";
    message.tokenId = object.tokenId ?? "";
    message.chain = object.chain ?? "";
    message.index = object.index ?? "";
    message.name = object.name ?? "";
    message.symbol = object.symbol ?? "";
    message.uri = object.uri ?? "";
    message.decimals = object.decimals ?? 0;
    message.seed = object.seed ?? "";
    message.imageUri = object.imageUri ?? "";
    message.imageHash = object.imageHash ?? "";
    message.tokenType = object.tokenType ?? 0;
    message.wrapped = object.wrapped ?? false;
    return message;
  },
};

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
