import http from './http'
export default {
    login (params) {
        return http.post('login', params)
    }
}