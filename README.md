# gql

A simple Graphql implementation recreating types from the QuickNode Trending Collection Graphql Endpoint.

## Prerequisites for Running the Application
* Ensure Go is installed

## How to Run the Application
* Run `go mod download` to install project dependencies
* Add your Quick Node API key to  the local environment with `export API_KEY=<API KEY>`
* Run `go run ./` to start up the Graphql server
* Visit http://localhost:8081/query/playground to access the Graphql playground
* Run the following query to get the trending collections
```graphql
query {
  getTrendingCollection(orderBy: "SALES", orderDirection: "DESC") {
    Edges {
      Node {
        Address
        Name
        Stats {
          TotalSales
          Average
          Ceiling
          Floor
          Volume
        }
        Symbol
      }
    }
  }
}
```
Sample output: ![ graphql query output](https://github.com/calebikhuohon/gql/blob/main/graphql-query-output.png?raw=true)


