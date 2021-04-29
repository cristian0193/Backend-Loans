package utils

import (
	"crypto/rand"
	"math"
	"net/http"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"gopkg.in/jeevatkm/go-model.v1"
)

const otpChars = "1234567890"

func StatusText(code int) string {
	return http.StatusText(code)
}

func GenerateCodeVerification(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
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

func RemoveHtmlTag(in string) string {
	const pattern = `(<\/?[a-zA-A]+?[^>]*\/?>)*`
	r := regexp.MustCompile(pattern)
	groups := r.FindAllString(in, -1)
	sort.Slice(groups, func(i, j int) bool {
		return len(groups[i]) > len(groups[j])
	})
	for _, group := range groups {
		if strings.TrimSpace(group) != "" {
			in = strings.ReplaceAll(in, group, "")
		}
	}
	return in
}

func FilterQueryParameters(structDefination interface{}) string {

	valueStruct := reflect.ValueOf(structDefination).Elem()

	var queryParams string

	for index := 0; index < valueStruct.NumField(); index++ {
		valueField := valueStruct.Field(index)
		typeField := valueStruct.Type().Field(index)

		valueInterface := valueField.Interface()
		value := reflect.ValueOf(valueInterface)

		if value.String() != "" {
			tag, _ := model.Tag(structDefination, typeField.Name)
			queryParams += tag.Get("json") + " = '" + value.String() + "' AND "
		}
	}

	if queryParams == "" {
		return ""
	}
	return queryParams[:len(queryParams)-4]
}
