// composables/useTheme.js
import { useAuthStore } from '@/services/authStore';
import { definePreset } from '@primevue/themes'
import Aura from '@primevue/themes/aura'
import { computed, ref } from 'vue'


export function useTheme() {
  const store = useAuthStore();
  const isDark = computed(() => store.getTheme() === 'dark');

  const toggleTheme = () => {
    if (!isDark.value) {
      document.documentElement.classList.add('dark-theme')
      store.setTheme('dark');
    } else {
      document.documentElement.classList.remove('dark-theme')
      store.setTheme('light');
    }
  }

  const setTheme = () => {
    if (isDark.value) {
      document.documentElement.classList.add('dark-theme')
    } else {
      document.documentElement.classList.remove('dark-theme')
    }
  }

  return {
    isDark,
    toggleTheme,
    setTheme
  }
}



export const MyPreset = definePreset(Aura, {
  semantic: {
    colorScheme: {
      light: {
        snow: '#1a4aba',
        darker: '#f3f2f2',
      },
      dark: {
        snow: '#575925',
        darker: '#101013',
      }
    }
  }
})
