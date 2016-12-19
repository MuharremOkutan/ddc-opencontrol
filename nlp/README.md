# Text Analysis Tool for DDC OpenControl

The NLP tool can be used to validate that the Docker Datacenter component narratives adhere to the NIST 800-53 security control specifications.

## Usage

The application is packaged using the included Dockerfile. In order to run the tool, you must provide your own Cognitive Services Text Analytics API key. This can be done via an environment variable or command-line flag passed to the container.

```sh
# Build image
docker build -t docker/ddc-opencontrol-nlp .

# Run command
docker run -e "TEXT_ANALYTICS_API_KEY=$TEXT_ANALYTICS_API_KEY" docker/ddc-opencontrol-nlp

Usage:
  nlp [command]

Available Commands:
  analyze     Run text analysis

Use "nlp [command] --help" for more information about a command.
```

## Developing

The text analysis tool is built with [Go 1.7+](https://golang.org/). Dependencies are vendored with [govendor](https://github.com/kardianos/govendor), and the application is packaged with Docker.

The tool uses the [Microsoft Cognitive Services Text Analytics API](https://www.microsoft.com/cognitive-services/en-us/text-analytics-api) to analyze the component narratives. [go-swagger](https://goswagger.io/) is used to generate the API client based on the Swagger definition [here](https://westus.dev.cognitive.microsoft.com/docs/services/TextAnalytics.V2.0/export?DocumentFormat=Swagger&ApiName=Azure%20Machine%20Learning%20-%20Text%20Analytics).

```sh
swagger generate client \
  -f https://westus.dev.cognitive.microsoft.com/docs/services/TextAnalytics.V2.0/export?DocumentFormat=Swagger&ApiName=Azure%20Machine%20Learning%20-%20Text%20Analytics \
  -t textanalytics/ \
  -A TextAnalytics

go get -u -f textanalytics/...
```

The `circle.yml` file runs the tool as part of the testing process and leverages Docker, Inc's existing Cognitive Services API keys.