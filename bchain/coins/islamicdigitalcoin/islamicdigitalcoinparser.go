package islamicdigitalcoin

import (
	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
	"github.com/ISLAMIC-DIGITAL-COIN/blockbook/bchain"
	"github.com/ISLAMIC-DIGITAL-COIN/blockbook/bchain/coins/btc"
)

// magic numbers
const (
	MainnetMagic wire.BitcoinNet = 0x3a7dd3b2
)

// chain parameters
var (
	MainNetParams chaincfg.Params
)

func init() {
	// islamicdigitalcoin mainnet Address encoding magics
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic
	MainNetParams.PubKeyHashAddrID = []byte{103} // starting with 'i'
	MainNetParams.ScriptHashAddrID = []byte{13}
	MainNetParams.PrivateKeyID = []byte{212}

}

// islamicdigitalcoinParser handle
type islamicdigitalcoinParser struct {
	*btc.BitcoinParser
	baseparser                         *bchain.BaseParser
}

// NewislamicdigitalcoinParser returns new islamicdigitalcoinParser instance
func NewislamicdigitalcoinParser(params *chaincfg.Params, c *btc.Configuration) *islamicdigitalcoinParser {
	p := &islamicdigitalcoinParser{
		BitcoinParser: btc.NewBitcoinParser(params, c),
		baseparser:    &bchain.BaseParser{},
	}
	return p
}

// GetChainParams contains network parameters for the main islamicdigitalcoincoin network
func GetChainParams(chain string) *chaincfg.Params {
	if !chaincfg.IsRegistered(&MainNetParams) {
		err := chaincfg.Register(&MainNetParams)
		if err != nil {
			panic(err)
		}
	}
	switch chain {
	default:
		return &MainNetParams
	}
}

// PackTx packs transaction to byte array using protobuf
func (p *islamicdigitalcoinParser) PackTx(tx *bchain.Tx, height uint32, blockTime int64) ([]byte, error) {
	return p.baseparser.PackTx(tx, height, blockTime)
}

// UnpackTx unpacks transaction from protobuf byte array
func (p *islamicdigitalcoinParser) UnpackTx(buf []byte) (*bchain.Tx, uint32, error) {
	return p.baseparser.UnpackTx(buf)
}
