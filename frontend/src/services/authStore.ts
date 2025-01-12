import { ref } from 'vue'
import { defineStore } from 'pinia'

const LOCAL_STORAGE_KEY_ID = 'userId'

export const useAuthStore = defineStore('auth', () => {
  const id = ref<string | null>(null)

  const setId = (newValue: string) => {
    id.value = newValue
    localStorage.setItem(LOCAL_STORAGE_KEY_ID, newValue)
  }

  const getId = () => {
    return id.value
  }

  const removeId = () => {
    id.value = null
    localStorage.removeItem(LOCAL_STORAGE_KEY_ID)
  }


  // Persist the token across page refreshes
  const idFromLocalStorage = localStorage.getItem(LOCAL_STORAGE_KEY_ID)
  if (idFromLocalStorage) {
    id.value = idFromLocalStorage
  }

  return {
    setId,
    getId,
    removeId,
  }
})