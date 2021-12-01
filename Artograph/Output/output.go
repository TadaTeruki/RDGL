/*
Artograph/Output/output.go
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
	artograph "../"
	terrain "../TerrainGeneration"
	cairo "github.com/ungerik/go-cairo"
	"math"
	//"fmt"
)

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

func DEMToPNG(file string, ats *artograph.ArtoDEM, image_pixel_w, image_pixel_h int,
				with_shadow bool, shadow_direction_z float64, shadow_direction_xy float64, shadow_width_Km float64,
				shadow_strength_01 float64){
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
				//fb := 0.0
				//count := 0.0
				/*
				for iy := y-1; iy <= y+1; iy++{
					for ix := x-1; ix <= x+1; ix++{
						if ix < 0 || iy <0 || ix >= fw || iy >= fh { continue }
						fb += brightness[int(iy)][int(ix)]
						count += 1.0
					}
				}
				*/

				//fb /= count
				fb := brightness[int(y)][int(x)]
				//fmt.Println(fb)

				fb = fb*shadow_strength_01+(1.0-shadow_strength_01)
				if fb < 0.0 { fb = 0.0 }
				if fb > 1.0 { fb = 1.0 }
				
				surface.SetSourceRGB(color.R*fb, color.G*fb, color.B*fb)
				//surface.SetSourceRGB(fb, fb, fb)
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

func WriteArtoDEMToPNG(file string, ats *artograph.ArtoDEM, image_pixel_w, image_pixel_h int){
	DEMToPNG(file, ats, image_pixel_w, image_pixel_h, false, 0.0, 0.0, 0.0, 0.0)
}

func WriteArtoDEMToPNGWithShadow(file string, ats *artograph.ArtoDEM, image_pixel_w, image_pixel_h int){
	DEMToPNG(file, ats, image_pixel_w, image_pixel_h, true, math.Pi/4.0, math.Pi/4.0, ats.UnitKm*5, 0.1)
}