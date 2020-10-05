import Vue from 'vue';
import Vuetify from 'vuetify/lib';

Vue.use(Vuetify);

export default new Vuetify({
    theme: {
        themes: {
          light: {
            primary: '#c71826',
            secondary: '#b0bec5',
            accent: '#f4f1ea',
            error: '#b71c1c',
          },
        },
      },
});
