//nolint:dupl // maybe we'll figure out how to make DB resource accessors generic in the future
package badger

import (
	"encoding/json"
	"github.com/khulnasoft/orchard/pkg/resource/v1"
	"github.com/dgraph-io/badger/v3"
	"path"
)

const SpaceVMs = "/vms"

func VMKey(name string) []byte {
	return []byte(path.Join(SpaceVMs, name))
}

func (txn *Transaction) GetVM(name string) (result *v1.VM, err error) {
	defer func() {
		err = mapErr(err)
	}()

	key := VMKey(name)

	item, err := txn.badgerTxn.Get(key)
	if err != nil {
		return nil, err
	}

	valueBytes, err := item.ValueCopy(nil)
	if err != nil {
		return nil, err
	}

	var vm v1.VM

	err = json.Unmarshal(valueBytes, &vm)
	if err != nil {
		return nil, err
	}

	return &vm, nil
}

func (txn *Transaction) SetVM(vm v1.VM) (err error) {
	defer func() {
		err = mapErr(err)
	}()

	key := VMKey(vm.Name)

	valueBytes, err := json.Marshal(vm)
	if err != nil {
		return err
	}

	return txn.badgerTxn.Set(key, valueBytes)
}

func (txn *Transaction) DeleteVM(name string) (err error) {
	defer func() {
		err = mapErr(err)
	}()

	key := VMKey(name)

	return txn.badgerTxn.Delete(key)
}

func (txn *Transaction) ListVMs() (result []v1.VM, err error) {
	defer func() {
		err = mapErr(err)
	}()

	it := txn.badgerTxn.NewIterator(badger.IteratorOptions{
		Prefix: []byte(SpaceVMs),
	})
	defer it.Close()

	for it.Rewind(); it.Valid(); it.Next() {
		item := it.Item()

		vmBytes, err := item.ValueCopy(nil)
		if err != nil {
			return nil, err
		}

		var vm v1.VM

		if err := json.Unmarshal(vmBytes, &vm); err != nil {
			return nil, err
		}

		result = append(result, vm)
	}

	return result, nil
}
