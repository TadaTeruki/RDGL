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

	/*	//コメントアウトにより、Oceanの範囲を確認可能
	for y := 0; y<len(obj.OceanTable); y++{
		for x := 0; x<len(obj.OceanTable[y]); x++{
			if obj.OceanTable[y][x].IsOcean == true {
				dx := obj.OceanTable[y][x].XKm/obj.WEKm*fw
				dy := obj.OceanTable[y][x].YKm/obj.NSKm*fh
				surface.SetSourceRGB(1.0, 0.5, 0.5)
				surface.Rectangle(dx, dy, 1, 1)
				surface.Fill()
			}
		}
	}
	*/
	/*	
	for y := 0; y<len(obj.LiverTable); y++{
		for x := 0; x<len(obj.LiverTable[y]); x++{

			if obj.LiverTable[y][x].Direction == DIRECTION_NONE {
				continue
			}

			liver_interval_km := obj.WorldTerrain.Config.LiverCheckIntervalKm
			dx := obj.LiverTable[y][x].XKm/obj.WEKm*fw
			dy := obj.LiverTable[y][x].YKm/obj.NSKm*fh
			din := liver_interval_km/obj.WEKm*fw

			surface.SetSourceRGBA(0, 0, 0, 1.0-obj.LiverTable[y][x].Cavity)
			surface.MoveTo(dx, dy)
			surface.SetLineWidth(1)

			switch obj.LiverTable[y][x].Direction{
				case DIRECTION_NORTH:{
					surface.LineTo(dx, dy-din)
				}
				case DIRECTION_WEST:{
					surface.LineTo(dx-din, dy)
				}
				case DIRECTION_SOUTH:{
					surface.LineTo(dx, dy+din)
				}
				case DIRECTION_EAST:{
					surface.LineTo(dx+din, dy)
				}
			}

			surface.Stroke()
			
			//surface.SetSourceRGB(1.0, 0.2, 0.2)
			//surface.Rectangle(dx, dy, 2, 2)
			//surface.Fill()
			
		}
	}
	*/
	
	
	
	

	surface.WriteToPNG("data.png")
	surface.Finish()

}