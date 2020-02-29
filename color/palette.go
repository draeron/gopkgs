package color

var (
	Palette []RGB
)

func init() {

	for cl := Red; cl < White; cl++ {
		Palette = append(Palette, HSV{uint16(cl * 30), 100, 100}.RGB())
	}

	// white
	Palette = append(Palette, RGB{255, 255, 255, 255})

	// black
	Palette = append(Palette, RGB{0, 0, 0, 255})

	// light gray
	gray := RGB{192, 192, 192, 255}
	Palette = append(Palette, gray)

	// gray
	gray = RGB{128, 128, 128, 255}
	Palette = append(Palette, gray)

	// dark gray
	gray = RGB{64, 64, 64, 255}
	Palette = append(Palette, gray)

	// transparent
	gray = RGB{0, 0, 0, 0}
	Palette = append(Palette, gray)

}
