package artif_terrain

import(
	cairo "github.com/ungerik/go-cairo"
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
			//_,_ = px, py
			color := GetColorFromElevation(obj.GetElevation(px*obj.WEKm, py*obj.NSKm))
			surface.SetSourceRGB(color.R, color.G, color.B)
			surface.Rectangle(x, y, 2, 2)
			surface.Fill()

		}
	} 

	surface.WriteToPNG("data.png")
	surface.Finish()

	/*
	for i := 0; i<1000; i++{
		surface.SetSourceRGB(1.0,0.0,0.0)
		iff := 0.001*float64(i)
		surface.Rectangle(
			float64(image_pixel_w)*(iff),
			float64(image_pixel_h)*f(iff), 2, 2)
		surface.Fill()
	}

	surface.WriteToPNG("data.png")
	surface.Finish()
	*/
}