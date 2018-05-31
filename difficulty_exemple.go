package main

import (
	"fmt"
	"math/big"

	"github.com/picfight/pfcd/blockchain"
)

func main() {

	printCompact("compact", 0x1b01ffff)
	printCompact("compact", 0x0404cb)
	printCompact("compact", 0x7fffff)
	printCompact("compact", 0x008000)
	printCompact("compact", 0x007fff)
	printCompact("compact", 32767)

	// bigOne is 1 represented as a big.Int.  It is defined here to avoid
	// the overhead of creating it multiple times.
	var bigOne = big.NewInt(1)

	// mainPowLimit is the highest proof of work value a PicFight block can
	// have for the main network.  It is the value 2^224 - 1.
	var mainPowLimit = new(big.Int).Sub(new(big.Int).Lsh(bigOne, 224), bigOne)

	printInt("   mainPowLimit", mainPowLimit)

	// testNetPowLimit is the highest proof of work value a PicFight block
	// can have for the test network.  It is the value 2^232 - 1.
	var testNetPowLimit = new(big.Int).Sub(new(big.Int).Lsh(bigOne, 232), bigOne)

	printInt("testNetPowLimit", testNetPowLimit)

	// simNetPowLimit is the highest proof of work value a PicFight block
	// can have for the simulation test network.  It is the value 2^255 - 1.
	var simNetPowLimit = new(big.Int).Sub(new(big.Int).Lsh(bigOne, 255), bigOne)

	printInt(" simNetPowLimit", simNetPowLimit)

}

func printCompact(tag string, compact uint32) {
	fmt.Println(tag+"  > ", fmt.Sprintf("%v", compact))
	fmt.Println("     hex > ", fmt.Sprintf("%04x", compact))
	targetDifficulty := blockchain.CompactToBig(compact)
	fmt.Println("         > ", fmt.Sprintf("%064x", targetDifficulty))
}

func printInt(tag string, big *big.Int) {
	fmt.Println(tag+" > ", fmt.Sprintf("%064x", big))
	compact := blockchain.BigToCompact(big)
	fmt.Println("     hex > ", fmt.Sprintf("%04x", compact))
}
