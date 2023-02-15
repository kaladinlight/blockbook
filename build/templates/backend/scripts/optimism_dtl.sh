#!/bin/sh

{{define "main" -}}

set -e

export DATA_TRANSPORT_LAYER__ADDRESS_MANAGER=0xdE1FCfB0851916CA5101820A69b13a4E276bd81F
export DATA_TRANSPORT_LAYER__DANGEROUSLY_CATCH_ALL_ERRORS=true
export DATA_TRANSPORT_LAYER__DB_PATH={{.Env.BackendDataPath}}/{{.Coin.Alias}}/backend
export DATA_TRANSPORT_LAYER__SYNC_FROM_L1=false
export DATA_TRANSPORT_LAYER__SYNC_FROM_L2=true
export DATA_TRANSPORT_LAYER__DEFAULT_BACKEND=l2
export DATA_TRANSPORT_LAYER__L1_GAS_PRICE_BACKEND=l2
export DATA_TRANSPORT_LAYER__ETH_NETWORK_NAME=mainnet
export DATA_TRANSPORT_LAYER__L2_CHAIN_ID=10
export DATA_TRANSPORT_LAYER__NODE_ENV=production
export DATA_TRANSPORT_LAYER__POLLING_INTERVAL=500
export DATA_TRANSPORT_LAYER__SERVER_HOSTNAME=127.0.0.1
export DATA_TRANSPORT_LAYER__SERVER_PORT={{.Ports.BackendHttp}}
export DATA_TRANSPORT_LAYER__L1_START_HEIGHT=13596466
export DATA_TRANSPORT_LAYER__L2_RPC_ENDPOINT=https://mainnet.optimism.io

node {{.Env.BackendInstallPath}}/{{.Coin.Alias}}/source/packages/data-transport-layer/dist/src/services/run.js

{{end}}