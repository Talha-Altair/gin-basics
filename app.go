package main

import (

	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/bxcodec/faker/v3"
)

type faker_struct struct {
	Name     string `faker:"name"`
	Time     string `faker:"time"`
	Sentence string `faker:"sentence"`
	Currency string `faker:"currency"`
}

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {

		a := faker_struct{}

		err := faker.FakeData(&a)

		if err != nil {
			fmt.Println(err)
		}

		c.JSON(200, gin.H{
			"data": "a",
		})
	})

	main2()

	r.Run()
}
