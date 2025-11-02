package vmt

import (
	"io"

	keyvalues "github.com/galaco/KeyValues"
)

// FromStream builds a material from a data stream.
// Please note that `include` keys cannot be automatically resolved
// using this method of parsing. In many cases it would be better to
// pass use FromFilesystem.
func FromStream(stream io.Reader, definition Material) (Material, error) {
	kvs, err := keyValuesFromStream(stream)
	if err != nil {
		return nil, err
	}

	return FromKeyValues(kvs, definition)
}

func keyValuesFromStream(stream io.Reader) (*keyvalues.KeyValue, error) {
	reader := keyvalues.NewReader(stream)
	kvs, err := reader.Read()

	return &kvs, err
}
