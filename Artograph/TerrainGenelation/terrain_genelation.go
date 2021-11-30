package terrain_generation

import(
	perlin "../PerlinNoise"
)

// WorldTerrain ... Whole area to generate terrains
type WorldTerrainObject struct{
	NoiseSrc, AdjNoiseSrc perlin.PerlinObject
	NSKm float64
	WEKm float64
	ElevationBaseM float64
	NEFPointList []NEFPoint
	Config GlobalConfig
}

// LocalTerrain ... A part of terrains generated in WorldTerrain
// (Effects are only applied for LocalTerrain)
type LocalTerrainObject struct{
	WorldTerrain *WorldTerrainObject
	xKm float64
	yKm float64
	NSKm float64
	WEKm float64
	OceanTable [][]OceanPoint
	OceanCheckIsAvailable bool
	LiverTable [][]LiverPoint
	LiverCheckIsAvailable bool
}

type GlobalConfig struct{ // details -> <globalconfig.go>
	Seed int64
	NoizeScaleKm float64
	NoizeOctave int
	NoizeMinPersistence float64
	NoizeMaxPersistence float64
	MinLand float64
	MaxLand float64
	LocalTerrainSelectionQuality int
	OceanCheckIntervalKm float64
	LiverCheckIntervalKm float64
	TerrainReverseScale float64
	VirtualOceanElevation float64
}


const DIRECTION_NONE  = 0
const DIRECTION_NORTH = 1
const DIRECTION_WEST  = 2
const DIRECTION_SOUTH = 3
const DIRECTION_EAST  = 4

type KmPoint struct{
	XKm float64
	YKm float64
}

type LiverPoint struct{
	XKm float64
	YKm float64
	Direction int
	Cavity float64
	BaseElevation float64
}

type NEFPoint struct {
	NoiseLevel float64
	Elevation float64
}

type OceanPoint struct{
	XKm float64
	YKm float64
	IsOcean bool
}

type PathPoint struct{
	point KmPoint
	parent KmPoint
	score float64
	flag int // 0:untreated 1:opened 2:closed
}

func MakeKmPoint(xKm, yKm float64) KmPoint{
	var r KmPoint
	r.XKm = xKm
	r.YKm = yKm
	return r
} 