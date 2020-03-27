package types

import (
	"fmt"
	"math/big"

	"github.com/BOPR/common"
	"github.com/BOPR/contracts/rollup"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

// UserAccount is the user data stored on the node per user
type UserAccount struct {
	DBModel
	// ID is the path of the user account in the PDA Tree
	// Cannot be changed once created
	AccountID uint64
	// Token type of the user account
	// Cannot be changed once creation
	TokenType uint64

	// Balance of the user account
	Balance uint64
	// Nonce of the account
	Nonce uint64

	// Path from root to leaf
	// NOTE: not a part of the leaf
	Path uint

	// Pending = 0 means has deposit but not merged to balance tree
	// Active = 1
	// InActive = 2 means user has withdrawn amount and is inactive
	Status int
}

func NewUserAccount(id, balance, tokenType, nonce uint64) UserAccount {
	return UserAccount{
		AccountID: id,
		Balance:   balance,
		TokenType: tokenType,
		Nonce:     nonce,
	}
}

func NewPendingUserAccount(id, balance, tokenType uint64) UserAccount {
	return UserAccount{
		AccountID: id,
		TokenType: tokenType,
		Balance:   balance,
		Nonce:     0,
		Path:      0,
		Status:    0,
	}
}

func (acc *UserAccount) ToABIAccount() rollup.DataTypesUserAccount {
	return rollup.DataTypesUserAccount{
		ID:        UintToBigInt(acc.AccountID),
		Balance:   UintToBigInt(acc.Balance),
		TokenType: UintToBigInt(acc.TokenType),
		Nonce:     UintToBigInt(acc.Nonce),
	}
}

func (acc *UserAccount) AccountInclusionProof(path int64) rollup.DataTypesAccountInclusionProof {
	return rollup.DataTypesAccountInclusionProof{
		PathToAccount: big.NewInt(path),
		Account:       acc.ToABIAccount(),
	}
}

func (acc *UserAccount) ABIEncode() ([]byte, error) {
	uint256Ty, err := abi.NewType("uint256", "uint256", nil)
	if err != nil {
		return []byte(""), err
	}

	arguments := abi.Arguments{
		{
			Type: uint256Ty,
		},
		{
			Type: uint256Ty,
		},
		{
			Type: uint256Ty,
		},
		{
			Type: uint256Ty,
		},
	}
	bytes, err := arguments.Pack(
		big.NewInt(int64(acc.AccountID)),
		big.NewInt(int64(acc.Balance)),
		big.NewInt(int64(acc.TokenType)),
		big.NewInt(int64(acc.Nonce)),
	)
	if err != nil {
		return []byte(""), err
	}

	return bytes, nil
}

func AccsToLeafHashes(accs []UserAccount) (result [][32]byte) {
	for i, acc := range accs {
		accEncoded, err := acc.ABIEncode()
		if err != nil {
			fmt.Println("Error while abi encoding accounts", err)
			return
		}
		result[i] = BytesToByteArray(common.Hash(accEncoded))
	}
	return
}
