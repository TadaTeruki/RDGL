/*
github.com/TadaTeruki/RDGL/Output/output.go
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

package output

import(
	rdg "github.com/TadaTeruki/RDGL"
	terrain "github.com/TadaTeruki/RDGL/TerrainGeneration"
	cairo "github.com/ungerik/go-cairo"
	"math"
	"os"
	"strconv"
)
type Shadow struct{
	DirectionZ 	float64
	DirectionXY float64
	WidthKm		float64
	StrengthLand01 	float64
	StrengthOcean01	float64
}

func WriteWorldToPNG(file string, obj *terrain.WorldTerrainObject, image_pixel_w, image_pixel_h int){

	fw := float64(image_pixel_w)
	fh := float64(image_pixel_h)

	if fw < 0 { fw = fh/obj.NSKm*obj.WEKm }
	if fh < 0 { fh = fw/obj.WEKm*obj.NSKm }

	surface := cairo.NewSurface(cairo.FORMAT_ARGB32, int(fw), int(fh))

	for y := 0.0; y < fh; y += 1.0{
		for x := 0.0; x < fw; x += 1.0{
			px := x/fw
			py := y/fh
			color := GetColorFromElevation(obj.GetElevationByKmPoint(px*obj.WEKm, py*obj.NSKm))
			surface.SetSourceRGB(color.R, color.G, color.B)
			surface.Rectangle(x, y, 2, 2)
			surface.Fill()

		}
	} 

	surface.WriteToPNG(file)
	surface.Finish()

}

func WriteLocalToPNG(file string, obj *terrain.LocalTerrainObject, image_pixel_w, image_pixel_h int){

	fw := float64(image_pixel_w)
	fh := float64(image_pixel_h)

	if fw < 0 { fw = fh/obj.NSKm*obj.WEKm }
	if fh < 0 { fh = fw/obj.WEKm*obj.NSKm }

	surface := cairo.NewSurface(cairo.FORMAT_ARGB32, int(fw), int(fh))

	for y := 0.0; y < fh; y += 1.0{
		for x := 0.0; x < fw; x += 1.0{
			px := x/fw
			py := y/fh
			color := GetColorFromElevation(obj.GetElevationByKmPoint(px*obj.WEKm, py*obj.NSKm))
			surface.SetSourceRGB(color.R, color.G, color.B)
			surface.Rectangle(x, y, 2, 2)
			surface.Fill()

		}
	}

	surface.WriteToPNG(file)
	surface.Finish()

}

func DEMToPNG(file string, ats *rdg.DEM, image_pixel_w, image_pixel_h int, with_shadow bool, shadow Shadow){

	shadow_direction_z  := shadow.DirectionZ
	shadow_direction_xy := shadow.DirectionXY
	shadow_width_Km		:= shadow.WidthKm
	shadow_strength_land_01	:= shadow.StrengthLand01
	shadow_strength_ocean_01	:= shadow.StrengthOcean01

	fw := float64(image_pixel_w)
	fh := float64(image_pixel_h)

	if fw < 0 { fw = fh/ats.VerticalKm*ats.HorizontalKm }
	if fh < 0 { fh = fw/ats.HorizontalKm*ats.VerticalKm }

	surface := cairo.NewSurface(cairo.FORMAT_ARGB32, int(fw), int(fh))

	brightness := make([][]float64, int(fh))

	for y := 0.0; y < fh; y += 1.0{
		brightness[int(y)] = make([]float64, int(fw))
		for x := 0.0; x < fw; x += 1.0{
			px := x/fw
			py := y/fh
			elevation, err := ats.GetElevationByKmPoint(px*ats.HorizontalKm, py*ats.VerticalKm)
			if err != nil {
				continue
			}
			if with_shadow == true {
				rxKm := px*ats.HorizontalKm
				ryKm := py*ats.VerticalKm
				x1Km := rxKm+math.Cos(shadow_direction_z)*shadow_width_Km
				y1Km := ryKm+math.Sin(shadow_direction_z)*shadow_width_Km
				x2Km := rxKm-math.Cos(shadow_direction_z)*shadow_width_Km
				y2Km := ryKm-math.Sin(shadow_direction_z)*shadow_width_Km
				if x1Km < 0 { x1Km = 0 }
				if y1Km < 0 { y1Km = 0 }
				if x2Km < 0 { x2Km = 0 }
				if y2Km < 0 { y2Km = 0 }
				if x1Km > ats.HorizontalKm { x1Km = ats.HorizontalKm }
				if y1Km > ats.VerticalKm   { y1Km = ats.VerticalKm   }
				if x2Km > ats.HorizontalKm { x2Km = ats.HorizontalKm }
				if y2Km > ats.VerticalKm   { y2Km = ats.VerticalKm   }
				elv1, err1 := ats.GetElevationByKmPoint(x1Km, y1Km)
				if err1 !=nil { elv1 = elevation }
				elv2, err2 := ats.GetElevationByKmPoint(x2Km, y2Km)
				if err2 !=nil { elv2 = elevation }
				elv_d := elv1-elv2
				dst_d := math.Sqrt((x1Km-x2Km)*(x1Km-x2Km)+(y1Km-y2Km)*(y1Km-y2Km))+0.0001
				dxt_xy := math.Atan(elv_d/dst_d)
				dxt_xy_d := dxt_xy-shadow_direction_xy
				brightness[int(y)][int(x)] = dxt_xy_d/(math.Pi/2.0)
			} else {
				brightness[int(y)][int(x)] = 1.0
			}

		}
	}

	for y := 0.0; y < fh; y += 1.0{
		for x := 0.0; x < fw; x += 1.0{
			px := x/fw
			py := y/fh
			elevation, err := ats.GetElevationByKmPoint(px*ats.HorizontalKm, py*ats.VerticalKm)
			if err != nil {
				continue
			}
			color := GetColorFromElevation(elevation)

			if with_shadow == true {
				
				fb := brightness[int(y)][int(x)]

				var shadow_strength float64
				
				if elevation >= 0.0{
					shadow_strength = shadow_strength_land_01 * (elevation/ats.ElevationAbsM)
				} else {
					shadow_strength = shadow_strength_ocean_01 * ((-elevation)/ats.ElevationAbsM)
				}
				
				fb = fb*shadow_strength+(1.0-shadow_strength)
				if fb < 0.0 { fb = 0.0 }
				if fb > 1.0 { fb = 1.0 }
				
				surface.SetSourceRGB(color.R*fb, color.G*fb, color.B*fb)
			} else {
				surface.SetSourceRGB(color.R, color.G, color.B)
			}

			
			surface.Rectangle(x, y, 2, 2)
			surface.Fill()

		}
	}

	surface.WriteToPNG(file)
	surface.Finish()
}

func WriteDEMtoPNG(file string, ats *rdg.DEM, image_pixel_w, image_pixel_h int){
	var shadow Shadow
	DEMToPNG(file, ats, image_pixel_w, image_pixel_h, false, shadow)
}

func DefaultShadow(ats *rdg.DEM) Shadow{
	var shadow Shadow
	shadow.DirectionZ = math.Pi/4.0
	shadow.DirectionXY = math.Pi/4.0
	shadow.WidthKm = ats.UnitKm*5.0
	shadow.StrengthLand01 = 0.5
	shadow.StrengthOcean01 = 0.05
	return shadow
}

func WriteDEMtoPNGwithShadow(file string, ats *rdg.DEM, image_pixel_w, image_pixel_h int, shadow Shadow){
	DEMToPNG(file, ats, image_pixel_w, image_pixel_h, true, shadow)
}

func FtoA(v float64) string{
	return strconv.FormatFloat(v, 'f', -1, 64)
}


func ItoA(v int) string{
	return strconv.Itoa(v)
}

func WriteDEMtoOBJ(filename string, ats *rdg.DEM, image_w float64, image_h float64, z_extend float64, z_is_vertical bool) error {
	var file *os.File
	var err error

	file, err = os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	

	if image_w < 0 { image_w = image_h/ats.VerticalKm*ats.HorizontalKm }
	if image_h < 0 { image_h = image_w/ats.HorizontalKm*ats.VerticalKm }

	write := func(s string){
		_, err = file.WriteString(s)
		if err != nil {
			panic(err)
		}
	}

	write("g DEM\n")
	
	data_w := 0
	data_h := 0

	for yKm := 0.0; yKm < ats.VerticalKm; yKm += ats.UnitKm {
		
		for xKm := 0.0; xKm < ats.HorizontalKm; xKm += ats.UnitKm {

			elevation, err := ats.GetElevationByKmPoint(xKm, yKm)
			if err != nil {
				continue
			}
			px := xKm/ats.HorizontalKm*image_w
			py := yKm/ats.VerticalKm*image_h
			pz := (elevation*0.001)/ats.HorizontalKm*image_w*z_extend
			if z_is_vertical == true {
				write("v "+FtoA(px)+" "+FtoA(py)+" "+FtoA(pz)+"\n")
			} else {
				write("v "+FtoA(px)+" "+FtoA(pz)+" "+FtoA(py)+"\n")
			}
			if yKm== 0.0 { data_w++ }
		}
		data_h++
	}

	ad := func(ix, iy int) int{
		return iy*data_w+ix+1
	}

	for y := 0; y < data_h-1; y++{
		for x := 0; x < data_w-1; x++{

			a := ItoA(ad(x,y))
			b := ItoA(ad(x+1,y))
			c := ItoA(ad(x,y+1))
			d := ItoA(ad(x+1,y+1))
			
			write("f "+a+" "+b+" "+d+" "+c+"\n")
		}
	}

	return nil
}


func WriteDEMtoTXT(filename string, ats *rdg.DEM, image_w float64, image_h float64) error {
	var file *os.File
	var err error

	file, err = os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	
	if image_w < 0 { image_w = image_h/ats.VerticalKm*ats.HorizontalKm }
	if image_h < 0 { image_h = image_w/ats.HorizontalKm*ats.VerticalKm }

	write := func(s string){
		_, err = file.WriteString(s)
		if err != nil {
			panic(err)
		}
	}
	write("#HorizontalKm\n")
	write(FtoA(ats.HorizontalKm)+"\n\n")
	write("#VerticalKm\n")
	write(FtoA(ats.VerticalKm)+"\n\n")
	write("#UnitKm\n")
	write(FtoA(ats.UnitKm)+"\n\n")


	for yKm := 0.0; yKm < ats.VerticalKm; yKm += ats.UnitKm {
		for xKm := 0.0; xKm < ats.HorizontalKm; xKm += ats.UnitKm {
			elevation, err := ats.GetElevationByKmPoint(xKm, yKm)
			if err != nil {
				continue
			}
			write(FtoA(elevation))
			if xKm+ats.UnitKm < ats.HorizontalKm {
				write(",")
			}
		}
		write("\n")
	}

	return nil
}