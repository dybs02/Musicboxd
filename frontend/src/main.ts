import 'primeicons/primeicons.css'
import './assets/main.css'


import { DefaultApolloClient, provideApolloClient } from '@vue/apollo-composable'
import { createPinia } from 'pinia'
import PrimeVue from 'primevue/config'
import ConfirmationService from 'primevue/confirmationservice'
import ToastService from 'primevue/toastservice'
import Tooltip from 'primevue/tooltip'
import { createApp, h, provide } from 'vue'
import apolloClient from './services/apollo'

import mitt from 'mitt'
import App from './App.vue'
import { MyPreset } from './primevue/theme'
import router from './router'

const app = createApp({
  setup () {
    provide(DefaultApolloClient, apolloClient)
  },

  render: () => h(App),
})
provideApolloClient(apolloClient);

app.use(createPinia())

app.use(PrimeVue, {
  theme: {
    preset: MyPreset,
    options: {
      darkModeSelector: '.dark-theme',
    }
  }
})

app.directive('tooltip', Tooltip);

app.use(router)

app.use(ConfirmationService);
app.use(ToastService);

app.provide('emitter', mitt())

app.mount('#app')
