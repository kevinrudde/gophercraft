package network

import (
	"fmt"
	"testing"
)

func TestCreateBuffer(t *testing.T) {
	buffer := CreateBuffer()
	bufLen := len(buffer.buf)
	if bufLen != 0 {
		t.Errorf("Buffer length incorrect, got: %d, want: %d", bufLen, 0)
	}

	if buffer.ReadIndex != 0 {
		t.Errorf("Buffer readIndex incorrect, got: %d, want: %d", buffer.ReadIndex, 0)
	}

	if buffer.WriteIndex != 0 {
		t.Errorf("Buffer writeIndex incorrect, got: %d, want: %d", buffer.WriteIndex, 0)
	}
}

func TestCreateBufferWithBuf(t *testing.T) {

}

func TestBuffer_Write_And_Read_Bool(t *testing.T) {
	buffer := CreateBuffer()
	buffer.WriteBool(true)

	value, err := buffer.ReadBool()
	if err != nil {
		t.Errorf("ReadBool returned the error: %s", err)
	}

	if !value {
		t.Errorf("Buffer ReadBool incorrect, got %t, want: %t", value, true)
	}
}

func TestBuffer_Write_And_Read_Int16(t *testing.T) {
	buffer := CreateBuffer()
	buffer.WriteInt16(64)
	fmt.Println(buffer.buf)

	value, err := buffer.ReadInt16()
	if err != nil {
		t.Errorf("ReadBool returned the error: %s", err)
	}

	if value != 64 {
		t.Errorf("Buffer ReadInt16 incorrect, got %d, want: %d", value, 64)
	}
}

func TestBuffer_Write_And_Read_UInt16(t *testing.T) {
	buffer := CreateBuffer()
	buffer.WriteUInt16(64)
	fmt.Println(buffer.buf)

	value, err := buffer.ReadUInt16()
	if err != nil {
		t.Errorf("ReadBool returned the error: %s", err)
	}

	if value != 64 {
		t.Errorf("Buffer ReadInt16 incorrect, got %d, want: %d", value, 64)
	}
}

func TestBuffer_Write_And_Read_Int32(t *testing.T) {
	buffer := CreateBuffer()
	buffer.WriteInt32(64)
	fmt.Println(buffer.buf)

	value, err := buffer.ReadInt32()
	if err != nil {
		t.Errorf("ReadBool returned the error: %s", err)
	}

	if value != 64 {
		t.Errorf("Buffer ReadInt16 incorrect, got %d, want: %d", value, 64)
	}
}
