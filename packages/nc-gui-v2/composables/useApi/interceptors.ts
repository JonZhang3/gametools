import type { Api } from 'nocodb-sdk'
import { navigateTo, useGlobal, useRoute, useRouter } from '#imports'

const DbNotFoundMsg = 'Database config not found'

export function addAxiosInterceptors(api: Api<any>) {
  const state = $(useGlobal())
  const router = useRouter()
  const route = useRoute()

  api.instance.interceptors.request.use((config) => {
    config.headers['xc-gui'] = 'true'

    if (state.token) config.headers['xc-auth'] = state.token

    if (!config.url?.endsWith('/user/me') && !config.url?.endsWith('/admin/roles')) {
      // config.headers['xc-preview'] = store.state.users.previewAs
    }

    if (!config.url?.endsWith('/user/me') && !config.url?.endsWith('/admin/roles')) {
      if (route && route.params && route.params.shared_base_id) config.headers['xc-shared-base-id'] = route.params.shared_base_id
    }

    return config
  })

  // Return a successful response back to the calling service
  api.instance.interceptors.response.use(
    (response) => response,
    // Handle Error
    (error) => {
      if (error.response && error.response.data && error.response.data.msg === DbNotFoundMsg) return router.replace('/project/0')

      // Return any error which is not due to authentication back to the calling service
      if (!error.response || error.response.status !== 401) {
        return Promise.reject(error)
      }

      // Logout user if token refresh didn't work or user is disabled
      if (error.config.url === '/auth/refresh-token') {
        state.signOut()

        return Promise.reject(error)
      }

      // Try request again with new token
      return api.instance
        .post('/auth/refresh-token', null, {
          withCredentials: true,
        })
        .then((token) => {
          // New request with new token
          const config = error.config
          config.headers['xc-auth'] = token.data.token
          state.signIn(token.data.token)

          return new Promise((resolve, reject) => {
            api.instance
              .request(config)
              .then((response) => {
                resolve(response)
              })
              .catch((error) => {
                reject(error)
              })
          })
        })
        .catch(async (error) => {
          state.signOut()
          // todo: handle new user

          navigateTo('/signIn')

          return Promise.reject(error)
        })
    },
  )

  return api
}
