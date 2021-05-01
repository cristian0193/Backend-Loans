package utils

import (
	"math"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const otpChars = "1234567890"

func StatusText(code int) string {
	return http.StatusText(code)
}

func GenerateDateExpirationCode(tiempo int) int64 {
	return time.Now().Add(time.Minute * time.Duration(tiempo)).Unix()
}

func IsInt(number string) bool {
	if number != "" {
		_, err := strconv.Atoi(number)
		if err != nil {
			return false
		}
		return true
	}
	return false
}

func GetInt(number string) int {
	if number != "" {
		value, err := strconv.Atoi(number)
		if err != nil {
			return 0
		}
		return value
	}
	return 0
}

func ValidateFieldEmpty(field string) string {
	if field != "" {
		return field
	}
	return "NULL"
}

func StructToMap(f interface{}) map[string]string {

	mapStruct := make(map[string]string)
	valueStruct := reflect.ValueOf(f).Elem()

	for i := 0; i < valueStruct.NumField(); i++ {
		valueField := valueStruct.Field(i)
		typeField := valueStruct.Type().Field(i)

		valueInterface := valueField.Interface()
		value := reflect.ValueOf(valueInterface)
		mapStruct[typeField.Name] = value.String()
	}
	return mapStruct
}

func GetDateToString(date string) string {
	if date != "" {
		result := strings.Split(date, "T")
		return result[0]
	}
	return ""
}

func GetDateWithTimeToString(date string) string {
	if date != "" {
		result := strings.Split(date, "T")
		return result[0] + " " + result[1]
	}
	return ""
}

func ABS(number int) int {
	return int(math.Abs(float64(number)))
}
