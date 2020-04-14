package btree1

import (
	"log"
	"testing"

	"github.com/google/btree"
)

func main() {

}

func TestName(t *testing.T) {

	b := btree.New(3)

	b.ReplaceOrInsert(btree.Int(1))
	b.ReplaceOrInsert(btree.Int(2))
	b.ReplaceOrInsert(btree.Int(3))

	var itemIterator btree.ItemIterator = func(i btree.Item) bool {

		if i.Less(btree.Int(10)) {
			log.Println(i)
			return true
		}
		return false
	}
	b.Ascend(itemIterator)

	b.Descend(itemIterator)

}

func TestName1(t *testing.T) {

	b := btree.New(3)

	for i := 0; i < 10; i++ {

		go func(i2 int) {

			for j := 10000 * i2; j < 10000*(i2+1); j++ {

				b.ReplaceOrInsert(btree.Int(j))

			}
		}(i)

	}

	log.Println(b.Min())
	log.Println(b.Max())

}

func TestName3(t *testing.T) {

	b := btree.New(3)

	for i := 0; i < 1000000; i++ {

		b.ReplaceOrInsert(btree.Int(i))

	}

	log.Println(b.Min())
	log.Println(b.Max())
}
