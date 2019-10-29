const { AddRequest, AddResponse } = require("./calculator_pb")
const { CalculatorClient } = require("./calculator_grpc_web_pb")
var client = new CalculatorClient('http://localhost:1960');

var request = new AddRequest()

request.setNum1(333)
request.setNum2(111)

client.add(request, {}, (err, response) => {
    console.log("Result of Add : ",response.getResult())
})

const { GetQuery, Customer } = require("./customers_pb")
const { CustomersClient } = require("./customers_grpc_web_pb")
var customersClient = new CustomersClient('http://localhost:1970');

var getQuery = new GetQuery()

getQuery.setCustomerid(12)

customersClient.get(getQuery, {}, (err, response) => {
    console.log("Result of Customer : ",response.array[0])
})