package outline

import(
	"image/png"
	"os"
	"strconv"
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

func LoadTerrainData(obj *terrain.LocalTerrainObject, config *terrain.GlobalConfig, file_name string){

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

	elevation_abs := obj.WorldTerrain.ElevationBaseM*0.5

	for y := 0; y < data_h; y++ {
		for x := 0; x < data_w; x++ {
			target := get_pixel_color(x, y)
			obj.ElevationTable[y][x] = (0.5-target)*elevation_abs
			//fmt.Println(target, obj.ElevationTable[y][x])
		}
	}

	Interpolate(obj,&obj.WorldTerrain.Config)

	obj.ElevationTableIsAvailable = true
}
