package nbt

import (
	"bytes"
	"testing"
)

func TestByteTag(t *testing.T) {
	// Create a ByteTag
	tag := NewByteTag("test_byte", 42)

	if tag.Name() != "test_byte" {
		t.Errorf("Expected name to be 'test_byte', got '%s'", tag.Name())
	}

	if tag.Value != 42 {
		t.Errorf("Expected value to be 42, got %d", tag.Value)
	}

	if tag.Type() != TagByte {
		t.Errorf("Expected tag type to be TagByte, got %d", tag.Type())
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

	byteTag, ok := readTag.(*ByteTag)
	if !ok {
		t.Fatalf("Expected ByteTag, got %T", readTag)
	}

	if byteTag.Name() != tag.Name() {
		t.Errorf("Expected name to be '%s', got '%s'", tag.Name(), byteTag.Name())
	}

	if byteTag.Value != tag.Value {
		t.Errorf("Expected value to be %d, got %d", tag.Value, byteTag.Value)
	}
}

func TestShortTag(t *testing.T) {
	// Create a ShortTag
	tag := NewShortTag("test_short", 12345)

	if tag.Name() != "test_short" {
		t.Errorf("Expected name to be 'test_short', got '%s'", tag.Name())
	}

	if tag.Value != 12345 {
		t.Errorf("Expected value to be 12345, got %d", tag.Value)
	}

	if tag.Type() != TagShort {
		t.Errorf("Expected tag type to be TagShort, got %d", tag.Type())
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

	shortTag, ok := readTag.(*ShortTag)
	if !ok {
		t.Fatalf("Expected ShortTag, got %T", readTag)
	}

	if shortTag.Name() != tag.Name() {
		t.Errorf("Expected name to be '%s', got '%s'", tag.Name(), shortTag.Name())
	}

	if shortTag.Value != tag.Value {
		t.Errorf("Expected value to be %d, got %d", tag.Value, shortTag.Value)
	}
}

func TestIntTag(t *testing.T) {
	// Create an IntTag
	tag := NewIntTag("test_int", 123456789)

	if tag.Name() != "test_int" {
		t.Errorf("Expected name to be 'test_int', got '%s'", tag.Name())
	}

	if tag.Value != 123456789 {
		t.Errorf("Expected value to be 123456789, got %d", tag.Value)
	}

	if tag.Type() != TagInt {
		t.Errorf("Expected tag type to be TagInt, got %d", tag.Type())
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

	intTag, ok := readTag.(*IntTag)
	if !ok {
		t.Fatalf("Expected IntTag, got %T", readTag)
	}

	if intTag.Name() != tag.Name() {
		t.Errorf("Expected name to be '%s', got '%s'", tag.Name(), intTag.Name())
	}

	if intTag.Value != tag.Value {
		t.Errorf("Expected value to be %d, got %d", tag.Value, intTag.Value)
	}
}

func TestLongTag(t *testing.T) {
	// Create a LongTag
	tag := NewLongTag("test_long", 1234567890123456789)

	if tag.Name() != "test_long" {
		t.Errorf("Expected name to be 'test_long', got '%s'", tag.Name())
	}

	if tag.Value != 1234567890123456789 {
		t.Errorf("Expected value to be 1234567890123456789, got %d", tag.Value)
	}

	if tag.Type() != TagLong {
		t.Errorf("Expected tag type to be TagLong, got %d", tag.Type())
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

	longTag, ok := readTag.(*LongTag)
	if !ok {
		t.Fatalf("Expected LongTag, got %T", readTag)
	}

	if longTag.Name() != tag.Name() {
		t.Errorf("Expected name to be '%s', got '%s'", tag.Name(), longTag.Name())
	}

	if longTag.Value != tag.Value {
		t.Errorf("Expected value to be %d, got %d", tag.Value, longTag.Value)
	}
}

func TestFloatTag(t *testing.T) {
	// Create a FloatTag
	tag := NewFloatTag("test_float", 3.14159)

	if tag.Name() != "test_float" {
		t.Errorf("Expected name to be 'test_float', got '%s'", tag.Name())
	}

	if tag.Value != 3.14159 {
		t.Errorf("Expected value to be 3.14159, got %f", tag.Value)
	}

	if tag.Type() != TagFloat {
		t.Errorf("Expected tag type to be TagFloat, got %d", tag.Type())
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

	floatTag, ok := readTag.(*FloatTag)
	if !ok {
		t.Fatalf("Expected FloatTag, got %T", readTag)
	}

	if floatTag.Name() != tag.Name() {
		t.Errorf("Expected name to be '%s', got '%s'", tag.Name(), floatTag.Name())
	}

	if floatTag.Value != tag.Value {
		t.Errorf("Expected value to be %f, got %f", tag.Value, floatTag.Value)
	}
}

func TestDoubleTag(t *testing.T) {
	// Create a DoubleTag
	tag := NewDoubleTag("test_double", 2.7182818284590452)

	if tag.Name() != "test_double" {
		t.Errorf("Expected name to be 'test_double', got '%s'", tag.Name())
	}

	if tag.Value != 2.7182818284590452 {
		t.Errorf("Expected value to be 2.7182818284590452, got %f", tag.Value)
	}

	if tag.Type() != TagDouble {
		t.Errorf("Expected tag type to be TagDouble, got %d", tag.Type())
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

	doubleTag, ok := readTag.(*DoubleTag)
	if !ok {
		t.Fatalf("Expected DoubleTag, got %T", readTag)
	}

	if doubleTag.Name() != tag.Name() {
		t.Errorf("Expected name to be '%s', got '%s'", tag.Name(), doubleTag.Name())
	}

	if doubleTag.Value != tag.Value {
		t.Errorf("Expected value to be %f, got %f", tag.Value, doubleTag.Value)
	}
}

func TestStringTag(t *testing.T) {
	// Create a StringTag
	tag := NewStringTag("test_string", "Hello, Minecraft!")

	if tag.Name() != "test_string" {
		t.Errorf("Expected name to be 'test_string', got '%s'", tag.Name())
	}

	if tag.Value != "Hello, Minecraft!" {
		t.Errorf("Expected value to be 'Hello, Minecraft!', got '%s'", tag.Value)
	}

	if tag.Type() != TagString {
		t.Errorf("Expected tag type to be TagString, got %d", tag.Type())
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

	stringTag, ok := readTag.(*StringTag)
	if !ok {
		t.Fatalf("Expected StringTag, got %T", readTag)
	}

	if stringTag.Name() != tag.Name() {
		t.Errorf("Expected name to be '%s', got '%s'", tag.Name(), stringTag.Name())
	}

	if stringTag.Value != tag.Value {
		t.Errorf("Expected value to be '%s', got '%s'", tag.Value, stringTag.Value)
	}
}

func TestEmptyStringTag(t *testing.T) {
	// Create a StringTag with empty value
	tag := NewStringTag("empty_string", "")

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

	stringTag, ok := readTag.(*StringTag)
	if !ok {
		t.Fatalf("Expected StringTag, got %T", readTag)
	}

	if stringTag.Value != "" {
		t.Errorf("Expected value to be empty string, got '%s'", stringTag.Value)
	}
}
