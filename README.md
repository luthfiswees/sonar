# Sonar

Service Liveness Checker

## Installation

Clone the repo. Remember to clone it in your *$GOPATH/src/github.com/luthfiswees*

```sh
git clone https://github.com/luthfiswees/sonar
```

And then copy the env

```sh
cp env.sample .env
```

And start the service

```sh
go run app/sonar/main.go
```

You could also compile it first and run

```sh
cd app/sonar
go build
chmod +x sonar
./sonar
```

## Usage

Sonar read configurations file in `config/config.json`. The default configuration file would look like this

```json
[
    {
      "name" : "Random Service",
      "base-url" : "http://127.0.0.1",
      "path" : "/",
      "port" : "3000"
    },
    {
      "name" : "Image Service",
      "base-url" : "http://127.0.0.1",
      "path" : "/",
      "port" : "4500"
    }
]
```

You could add or change it to your needs. For example here is how I use it

```json
[
    {
      "name" : "My Rails App",
      "base-url" : "http://127.0.0.1",
      "path" : "/",
      "port" : "5000"
    },
    {
      "name" : "My Image Microservice",
      "base-url" : "http://127.0.0.1",
      "path" : "/",
      "port" : "3000"
    },
  	{
      "name" : "My Elasticsearch",
      "base-url" : "http://127.0.0.1",
      "path" : "/_cat/indices?v",
      "port" : "9200"
    }
]
```

And then, restart the app. It will check if `My Rails App`, `My Image Microservice`, and `My Elasticsearch`is up on their respective machine or not.