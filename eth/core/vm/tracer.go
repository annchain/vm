package vm

import (
	ogTypes "github.com/annchain/OG/og_interface"
	"github.com/annchain/vm/instruction"
	vmtypes "github.com/annchain/vm/types"
	"math/big"
	"time"

	"io"
)

// Tracer is used to collect execution traces from an OVM transaction
// execution. CaptureState is called for each step of the VM with the
// current VM state.
// Note that reference types are actual VM data structures; make copies
// if you need to retain them beyond the current call.
type Tracer interface {
	CaptureStart(from ogTypes.Address, to ogTypes.Address, call bool, input []byte, gas uint64, value *big.Int) error
	CaptureState(ctx *vmtypes.Context, pc uint64, op instruction.OpCode, gas, cost uint64, memory *Memory, stack *Stack, contract *vmtypes.Contract, depth int, err error) error
	CaptureFault(ctx *vmtypes.Context, pc uint64, op instruction.OpCode, gas, cost uint64, memory *Memory, stack *Stack, contract *vmtypes.Contract, depth int, err error) error
	CaptureEnd(output []byte, gasUsed uint64, t time.Duration, err error) error
	Write(writer io.Writer)
}
