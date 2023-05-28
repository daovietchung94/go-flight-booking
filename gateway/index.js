const {ApolloServer} = require('apollo-server')
const { ApolloGateway } = require("@apollo/gateway")

const gateway = new ApolloGateway({
    serviceList: [
        {name: "customers", url: "http://127.0.0.1:4001/query"},
        // {name: "flight_manager", url: "http://127.0.0.1:4001/query"},
        // {name: "logistics_manager", url: "http://127.0.0.1:4002/query"}
    ]
})

const server = new ApolloServer({
    gateway,
})

server.listen().then(({ url }) => {
    console.log(`🚀 Server ready at ${url}`);
});
