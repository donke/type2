package type2

import (
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

type Type2 struct {
	Name     string
	Typeable bool
	File     io.Reader
	file     *os.File
}

func New(name string) *Type2 {
	var typeable bool = true
	fi, err := os.Stat(name)
	if err != nil && os.IsNotExist(err) {
		typeable = false
	}
	if err == nil && fi.IsDir() {
		typeable = false
	}
	if !typeable {
		return &Type2{
			Name:     name,
			Typeable: false,
		}
	}

	file, err := os.Open(name)
	if err != nil {
		typeable = false
	}
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		typeable = false
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		typeable = false
	}
	if !typeable {
		return &Type2{
			Name:     name,
			Typeable: false,
		}
	}

	var reader io.Reader
	e := guess_jp(buf)
	switch e {
	case ShiftJis:
		reader = transform.NewReader(file, japanese.ShiftJIS.NewDecoder())
	case EucJp:
		reader = transform.NewReader(file, japanese.EUCJP.NewDecoder())
	default:
		reader = file
	}
	return &Type2{
		Name:     name,
		Typeable: typeable,
		File:     reader,
	}
}

func (t2 *Type2) Close() {
	t2.file.Close()
}
