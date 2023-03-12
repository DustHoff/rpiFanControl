package fancontrol

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var _ http.Handler = &Api{}

type Api struct {
}

func (a Api) NewApi() Api {
	return Api{}
}

func (a Api) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		buf := new(strings.Builder)
		_, err := io.Copy(buf, request.Body)
		if err != nil {
			log.Fatalln(err)
		}
		s, err := strconv.Atoi(buf.String())
		if err != nil {
			log.Println(err)
		}
		SetSpeed(s)
	}

	fmt.Fprintf(writer, "ok")
}
