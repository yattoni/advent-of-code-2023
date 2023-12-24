package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	util "github.com/yattoni/advent-of-code-2023"
)

type mapping struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeLength           int
}

func (mapping mapping) String() string {
	return fmt.Sprintf("{sourceRangeStart=%d,destinationRangeStart=%d,rangeLength=%d}",
		mapping.sourceRangeStart, mapping.destinationRangeStart, mapping.rangeLength)
}

func makeMapping(lines []string) []mapping {
	mappings := make([]mapping, len(lines))
	for i, line := range lines {
		numStrs := strings.Split(line, " ")
		destinationRangeStart := util.MustAtoi(numStrs[0])
		sourceRangeStart := util.MustAtoi(numStrs[1])
		rangeLength := util.MustAtoi(numStrs[2])
		mappings[i] = mapping{destinationRangeStart, sourceRangeStart, rangeLength}
	}
	slices.SortStableFunc[[]mapping](mappings, func(a, b mapping) int {
		if a.sourceRangeStart > b.sourceRangeStart {
			return 1
		}
		return -1
	})
	return mappings
}

func findDest(source int, sourceToDestMapping []mapping) int {
	for _, cur := range sourceToDestMapping {
		if cur.sourceRangeStart <= source && cur.sourceRangeStart+cur.rangeLength > source {
			return cur.destinationRangeStart + (source - cur.sourceRangeStart)
		}
	}
	return source
}

// func findDests(sources []int, sourceToDestMapping []mapping) []int {
// 	m := make(map[int]int)
// 	for _, cur := range sourceToDestMapping {
// 		for i := cur.sourceRangeStart; i < cur.sourceRangeStart+cur.rangeLength; i++ {
// 			m[i] = cur.destinationRangeStart + i
// 		}
// 	}
// 	dests := make([]int, len(sources))
// 	for i, s := range sources {
// 		d, ok := m[s]
// 		if ok {
// 			dests[i] = d
// 		} else {
// 			dests[i] = s
// 		}
// 	}
// 	return dests
// }

func findSrc(destination int, sourceToDestMapping []mapping) int {
	for _, cur := range sourceToDestMapping {
		if cur.destinationRangeStart <= destination && cur.destinationRangeStart+cur.rangeLength > destination {
			return cur.sourceRangeStart + (destination - cur.destinationRangeStart)
		}
	}
	return destination
}

func mapSeedToLocation(seed int, seedToSoilMapping []mapping, soilToFertilizerMapping []mapping, fertilizerToWaterMapping []mapping, waterToLightMapping []mapping, lightToTemperatureMapping []mapping, temperatureToHumidityMapping []mapping, humidityToLocationMapping []mapping) int {
	fmt.Print("seed=", seed)
	soil := findDest(seed, seedToSoilMapping)
	fmt.Print(",soil=", soil)
	fertilizer := findDest(soil, soilToFertilizerMapping)
	fmt.Print(",fertilizer=", fertilizer)
	water := findDest(fertilizer, fertilizerToWaterMapping)
	fmt.Print(",water=", water)
	light := findDest(water, waterToLightMapping)
	fmt.Print(",light=", light)
	temperature := findDest(light, lightToTemperatureMapping)
	fmt.Print(",temperature=", temperature)
	humidity := findDest(temperature, temperatureToHumidityMapping)
	fmt.Print(",humidity=", humidity)
	location := findDest(humidity, humidityToLocationMapping)
	fmt.Print(",location=", location)
	fmt.Println()
	return location
}

func mapLocationToSeed(location int, seedToSoilMapping []mapping, soilToFertilizerMapping []mapping, fertilizerToWaterMapping []mapping, waterToLightMapping []mapping, lightToTemperatureMapping []mapping, temperatureToHumidityMapping []mapping, humidityToLocationMapping []mapping) int {
	humidity := findSrc(location, humidityToLocationMapping)
	temperature := findSrc(humidity, temperatureToHumidityMapping)
	light := findSrc(temperature, lightToTemperatureMapping)
	water := findSrc(light, waterToLightMapping)
	fertilizer := findSrc(water, fertilizerToWaterMapping)
	soil := findSrc(fertilizer, soilToFertilizerMapping)
	seed := findSrc(soil, seedToSoilMapping)
	// fmt.Print("seed=", seed)
	// fmt.Print(",soil=", soil)
	// fmt.Print(",fertilizer=", fertilizer)
	// fmt.Print(",water=", water)
	// fmt.Print(",light=", light)
	// fmt.Print(",temperature=", temperature)
	// fmt.Print(",humidity=", humidity)
	// fmt.Print(",location=", location)
	// fmt.Println()
	return seed
}

// func mapSeedsToLocations(seed []int, seedToSoilMapping []mapping, soilToFertilizerMapping []mapping, fertilizerToWaterMapping []mapping, waterToLightMapping []mapping, lightToTemperatureMapping []mapping, temperatureToHumidityMapping []mapping, humidityToLocationMapping []mapping) []int {
// 	soil := findDests(seed, seedToSoilMapping)
// 	fmt.Println("done with soil")
// 	fertilizer := findDests(soil, soilToFertilizerMapping)
// 	fmt.Println("done with fertilizer")
// 	water := findDests(fertilizer, fertilizerToWaterMapping)
// 	fmt.Println("done with water")
// 	light := findDests(water, waterToLightMapping)
// 	fmt.Println("done with light")
// 	temperature := findDests(light, lightToTemperatureMapping)
// 	fmt.Println("done with temperature")
// 	humidity := findDests(temperature, temperatureToHumidityMapping)
// 	fmt.Println("done with humidity")
// 	location := findDests(humidity, humidityToLocationMapping)
// 	fmt.Println("done with location")
// 	return location
// }

func main() {
	lines := util.ReadFileWithSpacesToLines("input")

	seedStrs := strings.Split(strings.TrimSpace(strings.TrimPrefix(lines[0], "seeds:")), " ")
	seeds := make([]int, len(seedStrs))
	for i := 0; i < len(seedStrs); i++ {
		seeds[i] = util.MustAtoi(seedStrs[i])
	}
	fmt.Println(seeds)

	seedToSoilIdx := slices.Index[[]string](lines, "seed-to-soil map:")
	// fmt.Println(seedToSoilIdx)
	soilToFertilizerIdx := slices.Index[[]string](lines, "soil-to-fertilizer map:")
	// fmt.Println(soilToFertilizerIdx)
	fertilizerToWaterIdx := slices.Index[[]string](lines, "fertilizer-to-water map:")
	// fmt.Println(fertilizerToWaterIdx)
	waterToLightIdx := slices.Index[[]string](lines, "water-to-light map:")
	// fmt.Println(waterToLightIdx)
	lightToTemperatureIdx := slices.Index[[]string](lines, "light-to-temperature map:")
	// fmt.Println(lightToTemperatureIdx)
	temperatureToHumidityIdx := slices.Index[[]string](lines, "temperature-to-humidity map:")
	// fmt.Println(temperatureToHumidityIdx)
	humidityToLocationIdx := slices.Index[[]string](lines, "humidity-to-location map:")
	// fmt.Println(humidityToLocationIdx)

	seedToSoilMapping := makeMapping(lines[seedToSoilIdx+1 : soilToFertilizerIdx-1])
	fmt.Println(seedToSoilMapping)
	soilToFertilizerMapping := makeMapping(lines[soilToFertilizerIdx+1 : fertilizerToWaterIdx-1])
	fmt.Println(soilToFertilizerMapping)
	fertilizerToWaterMapping := makeMapping(lines[fertilizerToWaterIdx+1 : waterToLightIdx-1])
	fmt.Println(fertilizerToWaterMapping)
	waterToLightMapping := makeMapping(lines[waterToLightIdx+1 : lightToTemperatureIdx-1])
	fmt.Println(waterToLightMapping)
	lightToTemperatureMapping := makeMapping(lines[lightToTemperatureIdx+1 : temperatureToHumidityIdx-1])
	fmt.Println(lightToTemperatureMapping)
	temperatureToHumidityMapping := makeMapping(lines[temperatureToHumidityIdx+1 : humidityToLocationIdx-1])
	fmt.Println(temperatureToHumidityMapping)
	humidityToLocationMapping := makeMapping(lines[humidityToLocationIdx+1:])
	fmt.Println(humidityToLocationMapping)

	lowestLocation := math.MaxInt
	for _, seed := range seeds {
		location := mapSeedToLocation(seed, seedToSoilMapping, soilToFertilizerMapping, fertilizerToWaterMapping, waterToLightMapping, lightToTemperatureMapping, temperatureToHumidityMapping, humidityToLocationMapping)
		if location < lowestLocation {
			lowestLocation = location
		}
	}
	fmt.Println("part1 lowest location:", lowestLocation) // 57075758

	part2Seeds := []int{}
	for i := 0; i < len(seeds); i += 2 {
		for j := seeds[i]; j < seeds[i]+seeds[i+1]; j++ {
			part2Seeds = append(part2Seeds, j)
		}
	}
	slices.Sort[[]int](part2Seeds)

	for i := 0; i < math.MaxInt; i++ {
		seed := mapLocationToSeed(i, seedToSoilMapping, soilToFertilizerMapping, fertilizerToWaterMapping, waterToLightMapping, lightToTemperatureMapping, temperatureToHumidityMapping, humidityToLocationMapping)
		_, found := slices.BinarySearch[[]int](part2Seeds, seed)
		if found {
			fmt.Println("part2: seed", seed, "location", i) // part2: seed 3267749434 location 31161857
			break
		}
	}
	// mapLocationToSeed(35, seedToSoilMapping, soilToFertilizerMapping, fertilizerToWaterMapping, waterToLightMapping, lightToTemperatureMapping, temperatureToHumidityMapping, humidityToLocationMapping)
	// seed 14647953 location 6295169 too low
	// part2: seed 3267749434 location 31161857

	// lowestLocation2 := math.MaxInt
	// locations := mapSeedsToLocations(seeds, seedToSoilMapping, soilToFertilizerMapping, fertilizerToWaterMapping, waterToLightMapping, lightToTemperatureMapping, temperatureToHumidityMapping, humidityToLocationMapping)
	// for _, l := range locations {
	// 	if l < lowestLocation2 {
	// 		lowestLocation2 = l
	// 	}
	// }
	// fmt.Println(lowestLocation2)
}
