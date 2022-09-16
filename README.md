# querystring

`querystring` finds a GraphQL query in your files.

```
go install github.com/gqlgo/querystring@latest
```

## Sample usage

```
$ cat testdata/test3.tsx
const Test3Query = gql(/* GraphQL */ `
    query Query {
        field
    }
`);

const Test3Mutation = gql(/* GraphQL */ `
    mutation Mutation {
        mutation(arg: "value") {
            field
        }
    }
`);

const Test3Text = `Hello`


$ querystring testdata/test3.tsx
query Query {
        field
}
mutation Mutation {
        mutation(arg: "value") {
                field
        }
}
```
