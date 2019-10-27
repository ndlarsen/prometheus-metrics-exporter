package server

import (
	"net/http"
	. "prometheus-metrics-exporter/e2etest/simpleTestServer/server/middleware"
	. "prometheus-metrics-exporter/e2etest/simpleTestServer/server/writers"
)

func Server(){

	jsonHandler := http.HandlerFunc(JsonWriter)
	htmlHandler := http.HandlerFunc(HtmlWriter)
	http.Handle("/jsonWithoutBasicAuth", MethodValidatorMiddleware(jsonHandler))
	http.Handle("/jsonWithBasicAuth", MethodValidatorMiddleware(BasicAuthMiddleware(jsonHandler)))
	http.Handle("/htmlWithoutBasicAuth", MethodValidatorMiddleware(htmlHandler))
	http.Handle("/htmlWithBasicAuth", MethodValidatorMiddleware(BasicAuthMiddleware(htmlHandler)))

	http.ListenAndServe(":8080", nil)
}
