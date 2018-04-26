package yagnats

type Logger interface {
	Fatal(string)
	Error(string)
	Warn(string)
	Info(string)
	Debug(string)

	Fatald(map[string]interface{}, string)
	Errord(map[string]interface{}, string)
	Warnd(map[string]interface{}, string)
	Infod(map[string]interface{}, string)
	Debugd(map[string]interface{}, string)
}

type DefaultLogger struct{}

func (dl *DefaultLogger) Fatal(string) {}
func (dl *DefaultLogger) Error(string) {}
func (dl *DefaultLogger) Warn(string)  {}
func (dl *DefaultLogger) Info(string)  {}
func (dl *DefaultLogger) Debug(string) {}

func (dl *DefaultLogger) Fatald(map[string]interface{}, string) {}
func (dl *DefaultLogger) Errord(map[string]interface{}, string) {}
func (dl *DefaultLogger) Warnd(map[string]interface{}, string)  {}
func (dl *DefaultLogger) Infod(map[string]interface{}, string)  {}
func (dl *DefaultLogger) Debugd(map[string]interface{}, string) {}

// import "fmt"

// type Logger interface {
// 	Fatal(string)
// 	Error(string)
// 	Warn(string)
// 	Info(string)
// 	Debug(string)

// 	Fatald(map[string]interface{}, string)
// 	Errord(map[string]interface{}, string)
// 	Warnd(map[string]interface{}, string)
// 	Infod(map[string]interface{}, string)
// 	Debugd(map[string]interface{}, string)
// }

// type DefaultLogger struct{}

// func (dl *DefaultLogger) Fatal(s string) { fmt.Println(s) }
// func (dl *DefaultLogger) Error(s string) { fmt.Println(s) }
// func (dl *DefaultLogger) Warn(s string)  { fmt.Println(s) }
// func (dl *DefaultLogger) Info(s string)  { fmt.Println(s) }
// func (dl *DefaultLogger) Debug(s string) { fmt.Println(s) }

// func (dl *DefaultLogger) Fatald(m map[string]interface{}, s string) { fmt.Println(s) }
// func (dl *DefaultLogger) Errord(m map[string]interface{}, s string) { fmt.Println(s) }
// func (dl *DefaultLogger) Warnd(m map[string]interface{}, s string)  { fmt.Println(s) }
// func (dl *DefaultLogger) Infod(m map[string]interface{}, s string)  { fmt.Println(s) }
// func (dl *DefaultLogger) Debugd(m map[string]interface{}, s string) { fmt.Println(s) }
