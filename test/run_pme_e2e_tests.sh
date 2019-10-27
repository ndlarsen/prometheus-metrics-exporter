#!/bin/bash

exit_error() {
    if [ $# -ne 0 ]
    then
      echo "$@"
    fi
    exit 1
}


JOB_NAME_JSON_NO_BASICAUTH="TestJsonNoBasicAuthJobName"
METRICS_HELP_JSON_NAME="test_json_help_name"
METRICS_HELP_JSON_VALUE="test json help value"
LBL_JSON_NO_BASICAUTH_NAME="TestJsonNoBasicAuthLabelName"
LBL_JSON_NO_BASICAUTH_VALUE="TestJsonNoBasicAuthLabelValue"
TYPE="gauge"
TYPE_DECLARATION="# TYPE $METRICS_HELP_JSON_NAME $TYPE"
HELP_DELACRATION="# HELP $METRICS_HELP_JSON_NAME $METRICS_HELP_JSON_VALUE"
EXPECTED_VALUE="1"

VALUES=(
  "$JOB_NAME_JSON_NO_BASICAUTH"
  "$METRICS_HELP_JSON_NAME"
  "$METRICS_HELP_JSON_VALUE"
  "$LBL_JSON_NO_BASICAUTH_NAME"
  "$LBL_JSON_NO_BASICAUTH_VALUE"
  "$TYPE_DECLARATION"
  "HELP_DELACRATION"
)



PG_NAME=pushgateway
PG_URL="http://$PG_NAME:9091"

DATE_STR=$(date +"%Y%m%d%H%M%S%N")
TEMP_CONF="tempconfig-$DATE_STR.json"

echo "executing PME E2 tests"

cat <<- EOF > ./"$TEMP_CONF" || exit_error "Could not create temporary test config file"
  {
    "pushGatewayUrl": "$PG_URL",
    "scrapeTargets": [
      {
        "url": "https://jsonplaceholder.typicode.com/posts/",
        "mimeType": "json",
        "jobName": "$JOB_NAME_JSON_NO_BASICAUTH",
        "timeoutInSecs": 10,
        "metrics": [
          {
            "name": "$METRICS_HELP_JSON_NAME",
            "help": "$METRICS_HELP_JSON_VALUE",
            "path": "0.id",
            "instrumentType": "gauge"
          }
        ],
        "labels": [
          {"name": "$LBL_JSON_NO_BASICAUTH_NAME", "value": "$LBL_JSON_NO_BASICAUTH_VALUE"}
        ]
      }
    ]
  }
EOF

if [ ! -f "./$TEMP_CONF" ]
then
  exit_error "config file not available"
fi

exit
./binaries/prometheus-metrics-exporter-linux-amd64 -config="./$TEMP_CONF" || exit_error "PME failed to run"

METRICS_FILE="$(date +"%Y%m%d%H%M%S%N").metrics"

curl -s --output "./$METRICS_FILE" "$PG_URL/metrics"

if [ ! -f "./$METRICS_FILE" ]
then
  exit_error "metrics file not available"
fi

cat "./$METRICS_FILE"

for v in "${VALUES[@]}"
do
  echo "$v"
  grep "$v" "$METRICS_FILE" || exit_error "Test failed: value \"$v\" not found"
done
