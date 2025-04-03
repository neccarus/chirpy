package handler

import (
	"strings"
)

func SwearReplacement(body string) string {
	replaceSwear := "****"
	swears := map[string]string{
		"kerfuffle": replaceSwear,
		"sharbert":  replaceSwear,
		"fornax":    replaceSwear,
	}
	words := strings.Fields(body)
	var newString string
	for _, word := range words {
		for swear := range swears {
			if strings.ToLower(word) == swear {
				word = swears[swear]
			}
		}
		newString = newString + word + " "
	}
	newString = strings.Trim(newString, " ")
	return newString
}

//func ValidateChirp(writer http.ResponseWriter, request *http.Request) {
//	var errorString string
//
//	type parameters struct {
//		Body string `json:"body"`
//	}
//
//	type returnVals struct {
//		Error       string `json:"error"`
//		Valid       bool   `json:"valid"`
//		CleanedBody string `json:"cleaned_body"`
//	}
//
//	//params := parameters{}
//	//status, err := crud.DecodeRequest(request, &params, http.StatusInternalServerError)
//	decoder := json.NewDecoder(request.Body)
//	params := parameters{}
//	err := decoder.Decode(&params)
//	if err != nil {
//		log.Printf("Error decoding parameters: %s", err)
//		errorString = "Something went wrong"
//		writer.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	valid := false
//	if len(params.Body) <= 140 {
//		valid = true
//	}
//	if !valid {
//		errorString = "Chirp is too long"
//		writer.WriteHeader(http.StatusBadRequest)
//	}
//	cleanedBody := swearReplacement(params.Body)
//	fmt.Println(cleanedBody)
//	responseBody := returnVals{
//		Valid:       valid,
//		Error:       errorString,
//		CleanedBody: cleanedBody,
//	}
//
//	data, err := json.Marshal(responseBody)
//	if err != nil {
//		log.Printf("Error marshalling JSON: %s", err)
//		writer.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	writer.Header().Add("Content-Type", "application/json")
//	writer.WriteHeader(http.StatusOK)
//	writer.Write(data)
//}
