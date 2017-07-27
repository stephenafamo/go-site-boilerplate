package config

type Config struct {
	SValue string
	IValue int
	FValue float64
	BValue bool
	XValue interface{}
}
var Settings = make(map[string]Config, 100)

func init() {
	setup()	
}

func GetS(name string) string {
	return Settings[name].SValue
}

func GetI(name string) int {
	return Settings[name].IValue
}

func GetF(name string) float64 {
	return Settings[name].FValue
}

func GetB(name string) bool {
	return Settings[name].BValue
}

func GetX(name string) interface{} {
	return Settings[name].XValue
}