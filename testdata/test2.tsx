const Test2Query = gql(/* GraphQL */ `
    query Query {
        field
    }
`);

const Test2Mutation = gql(/* GraphQL */ `
    mutation Mutation {
        mutation(arg: "value") {
            field
        }
    }
`);
