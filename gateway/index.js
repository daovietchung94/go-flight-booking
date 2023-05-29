const {ApolloServer} = require('apollo-server')
const { ApolloGateway } = require("@apollo/gateway")

const fs = require("fs")
let configFile = "config.json"
if(process.argv.length >2 && process.argv[2] == "DOCKER"){
    configFile = "config-docker.json"
}
let configData = fs.readFileSync(configFile, "utf8")
let config = JSON.parse(configData)

let graphqlCustomer = `http://${config.graphql.customer.host}:${config.graphql.customer.port}/query`
let graphqlFlightManager = `http://${config.graphql.flightManager.host}:${config.graphql.flightManager.port}/query`
let graphqlLogisticsManager = `http://${config.graphql.logisticsManager.host}:${config.graphql.logisticsManager.port}/query`

const gateway = new ApolloGateway({
    serviceList: [
        {name: "customers", url: graphqlCustomer},
        {name: "flight_manager", url: graphqlFlightManager},
        {name: "logistics_manager", url: graphqlLogisticsManager}
    ]
})

const server = new ApolloServer({
    gateway,
})

server.listen().then(({ url }) => {
    console.log(`🚀 Server ready at ${url}`);
});
