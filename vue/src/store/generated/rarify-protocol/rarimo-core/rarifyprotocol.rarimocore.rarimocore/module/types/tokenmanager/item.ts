/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "rarifyprotocol.rarimocore.tokenmanager";

export enum type {
  NATIVE = 0,
  ERC20 = 1,
  ERC721 = 2,
  ERC1155 = 3,
  METAPLEX_NFT = 4,
  METAPLEX_FT = 5,
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
    default:
      return "UNKNOWN";
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
  /** Seed for deriving address for Solana networks. Encoded into hex string */
  seed: string;
  tokenType: type;
}

const baseItem: object = {
  tokenAddress: "",
  tokenId: "",
  chain: "",
  index: "",
  name: "",
  symbol: "",
  uri: "",
  decimals: 0,
  seed: "",
  tokenType: 0,
};

export const Item = {
  encode(message: Item, writer: Writer = Writer.create()): Writer {
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
    if (message.tokenType !== 0) {
      writer.uint32(80).int32(message.tokenType);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Item {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseItem } as Item;
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
          message.tokenType = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Item {
    const message = { ...baseItem } as Item;
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
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.symbol !== undefined && object.symbol !== null) {
      message.symbol = String(object.symbol);
    } else {
      message.symbol = "";
    }
    if (object.uri !== undefined && object.uri !== null) {
      message.uri = String(object.uri);
    } else {
      message.uri = "";
    }
    if (object.decimals !== undefined && object.decimals !== null) {
      message.decimals = Number(object.decimals);
    } else {
      message.decimals = 0;
    }
    if (object.seed !== undefined && object.seed !== null) {
      message.seed = String(object.seed);
    } else {
      message.seed = "";
    }
    if (object.tokenType !== undefined && object.tokenType !== null) {
      message.tokenType = typeFromJSON(object.tokenType);
    } else {
      message.tokenType = 0;
    }
    return message;
  },

  toJSON(message: Item): unknown {
    const obj: any = {};
    message.tokenAddress !== undefined &&
      (obj.tokenAddress = message.tokenAddress);
    message.tokenId !== undefined && (obj.tokenId = message.tokenId);
    message.chain !== undefined && (obj.chain = message.chain);
    message.index !== undefined && (obj.index = message.index);
    message.name !== undefined && (obj.name = message.name);
    message.symbol !== undefined && (obj.symbol = message.symbol);
    message.uri !== undefined && (obj.uri = message.uri);
    message.decimals !== undefined && (obj.decimals = message.decimals);
    message.seed !== undefined && (obj.seed = message.seed);
    message.tokenType !== undefined &&
      (obj.tokenType = typeToJSON(message.tokenType));
    return obj;
  },

  fromPartial(object: DeepPartial<Item>): Item {
    const message = { ...baseItem } as Item;
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
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.symbol !== undefined && object.symbol !== null) {
      message.symbol = object.symbol;
    } else {
      message.symbol = "";
    }
    if (object.uri !== undefined && object.uri !== null) {
      message.uri = object.uri;
    } else {
      message.uri = "";
    }
    if (object.decimals !== undefined && object.decimals !== null) {
      message.decimals = object.decimals;
    } else {
      message.decimals = 0;
    }
    if (object.seed !== undefined && object.seed !== null) {
      message.seed = object.seed;
    } else {
      message.seed = "";
    }
    if (object.tokenType !== undefined && object.tokenType !== null) {
      message.tokenType = object.tokenType;
    } else {
      message.tokenType = 0;
    }
    return message;
  },
};

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
