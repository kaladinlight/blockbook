package gnosis

import (
	"encoding/json"
	"errors"
	"io"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
)

var hasherPool = sync.Pool{
	New: func() interface{} { return sha3.NewLegacyKeccak256() },
}

// rlpHash encodes x and hashes the encoded bytes.
func rlpHash(x interface{}) (h common.Hash) {
	sha := hasherPool.Get().(crypto.KeccakState)
	defer hasherPool.Put(sha)
	sha.Reset()
	_ = rlp.Encode(sha, x)
	_, _ = sha.Read(h[:])
	return h
}

type Header struct {
	ParentHash  common.Hash    `json:"parentHash"       gencodec:"required"`
	UncleHash   common.Hash    `json:"sha3Uncles"       gencodec:"required"`
	Coinbase    common.Address `json:"author"`
	Root        common.Hash    `json:"stateRoot"        gencodec:"required"`
	TxHash      common.Hash    `json:"transactionsRoot" gencodec:"required"`
	ReceiptHash common.Hash    `json:"receiptsRoot"     gencodec:"required"`
	Bloom       types.Bloom    `json:"logsBloom"        gencodec:"required"`
	Difficulty  *big.Int       `json:"difficulty"       gencodec:"required"`
	Number      *big.Int       `json:"number"           gencodec:"required"`
	GasLimit    uint64         `json:"gasLimit"         gencodec:"required"`
	GasUsed     uint64         `json:"gasUsed"          gencodec:"required"`
	Time        uint64         `json:"timestamp"        gencodec:"required"`
	Extra       []byte         `json:"extraData"        gencodec:"required"`

	MixDigest *common.Hash      `json:"mixHash" rlp:"optional"`
	Nonce     *types.BlockNonce `json:"nonce" rlp:"optional"`

	// AuRa POA Consensus
	AuraStep      *big.Int `json:"step" rlp:"optional"`
	AuraSignature []byte   `json:"signature" rlp:"optional"`

	// BaseFee was added by EIP-1559 and is ignored in legacy headers.
	BaseFee *big.Int `json:"baseFeePerGas" rlp:"optional"`

	// WithdrawalsHash was added by EIP-4895 and is ignored in legacy headers.
	WithdrawalsHash *common.Hash `json:"withdrawalsRoot" rlp:"optional"`

	// ExcessDataGas was added by EIP-4844 and is ignored in legacy headers.
	ExcessDataGas *big.Int `json:"excessDataGas" rlp:"optional"`
}

// MarshalJSON marshals as JSON.
func (h Header) MarshalJSON() ([]byte, error) {
	type Header struct {
		ParentHash      common.Hash       `json:"parentHash"       gencodec:"required"`
		UncleHash       common.Hash       `json:"sha3Uncles"       gencodec:"required"`
		Coinbase        common.Address    `json:"author"`
		Root            common.Hash       `json:"stateRoot"        gencodec:"required"`
		TxHash          common.Hash       `json:"transactionsRoot" gencodec:"required"`
		ReceiptHash     common.Hash       `json:"receiptsRoot"     gencodec:"required"`
		Bloom           types.Bloom       `json:"logsBloom"        gencodec:"required"`
		Difficulty      *hexutil.Big      `json:"difficulty"       gencodec:"required"`
		Number          *hexutil.Big      `json:"number"           gencodec:"required"`
		GasLimit        hexutil.Uint64    `json:"gasLimit"         gencodec:"required"`
		GasUsed         hexutil.Uint64    `json:"gasUsed"          gencodec:"required"`
		Time            hexutil.Uint64    `json:"timestamp"        gencodec:"required"`
		Extra           hexutil.Bytes     `json:"extraData"        gencodec:"required"`
		MixDigest       *common.Hash      `json:"mixHash" rlp:"optional"`
		Nonce           *types.BlockNonce `json:"nonce" rlp:"optional"`
		AuraStep        *uint64           `json:"step" rlp:"optional"`
		AuraSignature   hexutil.Bytes     `json:"signature" rlp:"optional"`
		BaseFee         *hexutil.Big      `json:"baseFeePerGas" rlp:"optional"`
		WithdrawalsHash *common.Hash      `json:"withdrawalsRoot" rlp:"optional"`
		ExcessDataGas   *hexutil.Big      `json:"excessDataGas" rlp:"optional"`
		Hash            common.Hash       `json:"hash"`
	}
	var enc Header
	enc.ParentHash = h.ParentHash
	enc.UncleHash = h.UncleHash
	enc.Coinbase = h.Coinbase
	enc.Root = h.Root
	enc.TxHash = h.TxHash
	enc.ReceiptHash = h.ReceiptHash
	enc.Bloom = h.Bloom
	enc.Difficulty = (*hexutil.Big)(h.Difficulty)
	enc.Number = (*hexutil.Big)(h.Number)
	enc.GasLimit = hexutil.Uint64(h.GasLimit)
	enc.GasUsed = hexutil.Uint64(h.GasUsed)
	enc.Time = hexutil.Uint64(h.Time)
	enc.Extra = h.Extra
	enc.MixDigest = h.MixDigest
	enc.Nonce = h.Nonce
	auraStep := h.AuraStep.Uint64()
	enc.AuraStep = &auraStep
	enc.AuraSignature = h.AuraSignature
	enc.BaseFee = (*hexutil.Big)(h.BaseFee)
	enc.WithdrawalsHash = h.WithdrawalsHash
	enc.ExcessDataGas = (*hexutil.Big)(h.ExcessDataGas)
	enc.Hash = h.Hash()
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (h *Header) UnmarshalJSON(input []byte) error {
	type Header struct {
		ParentHash      *common.Hash      `json:"parentHash"       gencodec:"required"`
		UncleHash       *common.Hash      `json:"sha3Uncles"       gencodec:"required"`
		Coinbase        *common.Address   `json:"author"`
		Root            *common.Hash      `json:"stateRoot"        gencodec:"required"`
		TxHash          *common.Hash      `json:"transactionsRoot" gencodec:"required"`
		ReceiptHash     *common.Hash      `json:"receiptsRoot"     gencodec:"required"`
		Bloom           *types.Bloom      `json:"logsBloom"        gencodec:"required"`
		Difficulty      *hexutil.Big      `json:"difficulty"       gencodec:"required"`
		Number          *hexutil.Big      `json:"number"           gencodec:"required"`
		GasLimit        *hexutil.Uint64   `json:"gasLimit"         gencodec:"required"`
		GasUsed         *hexutil.Uint64   `json:"gasUsed"          gencodec:"required"`
		Time            *hexutil.Uint64   `json:"timestamp"        gencodec:"required"`
		Extra           *hexutil.Bytes    `json:"extraData"        gencodec:"required"`
		MixDigest       *common.Hash      `json:"mixHash" rlp:"optional"`
		Nonce           *types.BlockNonce `json:"nonce" rlp:"optional"`
		AuraStep        *uint64           `json:"step" rlp:"optional"`
		AuraSignature   *hexutil.Bytes    `json:"signature" rlp:"optional"`
		BaseFee         *hexutil.Big      `json:"baseFeePerGas" rlp:"optional"`
		WithdrawalsHash *common.Hash      `json:"withdrawalsRoot" rlp:"optional"`
		ExcessDataGas   *hexutil.Big      `json:"excessDataGas" rlp:"optional"`
	}
	var dec Header
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.ParentHash == nil {
		return errors.New("missing required field 'parentHash' for Header")
	}
	h.ParentHash = *dec.ParentHash
	if dec.UncleHash == nil {
		return errors.New("missing required field 'sha3Uncles' for Header")
	}
	h.UncleHash = *dec.UncleHash
	if dec.Coinbase != nil {
		h.Coinbase = *dec.Coinbase
	}
	if dec.Root == nil {
		return errors.New("missing required field 'stateRoot' for Header")
	}
	h.Root = *dec.Root
	if dec.TxHash == nil {
		return errors.New("missing required field 'transactionsRoot' for Header")
	}
	h.TxHash = *dec.TxHash
	if dec.ReceiptHash == nil {
		return errors.New("missing required field 'receiptsRoot' for Header")
	}
	h.ReceiptHash = *dec.ReceiptHash
	if dec.Bloom == nil {
		return errors.New("missing required field 'logsBloom' for Header")
	}
	h.Bloom = *dec.Bloom
	if dec.Difficulty == nil {
		return errors.New("missing required field 'difficulty' for Header")
	}
	h.Difficulty = (*big.Int)(dec.Difficulty)
	if dec.Number == nil {
		return errors.New("missing required field 'number' for Header")
	}
	h.Number = (*big.Int)(dec.Number)
	if dec.GasLimit == nil {
		return errors.New("missing required field 'gasLimit' for Header")
	}
	h.GasLimit = uint64(*dec.GasLimit)
	if dec.GasUsed == nil {
		return errors.New("missing required field 'gasUsed' for Header")
	}
	h.GasUsed = uint64(*dec.GasUsed)
	if dec.Time == nil {
		return errors.New("missing required field 'timestamp' for Header")
	}
	h.Time = uint64(*dec.Time)
	if dec.Extra == nil {
		return errors.New("missing required field 'extraData' for Header")
	}
	h.Extra = *dec.Extra
	if dec.AuraSignature != nil {
		if dec.AuraStep == nil {
			return errors.New("missing required field 'step' for Header")
		}
		h.AuraStep = new(big.Int).SetUint64(*dec.AuraStep)
		h.AuraSignature = *dec.AuraSignature
	} else {
		if dec.MixDigest == nil {
			return errors.New("missing required field 'mixHash' for Header")
		}
		if dec.Nonce == nil {
			return errors.New("missing required field 'nonce' for Header")
		}
		h.MixDigest = dec.MixDigest
		h.Nonce = dec.Nonce
	}
	if dec.BaseFee != nil {
		h.BaseFee = (*big.Int)(dec.BaseFee)
	}
	if dec.WithdrawalsHash != nil {
		h.WithdrawalsHash = dec.WithdrawalsHash
	}
	if dec.ExcessDataGas != nil {
		h.ExcessDataGas = (*big.Int)(dec.ExcessDataGas)
	}
	return nil
}

func (h *Header) EncodeRLP(_w io.Writer) error {
	w := rlp.NewEncoderBuffer(_w)
	_tmp0 := w.List()
	w.WriteBytes(h.ParentHash[:])
	w.WriteBytes(h.UncleHash[:])
	w.WriteBytes(h.Coinbase[:])
	w.WriteBytes(h.Root[:])
	w.WriteBytes(h.TxHash[:])
	w.WriteBytes(h.ReceiptHash[:])
	w.WriteBytes(h.Bloom[:])
	if h.Difficulty == nil {
		_, _ = w.Write(rlp.EmptyString)
	} else {
		if h.Difficulty.Sign() == -1 {
			return rlp.ErrNegativeBigInt
		}
		w.WriteBigInt(h.Difficulty)
	}
	if h.Number == nil {
		_, _ = w.Write(rlp.EmptyString)
	} else {
		if h.Number.Sign() == -1 {
			return rlp.ErrNegativeBigInt
		}
		w.WriteBigInt(h.Number)
	}
	w.WriteUint64(h.GasLimit)
	w.WriteUint64(h.GasUsed)
	w.WriteUint64(h.Time)
	w.WriteBytes(h.Extra)
	_tmp1 := h.AuraSignature != nil
	_tmp2 := h.AuraStep != nil
	if _tmp1 || _tmp2 {
		if h.AuraStep == nil {
			_, _ = w.Write(rlp.EmptyString)
		} else {
			if h.AuraStep.Sign() == -1 {
				return rlp.ErrNegativeBigInt
			}
			w.WriteBigInt(h.AuraStep)
		}
		if h.AuraSignature == nil {
			_, _ = w.Write(rlp.EmptyString)
		} else {
			w.WriteBytes(h.AuraSignature[:])
		}
	}
	_tmp3 := h.MixDigest != nil
	_tmp4 := h.Nonce != nil
	if _tmp3 || _tmp4 {
		if h.MixDigest == nil {
			_, _ = w.Write([]byte{0x80})
		} else {
			w.WriteBytes(h.MixDigest[:])
		}
		if h.Nonce == nil {
			_, _ = w.Write(rlp.EmptyString)
		} else {
			w.WriteBytes(h.Nonce[:])
		}
	}
	_tmp5 := h.BaseFee != nil
	_tmp6 := h.WithdrawalsHash != nil
	if _tmp5 || _tmp6 {
		if h.BaseFee == nil {
			_, _ = w.Write(rlp.EmptyString)
		} else {
			if h.BaseFee.Sign() == -1 {
				return rlp.ErrNegativeBigInt
			}
			w.WriteBigInt(h.BaseFee)
		}
	}
	if _tmp6 {
		if h.WithdrawalsHash == nil {
			_, _ = w.Write([]byte{0x80})
		} else {
			w.WriteBytes(h.WithdrawalsHash[:])
		}
	}
	_tmp7 := h.ExcessDataGas != nil
	if _tmp7 {
		if h.ExcessDataGas == nil {
			_, _ = w.Write(rlp.EmptyString)
		} else {
			if h.ExcessDataGas.Sign() == -1 {
				return rlp.ErrNegativeBigInt
			}
			w.WriteBigInt(h.ExcessDataGas)
		}
	}
	w.ListEnd(_tmp0)
	return w.Flush()
}

func (h *Header) Hash() common.Hash {
	return rlpHash(h)
}
