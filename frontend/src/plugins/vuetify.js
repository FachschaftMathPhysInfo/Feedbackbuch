import Vue from 'vue';
import '@mdi/font/css/materialdesignicons.css' // Ensure you are using css-loader
import Vuetify from 'vuetify/lib';

Vue.use(Vuetify);

export default new Vuetify({
    theme: {
        themes: {
          light: {
            primary: '#c71826',
            secondary: '#f4f1ea',
            accent: '#590d08',
            error: '#b71c1c',
          },
          dark: {
            primary: "#590d08",
          },
        },
        
      },
      icons: {
        iconfont: 'mdi', // default - only for display purposes
      },
});
