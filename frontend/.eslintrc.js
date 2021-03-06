// In a file called .eslintrc.js
module.exports = {
    parser: "babel-eslint",
    rules: {
      "graphql/template-strings": ['error', {
        // Import default settings for your GraphQL client. Supported values:
        // 'apollo', 'relay', 'lokka', 'fraql', 'literal'
        env: 'apollo',
  
        // Import your schema JSON here
        schemaJson: require('./schema.json'),
  
        // OR provide absolute path to your schema JSON (but not if using `eslint --cache`!)
        // schemaJsonFilepath: path.resolve(__dirname, './schema.json'),
  
        // OR provide the schema in the Schema Language format
        // schemaString: printSchema(schema),
  
        // tagName is gql by default
      }]
    },
    plugins: [
      'graphql'
    ]
  }