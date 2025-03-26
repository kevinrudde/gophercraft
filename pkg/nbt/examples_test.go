package nbt

import (
	"bytes"
	"testing"
)

func TestBiomeRegistryExample(t *testing.T) {
	// Create a biome entry using our example function
	biome := CreateBiomeEntry("minecraft:plains")

	// Make sure the biome has the expected structure
	if biome.Name() != "biome" {
		t.Errorf("Expected name to be 'biome', got '%s'", biome.Name())
	}

	// Check basic properties
	nameTag, ok := biome.GetTag("name").(*StringTag)
	if !ok {
		t.Fatalf("Expected StringTag for 'name', got %T", biome.GetTag("name"))
	}

	if nameTag.Value != "minecraft:plains" {
		t.Errorf("Expected name value to be 'minecraft:plains', got '%s'", nameTag.Value)
	}

	// Check if has_precipitation is correctly set
	hasPrecipTag, ok := biome.GetTag("has_precipitation").(*ByteTag)
	if !ok {
		t.Fatalf("Expected ByteTag for 'has_precipitation', got %T", biome.GetTag("has_precipitation"))
	}

	if hasPrecipTag.Value != 1 {
		t.Errorf("Expected has_precipitation to be 1, got %d", hasPrecipTag.Value)
	}

	// Check effects compound
	effects, ok := biome.GetTag("effects").(*CompoundTag)
	if !ok {
		t.Fatalf("Expected CompoundTag for 'effects', got %T", biome.GetTag("effects"))
	}

	// Check colors in effects
	skyColorTag, ok := effects.GetTag("sky_color").(*IntTag)
	if !ok {
		t.Fatalf("Expected IntTag for 'sky_color', got %T", effects.GetTag("sky_color"))
	}

	if skyColorTag.Value != 7907327 {
		t.Errorf("Expected sky_color to be 7907327, got %d", skyColorTag.Value)
	}

	// Check nested mood sound
	moodSound, ok := effects.GetTag("mood_sound").(*CompoundTag)
	if !ok {
		t.Fatalf("Expected CompoundTag for 'mood_sound', got %T", effects.GetTag("mood_sound"))
	}

	soundTag, ok := moodSound.GetTag("sound").(*StringTag)
	if !ok {
		t.Fatalf("Expected StringTag for 'sound', got %T", moodSound.GetTag("sound"))
	}

	if soundTag.Value != "minecraft:ambient.cave" {
		t.Errorf("Expected sound to be 'minecraft:ambient.cave', got '%s'", soundTag.Value)
	}

	// Test that the structure can be correctly serialized and deserialized
	var buf bytes.Buffer
	err := WriteNamedTag(&buf, biome)
	if err != nil {
		t.Fatalf("Failed to write biome tag: %v", err)
	}

	readTag, err := ReadNamedTag(bytes.NewReader(buf.Bytes()))
	if err != nil {
		t.Fatalf("Failed to read biome tag: %v", err)
	}

	readBiome, ok := readTag.(*CompoundTag)
	if !ok {
		t.Fatalf("Expected CompoundTag, got %T", readTag)
	}

	// Verify at least one deeply nested structure was preserved
	readEffects, ok := readBiome.GetTag("effects").(*CompoundTag)
	if !ok {
		t.Fatalf("Expected CompoundTag for 'effects' after read, got %T", readBiome.GetTag("effects"))
	}

	readMoodSound, ok := readEffects.GetTag("mood_sound").(*CompoundTag)
	if !ok {
		t.Fatalf("Expected CompoundTag for 'mood_sound' after read, got %T", readEffects.GetTag("mood_sound"))
	}

	readSoundTag, ok := readMoodSound.GetTag("sound").(*StringTag)
	if !ok {
		t.Fatalf("Expected StringTag for 'sound' after read, got %T", readMoodSound.GetTag("sound"))
	}

	if readSoundTag.Value != "minecraft:ambient.cave" {
		t.Errorf("Expected sound to be 'minecraft:ambient.cave' after read, got '%s'", readSoundTag.Value)
	}

	// Check that complex nested structures like spawners were preserved
	readSpawners, ok := readBiome.GetTag("spawners").(*CompoundTag)
	if !ok {
		t.Fatalf("Expected CompoundTag for 'spawners', got %T", readBiome.GetTag("spawners"))
	}

	readMonsterList, ok := readSpawners.GetTag("monster").(*ListTag)
	if !ok {
		t.Fatalf("Expected ListTag for 'monster', got %T", readSpawners.GetTag("monster"))
	}

	// Check the skeleton entry
	readSkeleton, ok := readMonsterList.Value[1].(*CompoundTag)
	if !ok {
		t.Fatalf("Expected CompoundTag for skeleton, got %T", readMonsterList.Value[1])
	}

	readSkeletonType, ok := readSkeleton.GetTag("type").(*StringTag)
	if !ok {
		t.Fatalf("Expected StringTag for skeleton type, got %T", readSkeleton.GetTag("type"))
	}

	if readSkeletonType.Value != "minecraft:skeleton" {
		t.Errorf("Expected skeleton type to be 'minecraft:skeleton', got '%s'", readSkeletonType.Value)
	}
}
