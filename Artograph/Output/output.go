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