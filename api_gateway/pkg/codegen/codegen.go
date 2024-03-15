package codegen

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
)

func GenerateCode() int64 {
	code := rand.Intn(1000000)

	temp := fmt.Sprintf("%06d", code)
	intCode, err := strconv.ParseInt(temp, 10, 64)
	if err != nil {
		log.Println("error while generating code in codegen")
		return 0
	}
	return intCode
}
