package requester

import (
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"prometheus-metrics-exporter/internal/pmeerrors/request"
	. "prometheus-metrics-exporter/internal/types"
	"strings"
	"time"
)

func GetContent(url string, basicAuth *BasicAuth, mimeType string, timeoutInSecs int) ([]byte, string, error) {

	client := http.Client{}
	client.Timeout = time.Second * time.Duration(timeoutInSecs)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, "", err
	}

	if basicAuth != nil {
		req.SetBasicAuth(basicAuth.Username, basicAuth.Password)
	}

	response, err := client.Do(req)

	if err != nil {
		errString := fmt.Sprintf("Requester: Client error with \"%s\"", err.Error())
		return nil, "", request.ErrorRequestClient{Err: errString}
	}

	responseStatus := response.StatusCode

	if responseStatus == http.StatusUnauthorized {
		errString := fmt.Sprintf(
			"Requester: Response status code was \"%d\" on \"%s\"",
			responseStatus, url)
		return nil, "", request.ErrorRequestResponseStatus401{Err: errString}
	} else if responseStatus == http.StatusForbidden {
		errString := fmt.Sprintf(
			"Requester: Response status code was \"%d\" on \"%s\"",
			responseStatus, url)
		return nil, "", request.ErrorRequestResponseStatus403{Err: errString}
	} else if responseStatus == http.StatusNotFound {

		errString := fmt.Sprintf(
			"Requester: Response status code was \"%d\" on \"%s\"",
			responseStatus, url)
		return nil, "", request.ErrorRequestResponseStatus404{Err: errString}

	} else if responseStatus == http.StatusInternalServerError {

		errString := fmt.Sprintf(
			"Requester: Response status code was \"%d\" on \"%s\"",
			responseStatus, url)
		return nil, "", request.ErrorRequestResponseStatus500{Err: errString}

	} else if responseStatus != http.StatusOK {

		errString := fmt.Sprintf(
			"Requester: Response status code was \"%d\" on \"%s\"",
			responseStatus, url)
		return nil, "", request.ErrorRequestResponseStatusNot200{Err: errString}

	}

	receivedContentType := response.Header.Get("Content-type")
	expectedContentType := mimeType

	actualContentType, _, err := mime.ParseMediaType(receivedContentType)

	if err != nil {
		return nil, "", request.ErrorRequestContentTypeParse{Err: err.Error()}
	}

	validContentType := func(contentType string, expectedContentType string) bool {
		strArr := strings.Split(contentType, "/")
		return len(strArr) == 2 && strArr[1] == expectedContentType
	}(actualContentType, expectedContentType)

	if !validContentType {
		errString := fmt.Sprintf(
			"Requester: Not a valid content type. Expected \"%s\" but got \"%s\"",
			expectedContentType, receivedContentType)
		return nil, "", request.ErrorRequestInvalidContentTypeFound{Err: errString}
	}

	body, err := ioutil.ReadAll(response.Body)

	defer func() {
		if _err := response.Body.Close(); _err != nil {
			err = _err
		}
	}()

	if err != nil {
		errStr := fmt.Sprintf("Requester: Unable to read from body \"%s\"", err)
		return nil, "", request.ErrorRequestUnableToReadBody{Err: errStr}
	}

	return body, mimeType, nil
}
