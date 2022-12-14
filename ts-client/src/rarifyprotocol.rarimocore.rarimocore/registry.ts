import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateConfirmation } from "./types/rarimocore/tx";
import { MsgCreateTransferOp } from "./types/rarimocore/tx";
import { MsgCreateChangeKeyECDSA } from "./types/rarimocore/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/rarifyprotocol.rarimocore.rarimocore.MsgCreateConfirmation", MsgCreateConfirmation],
    ["/rarifyprotocol.rarimocore.rarimocore.MsgCreateTransferOp", MsgCreateTransferOp],
    ["/rarifyprotocol.rarimocore.rarimocore.MsgCreateChangeKeyECDSA", MsgCreateChangeKeyECDSA],
    
];

export { msgTypes }