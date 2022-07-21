package ginie

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

func trace(msg string) string {
	var pcs [32]uintptr
	// skip first 3 caller
	n := runtime.Callers(3, pcs[:])

	var sb strings.Builder
	sb.WriteString(msg + "\nTraceback:")
	for _, pc := range pcs[:n] {
		// get running function
		fn := runtime.FuncForPC(pc)
		// get the file name and line number of calling the function
		file, line := fn.FileLine(pc)
		sb.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return sb.String()
}

func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				msg := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(msg))
				c.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()
		c.Next()
	}
}
