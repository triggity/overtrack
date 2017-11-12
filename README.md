

## Developing

### Generating Code (enums)
This project uses generated functions for certain types
#### Enums
Enums have both [stringer](https://godoc.org/golang.org/x/tools/cmd/stringer) generated code to match `stringer` interface, as well as [jsonenums](https://github.com/campoy/jsonenums) generated code for json marshal/unmarshal of enums. In order to create new functions when a change is made, run: 
```
cd models
stringer -type=${Name}
jsonenums -type=${Name}
```

### Migrations
To create a migration, run
```
make NAME=foo migration
```