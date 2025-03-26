package network

import (
	"testing"

	"github.com/kevinrudde/gophercraft/pkg/nbt"
)

func TestCreateBuffer(t *testing.T) {
	buffer := GetBufferFromPool()
	defer PutBufferToPool(buffer)
	bufLen := buffer.buf.Len()
	if bufLen != 0 {
		t.Errorf("Buffer length incorrect, got: %d, want: %d", bufLen, 0)
	}
}

func TestCreateBufferWithBuf(t *testing.T) {

}

func TestBuffer_Write_And_Read_Bool(t *testing.T) {
	buffer := GetBufferFromPool()
	defer PutBufferToPool(buffer)
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
	buffer := GetBufferFromPool()
	defer PutBufferToPool(buffer)
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
	buffer := GetBufferFromPool()
	defer PutBufferToPool(buffer)
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
	buffer := GetBufferFromPool()
	defer PutBufferToPool(buffer)
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
	buf := GetBufferFromPool()
	defer PutBufferToPool(buf)

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

func TestBuffer_Write_And_Read_NBT(t *testing.T) {
	buffer := GetBufferFromPool()
	defer PutBufferToPool(buffer)

	// Create a simple NBT compound tag for testing
	compoundTag := nbt.NewCompoundTag("TestCompound")

	// Add some values
	stringTag := &nbt.StringTag{}
	stringTag.SetName("testString")
	stringTag.Value = "Hello World"
	compoundTag.PutTag(stringTag)

	intTag := &nbt.IntTag{}
	intTag.SetName("testInt")
	intTag.Value = 42
	compoundTag.PutTag(intTag)

	// Write the tag to the buffer
	err := buffer.WriteNBT(compoundTag)
	if err != nil {
		t.Errorf("WriteNBT returned the error: %s", err)
	}

	// Read it back
	readTag, err := buffer.ReadNBT()
	if err != nil {
		t.Errorf("ReadNBT returned the error: %s", err)
	}

	// Verify the tag type and name
	if readTag.Type() != nbt.TagCompound {
		t.Errorf("Read tag has incorrect type, got %d, want: %d", readTag.Type(), nbt.TagCompound)
	}

	if readTag.Name() != "TestCompound" {
		t.Errorf("Read tag has incorrect name, got %s, want: %s", readTag.Name(), "TestCompound")
	}

	// Cast to compound tag and verify contents
	readCompoundTag, ok := readTag.(*nbt.CompoundTag)
	if !ok {
		t.Errorf("Failed to cast read tag to CompoundTag")
		return
	}

	// Check string tag
	readStringTag := readCompoundTag.GetTag("testString")
	if readStringTag == nil {
		t.Errorf("Failed to get testString tag from compound")
		return
	}

	stringTagValue, ok := readStringTag.(*nbt.StringTag)
	if !ok {
		t.Errorf("testString tag is not a StringTag")
		return
	}

	if stringTagValue.Value != "Hello World" {
		t.Errorf("testString has incorrect value, got %s, want: %s", stringTagValue.Value, "Hello World")
	}

	// Check int tag
	readIntTag := readCompoundTag.GetTag("testInt")
	if readIntTag == nil {
		t.Errorf("Failed to get testInt tag from compound")
		return
	}

	intTagValue, ok := readIntTag.(*nbt.IntTag)
	if !ok {
		t.Errorf("testInt tag is not an IntTag")
		return
	}

	if intTagValue.Value != 42 {
		t.Errorf("testInt has incorrect value, got %d, want: %d", intTagValue.Value, 42)
	}
}

func TestBuffer_Write_And_Read_CompoundTag(t *testing.T) {
	buffer := GetBufferFromPool()
	defer PutBufferToPool(buffer)

	// Create a simple NBT compound tag for testing
	compoundTag := nbt.NewCompoundTag("TestCompound")

	// Add some nested values to test complex structures
	compoundTag.PutTag(nbt.NewStringTag("testString", "Hello World"))
	compoundTag.PutTag(nbt.NewIntTag("testInt", 42))

	// Create a nested compound
	nestedCompound := nbt.NewCompoundTag("nested")
	nestedCompound.PutTag(nbt.NewLongTag("testLong", 9223372036854775807))
	nestedCompound.PutTag(nbt.NewFloatTag("testFloat", 3.14159))

	// Add the nested compound to the parent
	compoundTag.PutTag(nestedCompound)

	// Write using the compound-specific method
	err := buffer.WriteCompoundTag(compoundTag)
	if err != nil {
		t.Errorf("WriteCompoundTag returned the error: %s", err)
	}

	// Read it back using the compound-specific method
	readCompoundTag, err := buffer.ReadCompoundTag()
	if err != nil {
		t.Errorf("ReadCompoundTag returned the error: %s", err)
	}

	// Verify the tag type and name
	if readCompoundTag.Type() != nbt.TagCompound {
		t.Errorf("Read tag has incorrect type, got %d, want: %d", readCompoundTag.Type(), nbt.TagCompound)
	}

	if readCompoundTag.Name() != "TestCompound" {
		t.Errorf("Read tag has incorrect name, got %s, want: %s", readCompoundTag.Name(), "TestCompound")
	}

	// Check string tag
	readStringTag := readCompoundTag.GetTag("testString")
	if readStringTag == nil {
		t.Errorf("Failed to get testString tag from compound")
		return
	}

	stringTagValue, ok := readStringTag.(*nbt.StringTag)
	if !ok {
		t.Errorf("testString tag is not a StringTag")
		return
	}

	if stringTagValue.Value != "Hello World" {
		t.Errorf("testString has incorrect value, got %s, want: %s", stringTagValue.Value, "Hello World")
	}

	// Check int tag
	readIntTag := readCompoundTag.GetTag("testInt")
	if readIntTag == nil {
		t.Errorf("Failed to get testInt tag from compound")
		return
	}

	intTagValue, ok := readIntTag.(*nbt.IntTag)
	if !ok {
		t.Errorf("testInt tag is not an IntTag")
		return
	}

	if intTagValue.Value != 42 {
		t.Errorf("testInt has incorrect value, got %d, want: %d", intTagValue.Value, 42)
	}

	// Check nested compound
	readNestedTag := readCompoundTag.GetTag("nested")
	if readNestedTag == nil {
		t.Errorf("Failed to get nested tag from compound")
		return
	}

	nestedCompoundTag, ok := readNestedTag.(*nbt.CompoundTag)
	if !ok {
		t.Errorf("nested tag is not a CompoundTag")
		return
	}

	// Verify nested long value
	readLongTag := nestedCompoundTag.GetTag("testLong")
	if readLongTag == nil {
		t.Errorf("Failed to get testLong tag from nested compound")
		return
	}

	longTagValue, ok := readLongTag.(*nbt.LongTag)
	if !ok {
		t.Errorf("testLong tag is not a LongTag")
		return
	}

	if longTagValue.Value != 9223372036854775807 {
		t.Errorf("testLong has incorrect value, got %d, want: %d", longTagValue.Value, 9223372036854775807)
	}

	// Verify nested float value
	readFloatTag := nestedCompoundTag.GetTag("testFloat")
	if readFloatTag == nil {
		t.Errorf("Failed to get testFloat tag from nested compound")
		return
	}

	floatTagValue, ok := readFloatTag.(*nbt.FloatTag)
	if !ok {
		t.Errorf("testFloat tag is not a FloatTag")
		return
	}

	if floatTagValue.Value != 3.14159 {
		t.Errorf("testFloat has incorrect value, got %f, want: %f", floatTagValue.Value, 3.14159)
	}
}

func TestBuffer_Write_And_Read_ListTag(t *testing.T) {
	buffer := GetBufferFromPool()
	defer PutBufferToPool(buffer)

	// Create a list tag for testing
	listTag := nbt.NewListTag("TestList", nbt.TagInt)

	// Add some values
	listTag.Add(nbt.NewIntTag("", 100))
	listTag.Add(nbt.NewIntTag("", 200))
	listTag.Add(nbt.NewIntTag("", 300))

	// Write the tag to the buffer
	err := buffer.WriteListTag(listTag)
	if err != nil {
		t.Errorf("WriteListTag returned the error: %s", err)
	}

	// Read it back
	readTag, err := buffer.ReadListTag()
	if err != nil {
		t.Errorf("ReadListTag returned the error: %s", err)
	}

	// Verify the tag type and name
	if readTag.Type() != nbt.TagList {
		t.Errorf("Read tag has incorrect type, got %d, want: %d", readTag.Type(), nbt.TagList)
	}

	if readTag.Name() != "TestList" {
		t.Errorf("Read tag has incorrect name, got %s, want: %s", readTag.Name(), "TestList")
	}

	// Check list contents
	if len(readTag.Value) != 3 {
		t.Errorf("List tag has incorrect length, got %d, want: %d", len(readTag.Value), 3)
		return
	}

	// Check individual elements
	intVal, ok := readTag.Value[0].(*nbt.IntTag)
	if !ok {
		t.Errorf("List element 0 is not an IntTag")
		return
	}

	if intVal.Value != 100 {
		t.Errorf("List element 0 has incorrect value, got %d, want: %d", intVal.Value, 100)
	}

	intVal, ok = readTag.Value[2].(*nbt.IntTag)
	if !ok {
		t.Errorf("List element 2 is not an IntTag")
		return
	}

	if intVal.Value != 300 {
		t.Errorf("List element 2 has incorrect value, got %d, want: %d", intVal.Value, 300)
	}
}

func TestBuffer_Write_And_Read_EmptyNameCompound(t *testing.T) {
	buffer := GetBufferFromPool()
	defer PutBufferToPool(buffer)

	// Create a compound tag with a name (will be temporarily set to empty)
	compoundTag := nbt.NewCompoundTag("OriginalName")

	// Add some values
	compoundTag.PutTag(nbt.NewStringTag("message", "Hello Minecraft"))
	compoundTag.PutTag(nbt.NewIntTag("version", 762))

	// Write with empty name
	err := buffer.WriteEmptyNameCompound(compoundTag)
	if err != nil {
		t.Errorf("WriteEmptyNameCompound returned the error: %s", err)
	}

	// Read it back
	readTag, err := buffer.ReadEmptyNameCompound()
	if err != nil {
		t.Errorf("ReadEmptyNameCompound returned the error: %s", err)
	}

	// Verify the tag type and that the name is empty
	if readTag.Type() != nbt.TagCompound {
		t.Errorf("Read tag has incorrect type, got %d, want: %d", readTag.Type(), nbt.TagCompound)
	}

	if readTag.Name() != "" {
		t.Errorf("Read tag should have empty name, got: %s", readTag.Name())
	}

	// Check that the original tag's name was preserved
	if compoundTag.Name() != "OriginalName" {
		t.Errorf("Original compound name was not preserved, got %s, want: %s",
			compoundTag.Name(), "OriginalName")
	}

	// Check contents
	stringTag := readTag.GetTag("message")
	if stringTag == nil {
		t.Errorf("Failed to get message tag from compound")
		return
	}

	stringValue, ok := stringTag.(*nbt.StringTag)
	if !ok {
		t.Errorf("message tag is not a StringTag")
		return
	}

	if stringValue.Value != "Hello Minecraft" {
		t.Errorf("message has incorrect value, got %s, want: %s",
			stringValue.Value, "Hello Minecraft")
	}

	intTag := readTag.GetTag("version")
	if intTag == nil {
		t.Errorf("Failed to get version tag from compound")
		return
	}

	intValue, ok := intTag.(*nbt.IntTag)
	if !ok {
		t.Errorf("version tag is not an IntTag")
		return
	}

	if intValue.Value != 762 {
		t.Errorf("version has incorrect value, got %d, want: %d",
			intValue.Value, 762)
	}
}
