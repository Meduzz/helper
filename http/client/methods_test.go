package client

import (
	"crypto/sha256"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/Meduzz/helper/hmac"
	"github.com/gin-gonic/gin"
)

type TestDTO struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestMain(m *testing.M) {
	server := gin.Default()

	server.GET("/hello/:name/:age", func(ctx *gin.Context) {
		name := ctx.Param("name")
		age := ctx.Param("age")

		iAge, err := strconv.Atoi(age)

		if err != nil {
			ctx.AbortWithError(500, err)
		}

		ctx.JSON(200, gin.H{"name": name, "age": iAge})
	})
	server.POST("/post/echo", func(ctx *gin.Context) {
		bs, err := ioutil.ReadAll(ctx.Request.Body)

		if err != nil {
			ctx.AbortWithError(500, err)
		}

		sign := ctx.GetHeader("x-sign")

		if sign != "" {
			log.Println("The request was signed.")
			if !hmac.Verify([]byte("key"), bs, []byte(sign), sha256.New) {
				log.Println("The signature was not valid.")
				ctx.Status(403)
				return
			}

			log.Println("The signature was valid.")
		}

		ctx.Data(200, "application/json", bs)
	})
	server.PUT("/put/echo", func(ctx *gin.Context) {
		bs, err := ioutil.ReadAll(ctx.Request.Body)

		if err != nil {
			ctx.AbortWithError(500, err)
		}

		contentType := ctx.GetHeader("Content-Type")

		ctx.Data(200, contentType, bs)
	})
	server.DELETE("/body", func(ctx *gin.Context) {
		bs, _ := ctx.GetRawData()

		if len(bs) == 0 {
			ctx.Status(200)
			return
		} else {
			ctx.Status(201)
			return
		}
	})

	go server.Run(":6007")

	os.Exit(m.Run())
}
func TestGet(t *testing.T) {
	req, err := GET("http://localhost:6007/hello/bosse/19")

	if err != nil {
		t.Error(err)
	}

	res, err := req.Do(http.DefaultClient)

	if err != nil {
		t.Error(err)
	}

	if res.Code() != 200 {
		t.Fail()
	}

	test := &TestDTO{}
	err = res.AsJson(test)

	if err != nil {
		t.Error(err)
	}

	if test.Name != "bosse" {
		t.Fail()
	}

	if test.Age != 19 {
		t.Fail()
	}
}

func TestPost(t *testing.T) {
	body := &TestDTO{"Sven", 64}
	req, err := POST("http://localhost:6007/post/echo", body)

	if err != nil {
		t.Error(err)
	}

	req.Sign(func(req *HttpRequest) error {
		val := hmac.Sign([]byte("key"), req.Body(), sha256.New)

		req.Header("x-sign", string(val))

		return nil
	})

	res, err := req.Do(http.DefaultClient)

	if err != nil {
		t.Error(err)
	}

	if res.Code() != 200 {
		t.Fail()
	}

	subject := &TestDTO{}

	err = res.AsJson(subject)

	if err != nil {
		t.Error(err)
	}

	if subject.Name != body.Name {
		t.Fail()
	}

	if subject.Age != body.Age {
		t.Fail()
	}
}

func TestPut(t *testing.T) {
	req, err := PUTText("http://localhost:6007/put/echo", "Hello there!", "text/plain")

	if err != nil {
		t.Error(err)
	}

	res, err := req.Do(http.DefaultClient)

	if err != nil {
		t.Error(err)
	}

	if res.Code() != 200 {
		t.Fail()
	}

	text, err := res.AsText()

	if err != nil {
		t.Error(err)
	}

	if text != "Hello there!" {
		t.Fail()
	}

	if res.Header("Content-Type") != "text/plain" {
		t.Fail()
	}
}

func TestDeleteBodyLess(t *testing.T) {
	req, err := DELETE("http://localhost:6007/body", nil)

	if err != nil {
		t.Error(err)
	}
	res, err := req.Do(http.DefaultClient)

	if err != nil {
		t.Error(err)
	}

	if res.Code() != 200 {
		t.Fail()
	}
}

func TestDeleteWithBody(t *testing.T) {
	req, err := DELETE("http://localhost:6007/body", "hello")

	if err != nil {
		t.Error(err)
	}
	res, err := req.Do(http.DefaultClient)

	if err != nil {
		t.Error(err)
	}

	if res.Code() != 201 {
		t.Fail()
	}
}
