package atype

func ConvStrings(ai []interface{}) []string {
	if len(ai) == 0 {
		return nil
	}
	a := make([]string, len(ai))
	for i, s := range ai {
		a[i] = String(s)
	}
	return a
}
func ConvStringsRaw(raw interface{}) []string {
	if raw == nil {
		return nil
	}
	ai, ok := raw.([]interface{})
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
func ConvStringMap(mi map[string]interface{}) map[string]string {
	if len(mi) == 0 {
		return nil
	}
	maps := make(map[string]string, len(mi))
	for k, m := range mi {
		maps[k] = String(m)
	}
	return maps
}

func ConvStringsMap(mi map[string]interface{}) map[string][]string {
	if len(mi) == 0 {
		return nil
	}
	maps := make(map[string][]string, len(mi))
	for k, m := range mi {
		maps[k] = ConvStringsRaw(m)
	}
	return maps
}
func ConvComplexStringMap(mi map[string]interface{}) map[string]map[string]string {
	if len(mi) == 0 {
		return nil
	}
	maps := make(map[string]map[string]string, len(mi))
	for k, m := range mi {
		if d, ok := m.(map[string]interface{}); ok {
			maps[k] = ConvStringMap(d)
		}
	}
	return maps
}
func ConvComplexStringsMap(mi map[string]interface{}) map[string][][]string {
	if len(mi) == 0 {
		return nil
	}
	maps := make(map[string][][]string, len(mi))
	for k, m := range mi {
		if d, ok := m.([]interface{}); ok {
			u := make([][]string, len(d))
			for i, v1 := range d {
				u[i] = ConvStringsRaw(v1)
			}
			maps[k] = u
		}
	}
	return maps
}
func ConvStringMaps(ai []interface{}) []map[string]string {
	if len(ai) == 0 {
		return nil
	}
	maps := make([]map[string]string, 0, len(ai))
	for _, m := range ai {
		if d, ok := m.(map[string]interface{}); ok {
			maps = append(maps, ConvStringMap(d))
		}
	}
	if len(maps) == 0 {
		return nil
	}
	return maps
}

func ConvComplexMaps(ai []interface{}) []map[string]interface{} {
	if len(ai) == 0 {
		return nil
	}
	maps := make([]map[string]interface{}, 0, len(ai))
	for _, m := range ai {
		if d, ok := m.(map[string]interface{}); ok {
			maps = append(maps, d)
		}
	}
	if len(maps) == 0 {
		return nil
	}
	return maps
}

// []interface{} -> []map[string]interface{} -> []map[string][]interface{}
// -> []map[string][]map[string]interface{} -> []map[string][]map[string]string
//func ConvComplexStringMaps(ai []interface{}) []map[string][]map[string]string {
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
//		if d, ok := m.(map[string]interface{}); ok {
//			maps = append(maps, d)
//		}
//	}
//	if len(maps) == 0 {
//		return nil
//	}
//	return maps
//
//}
