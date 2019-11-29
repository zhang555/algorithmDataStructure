package trie

import (
	"testing"
)

func TestTrie(t *testing.T) {

	trie := Constructor()
	trie.Insert(`apple`)

	f := trie.Search(`apple`)
	if !f {
		t.Error()
	}

	f = trie.Search(`app`)
	if f {
		t.Error()
	}

	f = trie.StartsWith(`app`)
	if !f {
		t.Error()
	}

	trie.Insert(`app`)

	f = trie.Search(`app`)
	if !f {
		t.Error()
	}

}
