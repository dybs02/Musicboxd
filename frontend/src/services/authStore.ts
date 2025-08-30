import { ref } from 'vue'
import { defineStore } from 'pinia'

export const LOCAL_STORAGE_KEY_ID = 'userId'
export const LOCAL_STORAGE_KEY_JWT = 'jwt'
export const LOCAL_STORAGE_KEY_ROLE = 'role'
export const LOCAL_STORAGE_KEY_AVATAR = 'avatarUrl'
export const LOCAL_STORAGE_KEY_THEME = 'light'
export const LOCAL_STORAGE_KEY_LOCALE = 'locale'

export const useAuthStore = defineStore('auth', () => {
  const id = ref<string | null>(null)
  const jwt = ref<string | null>(null)
  const role = ref<string | null>(null)
  const avatarUrl = ref<string | null>(null)
  const theme = ref<string | null>(null)

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

  const isModerator = () => {
    return role.value === 'moderator'
  }

  const setAvatarUrl = (newValue: string) => {
    avatarUrl.value = newValue
    localStorage.setItem(LOCAL_STORAGE_KEY_AVATAR, newValue)
  }

  const getAvatarUrl = () => {
    return avatarUrl.value
  }

  const setTheme = (newValue: string) => {
    theme.value = newValue
    localStorage.setItem(LOCAL_STORAGE_KEY_THEME, newValue)
  }

  const getTheme = () => {
    return theme.value
  }

  const logout = () => {
    removeId()
    removeJWT()
    removeRole()
  }

  const isLoggedIn = () => {
    return !!id.value && !!jwt.value && !!role.value
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

  const avatarFromLocalStorage = localStorage.getItem(LOCAL_STORAGE_KEY_AVATAR)
  if (avatarFromLocalStorage) {
    avatarUrl.value = avatarFromLocalStorage
  }

  const themeFromLocalStorage = localStorage.getItem(LOCAL_STORAGE_KEY_THEME)
  if (themeFromLocalStorage) {
    theme.value = themeFromLocalStorage
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
    isModerator,
    setAvatarUrl,
    getAvatarUrl,
    setTheme,
    getTheme,
    logout,
    isLoggedIn,
  }
})