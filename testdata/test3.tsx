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
