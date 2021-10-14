package assignment02IBC

const miningReward = 100
const rootUser = "Satoshi"

type BlockData struct {
	Title    string
	Sender   string
	Receiver string
	Amount   int
}
type Block struct {
	Data        []BlockData
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
}

func CalculateBalance(userName string, chainHead *Block) int {
	return 0
}
func CalculateHash(inputBlock *Block) string {
	return "123"
}
func VerifyTransaction(transaction *BlockData, chainHead *Block) bool {
	return false
}
func InsertBlock(blockData []BlockData, chainHead *Block) *Block {
	return chainHead
}
func ListBlocks(chainHead *Block) {

}
func VerifyChain(chainHead *Block) {

}
func PremineChain(chainHead *Block, numBlocks int) *Block {
	return chainHead
}
