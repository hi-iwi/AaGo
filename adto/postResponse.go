package adto

func PostResp(id interface{}) map[string]interface{} {
	return map[string]interface{}{"id": id}
}

func CountResp(cnt interface{}) map[string]interface{} {
	return map[string]interface{}{"total": cnt}
}
