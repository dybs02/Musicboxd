import 'primeicons/primeicons.css'
import './assets/main.css'


import Aura from '@primevue/themes/aura'
import { DefaultApolloClient, provideApolloClient } from '@vue/apollo-composable'
import { createPinia } from 'pinia'
import PrimeVue from 'primevue/config'
import ConfirmationService from 'primevue/confirmationservice'
import ToastService from 'primevue/toastservice';
import Tooltip from 'primevue/tooltip'
import { createApp, h, provide } from 'vue'
import apolloClient from './services/apollo'

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

app.directive('tooltip', Tooltip);

app.use(createPinia())
app.use(router)

app.use(ConfirmationService);
app.use(ToastService);

app.mount('#app')
