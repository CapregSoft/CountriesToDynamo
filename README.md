# CountriesToDynamo
Had to upload the countries [data](https://github.com/russ666/all-countries-and-cities-json) to Dynamo Db but the format was not correct so wrote a custom module to convert it into different rows and upload it to dynamo.

Firstly I have to translate the data from the unordered json so that I can store it into the DynamoDb