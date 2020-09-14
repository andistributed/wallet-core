package internal

import (
	"encoding/hex"
	"encoding/json"
	"log"
	"testing"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/dabankio/wallet-core/bip44"
	"github.com/dabankio/wallet-core/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	testMnemonic = "lecture leg select like delay limit spread retire toward west grape bachelor"
	btc          = &BTC{}
	testNet      = true
)

func init() {
	var err error
	seed, err := core.NewSeedFromMnemonic(testMnemonic, "")
	btc, err = New(bip44.PathFormat, seed, ChainRegtest)
	if err != nil {
		log.Fatal(err)
	}
}

func TestBTC_DerivePrivateKey(t *testing.T) {
	pk, err := btc.DerivePrivateKey()
	assert.NoError(t, err)
	t.Log(pk)
}

func TestBTC_DerivePublicKey(t *testing.T) {
	pk, err := btc.DerivePublicKey()
	assert.NoError(t, err)
	t.Log(pk)
}

func TestBTC_DeriveAddress(t *testing.T) {
	pk, err := btc.DeriveAddress()
	assert.NoError(t, err)
	t.Log(pk)
}

func TestBTC_DecodeTx(t *testing.T) {
	msg, err := btc.DecodeTx("7b0a2020225261775478223a2022303230303030303030316137393862313636356331313565303230386663356137336163613035353937373033636532663036336538316639383030326336653264353231663836326130303030303030303030666666666666666630323430623536343030303030303030303031376139313435376166323835333465353062663866386638616130323265666162643461373463303261333032383738303834316530303030303030303030313937366139313430663439373161346166313261353462393935316335366239326462636663653238373032373035383861633030303030303030222c0a202022496e70757473223a205b0a202020207b0a2020202020202274786964223a202232613836316635323264366532633030393831666538363366306532336337303937353561306163373335616663303830323565313135633636623139386137222c0a20202020202022766f7574223a20302c0a202020202020227363726970745075624b6579223a202261393134353761663238353334653530626638663866386161303232656661626434613734633032613330323837222c0a2020202020202272656465656d536372697074223a2022353232313033613436353730313234633864393766643534323531333432333966666566303531613238663238343561313265303566656363363666643639393737366366333231303362396665653963363232383665656664363035326538366262623030363539653566373436623164363366376337333035333965666364336535633833316362323130336365366234343066316630663139623931386139636235623462333665393263636335323039316431633664663266396630353961363434356333613331326235336165220a202020207d0a20205d2c0a202022507269764b657973223a206e756c6c2c0a202022466c616773223a2022414c4c220a7d")
	require.NoError(t, err, "btc.DecodeTx")
	t.Log(msg)

	msg, err = btc.DecodeTx("7b0a2020225261775478223a2022303230303030303030316137393862313636356331313565303230386663356137336163613035353937373033636532663036336538316639383030326336653264353231663836326130303030303030306234303034373330343430323230326162643131663332393865363661303830363133353463613334393431613333343338333261366231643035336535346231333036303839373361636438363032323034346461643034643361653038383266666534336161643332393963653165326533303863326631636566613166333036616566653133356533623439313464303134633639353232313033613436353730313234633864393766643534323531333432333966666566303531613238663238343561313265303566656363363666643639393737366366333231303362396665653963363232383665656664363035326538366262623030363539653566373436623164363366376337333035333965666364336535633833316362323130336365366234343066316630663139623931386139636235623462333665393263636335323039316431633664663266396630353961363434356333613331326235336165666666666666666630323430623536343030303030303030303031376139313435376166323835333465353062663866386638616130323265666162643461373463303261333032383738303834316530303030303030303030313937366139313430663439373161346166313261353462393935316335366239326462636663653238373032373035383861633030303030303030222c0a202022496e70757473223a205b0a202020207b0a2020202020202274786964223a202232613836316635323264366532633030393831666538363366306532336337303937353561306163373335616663303830323565313135633636623139386137222c0a20202020202022766f7574223a20302c0a202020202020227363726970745075624b6579223a202261393134353761663238353334653530626638663866386161303232656661626434613734633032613330323837222c0a2020202020202272656465656d536372697074223a2022353232313033613436353730313234633864393766643534323531333432333966666566303531613238663238343561313265303566656363363666643639393737366366333231303362396665653963363232383665656664363035326538366262623030363539653566373436623164363366376337333035333965666364336535633833316362323130336365366234343066316630663139623931386139636235623462333665393263636335323039316431633664663266396630353961363434356333613331326235336165220a202020207d0a20205d2c0a202022507269764b657973223a205b0a20202020226351524d616b554b57527678737957573532354e414e663668334c4e647770634a6b6e7465694152334d6a755135795670593643220a20205d2c0a202022466c616773223a2022414c4c220a7d")
	require.NoError(t, err, "btc.DecodeTx")
	t.Log(msg)
}

func TestBTC_Sign(t *testing.T) {
	msg := new(btcjson.SignRawTransactionCmd)
	// createhex
	msg.RawTx = "0200000001a798b1665c115e0208fc5a73aca05597703ce2f063e81f98002c6e2d521f862a0000000000ffffffff0240b564000000000017a91457af28534e50bf8f8f8aa022efabd4a74c02a3028780841e00000000001976a9140f4971a4af12a54b9951c56b92dbcfce2870270588ac00000000"
	msg.Inputs = &[]btcjson.RawTxInput{
		{
			Txid:         "2a861f522d6e2c00981fe863f0e23c709755a0ac735afc08025e115c66b198a7",
			Vout:         0,
			ScriptPubKey: "a91457af28534e50bf8f8f8aa022efabd4a74c02a30287",
			RedeemScript: "522103a46570124c8d97fd5425134239ffef051a28f2845a12e05fecc66fd699776cf32103b9fee9c62286eefd6052e86bbb00659e5f746b1d63f7c730539efcd3e5c831cb2103ce6b440f1f0f19b918a9cb5b4b36e92ccc52091d1c6df2f9f059a6445c3a312b53ae",
		},
	}
	data, err := json.Marshal(msg)
	assert.NoError(t, err)
	sign, err := btc.Sign(hex.EncodeToString(data), "cV4nrs2iHooPTayGs3zcUYKW4wyG4gLQrFhUYXZDNswN3CDeRaKN")
	assert.NoError(t, err)
	assert.Equal(t, "0200000001a798b1665c115e0208fc5a73aca05597703ce2f063e81f98002c6e2d521f862a00000000b40047304402202abd11f3298e66a08061354ca34941a3343832a6b1d053e54b130608973acd86022044dad04d3ae0882ffe43aad3299ce1e2e308c2f1cefa1f306aefe135e3b4914d014c69522103a46570124c8d97fd5425134239ffef051a28f2845a12e05fecc66fd699776cf32103b9fee9c62286eefd6052e86bbb00659e5f746b1d63f7c730539efcd3e5c831cb2103ce6b440f1f0f19b918a9cb5b4b36e92ccc52091d1c6df2f9f059a6445c3a312b53aeffffffff0240b564000000000017a91457af28534e50bf8f8f8aa022efabd4a74c02a3028780841e00000000001976a9140f4971a4af12a54b9951c56b92dbcfce2870270588ac00000000", sign)

	// sign once hex
	msg.RawTx = sign
	data, err = json.Marshal(msg)
	assert.NoError(t, err)
	sign, err = btc.Sign(hex.EncodeToString(data), "cQUzSaxub2pHRg418zCwqJWDhGfShn21kqbXq6UoPRFRKpZsdfHn")
	assert.NoError(t, err)
	assert.Equal(t, "0200000001a798b1665c115e0208fc5a73aca05597703ce2f063e81f98002c6e2d521f862a00000000fc00473044022003037bd2f3a58d3e2fc6a1f6fd390d2a792a06bff607f468cf7d04dfc1fdc78c022015fd15af9710dc48a2ec83fa33b100ea15b4ea56d8adcc94353bc6dd6da687e00147304402202abd11f3298e66a08061354ca34941a3343832a6b1d053e54b130608973acd86022044dad04d3ae0882ffe43aad3299ce1e2e308c2f1cefa1f306aefe135e3b4914d014c69522103a46570124c8d97fd5425134239ffef051a28f2845a12e05fecc66fd699776cf32103b9fee9c62286eefd6052e86bbb00659e5f746b1d63f7c730539efcd3e5c831cb2103ce6b440f1f0f19b918a9cb5b4b36e92ccc52091d1c6df2f9f059a6445c3a312b53aeffffffff0240b564000000000017a91457af28534e50bf8f8f8aa022efabd4a74c02a3028780841e00000000001976a9140f4971a4af12a54b9951c56b92dbcfce2870270588ac00000000", sign)
}
