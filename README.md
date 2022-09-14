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
Sample output: ![ graphql query output](https://github.com/calebikhuohon/gql/blob/main/img/graphql-query-output.png?raw=true)

_________
## Technical Question

Here, we describe a computing system that aims to be secure and serve the needs of millions of users. 
The following considerations will be made:
* How we can achieve user authentication and the authentication mechanisms to be used
* The database in use and why it was picked
* How we can efficiently scale the system to handle multiple concurrent requests
* How we can achieve real-time updates on new data

For this application, we will be utilizing a centralized authentication system in which JWT (JSON Web Tokens) are created for each authenticated user. These tokens are also verified by individual services before data access is allowed. This ensures that trust is defined at every border and a single failure point is avoided.

We will also utilize authorization schemes like Role-Based Access Control (RBAC) schemes to ensure that authenticated users can only access data resources within their permission levels. RBAC schemes have the advantage of been easy to implement, well-supported by major web platforms, and having an intuitive approach which is easy to understand for developers and users alike.

![ authentication layer](https://github.com/calebikhuohon/gql/blob/main/img/auth.jpg?raw=true)

To allow for authentication verification by individual services, authentication information needs to be stored in a database. An SQL-type database will be used as they are ACID (Atomic, consistency, isolation, durability) compliant and hence authentication data is guaranteed to always be consistent. PostgreSQL will be the SQL database of choice as it is suitable for high volume reads and writes, which would be the case for our authentication system.

Scaling out the system to support several requests per second will involve the addition of several layers. My top three are discussed below:

* **A server auto-scaling setup**: In other to support a large number of concurrent requests, there will be a need for multiple servers attending to various simultaneous requests. The server setup will need to horizontally scale as server load increases, hence the need for this process to be automated.

* **A Load Balancer**: The load balancer is linked to the auto-scaling group and ensures an even distribution of web traffic among these servers. The load balancer will ensure that web traffic is routed to only healthy servers, hence keeping the system online.

* **Caches**: Caches will store the result of expensive data operations or frequently accessed data and provide quick data retrievals. This greatly improves the platform performance as repeated database calls (which are expensive) are avoided and the database workload is reduced. Data not present in the cache is gotten from the database, stored in the cache, and sent to the client.

To integrate support for real-time client updates, the platform has to be event-driven. An event is fired when new data is available. The client subscribes to this event stream and listens for the update event. Once an update event is detected, the client could carry out other business logic based off the new data.