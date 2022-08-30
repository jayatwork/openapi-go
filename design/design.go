package design

import (
	"fmt"

	. "goa.design/goa/v3/dsl"
)

var _ = API("users", func() {
	Title("users API")
	Version("0.0.1")
	Server("http", func() {
		Host("development", func() {
			URI("http://localhost:8085")
		})
	})
})

var User = Type("User", func() {
	Attribute("name", String, func() {
		MinLength(5)
	})

	Attribute("birth_year", Int32, func() {
		Minimum("1900")
		Maximum("2022")
	})

	Required("name", "birth_year")
})

var _ = Service("users", func() {
	Description("API for users")

	Method("create", func() {
		Payload(User)
		Result(User)
		HTTP(func() {
			POST("/users")
		})
	})

	Method("all", func() {
		Result(ArrayOf(User))
		HTTP(func() {
			GET("/users")
		})
	})

	Files("/openapi3.json", "./gen/http/openapi3.json")
})

func main() {
	fmt.Println("begin main block goa design startup")
}
