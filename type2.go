package type2

import (
	"os"
)

type Type2 struct {
	Name     string
	Typeable bool
	File     *os.File
}

func New(name string) *Type2 {
	var typeable bool = true
	_, err := os.Stat(name)
	if err != nil && os.IsNotExist(err) {
		typeable = false
	}

	file, err := os.Open(name)
	if err != nil {
		typeable = false
	}
	return &Type2{
		Name:     name,
		Typeable: typeable,
		File:     file,
	}
}

func (t2 *Type2) Close() {
	t2.File.Close()
}
