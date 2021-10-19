package artif_terrain

import(
	perlin "../PerlinNoise"
	//utility "../Utility"
)

type WorldTerrainObject struct{
	NoiseSrc, AdjNoiseSrc perlin.PerlinObject
	NSKm float64
	WEKm float64
	ElevationBaseM float64
	NEFPointList []NEFPoint
	Config GlobalConfig
}

