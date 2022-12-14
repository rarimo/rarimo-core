/* eslint-disable */
import { Params } from "../tokenmanager/params";
import { Item } from "../tokenmanager/item";
import { Info } from "../tokenmanager/info";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "rarimo.rarimocore.tokenmanager";

/** GenesisState defines the tokenmanager module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  itemList: Item[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  infoList: Info[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.itemList) {
      Item.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.infoList) {
      Info.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.itemList = [];
    message.infoList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.itemList.push(Item.decode(reader, reader.uint32()));
          break;
        case 3:
          message.infoList.push(Info.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.itemList = [];
    message.infoList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.itemList !== undefined && object.itemList !== null) {
      for (const e of object.itemList) {
        message.itemList.push(Item.fromJSON(e));
      }
    }
    if (object.infoList !== undefined && object.infoList !== null) {
      for (const e of object.infoList) {
        message.infoList.push(Info.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.itemList) {
      obj.itemList = message.itemList.map((e) =>
        e ? Item.toJSON(e) : undefined
      );
    } else {
      obj.itemList = [];
    }
    if (message.infoList) {
      obj.infoList = message.infoList.map((e) =>
        e ? Info.toJSON(e) : undefined
      );
    } else {
      obj.infoList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.itemList = [];
    message.infoList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.itemList !== undefined && object.itemList !== null) {
      for (const e of object.itemList) {
        message.itemList.push(Item.fromPartial(e));
      }
    }
    if (object.infoList !== undefined && object.infoList !== null) {
      for (const e of object.infoList) {
        message.infoList.push(Info.fromPartial(e));
      }
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
