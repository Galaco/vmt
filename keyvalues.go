package vmt

import (
	keyvalues "github.com/galaco/KeyValues"
	"reflect"
)

// FromKeyValues builds a material from keyvalue definitions.
// Please note that `include` keys cannot be automatically resolved
// using this method of parsing.
func FromKeyValues(kvs *keyvalues.KeyValue, definition Material) (Material, error) {
	return unmarshalKeyvalues(kvs, definition)
}

func unmarshalKeyvalues(kvs *keyvalues.KeyValue, definition Material) (mat Material, err error) {
	mat = definition
	defer func() {
		if r := recover(); r != nil {
			if je, ok := r.(error); ok {
				err = je
			} else {
				panic(r)
			}
		}
	}()
	defType := reflect.TypeOf(definition).Elem()
	valueType := reflect.ValueOf(definition).Elem()

	root := *kvs

	fastMap, err := createKeyvalueFastMap(&root)
	if err != nil {
		return nil, err
	}
	fastMap["__SHADER_NAME__"] = keyvalues.NewKeyValuePair("__SHADER_NAME__", root.Key(), keyvalues.ValueString)

	for i := 0; i < valueType.NumField(); i++ {
		tag, ok := defType.Field(i).Tag.Lookup("vmt")
		if !ok {
			continue
		}
		if kv, ok := fastMap[tag]; ok {
			switch kv.Type() {
			case keyvalues.ValueInt:
				val, _ := kv.AsInt()
				valueType.Field(i).SetInt(int64(val))
			case keyvalues.ValueFloat:
				val, _ := kv.AsFloat()
				valueType.Field(i).SetFloat(float64(val))
			case keyvalues.ValueString:
				val, _ := kv.AsString()
				valueType.Field(i).SetString(val)
			default:
			}
		}
	}

	return definition, nil
}

func createKeyvalueFastMap(kvs *keyvalues.KeyValue) (map[string]*keyvalues.KeyValue, error) {
	if !kvs.HasChildren() {
		return nil, ErrorMaterialKeyValuesEmpty
	}

	children, err := kvs.Children()
	if err != nil {
		return nil, err
	}

	fastmap := make(map[string]*keyvalues.KeyValue)

	for _, c := range children {
		fastmap[c.Key()] = c
	}

	return fastmap, nil
}
