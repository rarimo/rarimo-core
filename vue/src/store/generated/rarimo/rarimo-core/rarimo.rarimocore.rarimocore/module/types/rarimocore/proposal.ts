/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "rarimo.rarimocore.rarimocore";

export interface AddSignerPartyProposal {
  title: string;
  description: string;
  account: string;
  address: string;
  trialPublicKey: string;
}

export interface RemoveSignerPartyProposal {
  title: string;
  description: string;
  account: string;
}

const baseAddSignerPartyProposal: object = {
  title: "",
  description: "",
  account: "",
  address: "",
  trialPublicKey: "",
};

export const AddSignerPartyProposal = {
  encode(
    message: AddSignerPartyProposal,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.title !== "") {
      writer.uint32(10).string(message.title);
    }
    if (message.description !== "") {
      writer.uint32(18).string(message.description);
    }
    if (message.account !== "") {
      writer.uint32(26).string(message.account);
    }
    if (message.address !== "") {
      writer.uint32(34).string(message.address);
    }
    if (message.trialPublicKey !== "") {
      writer.uint32(42).string(message.trialPublicKey);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): AddSignerPartyProposal {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseAddSignerPartyProposal } as AddSignerPartyProposal;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.title = reader.string();
          break;
        case 2:
          message.description = reader.string();
          break;
        case 3:
          message.account = reader.string();
          break;
        case 4:
          message.address = reader.string();
          break;
        case 5:
          message.trialPublicKey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AddSignerPartyProposal {
    const message = { ...baseAddSignerPartyProposal } as AddSignerPartyProposal;
    if (object.title !== undefined && object.title !== null) {
      message.title = String(object.title);
    } else {
      message.title = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.account !== undefined && object.account !== null) {
      message.account = String(object.account);
    } else {
      message.account = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.trialPublicKey !== undefined && object.trialPublicKey !== null) {
      message.trialPublicKey = String(object.trialPublicKey);
    } else {
      message.trialPublicKey = "";
    }
    return message;
  },

  toJSON(message: AddSignerPartyProposal): unknown {
    const obj: any = {};
    message.title !== undefined && (obj.title = message.title);
    message.description !== undefined &&
      (obj.description = message.description);
    message.account !== undefined && (obj.account = message.account);
    message.address !== undefined && (obj.address = message.address);
    message.trialPublicKey !== undefined &&
      (obj.trialPublicKey = message.trialPublicKey);
    return obj;
  },

  fromPartial(
    object: DeepPartial<AddSignerPartyProposal>
  ): AddSignerPartyProposal {
    const message = { ...baseAddSignerPartyProposal } as AddSignerPartyProposal;
    if (object.title !== undefined && object.title !== null) {
      message.title = object.title;
    } else {
      message.title = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.account !== undefined && object.account !== null) {
      message.account = object.account;
    } else {
      message.account = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.trialPublicKey !== undefined && object.trialPublicKey !== null) {
      message.trialPublicKey = object.trialPublicKey;
    } else {
      message.trialPublicKey = "";
    }
    return message;
  },
};

const baseRemoveSignerPartyProposal: object = {
  title: "",
  description: "",
  account: "",
};

export const RemoveSignerPartyProposal = {
  encode(
    message: RemoveSignerPartyProposal,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.title !== "") {
      writer.uint32(10).string(message.title);
    }
    if (message.description !== "") {
      writer.uint32(18).string(message.description);
    }
    if (message.account !== "") {
      writer.uint32(26).string(message.account);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): RemoveSignerPartyProposal {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseRemoveSignerPartyProposal,
    } as RemoveSignerPartyProposal;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.title = reader.string();
          break;
        case 2:
          message.description = reader.string();
          break;
        case 3:
          message.account = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RemoveSignerPartyProposal {
    const message = {
      ...baseRemoveSignerPartyProposal,
    } as RemoveSignerPartyProposal;
    if (object.title !== undefined && object.title !== null) {
      message.title = String(object.title);
    } else {
      message.title = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.account !== undefined && object.account !== null) {
      message.account = String(object.account);
    } else {
      message.account = "";
    }
    return message;
  },

  toJSON(message: RemoveSignerPartyProposal): unknown {
    const obj: any = {};
    message.title !== undefined && (obj.title = message.title);
    message.description !== undefined &&
      (obj.description = message.description);
    message.account !== undefined && (obj.account = message.account);
    return obj;
  },

  fromPartial(
    object: DeepPartial<RemoveSignerPartyProposal>
  ): RemoveSignerPartyProposal {
    const message = {
      ...baseRemoveSignerPartyProposal,
    } as RemoveSignerPartyProposal;
    if (object.title !== undefined && object.title !== null) {
      message.title = object.title;
    } else {
      message.title = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.account !== undefined && object.account !== null) {
      message.account = object.account;
    } else {
      message.account = "";
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
