package vmt

import (
	keyvalues "github.com/galaco/KeyValues"
	"io"
	"strings"
)

const (
	materialRootDir = "materials/"
	vmtExtension    = ".vmt"
)

type virtualFilesystem interface {
	GetFile(string) (io.Reader, error)
}

// FromFilesystem loads a material from filesystem
// Its important to note this is the ONLY way to automatically
// resolve `import` properties present in a vmt
func FromFilesystem(filePath string, fs virtualFilesystem, definition Material) (Material, error) {
	// ensure proper file path
	validatedPath := sanitizeFilePath(filePath)

	kvs, err := readVmtFromFilesystem(validatedPath, fs)
	if err != nil {
		return nil, err
	}

	return FromKeyValues(kvs, definition)
}

func sanitizeFilePath(filePath string) string {
	if !strings.HasPrefix(filePath, materialRootDir) {
		filePath = materialRootDir + filePath
	}

	if !strings.HasSuffix(filePath, vmtExtension) {
		filePath += vmtExtension
	}

	return filePath
}

func readVmtFromFilesystem(path string, fs virtualFilesystem) (*keyvalues.KeyValue, error) {
	root, err := readKeyValuesFromFilesystem(path, fs)
	if err != nil {
		return nil, err
	}

	include, err := root.Find("include")
	if include != nil && err == nil {
		includePath, _ := include.AsString()
		root, err = mergeIncludedVmtRecursive(root, includePath, fs)
		if err != nil {
			return nil, err
		}
	}

	return root, nil
}

func mergeIncludedVmtRecursive(base *keyvalues.KeyValue, includePath string, fs virtualFilesystem) (*keyvalues.KeyValue, error) {
	parent, err := readKeyValuesFromFilesystem(includePath, fs)
	if err != nil {
		return base, err
	}
	result, err := base.Patch(parent)
	if err != nil {
		return base, err
	}
	include, err := result.Find("include")
	if err == nil {
		newIncludePath, _ := include.AsString()
		if newIncludePath == includePath {
			err = result.RemoveChild("include")
			return &result, err
		}
		return mergeIncludedVmtRecursive(&result, newIncludePath, fs)
	}
	return &result, nil
}

// ReadKeyValues loads a keyvalues file.
// Its just a simple wrapper that combines the KeyValues library and
// the filesystem module.
func readKeyValuesFromFilesystem(filePath string, fs virtualFilesystem) (*keyvalues.KeyValue, error) {
	stream, err := fs.GetFile(filePath)
	if err != nil {
		return nil, err
	}

	reader := keyvalues.NewReader(stream)
	kvs, err := reader.Read()

	return &kvs, err
}
