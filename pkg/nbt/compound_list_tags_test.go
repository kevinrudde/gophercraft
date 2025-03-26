package nbt

import (
	"bytes"
	"testing"
)

func TestListTag(t *testing.T) {
	// Create a list of string tags
	list := NewListTag("test_list", TagString)

	// Add some string tags to the list
	if err := list.Add(NewStringTag("", "first")); err != nil {
		t.Fatalf("Failed to add tag to list: %v", err)
	}

	if err := list.Add(NewStringTag("", "second")); err != nil {
		t.Fatalf("Failed to add tag to list: %v", err)
	}

	if err := list.Add(NewStringTag("", "third")); err != nil {
		t.Fatalf("Failed to add tag to list: %v", err)
	}

	if list.Name() != "test_list" {
		t.Errorf("Expected name to be 'test_list', got '%s'", list.Name())
	}

	if list.Type() != TagList {
		t.Errorf("Expected tag type to be TagList, got %d", list.Type())
	}

	if len(list.Value) != 3 {
		t.Fatalf("Expected list to have 3 elements, got %d", len(list.Value))
	}

	// Check first element
	strTag, ok := list.Value[0].(*StringTag)
	if !ok {
		t.Fatalf("Expected StringTag at index 0, got %T", list.Value[0])
	}

	if strTag.Value != "first" {
		t.Errorf("Expected value at index 0 to be 'first', got '%s'", strTag.Value)
	}

	// Test adding a wrong type tag
	if err := list.Add(NewIntTag("", 123)); err == nil {
		t.Errorf("Expected error when adding int tag to string list, but got none")
	}

	// Test read/write
	var buf bytes.Buffer
	err := WriteNamedTag(&buf, list)
	if err != nil {
		t.Fatalf("Failed to write tag: %v", err)
	}

	readTag, err := ReadNamedTag(bytes.NewReader(buf.Bytes()))
	if err != nil {
		t.Fatalf("Failed to read tag: %v", err)
	}

	readList, ok := readTag.(*ListTag)
	if !ok {
		t.Fatalf("Expected ListTag, got %T", readTag)
	}

	if readList.Name() != list.Name() {
		t.Errorf("Expected name to be '%s', got '%s'", list.Name(), readList.Name())
	}

	if len(readList.Value) != len(list.Value) {
		t.Errorf("Expected %d elements, got %d", len(list.Value), len(readList.Value))
	}

	// Check that first element is preserved
	readStrTag, ok := readList.Value[0].(*StringTag)
	if !ok {
		t.Fatalf("Expected StringTag at index 0, got %T", readList.Value[0])
	}

	if readStrTag.Value != "first" {
		t.Errorf("Expected value at index 0 to be 'first', got '%s'", readStrTag.Value)
	}
}

func TestEmptyListTag(t *testing.T) {
	// Create an empty list
	list := NewListTag("empty_list", TagByte)

	// Test read/write
	var buf bytes.Buffer
	err := WriteNamedTag(&buf, list)
	if err != nil {
		t.Fatalf("Failed to write tag: %v", err)
	}

	readTag, err := ReadNamedTag(bytes.NewReader(buf.Bytes()))
	if err != nil {
		t.Fatalf("Failed to read tag: %v", err)
	}

	readList, ok := readTag.(*ListTag)
	if !ok {
		t.Fatalf("Expected ListTag, got %T", readTag)
	}

	if len(readList.Value) != 0 {
		t.Errorf("Expected empty list, got %d elements", len(readList.Value))
	}
}

func TestNestedListTag(t *testing.T) {
	// Create a list of lists
	outerList := NewListTag("outer_list", TagList)

	// Create first inner list
	innerList1 := NewListTag("", TagInt)
	innerList1.Add(NewIntTag("", 1))
	innerList1.Add(NewIntTag("", 2))

	// Create second inner list
	innerList2 := NewListTag("", TagInt)
	innerList2.Add(NewIntTag("", 3))
	innerList2.Add(NewIntTag("", 4))

	// Add inner lists to outer list
	outerList.Add(innerList1)
	outerList.Add(innerList2)

	// Test read/write
	var buf bytes.Buffer
	err := WriteNamedTag(&buf, outerList)
	if err != nil {
		t.Fatalf("Failed to write tag: %v", err)
	}

	readTag, err := ReadNamedTag(bytes.NewReader(buf.Bytes()))
	if err != nil {
		t.Fatalf("Failed to read tag: %v", err)
	}

	readOuterList, ok := readTag.(*ListTag)
	if !ok {
		t.Fatalf("Expected ListTag, got %T", readTag)
	}

	// Check first inner list
	readInnerList1, ok := readOuterList.Value[0].(*ListTag)
	if !ok {
		t.Fatalf("Expected ListTag at index 0, got %T", readOuterList.Value[0])
	}

	if len(readInnerList1.Value) != 2 {
		t.Errorf("Expected inner list to have 2 elements, got %d", len(readInnerList1.Value))
	}

	// Check first element of first inner list
	readIntTag, ok := readInnerList1.Value[0].(*IntTag)
	if !ok {
		t.Fatalf("Expected IntTag, got %T", readInnerList1.Value[0])
	}

	if readIntTag.Value != 1 {
		t.Errorf("Expected value to be 1, got %d", readIntTag.Value)
	}
}

func TestCompoundTag(t *testing.T) {
	// Create a compound tag
	compound := NewCompoundTag("test_compound")

	// Add various tags to the compound
	compound.PutTag(NewStringTag("name", "TestEntity"))
	compound.PutTag(NewIntTag("id", 12345))
	compound.PutTag(NewFloatTag("health", 20.0))

	// Check values
	if compound.Name() != "test_compound" {
		t.Errorf("Expected name to be 'test_compound', got '%s'", compound.Name())
	}

	if compound.Type() != TagCompound {
		t.Errorf("Expected tag type to be TagCompound, got %d", compound.Type())
	}

	if len(compound.Tags) != 3 {
		t.Errorf("Expected 3 tags in compound, got %d", len(compound.Tags))
	}

	// Check specific tag
	idTag, ok := compound.GetTag("id").(*IntTag)
	if !ok {
		t.Fatalf("Expected IntTag for 'id', got %T", compound.GetTag("id"))
	}

	if idTag.Value != 12345 {
		t.Errorf("Expected 'id' value to be 12345, got %d", idTag.Value)
	}

	// Test read/write
	var buf bytes.Buffer
	err := WriteNamedTag(&buf, compound)
	if err != nil {
		t.Fatalf("Failed to write tag: %v", err)
	}

	readTag, err := ReadNamedTag(bytes.NewReader(buf.Bytes()))
	if err != nil {
		t.Fatalf("Failed to read tag: %v", err)
	}

	readCompound, ok := readTag.(*CompoundTag)
	if !ok {
		t.Fatalf("Expected CompoundTag, got %T", readTag)
	}

	if readCompound.Name() != compound.Name() {
		t.Errorf("Expected name to be '%s', got '%s'", compound.Name(), readCompound.Name())
	}

	if len(readCompound.Tags) != len(compound.Tags) {
		t.Errorf("Expected %d tags, got %d", len(compound.Tags), len(readCompound.Tags))
	}

	// Check specific tag after reading
	readIdTag, ok := readCompound.GetTag("id").(*IntTag)
	if !ok {
		t.Fatalf("Expected IntTag for 'id' after reading, got %T", readCompound.GetTag("id"))
	}

	if readIdTag.Value != 12345 {
		t.Errorf("Expected 'id' value to be 12345 after reading, got %d", readIdTag.Value)
	}
}

func TestEmptyCompoundTag(t *testing.T) {
	// Create an empty compound
	compound := NewCompoundTag("empty_compound")

	// Test read/write
	var buf bytes.Buffer
	err := WriteNamedTag(&buf, compound)
	if err != nil {
		t.Fatalf("Failed to write tag: %v", err)
	}

	readTag, err := ReadNamedTag(bytes.NewReader(buf.Bytes()))
	if err != nil {
		t.Fatalf("Failed to read tag: %v", err)
	}

	readCompound, ok := readTag.(*CompoundTag)
	if !ok {
		t.Fatalf("Expected CompoundTag, got %T", readTag)
	}

	if len(readCompound.Tags) != 0 {
		t.Errorf("Expected empty compound, got %d tags", len(readCompound.Tags))
	}
}

func TestNestedCompoundTag(t *testing.T) {
	// Create a nested compound structure
	root := NewCompoundTag("root")

	// Add a simple tag
	root.PutTag(NewStringTag("version", "1.0"))

	// Create and add a nested compound
	player := NewCompoundTag("player")
	player.PutTag(NewStringTag("username", "Notch"))
	player.PutTag(NewIntTag("level", 30))

	// Create and add a nested inventory compound
	inventory := NewCompoundTag("inventory")

	// Create a list of items
	items := NewListTag("items", TagCompound)

	// Create some item compounds
	diamond := NewCompoundTag("")
	diamond.PutTag(NewStringTag("id", "minecraft:diamond"))
	diamond.PutTag(NewIntTag("count", 64))

	sword := NewCompoundTag("")
	sword.PutTag(NewStringTag("id", "minecraft:diamond_sword"))
	sword.PutTag(NewIntTag("count", 1))
	sword.PutTag(NewIntTag("damage", 10))

	// Add items to list
	items.Add(diamond)
	items.Add(sword)

	// Add items to inventory
	inventory.PutTag(items)

	// Add inventory to player
	player.PutTag(inventory)

	// Add player to root
	root.PutTag(player)

	// Test read/write
	var buf bytes.Buffer
	err := WriteNamedTag(&buf, root)
	if err != nil {
		t.Fatalf("Failed to write tag: %v", err)
	}

	readTag, err := ReadNamedTag(bytes.NewReader(buf.Bytes()))
	if err != nil {
		t.Fatalf("Failed to read tag: %v", err)
	}

	readRoot, ok := readTag.(*CompoundTag)
	if !ok {
		t.Fatalf("Expected CompoundTag, got %T", readTag)
	}

	// Check version tag
	versionTag, ok := readRoot.GetTag("version").(*StringTag)
	if !ok {
		t.Fatalf("Expected StringTag for 'version', got %T", readRoot.GetTag("version"))
	}

	if versionTag.Value != "1.0" {
		t.Errorf("Expected version to be '1.0', got '%s'", versionTag.Value)
	}

	// Check player tag
	readPlayer, ok := readRoot.GetTag("player").(*CompoundTag)
	if !ok {
		t.Fatalf("Expected CompoundTag for 'player', got %T", readRoot.GetTag("player"))
	}

	// Check username
	usernameTag, ok := readPlayer.GetTag("username").(*StringTag)
	if !ok {
		t.Fatalf("Expected StringTag for 'username', got %T", readPlayer.GetTag("username"))
	}

	if usernameTag.Value != "Notch" {
		t.Errorf("Expected username to be 'Notch', got '%s'", usernameTag.Value)
	}

	// Check inventory
	readInventory, ok := readPlayer.GetTag("inventory").(*CompoundTag)
	if !ok {
		t.Fatalf("Expected CompoundTag for 'inventory', got %T", readPlayer.GetTag("inventory"))
	}

	// Check items list
	readItems, ok := readInventory.GetTag("items").(*ListTag)
	if !ok {
		t.Fatalf("Expected ListTag for 'items', got %T", readInventory.GetTag("items"))
	}

	if len(readItems.Value) != 2 {
		t.Errorf("Expected 2 items, got %d", len(readItems.Value))
	}

	// Check sword item
	readSword, ok := readItems.Value[1].(*CompoundTag)
	if !ok {
		t.Fatalf("Expected CompoundTag for sword, got %T", readItems.Value[1])
	}

	swordIdTag, ok := readSword.GetTag("id").(*StringTag)
	if !ok {
		t.Fatalf("Expected StringTag for sword id, got %T", readSword.GetTag("id"))
	}

	if swordIdTag.Value != "minecraft:diamond_sword" {
		t.Errorf("Expected sword id to be 'minecraft:diamond_sword', got '%s'", swordIdTag.Value)
	}

	swordDamageTag, ok := readSword.GetTag("damage").(*IntTag)
	if !ok {
		t.Fatalf("Expected IntTag for sword damage, got %T", readSword.GetTag("damage"))
	}

	if swordDamageTag.Value != 10 {
		t.Errorf("Expected sword damage to be 10, got %d", swordDamageTag.Value)
	}
}
