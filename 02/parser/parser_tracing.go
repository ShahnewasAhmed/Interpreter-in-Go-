package parser

var traceLevel int

const traceIdentPlaceholder string = "\t"

func identLevel() string {
	return ""
}

func tracePrint(fs string) {
}

func incIdent() { traceLevel = traceLevel + 1 }
func decIdent() { traceLevel = traceLevel - 1 }

func trace(msg string) string {
	incIdent()
	return msg
}

func untrace(msg string) {
	tracePrint("END " + msg)
	decIdent()
}
