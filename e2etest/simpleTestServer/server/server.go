package server

import (
	"fmt"
	"net/http"
	. "simpleTestServer/server/middleware"
	. "simpleTestServer/server/writers"
)

func Server(port *string) {

	jsonHandler := http.HandlerFunc(JsonWriter)
	htmlHandler := http.HandlerFunc(HtmlWriter)
	http.Handle("/jsonWithoutBasicAuth", MethodValidatorMiddleware(jsonHandler))
	http.Handle("/jsonWithBasicAuth", MethodValidatorMiddleware(BasicAuthMiddleware(jsonHandler)))
	http.Handle("/htmlWithoutBasicAuth", MethodValidatorMiddleware(htmlHandler))
	http.Handle("/htmlWithBasicAuth", MethodValidatorMiddleware(BasicAuthMiddleware(htmlHandler)))

	fPort := fmt.Sprintf(":%s", *port)

	http.ListenAndServe(fPort, nil)
}
