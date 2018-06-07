package genesis

import (
	"sync"
	"time"

	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/core/vm"
)

type vmTracer struct {
	sync.Mutex
	data map[common.Address]map[common.Hash]common.Hash
}

const defaultSize = 1024

func newVmTracer() *vmTracer {
	return &vmTracer{
		data: make(map[common.Address]map[common.Hash]common.Hash, defaultSize),
	}
}

func (vmt *vmTracer) CaptureState(env *vm.EVM, pc uint64, op vm.OpCode, gas, cost uint64, memory *vm.Memory, stack *vm.Stack, contract *vm.Contract, depth int, err error) error {
	if err != nil {
		return err
	}
	if op == vm.SSTORE {
		s := stack.Data()

		vmt.Lock()
		addrStorage, ok := vmt.getAddrStorage(contract.Address())
		if !ok {
			addrStorage = make(map[common.Hash]common.Hash, defaultSize)
			vmt.setAddrStorage(contract.Address(), addrStorage)
		}
		addrStorage[common.BigToHash(s[len(s)-1])] = common.BigToHash(s[len(s)-2])
		vmt.Unlock()
	}
	return nil
}

func (vmt *vmTracer) getAddrStorage(contractAddress common.Address) (addrStorage map[common.Hash]common.Hash, ok bool) {
	addrStorage, ok = vmt.data[contractAddress]
	return
}

func (vmt *vmTracer) setAddrStorage(contractAddress common.Address, addrStorage map[common.Hash]common.Hash) {
	vmt.data[contractAddress] = addrStorage
}

func (vmt *vmTracer) CaptureEnd(output []byte, gasUsed uint64, t time.Duration, err error) error {
	return nil
}
