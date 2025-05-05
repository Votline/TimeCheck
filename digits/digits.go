package digits

func CreateVertexDigits(number int, offset float32) ([]float32, int32) {
	switch number{
	case 1:
		return []float32 {	
			-0.8 + offset, 0.0, 0.0,
			-0.7 + offset, 0.3, 0.0,
			-0.7 + offset, -0.275, 0.0,
		}, 3
	case 2:
		return []float32 {
			-0.8 + offset, 0.2, 0.0,
			-0.75 + offset, 0.3, 0.0,
			-0.7 + offset, 0.2, 0.0,
		
			-0.8 + offset, -0.2, 0.0,
			-0.68 + offset, -0.2, 0.0,
		}, 5
	case 3:
		return []float32 {
			-0.8 + offset, 0.3, 0.0,
			-0.7 + offset, 0.3, 0.0,

			-0.7 + offset, 0.05, 0.0,
			-0.8 + offset, 0.05, 0.0,
			-0.7 + offset, 0.05, 0.0,

			-0.7 + offset, -0.2, 0.0,
			-0.8 + offset, -0.2, 0.0,
		}, 7
	case 4:
		return []float32 {
			-0.8 + offset, 0.3, 0.0,
			-0.8 + offset, 0.1, 0.0,
		
			-0.7 + offset, 0.1, 0.0,
			-0.7 + offset, 0.3, 0.0,

			-0.7 + offset, -0.25, 0.0,	
		}, 5
	case 5:
		return []float32 {
			-0.7 + offset, 0.3, 0.0,
			-0.8 + offset, 0.3, 0.0,

			-0.8 + offset, 0.05, 0.0,
			-0.7 + offset, 0.05, 0.0,
			
			-0.7 + offset, -0.2, 0.0,
			-0.8 + offset, -0.2, 0.0,
		}, 6
	case 6:
		return []float32 {
			-0.7 + offset, 0.3, 0.0,
			-0.8 + offset, 0.3, 0.0,

			-0.8 + offset, -0.2, 0.0,
			-0.7 + offset, -0.2, 0.0,
			
			-0.7 + offset, 0.05, 0.0,
			-0.8 + offset, 0.05, 0.0,
		}, 6
	case 7:
		return []float32 {
			-0.8 + offset, 0.3, 0.0,
			-0.7 + offset, 0.3, 0.0,

			-0.75 + offset, -0.25, 0.0,
		}, 3
	case 8:
		return []float32 {
			-0.8 + offset, 0.3, 0.0,
			-0.7 + offset, 0.3, 0.0,

			-0.7 + offset, -0.2, 0.0,
			-0.8 + offset, -0.2, 0.0,

			-0.8 + offset, 0.3, 0.0,
			-0.8 + offset, 0.05, 0.0,
			-0.7 + offset, 0.05, 0.0,
		}, 7
	case 9:
		return []float32 {
			-0.7 + offset, 0.3, 0.0,
			-0.8 + offset, 0.3, 0.0,
			
			-0.8 + offset, 0.05, 0.0,
			-0.7 + offset, 0.05, 0.0,
			-0.7 + offset, 0.3, 0.0,
		
			-0.7 + offset, -0.27, 0.0,
		}, 6
	case 0:
		return []float32 {
			-0.8 + offset, 0.3, 0.0,
			-0.7 + offset, 0.3, 0.0,

			-0.7 + offset, -0.2, 0.0,
			-0.8 + offset, -0.2, 0.0,

			-0.8 + offset, 0.3, 0.0,
		}, 5
	case 10:
		return []float32 {
			-0.8 + offset, 0.3, 0.0,
			-0.8 + offset, 0.25, 0.0,
		}, 2
	case 11:
		return []float32 {
			-0.8 + offset, -0.2, 0.0,
			-0.8 + offset, -0.25, 0.0,
		}, 2
	default:
		return []float32 {
			0.0, 0.0, 0.0,
		}, 1 
	}
}
