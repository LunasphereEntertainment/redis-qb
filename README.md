# Redis Query Builder (Golang)
A golang library for constructing FT.SEARCH queries.

## How to use
1. Construct a query builder.
    ```go
    import "github.com/LunasphereEntertainment/redis-qb/builder"
    
    func main() {
        qb := &builder.QueryBuilder{}
    }
    ```
2. Make your query!
    ```go
    query := qb.NewQuery("my_index_name")
    ```
The query contains chainable methods to fully construct your query.

## Examples
- Text Search
    ```go
    query := qb.NewQuery("test").Equals("title", "hello world")
    ```
  will result in a query like:
    `FT.SEARCH test @title:(hello world)`
- Range search
    ```go
    query := qb.NewQuery("test").InRange("age", 18, 100)
    ```
  will result in a query like:
    `FT.SEARCH test @age:[18 100]`
- Not in that range?
    ```go
    query := qb.NewQuery("test").NotInRange("age", 18, 100)
    ```
    `FT.SEARCH test @age:-[18 100]`
- What about grouping?
    ```go
    query := qb.NewQuery("test").Equals("title", "kafka")
    
    query.CaptureGroup(builder.Or).Equals("test", "hello world")
    ```
  will create the below query:
  `FT.SEARCH test @title:(kafka) | (@test:(hello world))`

### Query completion
Once you've completed the building of your complex query. Simply call `.QueryString()` on the query object.
E.g.
```go
query := qb.NewQuery("test").
	Equals("title", "kafka").
	QueryString()
```

Your query string is now runnable through your chosen Redis Search library.