import { ethers } from "ethermint/ethminttests/tooling/ethers/ethers";

// connects to localhost:8545
const provider = new ethers.providers.JsonRpcProvider()

// const signer = provider.getSigner();

const blockNumber = await provider.getBlockNumber()