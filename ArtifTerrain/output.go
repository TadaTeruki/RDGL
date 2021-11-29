package artif_terrain

import(
	cairo "github.com/ungerik/go-cairo"
	color "../Color"
)

func (obj *WorldTerrainObject) WriteWorldToPNG(image_pixel_w, image_pixel_h int){

	fw := float64(image_pixel_w)
	fh := float64(image_pixel_h)

	if fw < 0 { fw = fh/obj.NSKm*obj.WEKm }
	if fh < 0 { fh = fw/obj.WEKm*obj.NSKm }

	surface := cairo.NewSurface(cairo.FORMAT_ARGB32, int(fw), int(fh))

	for y := 0.0; y < fh; y += 1.0{
		for x := 0.0; x < fw; x += 1.0{
			px := x/fw
			py := y/fh
			color := color.GetColorFromElevation(obj.GetElevationByKmPoint(px*obj.WEKm, py*obj.NSKm))
			surface.SetSourceRGB(color.R, color.G, color.B)
			surface.Rectangle(x, y, 2, 2)
			surface.Fill()

		}
	} 

	surface.WriteToPNG("data.png")
	surface.Finish()

}
func (obj *LocalTerrainObject) WriteLocalToPNG(image_pixel_w, image_pixel_h int){

	fw := float64(image_pixel_w)
	fh := float64(image_pixel_h)

	if fw < 0 { fw = fh/obj.NSKm*obj.WEKm }
	if fh < 0 { fh = fw/obj.WEKm*obj.NSKm }

	surface := cairo.NewSurface(cairo.FORMAT_ARGB32, int(fw), int(fh))

	for y := 0.0; y < fh; y += 1.0{
		for x := 0.0; x < fw; x += 1.0{
			px := x/fw
			py := y/fh
			color := color.GetColorFromElevation(obj.GetElevationByKmPoint(px*obj.WEKm, py*obj.NSKm))
			surface.SetSourceRGB(color.R, color.G, color.B)
			surface.Rectangle(x, y, 2, 2)
			surface.Fill()

		}
	}

	/*	コメントアウトにより、Oceanの範囲を確認可能
	for y := 0; y<len(obj.OceanTable); y++{
		for x := 0; x<len(obj.OceanTable[y]); x++{
			if obj.OceanTable[y][x].IsOcean == true {
				dx := obj.OceanTable[y][x].XKm/obj.WEKm*fw
				dy := obj.OceanTable[y][x].YKm/obj.NSKm*fh
				surface.SetSourceRGB(1.0, 0.0, 0.0)
				surface.Rectangle(dx, dy, 1, 1)
				surface.Fill()
			}
		}
	}
	*/

	surface.WriteToPNG("data.png")
	surface.Finish()

}