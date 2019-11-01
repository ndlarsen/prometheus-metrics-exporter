# Prometheus Metrics Exporter

PME is a simple application to scrape values from JSON or HTML HTTP endpoints and push metrics to a prometheus 
pushgateway.
The application might come in handy when for example:
- your scrape targets are behind a NAT firewall and your prometheus instance cannot reach them from the outside.
- for some reason you cannot or are not allowed to set up a prometheus client on the systems you wish to monitor.

The tool is neither a service nor a daemon. If you want continuous pull/push, set up something like a CRON job.
Values can be scraped from JSON by [dot-notation](https://docs.oracle.com/en/database/oracle/oracle-database/12.2/adjsn/simple-dot-notation-access-to-json-data.html#GUID-7249417B-A337-4854-8040-192D5CEFD576)
 or from HTML by [XPath](https://en.wikipedia.org/wiki/XPath).

Currently PME supports following prometheus instruments:
 - counter
 - gauge

## Command line flags
`-config=path/to/configFile`

## Configuration
The configuration format is JSON.

Configuration overview:

![Configuration diagram](http://www.plantuml.com/plantuml/proxy?cache=no&src=https://raw.githubusercontent.com/ndlarsen/prometheus-metrics-exporter/master/docs/configuration_overview.puml)

Configuration example:

```json
{
  "pushGatewayUrl": "http://pushgateway.url:port",
  "scrapeTargets": [
    {
      "url": "http://some.url/path",
      "basicAuth": {
                    "username": "username1",
                    "password": "password123"
                   },
      "mimeType": "json",
      "jobName": "jobNameGoesHere",
      "timeoutInSecs": 10,
      "metrics": [
        {
          "name": "name_goes_here",
          "help": "help text goes here",
          "path": "json.path.to.value",
          "instrumentType": "gauge",
          "regex": "(\\d+)"
        },
        {
          "name": "name_goes_here",
          "help": "help text goes here",
          "path": "some.other.path",
          "instrumentType": "counter",
          "regex": "(\\d+)"
        }
      ],
      "labels": [
        {
          "name": "nameOfLabel",
          "value": "valueOfLabel"
        },
        {
          "name": "nameOfLabel2",
          "value": "valueOfLabel2"
        }
      ]
    },
    {
      "url": "http://another.url/path",
      "mimeType": "html",
      "jobName": "anotherJobName",
      "timeoutInSecs": 10,
      "metrics": [
        {
          "name": "name_goes_here",
          "help": "help text goes here",
          "path": "/xpath/to/some[1]/value",
          "instrumentType": "gauge",
          "regex": "*(\\d+)"
        },
        {
          "name": "name_goes_here",
          "help": "help text goes here",
          "path": "/xpath/to/another/value",
          "instrumentType": "counter",
          "regex": "*(\\d+)"
        }
      ],
      "labels": [
        {
          "name": "nameOfLabel",
          "value": "valueOfLabel"
        },
        {
          "name": "nameOfLabel2",
          "value": "valueOfLabel2"
        }
      ]
    }
  ]
}
```
### Global entries
- pushGatewayUrl: the url of the pushgateway. It is global for the entire configuration.
- scrapeTargets: a list of targets to scrape from. (Documented below)

### ScrapeTargets
- url: the url to scrape from
- basicAuth (optional) : information for authentication with the specific scrape target (Documented below)
- mimeType: the content type of the scrape url. JSON and HTML are supported.
- jobName: the job name to display in the pushgateway.
- timeoutInSecs: the time to wait for response from the scrape target before timing out.
- metrics: a list of metrics to pull from the scrape target response. (Documented below)
- labels: a list of labels for the current job. (Documented below)

### BasicAuth (optional)
- username: the basic auth username
- password: the basic auth password

### Metrics
- name: the name for the instrument.
- help: the help text for the instrument.
- path: the path of the value for the instrument. (dot notation or xpath)
- instrumentType: the type of the instrument. Counter and gauge are supported.
- regex (optional): the regex to apply to the path content.
  - An empty regex will return the entire path content as a float.
  - An omitted regex will return the entire path content as a float.
  - A nonempty regex must include a single capture group. 

### Labels
- name: the name of the label.
- value: the value of the label.

## Go version
The application was originally written in go v. 1.11.2. Latest tested in 1.11.13

## External libraries and/or modules
The following libraries and/or modules are directly used in the project.
Note that other modules might be indirectly included. 

- [gjson](https://github.com/tidwall/gjson)
- [htmlquery](https://github.com/antchfx/htmlquery)
- [prometheus go client library](https://github.com/prometheus/client_golang)
