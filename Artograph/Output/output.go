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
	terrain "../TerrainGenelation"
	cairo "github.com/ungerik/go-cairo"
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

func WriteArtoTerrainSurfaceToPNG(file string, ats *artograph.ArtoTerrainSurface, image_pixel_w, image_pixel_h int){
	fw := float64(image_pixel_w)
	fh := float64(image_pixel_h)

	if fw < 0 { fw = fh/ats.VerticalKm*ats.HorizontalKm }
	if fh < 0 { fh = fw/ats.HorizontalKm*ats.VerticalKm }

	surface := cairo.NewSurface(cairo.FORMAT_ARGB32, int(fw), int(fh))

	for y := 0.0; y < fh; y += 1.0{
		for x := 0.0; x < fw; x += 1.0{
			px := x/fw
			py := y/fh
			elevation, err := ats.GetElevationByKmPoint(px*ats.HorizontalKm, py*ats.VerticalKm)
			if err != nil {
				continue
			}
			color := GetColorFromElevation(elevation)

			surface.SetSourceRGB(color.R, color.G, color.B)
			surface.Rectangle(x, y, 2, 2)
			surface.Fill()

		}
	}

	surface.WriteToPNG(file)
	surface.Finish()
}