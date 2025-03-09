// greeting/greeting.go
package greeting

func Greet(name string, isMorning bool) string {
	if isMorning {
		return "Good Morning, " + name
	}
	return "Hello, " + name
}
