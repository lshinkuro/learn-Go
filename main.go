package main

import (
	"fmt"

	. "example.com/latihanGo/library"
)

//maksud dari titik adalah menjadikanya sejajar sehingga kita tidak perlu menuliskan nama librarynya lagi

func main() {
	// fmt.Println("coba import")
	// library.SayHello("ahmad")
	var s1 = Student{"ethan", 21}
	fmt.Println("name ", s1.Name)

}

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func sqrt(x float64) string {
// 	if x < 0 {
// 		return sqrt(-x) + "i"
// 	}
// 	return fmt.Sprint(math.Sqrt(x))
// }

// func main() {
// 	fmt.Println(sqrt(5), sqrt(-9))
// }

// package main

// import "github.com/gofiber/fiber"

// func main() {
// 	app := fiber.New()

// 	app.Get("/", func(c *fiber.Ctx) {
// 		c.Send("hallo semuanya!")
// 	})

// 	app.Listen(3000)
// }
