
/*
Artograph/Outline/outline.go
Copyright (C) 2021 Tada Teruki

This program is free software; you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation; either version 3 of the License, or 
(at your option) any later version.

This program is distributed in the hope that it will be useful, 
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the 
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.
*/


package outline

import(
	"image/png"
	"os"
	"strconv"
	//"math"
	terrain "../TerrainGeneration"
	//"fmt"
)

const(
	PointDataQuality = 10000000
	ScoreIDQuality = 10000000
)


func Atof(tar string) float64{
	res, err := strconv.ParseFloat(tar, 64)
    if err != nil {
		return 0.0
    }
	return res
}

func Atoi(tar string) int{
	for i := 0; i<len(tar); i++{
		c := tar[i]
		if int(c)<int('0') || int(c)>int('9') {
			tar = tar[:i]+tar[i+1:]
			i--
		}
	}
	res, err := strconv.Atoi(tar)
    if err != nil {
		return 0
    }
	return res
}


func GetImageScale(file_name string) (float64, float64){

    png_file, err := os.Open(file_name)
    if err != nil {
		panic(err)
    }

    image, err := png.Decode(png_file)
    if err != nil {
		panic(err)
    }

	bounds := image.Bounds()

	data_w := bounds.Max.X - bounds.Min.X
	data_h := bounds.Max.Y - bounds.Min.Y

	return float64(data_w), float64(data_h)

	
}

func LoadTerrainData(obj *terrain.LocalTerrainObject, config *terrain.InternalConfig, file_name string){

    png_file, err := os.Open(file_name)
    if err != nil {
		panic(err)
    }

    image, err := png.Decode(png_file)
    if err != nil {
		panic(err)
    }

	bounds := image.Bounds()

	data_w := bounds.Max.X - bounds.Min.X
	data_h := bounds.Max.Y - bounds.Min.Y

	obj.ElevationTable = make([][]float64, data_h)
	for i := 0; i<data_h; i++{
		obj.ElevationTable[i] = make([]float64, data_w)
	}

	get_pixel_color := func(x, y int) float64 {
		
		r, g, b, a := image.At(x, y).RGBA()

		r /= 256
		g /= 256
		b /= 256
		a /= 256

		if a != 255 { return 0 }

		return (float64(r+g+b))/(255.0*3.0)
	}

	for y := 0; y < data_h; y++ {
		for x := 0; x < data_w; x++ {
			target := get_pixel_color(x, y)
			obj.ElevationTable[y][x] = obj.WorldTerrain.GetElevationFromNoiseLevel(obj.WorldTerrain.Config.StandardLandProportion-(target-0.5)*0.5)
			
		}
	}

	Interpolate(obj,&obj.WorldTerrain.Config)

	obj.ElevationTableIsAvailable = true
}
