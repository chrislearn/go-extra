package shortid

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/titanous/guid"
)

func NextID() string {
	id, _ := guid.NextId()
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, id)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	bytes := buf.Bytes()
	id := ""
	for i < 8 {
		v = bytes[i]<<8 + bytes[i+1]
		id += chars[v]
	}
	return id
}
