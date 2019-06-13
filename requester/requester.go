package requester

import (
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	. "prometheus-metrics-exporter/pmeerrors"
	"strings"
	"time"
)

func GetContent(url string, mimeType string, timeoutInSecs int) ([]byte, string, error) {

	var httpClient http.Client

	httpClient.Timeout = time.Second * time.Duration(timeoutInSecs)
	response, err := httpClient.Get(url)

	if err != nil {
		errString := fmt.Sprintf("Requester: Timeout at \"%s\"", url)
		return nil, "", ErrorRequestTimeOut{Err: errString}
	}

	responseStatus := response.StatusCode

	if responseStatus == http.StatusNotFound {

		errString := fmt.Sprintf(
			"Requester: Response status code was \"%d\" on \"%s\"",
			responseStatus, url)
		return nil, "", ErrorRequestResponseStatus404{Err: errString}

	} else if responseStatus == http.StatusInternalServerError {

		errString := fmt.Sprintf(
			"Requester: Response status code was \"%d\" on \"%s\"",
			responseStatus, url)
		return nil, "", ErrorRequestResponseStatus500{Err: errString}

	} else if responseStatus != http.StatusOK {

		errString := fmt.Sprintf(
			"Requester: Response status code was \"%d\" on \"%s\"",
			responseStatus, url)
		return nil, "", ErrorRequestResponseStatusNot200{Err: errString}

	}

	receivedContentType := response.Header.Get("Content-type")
	expectedContentType := mimeType

	actualContentType, _, err := mime.ParseMediaType(receivedContentType)

	if err != nil {
		return nil, "", ErrorRequestContentTypeParse{Err: err.Error()}
	}

	validContentType := func(contentType string, expectedContentType string) bool {
		strArr := strings.Split(contentType, "/")
		if len(strArr) == 2 && strArr[1] == expectedContentType {
			return true
		}
		return false
	}(actualContentType, expectedContentType)

	if !validContentType {
		errString := fmt.Sprintf(
			"Requester: Not a valid content type. Expected \"%s\" but got \"%s\"",
			expectedContentType, receivedContentType)
		return nil, "", ErrorRequestInvalidContentTypeFound{Err: errString}
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		errStr := fmt.Sprintf("Requester: Unable to read from body \"%s\"", err)
		return nil, "", ErrorRequestUnableToReadBody{Err: errStr}
	}

	return body, mimeType, nil
}
