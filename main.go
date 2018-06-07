package main

import "errors"
import "strconv"
import "strings"
import "github.com/gin-gonic/gin"

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/numbers", func(c *gin.Context) {

		begin := c.DefaultQuery("begin", "1")
		end := c.DefaultQuery("end", "100")

		message, err := processNumbers(begin, end)

		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"numbers": message,
			})
		}
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}


func processNumbers(beginStr string, endStr string) (string, error) {

		message := ""

		begin, beginErr := strconv.ParseInt(beginStr, 10, 32)
		if beginErr != nil {
			return message, errors.New("invalid begin value")
		}

		end, endErr := strconv.ParseInt(endStr, 10, 32)
		if endErr != nil {
			return message, errors.New("invalid end value")
		}

		if begin < 1 || end < 1 {
			return message, errors.New("parameters must be unsigned and greater than 0")
		}

		if begin > end {
			return message, errors.New("begin must be lower than end")
		}

		if begin > 100 || end > 100 {
			return message, errors.New("parameters must be lower than 100")
		}

		for i := begin; i <= end; i++ {
			tmpStr := ""
			if i % 3 == 0 {
				tmpStr += "Pé"
			}
			if i % 5 == 0 {
				tmpStr += "Do"
			}
			if tmpStr == "" {
				tmpStr = strconv.FormatInt(i, 10)
			}
			message += tmpStr + " "
		}

		return strings.TrimSpace(message), nil
}
