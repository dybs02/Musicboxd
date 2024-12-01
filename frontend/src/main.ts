import './assets/main.css'
import 'primeicons/primeicons.css'


import { createApp, provide, h } from 'vue'
import { createPinia } from 'pinia'
import { DefaultApolloClient, provideApolloClient } from '@vue/apollo-composable'
import apolloClient from './services/apollo'
import PrimeVue from 'primevue/config';
import Aura from '@primevue/themes/aura';

import App from './App.vue'
import router from './router'

const app = createApp({
  setup () {
    provide(DefaultApolloClient, apolloClient)
  },

  render: () => h(App),
})
provideApolloClient(apolloClient);


app.use(PrimeVue, {
  theme: {
      preset: Aura
  }
});

app.use(createPinia())
app.use(router)

app.mount('#app')
