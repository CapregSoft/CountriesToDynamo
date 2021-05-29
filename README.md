# CountriesToDynamo

## Why Needed
I had to upload the countries [data](https://github.com/russ666/all-countries-and-cities-json) to Dynamo Db but the format was not correct so had to wrote a custom module to convert it into individual row and upload it to dynamo.

## Purpose
To store the data into the DynamoDb

## Requirements
- Golang 
- DynamoDb
- AWS

## Environment and Data
- The countries data are in the [/data/countries.json](/data/countries.json)

- Config variables are in the [/env/default.example.yml](/env/default.example.yml).


## Build

```
$ go build
$ ./CountriesToDynamo
```

## Contribution
You're welcome to contribute to this project.

You should follow contribution guideline [Contribution guideline](/CONTRIBUTING.md)


## License

Apache License
