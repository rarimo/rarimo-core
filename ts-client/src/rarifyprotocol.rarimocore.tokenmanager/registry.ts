import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgAddChain } from "./types/tokenmanager/tx";
import { MsgCreateInfo } from "./types/tokenmanager/tx";
import { MsgDeleteInfo } from "./types/tokenmanager/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/rarifyprotocol.rarimocore.tokenmanager.MsgAddChain", MsgAddChain],
    ["/rarifyprotocol.rarimocore.tokenmanager.MsgCreateInfo", MsgCreateInfo],
    ["/rarifyprotocol.rarimocore.tokenmanager.MsgDeleteInfo", MsgDeleteInfo],
    
];

export { msgTypes }