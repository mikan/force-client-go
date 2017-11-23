force-client-go
===============

A simple client library for Force.com REST API written by [Go](https://golang.org/).

### Key features

- Create, read, update, delete, and query operations through a simple API
- Pre-defined and fully extendable SObject structs
- Only using Golang standard libraries, no additional dependencies required 

### Sample

#### Authentication

```go
cred := force.Credential{ClientID: "xxx", ClientSecret: "xxx", Username: "xxx", Password: "xxx", APIToken: "xxx"}
client, _ := force.NewClient("xxx.salesforce.com", force.Production, "41.0", nil)
err := client.Login(context.Background(), &cred)
```

##### SOQL query

```
var set sobject.ContactSet
client.Query(context.Background(), "SELECT Id,FirstName,LastName FROM Contact", &set)
for _, c := range set.Records {
    fmt.Printf("%s: %s %s\n", c.Id, c.FirstName, c.LastName)
}
```

Output of the above code:

```
0037F000004g6hzQAA: Liz D'Cruz
0037F000004g6i0QAA: Edna Frank
0037F000004g6i1QAA: Avi Green
...
```

##### CRUD

```go
// Create
contact := sobject.Contact{FirstName: "Test", LastName: "User"}
id, err := client.Create(context.Background(), sobject.ContactObjectName, &contact)

// Read
var readResult sobject.Contact
err = client.Read(context.Background(), sobject.ContactObjectName, id, &readResult)

// Update
update := sobject.Contact{FirstName: "Test2"}
err = client.Update(context.Background(), sobject.ContactObjectName, id, &update)

// Delete
err = client.Delete(context.Background(), sobject.ContactObjectName, id)
```

### CLI

Basics:

```bash
go run cmd\force-client\main.go -f config.json <OPERATION>
```

`<OPERATION>` by samples:

- Query: `-q "SELECT Name FROM Contact LIMIT 3"`
- Create: `-c Contact -j {\"FirstName\":\"Test\",\"LastName\":\"User\"}`
- Read: `-r Contact -i 0037F00000Hc2GyQAJ`
- Update: `-u Contact -i 0037F00000Hc2GyQAJ -j {\"FirstName\":\"Test2\"}`
- Delete: `-d Contact -i 0037F00000Hc2GyQAJ`

See `-h` (`--help`) for more information.

### Future works

- BLOB in/out
- SOSL Search
- Composite resources
  - Composite requests
  - Batch requests
  - SObject Tree requests
- Other special endpoints

### Author

- [mikan](https://github.com/mikan)

### License

BSD 3-Clause
