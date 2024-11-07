#!/usr/bin/env node

import { Command, program } from "commander";
import dotenv from "dotenv";
import { ethers } from "ethers";
import axios from "axios";

dotenv.config();

async function getSafeMessageDigest(message, chainId, safeAddress) {
  const domain = {
    chainId: chainId,
    verifyingContract: safeAddress,
  };

  const types = {
    SafeMessage: [{ name: "message", type: "bytes" }],
  };

  const value = {
    message: message,
  };

  const signature = await ethers.utils._TypedDataEncoder.hash(
    domain,
    types,
    value
  );

  return signature;
}
function adjustSignatureV(signature) {
  if (signature.length !== 132) {
    throw new Error(
      "Invalid signature length. Expected 132 characters (66 bytes)"
    );
  }

  // Extract the V value (last 2 characters of the hex string)
  let v = parseInt(signature.slice(-2), 16);

  // Adjust V if it's 0 or 1
  if (v === 0 || v === 1) {
    v += 27;
    // Convert back to hex, pad to 2 characters, and ensure lowercase
    const newV = v.toString(16).padStart(2, "0").toLowerCase();
    // Replace the last 2 characters of the signature with the new V value
    signature = signature.slice(0, -2) + newV;
  }

  return signature;
}

const initTimeStamp = new Date().getTime();
// const initTimeStamp = 1;

const metadata = {
  addressTags: JSON.parse(process.env.ADDRESS_TAGS || "{}"),
};

const executor = {
  domain: {
    chainId: parseInt(process.env.CHAIN_ID || "1"),
  },
  message: {
    timestamp: initTimeStamp,
    executor: process.env.EXECUTOR_ADDRESS,
    inputTokens: JSON.parse(process.env.INPUT_TOKENS || "[]"),
    hopAddresses: JSON.parse(process.env.HOP_ADDRESSES || "[]"),
    feeInBPS: parseInt(process.env.FEE_IN_BPS || "0"),
    feeToken: process.env.FEE_TOKEN,
    feeReceiver: process.env.FEE_RECEIVER,
    limitPerExecution: process.env.LIMIT_PER_EXECUTION === "true",
    clientId: process.env.CLIENT_ID,
  },
  primaryType: "RegisterExecutor",
  types: {
    RegisterExecutor: [
      { name: "timestamp", type: "uint256" },
      { name: "executor", type: "address" },
      { name: "inputTokens", type: "address[]" },
      { name: "hopAddresses", type: "address[]" },
      { name: "feeInBPS", type: "uint256" },
      { name: "feeToken", type: "address" },
      { name: "feeReceiver", type: "address" },
      { name: "limitPerExecution", type: "bool" },
      { name: "clientId", type: "string" },
    ],
  },
};

console.log(executor);

const payload = {
  config: {
    inputTokens: executor.message.inputTokens,
    hopAddresses: executor.message.hopAddresses,
    feeInBPS: executor.message.feeInBPS,
    feeToken: executor.message.feeToken,
    feeReceiver: executor.message.feeReceiver,
    limitPerExecution: executor.message.limitPerExecution,
  },
  executor: executor.message.executor,
  signature: "",
  chainId: executor.domain.chainId,
  timestamp: 0,
  executorMetadata: {
    id: executor.message.clientId,
    name: process.env.EXECUTOR_NAME,
    logo: process.env.EXECUTOR_LOGO,
    metadata: metadata,
  },
};

program
  .version("1.0.0")
  .addCommand(
    new Command().name("generate").action(async () => {
      const dataHash = ethers.utils._TypedDataEncoder.hash(
        executor.domain,
        executor.types,
        executor.message
      );

      console.log("actual", dataHash);
      console.log(
        dataHash,
        executor.domain.chainId,
        process.env.EXECUTOR_ADDRESS
      );

      console.log({
        timestamp: initTimeStamp,
        dataHash,
        cmd: `vault write ethereum/key-managers/brahma-builder/sign address='${
          payload.executor
        }' hash='${await getSafeMessageDigest(
          dataHash,
          executor.domain.chainId,
          process.env.EXECUTOR_ADDRESS
        )}'`,
      });
    })
  )
  .description("Generate Executor Digest To Sign")
  .addCommand(
    new Command()
      .name("submit")
      .description("Submit signed payload to API")
      .requiredOption("-s, --signature <string>", "Signature string")
      .requiredOption("-t, --timestamp <number>", "Timestamp")
      .action(async (options) => {
        const { signature, timestamp } = options;

        // Update payload with signature and timestamp
        payload.signature = adjustSignatureV(signature);
        payload.timestamp = parseInt(timestamp);

        console.log(payload);

        // Get API base URL from environment variable
        const apiBaseUrl = process.env.API_BASE_URL;
        if (!apiBaseUrl) {
          console.error("API_BASE_URL is not set in the environment variables");
          process.exit(1);
        }

        try {
          const response = await axios.post(
            `${apiBaseUrl}/v1/automations/executor`,
            payload
          );
          console.log("registered executor:", response.data);
        } catch (error) {
          console.error(
            "error registering executor:",
            error.response ? error.response.data : error.message
          );
        }
      })
  );
program.parse(process.argv);
