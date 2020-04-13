package main

func notValidRequestError(method string) string {
	return `{"error": "true","message":"` + method + ` yöntemi bu hizmette geçersizdir."}`
}
func notFindRecordError() string {
	return `{"error": "true","message":"Kayıt bulunamamıştır."}`
}
func requiredInputError(inputName string) string {
	return `{"error": "true","message":"` + inputName + ` alanı doldurulmalıdır."}`
}
func succesfullyRecordedError() string {
	return `{"error": "false","message":"Başarılı şekilde kaydedilmiştir."}`
}
func failedRecordError() string {
	return `{"error": "true","message":"Kayıt başarısız olmuştur."}`
}
func invalidPermission() string {
	return `{"error": "true","message":"Bu servisi kullanma yetkisine sahip değilsiniz."}`
}
func incorrectInput(inputName string) string {
	return `{"error": "true","message":"` + inputName + ` hatalı ."}`
}
func dataBaseSaveError() string {
	return `{"error": "true","message":"Veri tabanına kayıtta hata oluştu."}`
}
func creditCardError() string {
	return `{"error": "true","message":"Hatalı kredi kartı."}`
}
func objectIDError() string {
	return `{"error": "true","message":"Hatalı ID."}`
}
func someThingWentWrong() string {
	return `{"error": "true","message":"Bir şeyler yanlış gitti."}`
}
func sendMailError() string {
	return `{"error": "true","message":"Mail gönderimi başarısız."}`
}
func invalidLoginRequest() string {
	return `{"error": "true","message":"Onaylanmamış kullanıcı"}`
}
func alreadyDefinedError(inputName string) string {
	return `{"error": "true","message":"` + inputName + ` zaten tanımlı ."}`
}
