/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "rarifyprotocol.rarimocore.rarimocore";

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

export interface Deposit {
  tx: string;
  fromChain: string;
  toChain: string;
  receiver: string;
  tokenAddress: string;
  tokenId: string;
  creator: string;
  signed: boolean;
  tokenType: type;
}

const baseDeposit: object = {
  tx: "",
  fromChain: "",
  toChain: "",
  receiver: "",
  tokenAddress: "",
  tokenId: "",
  creator: "",
  signed: false,
  tokenType: 0,
};

export const Deposit = {
  encode(message: Deposit, writer: Writer = Writer.create()): Writer {
    if (message.tx !== "") {
      writer.uint32(10).string(message.tx);
    }
    if (message.fromChain !== "") {
      writer.uint32(18).string(message.fromChain);
    }
    if (message.toChain !== "") {
      writer.uint32(26).string(message.toChain);
    }
    if (message.receiver !== "") {
      writer.uint32(34).string(message.receiver);
    }
    if (message.tokenAddress !== "") {
      writer.uint32(42).string(message.tokenAddress);
    }
    if (message.tokenId !== "") {
      writer.uint32(50).string(message.tokenId);
    }
    if (message.creator !== "") {
      writer.uint32(58).string(message.creator);
    }
    if (message.signed === true) {
      writer.uint32(64).bool(message.signed);
    }
    if (message.tokenType !== 0) {
      writer.uint32(72).int32(message.tokenType);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Deposit {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseDeposit } as Deposit;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.tx = reader.string();
          break;
        case 2:
          message.fromChain = reader.string();
          break;
        case 3:
          message.toChain = reader.string();
          break;
        case 4:
          message.receiver = reader.string();
          break;
        case 5:
          message.tokenAddress = reader.string();
          break;
        case 6:
          message.tokenId = reader.string();
          break;
        case 7:
          message.creator = reader.string();
          break;
        case 8:
          message.signed = reader.bool();
          break;
        case 9:
          message.tokenType = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Deposit {
    const message = { ...baseDeposit } as Deposit;
    if (object.tx !== undefined && object.tx !== null) {
      message.tx = String(object.tx);
    } else {
      message.tx = "";
    }
    if (object.fromChain !== undefined && object.fromChain !== null) {
      message.fromChain = String(object.fromChain);
    } else {
      message.fromChain = "";
    }
    if (object.toChain !== undefined && object.toChain !== null) {
      message.toChain = String(object.toChain);
    } else {
      message.toChain = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = String(object.receiver);
    } else {
      message.receiver = "";
    }
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
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.signed !== undefined && object.signed !== null) {
      message.signed = Boolean(object.signed);
    } else {
      message.signed = false;
    }
    if (object.tokenType !== undefined && object.tokenType !== null) {
      message.tokenType = typeFromJSON(object.tokenType);
    } else {
      message.tokenType = 0;
    }
    return message;
  },

  toJSON(message: Deposit): unknown {
    const obj: any = {};
    message.tx !== undefined && (obj.tx = message.tx);
    message.fromChain !== undefined && (obj.fromChain = message.fromChain);
    message.toChain !== undefined && (obj.toChain = message.toChain);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.tokenAddress !== undefined &&
      (obj.tokenAddress = message.tokenAddress);
    message.tokenId !== undefined && (obj.tokenId = message.tokenId);
    message.creator !== undefined && (obj.creator = message.creator);
    message.signed !== undefined && (obj.signed = message.signed);
    message.tokenType !== undefined &&
      (obj.tokenType = typeToJSON(message.tokenType));
    return obj;
  },

  fromPartial(object: DeepPartial<Deposit>): Deposit {
    const message = { ...baseDeposit } as Deposit;
    if (object.tx !== undefined && object.tx !== null) {
      message.tx = object.tx;
    } else {
      message.tx = "";
    }
    if (object.fromChain !== undefined && object.fromChain !== null) {
      message.fromChain = object.fromChain;
    } else {
      message.fromChain = "";
    }
    if (object.toChain !== undefined && object.toChain !== null) {
      message.toChain = object.toChain;
    } else {
      message.toChain = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = object.receiver;
    } else {
      message.receiver = "";
    }
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
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.signed !== undefined && object.signed !== null) {
      message.signed = object.signed;
    } else {
      message.signed = false;
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
