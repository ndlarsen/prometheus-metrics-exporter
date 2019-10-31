package writers

import (
	"encoding/json"
	"net/http"
)

const JsonString = `{
    "_id": "5db5ebb0a48c30b56c8ffe7d",
    "index": 0,
    "guid": "5fac063c-54e3-4a0d-b518-5aa97a417f1c",
    "age": 25,
    "email": "mullinsgregory@harmoney.com",
    "phone": "+1 (996) 489-3894",
    "registered": "2015-11-07T06:09:25 -01:00",
    "friends": [
      {
        "id": 25,
        "name": "Parrish Mills"
      },
      {
        "id": 65,
        "name": "Patti Walter"
      }
    ],
    "greeting": "Hello, Mullins Gregory! You have 5 unread messages."
}`

const HtmlString = `<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Test Page</title>
	</head>
	<body>
		<div>
			<ul>
				<li>
					This item has no numeric value.
				</li>
				<li>
					This item has a numeric value somewhere. 567MB is an odd amount of RAM.
				</li>
			</ul>
		</div>
	</body>
</html>`

func JsonWriter(w http.ResponseWriter, r *http.Request) {
	in := []byte(JsonString)
	var raw map[string]interface{}
	err := json.Unmarshal(in, &raw)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	out, _ := json.Marshal(raw)
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func HtmlWriter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(HtmlString))
}
