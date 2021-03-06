package infinispan

import (
	"bytes"
	"testing"
)

func TestCreateGet(t *testing.T) {

	expectedGet := []byte{
		0xA0, //MAGIC BYTE
		0x00, //Message Id
		0x19, //Protocol version
		0x03, //Operation
		0x00, //Cache name length ("" = default)
		0x00, //Flags
		0x01, //Client intelligence
		0x00, //Client Topology ID
		0x01, //Key length
		0x31, //Key: 1
	}

	get := createGet([]byte("1"), 0, "")

	if !bytes.Equal(expectedGet, get) {
		t.Errorf("Expected %v, was %v", expectedGet, get)
	}

}
