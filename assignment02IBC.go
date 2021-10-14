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

}
func CalculateHash(inputBlock *Block) string {

}
func VerifyTransaction(transaction *BlockData, chainHead *Block) bool {

}
func InsertBlock(blockData []BlockData, chainHead *Block) *Block {

}
func ListBlocks(chainHead *Block) {

}
func VerifyChain(chainHead *Block) {

}
func PremineChain(chainHead *Block, numBlocks int) *Block {

}
