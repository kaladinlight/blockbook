#!/bin/sh

{{define "main" -}}

set -e

INSTALL_DIR={{.Env.BackendInstallPath}}/{{.Coin.Alias}}
DATA_DIR={{.Env.BackendDataPath}}/{{.Coin.Alias}}/backend

NITRO_BIN=$INSTALL_DIR/nitro

$NITRO_BIN \
  --init.url https://snapshot.arbitrum.io/mainnet/nitro.tar \
  --init.download-path $DATA_DIR/.arbitrum \
  --auth.jwtsecret $DATA_DIR/jwtsecret \
  --l1.url http://127.0.0.1:8136 \
  --l2.chain-id 42161 \
  --http.addr 127.0.0.1 \
  --http.port {{.Ports.BackendHttp}} \
  --http.api eth,net,web3,debug,txpool,arb \
  --http.vhosts '*' \
  --http.corsdomain '*' \
  --ws.addr 127.0.0.1 \
  --ws.api eth,net,web3,debug,txpool,arb \
  --ws.port {{.Ports.BackendRPC}} \
  --ws.origins '*' &

{{end}}