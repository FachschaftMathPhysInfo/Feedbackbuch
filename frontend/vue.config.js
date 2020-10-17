module.exports = {
  "transpileDependencies": [
    "vuetify"
  ],
  chainWebpack: config => {
    config
      .plugin('html')
      .tap(args => {
        args[0].title = 'Feedbackbuch'
        return args
      })
  },
  pluginOptions: {
    // Apollo-related options
    apollo: {
      // Enable ESLint for `.gql` files
      lintGQL: true,
    }
  }
}