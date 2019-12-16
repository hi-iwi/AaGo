package adto

func AddOkResp(id interface{}) map[string]interface{} {
	return map[string]interface{}{"id": id}
}

func CountResp(cnt interface{}) map[string]interface{} {
	return map[string]interface{}{"total": cnt}
}
