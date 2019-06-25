package btc

import (
	"encoding/hex"
	"testing"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/stretchr/testify/assert"
)

func TestCustomHexMsg_MarshalJSON(t *testing.T) {
	a := new(CustomHexMsg)
	a.RawTx = "02000000016343a847f711d0eca0000526822d4fb1d552dcb0aa1bbebf170157670800bd160100000000ffffffff019ce0f505000000001976a914e2ba07020935672d20b9af2eb28d1b358ab4e2b488ac00000000"
	a.Inputs = &[]RawTxInput{
		{
			Txid:         "16bd000867570117bfbe1baab0dc52d5b14f2d82260500a0ecd011f747a84363",
			Vout:         1,
			ScriptPubKey: "a914611ae902f14f4d1c88a0f06bbb9c6b3c1091fdeb87",
			RedeemScript: "52210281af84e3d70d4440e478ea7281bb06b28f8fe0ced72d5a86137c4161439c85642103d987257192c1b2782dcc9443df372ee2a6b509988ba089e84fbed43af54f33c421037d3c8dc27ce1386d4e526454b872f93bd57c4cb9907afd954134e7be550d1f4d53ae",
			Amount:       0.1,
		},
	}
	a.DecodeTransaction = DecodeRawTransaction
	data, err := a.MarshalJSON()
	assert.NoError(t, err)
	t.Log(hex.EncodeToString(data))
}

func TestCustomHexMsg_UnmarshalJSON(t *testing.T) {
	var a = new(CustomHexMsg)
	var hexStr = "7b225261775478223a223032303030303030303136333433613834376637313164306563613030303035323638323264346662316435353264636230616131626265626631373031353736373038303062643136303130303030303030306666666666666666303139636530663530353030303030303030313937366139313465326261303730323039333536373264323062396166326562323864316233353861623465326234383861633030303030303030222c22496e70757473223a5b7b2274786964223a2231366264303030383637353730313137626662653162616162306463353264356231346632643832323630353030613065636430313166373437613834333633222c22766f7574223a312c227363726970745075624b6579223a2261393134363131616539303266313466346431633838613066303662626239633662336331303931666465623837222c2272656465656d536372697074223a22353232313032383161663834653364373064343434306534373865613732383162623036623238663866653063656437326435613836313337633431363134333963383536343231303364393837323537313932633162323738326463633934343364663337326565326136623530393938386261303839653834666265643433616635346633336334323130333764336338646332376365313338366434653532363435346238373266393362643537633463623939303761666439353431333465376265353530643166346435336165222c22616d6f756e74223a302e317d5d2c22507269764b657973223a6e756c6c2c22466c616773223a6e756c6c7d"
	err := a.UnmarshalJSON(hexStr)
	assert.NoError(t, err)
	t.Logf("%+v", a.RawTx)
	t.Logf("%+v", a.Inputs)
	a.DecodeTransaction = DecodeRawTransaction
	tx, err := a.MarshalToWalletTxJSON(&chaincfg.MainNetParams)
	assert.NoError(t, err)
	t.Logf("%+v", tx)
	t.Log(hex.EncodeToString([]byte(tx)))
}
