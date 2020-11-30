package basicadapter

// NewRectangle is the function generate vector
func NewRectangle(width, height int) *VectorImage {
	width--
	height--

	return &VectorImage{
		[]Line{
			{0, 0, width, 0},
			{0, 0, 0, height},
			{width, 0, width, height},
			{0, height, width, height},
		},
	}
}
