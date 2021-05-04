package utils

import (
	"Backend-Loans/domain/dto"
	"math"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const otpChars = "1234567890"

var responseDto = dto.Response{}

func StatusText(code int) string {
	return http.StatusText(code)
}

func GenerateDateExpirationCode(tiempo int) int64 {
	return time.Now().Add(time.Minute * time.Duration(tiempo)).Unix()
}

func ConvertInt32(number string) int32 {
	if number != "" {
		value, err := strconv.Atoi(number)
		if err != nil {
			return 0
		}
		return int32(value)
	}
	return 0
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

func ResponseError(code int, err error) dto.Response {
	if err != nil {
		responseDto.Status = code
		responseDto.Description = StatusText(code)
		responseDto.Message = err.Error()
		return responseDto
	}
	responseDto.Status = http.StatusOK
	return responseDto
}

func ResponseValidation(code int, headers dto.Headers, message string) dto.Response {
	responseDto.Status = code
	responseDto.Description = StatusText(code)
	responseDto.Message = Lenguage(headers.Lenguage, message)
	return responseDto
}

func TimeElapsed(a, b time.Time) (year, month, day int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)

	if day < 0 {
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}

func CalculateMonthsArrears(year, month int) int {
	if year > 0 || month > 0 {
		return month + (year * 12)
	}
	return 0
}

func ResultCalculateMonthsArrears(dateCreation time.Time, state int32) int {
	if state == 2 || state == 3 {
		var creationDate = dateCreation
		var dateNow = time.Now()

		var newCreationDate = time.Date(creationDate.Year(), creationDate.Month(), creationDate.Day(), 0, 0, 0, 0, time.UTC)
		var newDateNow = time.Date(dateNow.Year(), dateNow.Month(), dateNow.Day(), 0, 0, 0, 0, time.UTC)

		year, month, _ := TimeElapsed(newCreationDate, newDateNow)
		return CalculateMonthsArrears(year, month)
	}
	return 0
}
