package cli

import "strconv"

func convertToFloat32(ans any) any {
	value, err := strconv.ParseFloat(ans.(string), 32)
	if err != nil {
		return float32(0.28)
	}
	return float32(value)
}
