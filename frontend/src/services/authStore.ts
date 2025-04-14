import { ref } from 'vue'
import { defineStore } from 'pinia'

const LOCAL_STORAGE_KEY_ID = 'userId'
const LOCAL_STORAGE_KEY_JWT = 'jwt'
const LOCAL_STORAGE_KEY_ROLE = 'role'

export const useAuthStore = defineStore('auth', () => {
  const id = ref<string | null>(null)
  const jwt = ref<string | null>(null)
  const role = ref<string | null>(null)

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

  const setJWT = (newValue: string) => {
    jwt.value = newValue
    localStorage.setItem(LOCAL_STORAGE_KEY_JWT, newValue)
  }

  const getJWT = () => {
    return jwt.value
  }

  const removeJWT = () => {
    jwt.value = null
    localStorage.removeItem(LOCAL_STORAGE_KEY_JWT)
  }

  const setRole = (newValue: string) => {
    role.value = newValue
    localStorage.setItem(LOCAL_STORAGE_KEY_ROLE, newValue)
  }

  const getRole = () => {
    return role.value
  }

  const removeRole = () => {
    role.value = null
    localStorage.removeItem(LOCAL_STORAGE_KEY_ROLE)
  }

  // Persist the token across page refreshes
  const idFromLocalStorage = localStorage.getItem(LOCAL_STORAGE_KEY_ID)
  if (idFromLocalStorage) {
    id.value = idFromLocalStorage
  }

  const jwtFromLocalStorage = localStorage.getItem(LOCAL_STORAGE_KEY_JWT)
  if (jwtFromLocalStorage) {
    jwt.value = jwtFromLocalStorage
  }

  const roleFromLocalStorage = localStorage.getItem(LOCAL_STORAGE_KEY_ROLE)
  if (roleFromLocalStorage) {
    role.value = roleFromLocalStorage
  }

  return {
    setId,
    getId,
    removeId,
    setJWT,
    getJWT,
    removeJWT,
    setRole,
    getRole,
    removeRole,
  }
})