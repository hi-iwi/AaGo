package atype

func ConvFloat64Map(mi map[string]any) (map[string]float64, error) {
	if len(mi) == 0 {
		return nil, nil
	}
	maps := make(map[string]float64, len(mi))
	var err error
	for k, m := range mi {
		maps[k], err = Float64(m, 64)
		if err != nil {
			return nil, err
		}
	}
	return maps, nil
}

func ConvStrings(ai []any) []string {
	if len(ai) == 0 {
		return nil
	}
	a := make([]string, len(ai))
	for i, s := range ai {
		a[i] = String(s)
	}
	return a
}
func ConvStringsRaw(raw any) []string {
	if raw == nil {
		return nil
	}
	ai, ok := raw.([]any)
	if !ok {
		return nil
	}
	if len(ai) == 0 {
		return nil
	}
	a := make([]string, len(ai))
	for i, s := range ai {
		a[i] = String(s)
	}
	return a
}
func ConvStringMap(mi map[string]any) map[string]string {
	if len(mi) == 0 {
		return nil
	}
	maps := make(map[string]string, len(mi))
	for k, m := range mi {
		maps[k] = String(m)
	}
	return maps
}

func ConvStringsMap(mi map[string]any) map[string][]string {
	if len(mi) == 0 {
		return nil
	}
	maps := make(map[string][]string, len(mi))
	for k, m := range mi {
		maps[k] = ConvStringsRaw(m)
	}
	return maps
}
func ConvComplexStringMap(mi map[string]any) map[string]map[string]string {
	if len(mi) == 0 {
		return nil
	}
	maps := make(map[string]map[string]string, len(mi))
	for k, m := range mi {
		if d, ok := m.(map[string]any); ok {
			maps[k] = ConvStringMap(d)
		}
	}
	return maps
}
func ConvComplexStringsMap(mi map[string]any) map[string][][]string {
	if len(mi) == 0 {
		return nil
	}
	maps := make(map[string][][]string, len(mi))
	for k, m := range mi {
		if d, ok := m.([]any); ok {
			u := make([][]string, len(d))
			for i, v1 := range d {
				u[i] = ConvStringsRaw(v1)
			}
			maps[k] = u
		}
	}
	return maps
}
func ConvStringMaps(ai []any) []map[string]string {
	if len(ai) == 0 {
		return nil
	}
	maps := make([]map[string]string, 0, len(ai))
	for _, m := range ai {
		if d, ok := m.(map[string]any); ok {
			maps = append(maps, ConvStringMap(d))
		}
	}
	if len(maps) == 0 {
		return nil
	}
	return maps
}

func ConvComplexMaps(ai []any) []map[string]any {
	if len(ai) == 0 {
		return nil
	}
	maps := make([]map[string]any, 0, len(ai))
	for _, m := range ai {
		if d, ok := m.(map[string]any); ok {
			maps = append(maps, d)
		}
	}
	if len(maps) == 0 {
		return nil
	}
	return maps
}

// []any -> []map[string]any -> []map[string][]any
// -> []map[string][]map[string]any -> []map[string][]map[string]string
//func ConvComplexStringMaps(ai []any) []map[string][]map[string]string {
//	if len(ai) == 0 {
//		return nil
//	}
//	a := ConvComplexMaps(ai)
//	if a == nil {
//		return nil
//	}
//
//	maps := make([]map[string][]map[string]string, 0, len(ai))
//	for _, m := range ai {
//		if d, ok := m.(map[string]any); ok {
//			maps = append(maps, d)
//		}
//	}
//	if len(maps) == 0 {
//		return nil
//	}
//	return maps
//
//}
