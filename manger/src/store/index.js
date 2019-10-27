let token = ''
export function setToken (t) {
    if (typeof t === 'string') {
        token = t
    }
}
export function getToken () {
    let t = token
    return t
}