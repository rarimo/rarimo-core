/* eslint-disable */
import { Params } from "../rarimocore/params";
import { Deposit } from "../rarimocore/deposit";
import { Confirmation } from "../rarimocore/confirmation";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "rarifyprotocol.rarimocore.rarimocore";

/** GenesisState defines the rarimocore module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  depositList: Deposit[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  confirmationList: Confirmation[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.depositList) {
      Deposit.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.confirmationList) {
      Confirmation.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.depositList = [];
    message.confirmationList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.depositList.push(Deposit.decode(reader, reader.uint32()));
          break;
        case 3:
          message.confirmationList.push(
            Confirmation.decode(reader, reader.uint32())
          );
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
    message.depositList = [];
    message.confirmationList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.depositList !== undefined && object.depositList !== null) {
      for (const e of object.depositList) {
        message.depositList.push(Deposit.fromJSON(e));
      }
    }
    if (
      object.confirmationList !== undefined &&
      object.confirmationList !== null
    ) {
      for (const e of object.confirmationList) {
        message.confirmationList.push(Confirmation.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.depositList) {
      obj.depositList = message.depositList.map((e) =>
        e ? Deposit.toJSON(e) : undefined
      );
    } else {
      obj.depositList = [];
    }
    if (message.confirmationList) {
      obj.confirmationList = message.confirmationList.map((e) =>
        e ? Confirmation.toJSON(e) : undefined
      );
    } else {
      obj.confirmationList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.depositList = [];
    message.confirmationList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.depositList !== undefined && object.depositList !== null) {
      for (const e of object.depositList) {
        message.depositList.push(Deposit.fromPartial(e));
      }
    }
    if (
      object.confirmationList !== undefined &&
      object.confirmationList !== null
    ) {
      for (const e of object.confirmationList) {
        message.confirmationList.push(Confirmation.fromPartial(e));
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
