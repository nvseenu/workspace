package list

import (
	"fmt"
	"testing"
)

var newListFuncs = map[string]func() List[int]{
	"SINGLE_LINKED_LIST": func() List[int] {
		return NewLinkedList[int]()
	},

	"DOUBLY_LINKED_LIST": func() List[int] {
		return NewDoublyLinkedList[int]()
	},
}

func TestNewList(t *testing.T) {

	for _, f := range newListFuncs {
		ls := f()
		if ls.Len() != 0 {
			t.Errorf("Expected 0 but got %v", ls.Len())
		}
	}
}

var listAddTestData = []struct {
	desc        string
	initialElms []int
	pos         int
	value       int
	expected    []int
}{
	{
		desc:        "Add an item to an empty list",
		initialElms: []int{},
		pos:         0,
		value:       1,
		expected:    []int{1},
	},

	{
		desc:        "Add an item at end of the list",
		initialElms: []int{1, 2, 3, 4},
		pos:         4,
		value:       5,
		expected:    []int{1, 2, 3, 4, 5},
	},

	{
		desc:        "Add an item at middle of the list",
		initialElms: []int{1, 2, 3, 4},
		pos:         2,
		value:       5,
		expected:    []int{1, 2, 5, 3, 4},
	},

	{
		desc:        "Add an item next to head",
		initialElms: []int{1, 2, 3, 4},
		pos:         1,
		value:       5,
		expected:    []int{1, 5, 2, 3, 4},
	},
	{
		desc:        "Add an item previous to the last one",
		initialElms: []int{1, 2, 3, 4},
		pos:         3,
		value:       5,
		expected:    []int{1, 2, 3, 5, 4},
	},
}

func TestListAdd(t *testing.T) {

	for listName, newListFunc := range newListFuncs {
		for _, td := range listAddTestData {
			ls := newListFunc()
			for i, v := range td.initialElms {
				ls.Add(i, v)
			}
			ls.Add(td.pos, td.value)
			if err := assertList(td.expected, ls); err != nil {
				t.Errorf("error: %s for test data: %v against list: %s", err, td, listName)
			}
		}
	}
}

var listDeleteTestData = []struct {
	desc         string
	initialElms  []int
	pos          int
	deletedValue int
	expected     []int
	error        error
}{
	{
		desc:         "Delete the item from an empty list",
		initialElms:  []int{},
		pos:          0,
		deletedValue: 0,
		expected:     []int{},
		error:        ErrInvalidPosition,
	},
	{
		desc:         "Delete the item at head",
		initialElms:  []int{1, 2, 3, 4, 5},
		pos:          0,
		deletedValue: 1,
		expected:     []int{2, 3, 4, 5},
	},

	{
		desc:         "Delete the item at end",
		initialElms:  []int{1, 2, 3, 4, 5},
		pos:          4,
		deletedValue: 5,
		expected:     []int{1, 2, 3, 4},
	},

	{
		desc:         "Delete the item at middle",
		initialElms:  []int{1, 2, 3, 4, 5},
		pos:          2,
		deletedValue: 3,
		expected:     []int{1, 2, 4, 5},
	},

	{
		desc:         "Delete the item next to head",
		initialElms:  []int{1, 2, 3, 4, 5},
		pos:          1,
		deletedValue: 2,
		expected:     []int{1, 3, 4, 5},
	},

	{
		desc:         "Delete the item previous to last one",
		initialElms:  []int{1, 2, 3, 4, 5},
		pos:          3,
		deletedValue: 4,
		expected:     []int{1, 2, 3, 5},
	},
}

func TestListDelete(t *testing.T) {
	for listName, newListFunc := range newListFuncs {
		for _, td := range listDeleteTestData {
			ls := newListFunc()
			for i, v := range td.initialElms {
				ls.Add(i, v)
			}
			elm, err := ls.Delete(td.pos)
			if err != td.error {
				t.Errorf("Error expected >%v<, but got >%v<for test data: %v in the list: %s", td.error, err, td, listName)
			}

			if elm != td.deletedValue {
				t.Errorf("Value expected >%v<, but got >%v<for test data: %v in the list: %s", td.deletedValue, elm, td, listName)
			}
			if err := assertList(td.expected, ls); err != nil {
				t.Errorf("error: %s for test data: %v in the list: %s", err, td, listName)
			}
		}
	}
}

func assertList(expList []int, ls List[int]) error {
	if ls.Len() != len(expList) {
		return fmt.Errorf("Length expected >%v<, but got >%v<", len(expList), ls.Len())
	}

	for expIndex, expVal := range expList {
		if actualVal, err := ls.Value(expIndex); err != nil {
			return fmt.Errorf("Some error happend while getting a value at %v from the list: %s", expIndex, ls)
		} else if expVal != actualVal {
			return fmt.Errorf("Expected Element >%v<, but got >%v<", expVal, actualVal)
		}
	}
	return nil
}
