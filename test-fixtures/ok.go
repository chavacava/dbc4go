package mains

// func main is the top
//@requires nonil
//@requires a > 0
//@ensures a==a
func mains(a, b string, c bool, d int) {
	println("This is the end.")
}

// NewCar returns a Car struct
//@requires wheels > 2
//@requires wheelsDrive <= wheels && wheelsDrive >= 2 && wheelsDrive%2 == 0
//@requires maxSpeedKmh > 0 && maxSpeedKmh <= maxAuthorizedSpeed
//@requires manufacturer != ""
func NewCar(wheels int, wheelsDrive int, maxSpeedKmh int, manufacturer string) Car { ... }

