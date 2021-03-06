package vm_test

import (
	ogTypes "github.com/annchain/OG/og_interface"
	"github.com/annchain/commongo/math"
	"github.com/annchain/vm/eth/core/vm"
	"github.com/annchain/vm/ovm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAsserts(t *testing.T) {
	from := ogTypes.HexToAddress20("0xABCDEF88")
	from2 := ogTypes.HexToAddress20("0xABCDEF87")
	coinBase := ogTypes.HexToAddress20("0x1234567812345678AABBCCDDEEFF998877665544")

	tracer := vm.NewStructLogger(&vm.LogConfig{
		Debug: true,
	})
	ldb := DefaultLDB(from, coinBase)
	ldb.CreateAccount(from2)
	ldb.AddBalance(from2, math.NewBigInt(10000000))

	rt := &Runtime{
		Tracer:    tracer,
		VmContext: ovm.NewOVMContext(&ovm.DefaultChainContext{}, &coinBase, ldb),
		TxContext: &ovm.TxContext{
			From:       ogTypes.HexToAddress20("0xABCDEF88"),
			Value:      math.NewBigInt(0),
			GasPrice:   math.NewBigInt(1),
			GasLimit:   DefaultGasLimit,
			Coinbase:   coinBase,
			SequenceID: 0,
		},
	}

	_, contractAddr, _, err := DeployContract("asserts.bin", from, coinBase, rt, nil)
	assert.NoError(t, err)

	{
		// op jumps to 0xfe and then raise a non-existing op
		ret, leftGas, err := CallContract(contractAddr, from, coinBase, rt, math.NewBigInt(3), "2911e7b2", nil)
		dump(t, ldb, ret, leftGas, err)
	}
	//vm.WriteTrace(os.Stdout, tracer.Logs)
	{
		ret, leftGas, err := CallContract(contractAddr, from2, coinBase, rt, math.NewBigInt(3), "0d43aaf2", nil)
		dump(t, ldb, ret, leftGas, err)
	}
	//vm.WriteTrace(os.Stdout, tracer.Logs)
}
