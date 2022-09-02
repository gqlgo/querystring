const Tes2Query = gql(/* GraphQL */ `
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

const Test3Text = `Hello`
