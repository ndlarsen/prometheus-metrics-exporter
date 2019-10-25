package htmlparser

import (
	"bytes"
	"io"
	//"io/ioutil"
	"os"
	. "prometheus-metrics-exporter/pmeerrors"
	"testing"
)

func TestFetchValue_NoSuchElement(t *testing.T) {

	handle, openErr := os.Open("../test_related/lorem_ipsum.html")
	defer handle.Close()

	if openErr != nil {
		t.Fatalf("Unable to open file: %s", openErr.Error())
	}

	path := "html/body/a"

	_, parseErr := FetchValue(path, handle, "")

	if parseErr != nil && parseErr == parseErr.(ErrorHtmlParserNoSuchElement) {
		t.Log("Test failed as expected.")
	} else {
		t.Fatalf("Test failed unexpectedly: %s", parseErr)
	}

}

func TestFetchValue_MoreThanOneElement(t *testing.T) {

	handle, openErr := os.Open("../test_related/lorem_ipsum.html")
	defer handle.Close()

	if openErr != nil {
		t.Fatalf("Unable to open file: %s", openErr.Error())
	}

	path := "html/body/ul/li"

	_, parseErr := FetchValue(path, handle, "")

	if parseErr != nil && parseErr == parseErr.(ErrorHtmlParserTooManyElements) {
		t.Log("Test failed as expected.")
	} else {
		t.Fatalf("Test failed unexpectedly: %s", parseErr)
	}

}

type fakeReader struct {
}

func (f fakeReader) Read(p []byte) (n int, err error) {
	return 1, io.ErrUnexpectedEOF
}

func Test_FetchValue_ReadError(t *testing.T) {

	path := "html/body/ul/li"

	_, parseErr := FetchValue(path, fakeReader{}, "")

	if parseErr != nil && parseErr == parseErr.(ErrorHtmlParserParsing) {
		t.Log("Test failed as expected.")
	} else {
		t.Fatalf("Test failed unexpectedly: %s", parseErr)
	}

}

func Test_extract_OK(t *testing.T) {

	const html = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>lorem ipsum</title>
</head>
<body>
<ul>
    <li>Lorem ipsum dolor sit amet, consectetuer adipiscing elit.</li>
    <li>Aliquam tincidunt mauris eu risus.</li>
    <li>Vestibulum auctor dapibus neque.</li>
</ul>
</body>
</html>`

	path := "html/body/ul/li[1]"
	expectedResult := "Lorem ipsum dolor sit amet, consectetuer adipiscing elit."

	result, parseErr := extract(path, bytes.NewBuffer([]byte(html)))

	if parseErr == nil && result == expectedResult {
		t.Log("Test succeeded as expected.")
	} else {
		t.Fatalf("Test failed unexpectedly: %s", parseErr)
	}

}

func Test_FetchValue_ConversionError(t *testing.T) {

	const html = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>lorem ipsum</title>
</head>
<body>
<ul>
    <li>Lorem ipsum dolor sit amet, consectetuer adipiscing elit.</li>
    <li>Aliquam tincidunt mauris eu risus.</li>
    <li>2.0</li>
</ul>
</body>
</html>`

	path := "html/body/ul/li[2]"

	_, parseErr := FetchValue(path, bytes.NewBuffer([]byte(html)), "")

	if parseErr != nil && parseErr == parseErr.(ErrorHtmlParserTypeConversion) {
		t.Log("Conversion failed as expected.")
	} else {
		t.Fatalf("Test failed unexpectedly: %s", parseErr)
	}

}

func Test_FetchValue_OK(t *testing.T) {

	const html = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>lorem ipsum</title>
</head>
<body>
<ul>
    <li>Lorem ipsum dolor sit amet, consectetuer adipiscing elit.</li>
    <li>Aliquam tincidunt mauris eu risus.</li>
    <li>2.0</li>
</ul>
</body>
</html>`

	path := "html/body/ul/li[3]"
	expectedResult := 2.0

	result, parseErr := FetchValue(path, bytes.NewBuffer([]byte(html)), "")

	if parseErr == nil && result == expectedResult {
		t.Log("Test succeeded as expected.")
	} else {
		t.Fatalf("Test failed unexpectedly: %s", parseErr)
	}

}

// ([0-9]+)Mb

func Test_FetchValue_Regex_OK(t *testing.T) {

	const html = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>lorem ipsum</title>
</head>
<body>
<ul>
    <li>Lorem ipsum dolor sit amet, consectetuer adipiscing elit.</li>
    <li>Aliquam tincidunt mauris eu risus.</li>
    <li>Vestibulum 650Mb auctor dapibus neque.</li>
</ul>
</body>
</html>`

	path := "html/body/ul/li[3]"
	regex := `([0-9]+)Mb`
	var expectedResult float64 = 650

	result, parseErr := FetchValue(path, bytes.NewBuffer([]byte(html)), regex)

	if parseErr == nil && result == expectedResult {
		t.Log("Test succeeded as expected.")
	} else {
		t.Fatalf("Test failed unexpectedly: %s", parseErr)
	}

}

func Test_FetchValue_Regex_Compile_Error(t *testing.T) {

	const html = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>lorem ipsum</title>
</head>
<body>
<ul>
    <li>Lorem ipsum dolor sit amet, consectetuer adipiscing elit.</li>
    <li>Aliquam tincidunt mauris eu risus.</li>
    <li>67 Vestibulum 650Mb auctor dapibus 12 neque.</li>
</ul>
</body>
</html>`

	path := "html/body/ul/li[3]"
	regex := `(\d+Mb`

	value, parseErr := FetchValue(path, bytes.NewBuffer([]byte(html)), regex)

	if parseErr != nil && parseErr == parseErr.(ErrorHtmlParserRegexCompileError) {
		t.Log("Test succeeded as expected.")
		t.Log("value: ", value, "Error: ", parseErr)
	} else {
		t.Log("value: ", value, "Error: ", parseErr)
		t.Fatalf("Test failed unexpectedly: %s", parseErr)
	}

}

func Test_FetchValue_Regex_No_Match(t *testing.T) {

	const html = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>lorem ipsum</title>
</head>
<body>
<ul>
    <li>Lorem ipsum dolor sit amet, consectetuer adipiscing elit.</li>
    <li>Aliquam tincidunt mauris eu risus.</li>
    <li>Vestibulum auctor dapibus neque.</li>
</ul>
</body>
</html>`

	path := "html/body/ul/li[3]"
	regex := `([0-9]+)`

	value, parseErr := FetchValue(path, bytes.NewBuffer([]byte(html)), regex)

	if parseErr != nil && parseErr == parseErr.(ErrorHtmlParserRegexNoMatch) {
		t.Log("Test succeeded as expected.")
		t.Log("value: ", value, "Error: ", parseErr)
	} else {
		t.Log("value: ", value, "Error: ", parseErr)
		t.Fatalf("Test failed unexpectedly: %s", parseErr)
	}

}

func Test_FetchValue_Regex_No_Capture_Group(t *testing.T) {

	const html = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>lorem ipsum</title>
</head>
<body>
<ul>
    <li>Lorem ipsum dolor sit amet, consectetuer adipiscing elit.</li>
    <li>Aliquam tincidunt mauris eu risus.</li>
    <li>Vestibulum auctor 600Mb dapibus neque.</li>
</ul>
</body>
</html>`

	path := "html/body/ul/li[3]"
	regex := `[0-9]+`

	value, parseErr := FetchValue(path, bytes.NewBuffer([]byte(html)), regex)

	if parseErr != nil && parseErr == parseErr.(ErrorHtmlParserRegexNoCaptureGroup) {
		t.Log("Test succeeded as expected.")
		t.Log("value: ", value, "Error: ", parseErr)
	} else {
		t.Log("value: ", value, "Error: ", parseErr)
		t.Fatalf("Test failed unexpectedly: %s", parseErr)
	}

}
