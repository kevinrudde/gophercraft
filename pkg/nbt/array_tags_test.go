package nbt

import (
	"bytes"
	"reflect"
	"testing"
)

func TestByteArrayTag(t *testing.T) {
	// Create a ByteArrayTag
	data := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tag := NewByteArrayTag("test_byte_array", data)

	if tag.Name() != "test_byte_array" {
		t.Errorf("Expected name to be 'test_byte_array', got '%s'", tag.Name())
	}

	if !reflect.DeepEqual(tag.Value, data) {
		t.Errorf("ByteArray value doesn't match expected value")
	}

	if tag.Type() != TagByteArray {
		t.Errorf("Expected tag type to be TagByteArray, got %d", tag.Type())
	}

	// Test read/write
	var buf bytes.Buffer
	err := WriteNamedTag(&buf, tag)
	if err != nil {
		t.Fatalf("Failed to write tag: %v", err)
	}

	readTag, err := ReadNamedTag(bytes.NewReader(buf.Bytes()))
	if err != nil {
		t.Fatalf("Failed to read tag: %v", err)
	}

	byteArrayTag, ok := readTag.(*ByteArrayTag)
	if !ok {
		t.Fatalf("Expected ByteArrayTag, got %T", readTag)
	}

	if byteArrayTag.Name() != tag.Name() {
		t.Errorf("Expected name to be '%s', got '%s'", tag.Name(), byteArrayTag.Name())
	}

	if !reflect.DeepEqual(byteArrayTag.Value, tag.Value) {
		t.Errorf("Expected value arrays to be equal")
	}
}

func TestEmptyByteArrayTag(t *testing.T) {
	// Create an empty ByteArrayTag
	tag := NewByteArrayTag("empty_byte_array", []byte{})

	// Test read/write
	var buf bytes.Buffer
	err := WriteNamedTag(&buf, tag)
	if err != nil {
		t.Fatalf("Failed to write tag: %v", err)
	}

	readTag, err := ReadNamedTag(bytes.NewReader(buf.Bytes()))
	if err != nil {
		t.Fatalf("Failed to read tag: %v", err)
	}

	byteArrayTag, ok := readTag.(*ByteArrayTag)
	if !ok {
		t.Fatalf("Expected ByteArrayTag, got %T", readTag)
	}

	if len(byteArrayTag.Value) != 0 {
		t.Errorf("Expected empty byte array, got %d bytes", len(byteArrayTag.Value))
	}
}

func TestIntArrayTag(t *testing.T) {
	// Create an IntArrayTag
	data := []int32{100, 200, 300, 400, 500}
	tag := NewIntArrayTag("test_int_array", data)

	if tag.Name() != "test_int_array" {
		t.Errorf("Expected name to be 'test_int_array', got '%s'", tag.Name())
	}

	if !reflect.DeepEqual(tag.Value, data) {
		t.Errorf("IntArray value doesn't match expected value")
	}

	if tag.Type() != TagIntArray {
		t.Errorf("Expected tag type to be TagIntArray, got %d", tag.Type())
	}

	// Test read/write
	var buf bytes.Buffer
	err := WriteNamedTag(&buf, tag)
	if err != nil {
		t.Fatalf("Failed to write tag: %v", err)
	}

	readTag, err := ReadNamedTag(bytes.NewReader(buf.Bytes()))
	if err != nil {
		t.Fatalf("Failed to read tag: %v", err)
	}

	intArrayTag, ok := readTag.(*IntArrayTag)
	if !ok {
		t.Fatalf("Expected IntArrayTag, got %T", readTag)
	}

	if intArrayTag.Name() != tag.Name() {
		t.Errorf("Expected name to be '%s', got '%s'", tag.Name(), intArrayTag.Name())
	}

	if !reflect.DeepEqual(intArrayTag.Value, tag.Value) {
		t.Errorf("Expected value arrays to be equal")
	}
}

func TestEmptyIntArrayTag(t *testing.T) {
	// Create an empty IntArrayTag
	tag := NewIntArrayTag("empty_int_array", []int32{})

	// Test read/write
	var buf bytes.Buffer
	err := WriteNamedTag(&buf, tag)
	if err != nil {
		t.Fatalf("Failed to write tag: %v", err)
	}

	readTag, err := ReadNamedTag(bytes.NewReader(buf.Bytes()))
	if err != nil {
		t.Fatalf("Failed to read tag: %v", err)
	}

	intArrayTag, ok := readTag.(*IntArrayTag)
	if !ok {
		t.Fatalf("Expected IntArrayTag, got %T", readTag)
	}

	if len(intArrayTag.Value) != 0 {
		t.Errorf("Expected empty int array, got %d integers", len(intArrayTag.Value))
	}
}

func TestLongArrayTag(t *testing.T) {
	// Create a LongArrayTag
	data := []int64{10000000000, 20000000000, 30000000000}
	tag := NewLongArrayTag("test_long_array", data)

	if tag.Name() != "test_long_array" {
		t.Errorf("Expected name to be 'test_long_array', got '%s'", tag.Name())
	}

	if !reflect.DeepEqual(tag.Value, data) {
		t.Errorf("LongArray value doesn't match expected value")
	}

	if tag.Type() != TagLongArray {
		t.Errorf("Expected tag type to be TagLongArray, got %d", tag.Type())
	}

	// Test read/write
	var buf bytes.Buffer
	err := WriteNamedTag(&buf, tag)
	if err != nil {
		t.Fatalf("Failed to write tag: %v", err)
	}

	readTag, err := ReadNamedTag(bytes.NewReader(buf.Bytes()))
	if err != nil {
		t.Fatalf("Failed to read tag: %v", err)
	}

	longArrayTag, ok := readTag.(*LongArrayTag)
	if !ok {
		t.Fatalf("Expected LongArrayTag, got %T", readTag)
	}

	if longArrayTag.Name() != tag.Name() {
		t.Errorf("Expected name to be '%s', got '%s'", tag.Name(), longArrayTag.Name())
	}

	if !reflect.DeepEqual(longArrayTag.Value, tag.Value) {
		t.Errorf("Expected value arrays to be equal")
	}
}

func TestEmptyLongArrayTag(t *testing.T) {
	// Create an empty LongArrayTag
	tag := NewLongArrayTag("empty_long_array", []int64{})

	// Test read/write
	var buf bytes.Buffer
	err := WriteNamedTag(&buf, tag)
	if err != nil {
		t.Fatalf("Failed to write tag: %v", err)
	}

	readTag, err := ReadNamedTag(bytes.NewReader(buf.Bytes()))
	if err != nil {
		t.Fatalf("Failed to read tag: %v", err)
	}

	longArrayTag, ok := readTag.(*LongArrayTag)
	if !ok {
		t.Fatalf("Expected LongArrayTag, got %T", readTag)
	}

	if len(longArrayTag.Value) != 0 {
		t.Errorf("Expected empty long array, got %d longs", len(longArrayTag.Value))
	}
}
