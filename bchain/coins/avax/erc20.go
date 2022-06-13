package avax

import (
	"context"
	"math/big"
	"sync"

	"github.com/golang/glog"
	"github.com/juju/errors"
	"github.com/trezor/blockbook/bchain"
	"github.com/trezor/blockbook/bchain/coins/eth"
)

// doing the parsing/processing without using go-ethereum/accounts/abi library, it is simple to get data from Transfer event
const erc20TransferEventSignature = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
const erc20NameSignature = "0x06fdde03"
const erc20SymbolSignature = "0x95d89b41"
const erc20DecimalsSignature = "0x313ce567"
const erc20BalanceOf = "0x70a08231"

var cachedContracts = make(map[string]*bchain.Erc20Contract)
var cachedContractsMux sync.Mutex

func (b *AvalancheRPC) ethCall(data, to string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), b.timeout)
	defer cancel()
	var r string
	err := b.rpc.CallContext(ctx, &r, "eth_call", map[string]interface{}{
		"data": data,
		"to":   to,
	}, "latest")
	if err != nil {
		return "", err
	}
	return r, nil
}

// EthereumTypeGetErc20ContractInfo returns information about ERC20 contract
func (b *AvalancheRPC) EthereumTypeGetErc20ContractInfo(contractDesc bchain.AddressDescriptor) (*bchain.Erc20Contract, error) {
	cds := string(contractDesc)
	cachedContractsMux.Lock()
	contract, found := cachedContracts[cds]
	cachedContractsMux.Unlock()
	if !found {
		address := eth.EIP55Address(contractDesc)
		data, err := b.ethCall(erc20NameSignature, address)
		if err != nil {
			// ignore the error from the eth_call - since geth v1.9.15 they changed the behavior
			// and returning error "execution reverted" for some non contract addresses
			// https://github.com/ethereum/go-ethereum/issues/21249#issuecomment-648647672
			glog.Warning(errors.Annotatef(err, "erc20NameSignature %v", address))
			return nil, nil
			// return nil, errors.Annotatef(err, "erc20NameSignature %v", address)
		}
		name := eth.ParseErc20StringProperty(contractDesc, data)
		if name != "" {
			data, err = b.ethCall(erc20SymbolSignature, address)
			if err != nil {
				glog.Warning(errors.Annotatef(err, "erc20SymbolSignature %v", address))
				return nil, nil
				// return nil, errors.Annotatef(err, "erc20SymbolSignature %v", address)
			}
			symbol := eth.ParseErc20StringProperty(contractDesc, data)
			data, err = b.ethCall(erc20DecimalsSignature, address)
			if err != nil {
				glog.Warning(errors.Annotatef(err, "erc20DecimalsSignature %v", address))
				// return nil, errors.Annotatef(err, "erc20DecimalsSignature %v", address)
			}
			contract = &bchain.Erc20Contract{
				Contract: address,
				Name:     name,
				Symbol:   symbol,
			}
			d := eth.ParseErc20NumericProperty(contractDesc, data)
			if d != nil {
				contract.Decimals = int(uint8(d.Uint64()))
			} else {
				contract.Decimals = eth.EtherAmountDecimalPoint
			}
		} else {
			contract = nil
		}
		cachedContractsMux.Lock()
		cachedContracts[cds] = contract
		cachedContractsMux.Unlock()
	}
	return contract, nil
}

// EthereumTypeGetErc20ContractBalance returns balance of ERC20 contract for given address
func (b *AvalancheRPC) EthereumTypeGetErc20ContractBalance(addrDesc, contractDesc bchain.AddressDescriptor) (*big.Int, error) {
	addr := eth.EIP55Address(addrDesc)
	contract := eth.EIP55Address(contractDesc)
	req := erc20BalanceOf + "0000000000000000000000000000000000000000000000000000000000000000"[len(addr)-2:] + addr[2:]
	data, err := b.ethCall(req, contract)
	if err != nil {
		return nil, err
	}
	r := eth.ParseErc20NumericProperty(contractDesc, data)
	if r == nil {
		return nil, errors.New("Invalid balance")
	}
	return r, nil
}
