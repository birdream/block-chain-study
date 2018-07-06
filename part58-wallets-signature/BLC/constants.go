package BLC

const walletFile = "Wallets.dat"
const dbFile = "blockchain.db"
const blocksBucket = "blocks"
const genesisCoinbaseData = "The Time 03/jan/2009 chancellor on brink of second bailout for banks"
const (
	version            = byte(0x00) // wallet version
	addressChecksumLen = 4          // wallet address len
	subsidy            = 10         // begin with how many
	targetBits         = 12         // 挖矿难度
)
