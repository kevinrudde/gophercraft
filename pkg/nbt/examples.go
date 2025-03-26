package nbt

import (
	"bytes"
	"log"
)

// CreateBiomeEntry creates an example biome entry as it would appear in a registry
func CreateBiomeEntry(biomeName string) *CompoundTag {
	// Create the root compound tag for the biome
	biome := NewCompoundTag("biome")

	// Add basic properties
	biome.PutTag(NewStringTag("name", biomeName))
	biome.PutTag(NewFloatTag("temperature", 0.8))
	biome.PutTag(NewFloatTag("downfall", 0.4))
	biome.PutTag(NewByteTag("has_precipitation", 1)) // true

	// Create effects compound
	effects := NewCompoundTag("effects")
	effects.PutTag(NewIntTag("sky_color", 7907327))
	effects.PutTag(NewIntTag("water_color", 4159204))
	effects.PutTag(NewIntTag("water_fog_color", 329011))
	effects.PutTag(NewIntTag("fog_color", 12638463))

	// Add grass colors
	effects.PutTag(NewIntTag("grass_color", 9551193))
	effects.PutTag(NewIntTag("foliage_color", 7842607))

	// Add mood sound
	moodSound := NewCompoundTag("mood_sound")
	moodSound.PutTag(NewStringTag("sound", "minecraft:ambient.cave"))
	moodSound.PutTag(NewIntTag("tick_delay", 6000))
	moodSound.PutTag(NewFloatTag("offset", 2.0))
	moodSound.PutTag(NewDoubleTag("block_search_extent", 8.0))

	effects.PutTag(moodSound)

	// Add effects to biome
	biome.PutTag(effects)

	// Create and add creature spawn probabilities
	spawners := NewCompoundTag("spawners")

	// Add monster spawners list
	monsterList := NewListTag("monster", TagCompound)

	// Create a zombie spawn entry
	zombie := NewCompoundTag("")
	zombie.PutTag(NewStringTag("type", "minecraft:zombie"))
	zombie.PutTag(NewIntTag("weight", 95))
	zombie.PutTag(NewIntTag("minCount", 4))
	zombie.PutTag(NewIntTag("maxCount", 4))

	// Create a skeleton spawn entry
	skeleton := NewCompoundTag("")
	skeleton.PutTag(NewStringTag("type", "minecraft:skeleton"))
	skeleton.PutTag(NewIntTag("weight", 100))
	skeleton.PutTag(NewIntTag("minCount", 4))
	skeleton.PutTag(NewIntTag("maxCount", 4))

	// Add entries to monster list
	monsterList.Add(zombie)
	monsterList.Add(skeleton)

	// Add monster list to spawners
	spawners.PutTag(monsterList)

	// Add creature list (empty example)
	creatureList := NewListTag("creature", TagCompound)
	spawners.PutTag(creatureList)

	// Add spawners to biome
	biome.PutTag(spawners)

	// Add spawn costs for certain entities
	spawnCosts := NewCompoundTag("spawn_costs")

	// Zombie spawn cost
	zombieCost := NewCompoundTag("minecraft:zombie")
	zombieCost.PutTag(NewDoubleTag("energy_budget", 0.15))
	zombieCost.PutTag(NewDoubleTag("charge", 0.7))

	// Add spawn costs to biome
	spawnCosts.PutTag(zombieCost)
	biome.PutTag(spawnCosts)

	return biome
}

// Example function showing how to use NBT in a registry
func ExampleCreateBiomeRegistry() {
	// Create a biome registry entry
	biome := CreateBiomeEntry("minecraft:plains")

	// Write it to bytes
	var buf bytes.Buffer
	err := WriteNamedTag(&buf, biome)
	if err != nil {
		log.Printf("Error writing NBT: %v", err)
		return
	}

	// You can now use these bytes in your registry data packets
	// nbtData := buf.Bytes()

	// To read the data back
	reader := bytes.NewReader(buf.Bytes())
	readTag, err := ReadNamedTag(reader)
	if err != nil {
		log.Printf("Error reading NBT: %v", err)
		return
	}

	// Print the tag structure
	log.Printf("Read NBT data: %s", readTag.String())
}
