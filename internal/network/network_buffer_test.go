package network

import (
	"testing"
)

func TestCreateBuffer(t *testing.T) {
	buffer := CreateBuffer()
	bufLen := buffer.buf.Len()
	if bufLen != 0 {
		t.Errorf("Buffer length incorrect, got: %d, want: %d", bufLen, 0)
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

	value, err := buffer.ReadInt32()
	if err != nil {
		t.Errorf("ReadBool returned the error: %s", err)
	}

	if value != 64 {
		t.Errorf("Buffer ReadInt16 incorrect, got %d, want: %d", value, 64)
	}
}

func TestBufferWithData(t *testing.T) {
	buf := CreateBuffer()

	// write test data to the buffer
	buf.WriteBool(true)
	buf.WriteByte(10)
	buf.WriteUInt16(1000)
	buf.WriteInt32(50000)
	buf.WriteString("hello world")
	buf.WriteVarInt(12345)

	// read data from the buffer and verify it matches the test data
	if value, err := buf.ReadBool(); err != nil || value != true {
		t.Errorf("Error reading bool: %v", err)
	}
	if value, err := buf.ReadByte(); err != nil || value != 10 {
		t.Errorf("Error reading byte: %v", err)
	}
	if value, err := buf.ReadUInt16(); err != nil || value != 1000 {
		t.Errorf("Error reading uint16: %v", err)
	}
	if value, err := buf.ReadInt32(); err != nil || value != 50000 {
		t.Errorf("Error reading int32: %v", err)
	}
	if value, err := buf.ReadString(); err != nil || value != "hello world" {
		t.Errorf("Error reading string: %v", err)
	}
	if value, err := buf.ReadVarInt(); err != nil || value != 12345 {
		t.Errorf("Error reading varint: %v", err)
	}
}
