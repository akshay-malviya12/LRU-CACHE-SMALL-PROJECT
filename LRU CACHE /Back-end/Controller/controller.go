package controllers

import (
	"encoding/json"
	"io/ioutil" // Added import for ioutil package
	"log"
	"main/model"
	"net/http"
	"strconv"
	"time"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Payload Payload
}

type Payload struct {
	Key        string
	Value      string
	ExpireTime string
}

var setGetCacheValue *models.CacheLRU

// Set LRU Cache Capacity by default set 5.
func SetCapacity(w http.ResponseWriter, r *http.Request) {
	// Get the capacity from the request URL.
	capacityInt := StringToInt(r.URL.Query().Get("capacity"))
	log.Println("capacityInt", capacityInt)
	if (capacityInt <= 0){
	//set default capacity.
           capacityInt = 5
	}
	
	// calling for set LRU cache function with capcity of key value pair(int).
	setGetCacheValue = models.NewCacheLRU(capacityInt)
}

// string value to int value convert.
func StringToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return num
}

// LRU Cache Capacity Iniatialization check.
func LRUCacheInitiallization(setGetCacheValue *models.CacheLRU) (bool, string) {
	if setGetCacheValue == nil {
		msg := "Error: LRU Cache capacity is not initialized."
		return true, msg
	}
	msg := "Success:LRU Cache capacity is initialized."
	return false, msg
}

// GetCache function is used to get the value from the cache.
func GetCache(w http.ResponseWriter, r *http.Request) {

	var cache models.CacheModel
	var resp Response
	// Get the key from the request URL.
	cache.Key = r.URL.Query().Get("key")

	// cacheGet := models.NewCacheLRU(10)
	//Check capacity of LRU cache provided or not.
	boolType, Msg := LRUCacheInitiallization(setGetCacheValue)
	if boolType {
		ResponseGenerateJson(w, r, "Failure", Msg, resp.Payload)
		return
	}
	// Get the value from the cache.
	value, boolType := setGetCacheValue.Get(cache.Key)
	if boolType {
		resp.Payload.Value = value.Value
		resp.Payload.ExpireTime = value.CacheExpireTime.String()
		resp.Payload.Key = value.Key
		ResponseGenerateJson(w, r, "Success", "", resp.Payload)
		return
	}
	msg := "Key is not avaible,Please enter anoter key."
	ResponseGenerateJson(w, r, "Failure", msg, resp.Payload)
}

// SetCache function is used to set the value in the cache.
func SetCache(w http.ResponseWriter, r *http.Request) {

	var resp Response
	var cache models.CacheModel

	//Check capacity of LRU cache provided or not.
	boolType, Msg := LRUCacheInitiallization(setGetCacheValue)
	if boolType {
		ResponseGenerateJson(w, r, "Failure", Msg, resp.Payload)
		return
	}
	// Get the key from the request Body.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		ResponseGenerateJson(w, r, "Failure", err.Error(), resp.Payload)
		return

	}
	// Unmarshal the JSON data from []byte into cache.
	err = json.Unmarshal(body, &cache)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		ResponseGenerateJson(w, r, "Failure", err.Error(), resp.Payload)
		return
	}
	setTime := time.Now().Add(450 * time.Second)
	// boolValue, msg := setGetCacheValue.Set(cache.Key, cache.Value, cache.CacheExpireTime)

	boolValue, msg := setGetCacheValue.Set(cache.Key, cache.Value, setTime)

	if boolValue {
		resp.Payload.Key = cache.Key
		resp.Payload.Value = cache.Value
		resp.Payload.ExpireTime = cache.CacheExpireTime.Format(time.RFC3339)
		//setGetCacheValue.GetAllKeys()
		ResponseGenerateJson(w, r, "Success", msg, resp.Payload)
		return
	}
	ResponseGenerateJson(w, r, "Failure", msg, resp.Payload)
}

// responseGenerateJson function is used to generate the response in JSON format.
func ResponseGenerateJson(w http.ResponseWriter, r *http.Request, status string, message string, payload Payload) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res, _ := json.Marshal(Response{Status: "Success", Message: message, Payload: payload})
	w.Write(res)
}

// DelelteCache function is used to delete the cache.
func DeleteCache(w http.ResponseWriter, r *http.Request) {

	var resp Response
	//Check capacity of LRU cache provided or not.
	boolType, Msg := LRUCacheInitiallization(setGetCacheValue)
	if boolType {
		ResponseGenerateJson(w, r, "Failure", Msg, resp.Payload)
		return
	}
	// Get the key from the request URL.
	boolType = setGetCacheValue.CacheDeleteByUserKeyDirectly(r.URL.Query().Get("key"))
	if boolType {
		ResponseGenerateJson(w, r, "Success", "LRU cache delete Success.", resp.Payload)
		return
	}
	ResponseGenerateJson(w, r, "Failure", "LRU cache delete  failure.", resp.Payload)
}

