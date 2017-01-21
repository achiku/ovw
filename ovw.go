package ovw

import (
	"errors"
	"log"

	"github.com/fatih/structs"
)

// OverWrite over write struct values
func OverWrite(dst, src interface{}) error {
	srcMap := structs.Map(src)
	s := structs.New(dst)
	for _, k := range s.Fields() {
		v, ok := srcMap[k.Name()]
		if !ok {
			return errors.New("no")
		}
		log.Printf("%s: %v", k.Name(), v)
		if err := k.Set(srcMap[k.Name()]); err != nil {
			return err
		}
		log.Printf("%s: %t, %v", k.Name(), k.IsZero(), k.Value())
	}
	return nil
}
