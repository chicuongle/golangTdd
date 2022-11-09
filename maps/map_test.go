package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	t.Run("basic function of maps", func(t *testing.T) {
		m := map[string]int{"Foo": 20, "Bar": 30}
		assert.Equal(t, m["Foo"], 20)
		assert.Equal(t, len(m), 2)
	})
}

func TestTypeDictionary(t *testing.T) {
	t.Run("map as customized type", func(t *testing.T) {
		d := Dictionary{"k1": "value1", "k2": "second value"}
		assert.Equal(t, d.Search("k1"), "value1")
	})
	t.Run("editing an existing dictionary", func(t *testing.T) {
		d := Dictionary{"k1": "value1", "k2": "second value"}
		d.Add("k3","new content")
		assert.Equal(t, d.Search("k3"), "new content")
	})
}
