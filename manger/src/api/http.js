import { message } from 'antd';
import { setToken, getToken } from '../store/index.js'
let baseUrl = 'http://127.0.0.1:8080'
function httpBox (url, params, conf = {}) {
    let config = {
        headers: { 
            'Accept': 'application/x-www-form-urlencoded',
            'token': getToken()
        },
        body: JSON.stringify(params),
        ...conf
    }
    console.log('config', config)
    return fetch(`${baseUrl}/${url}`, config)
    .then((response) => {
        let result = response.json()
        console.log('fanui1', result)
        return result
    })
    .then(v => {
        if (!v.success) {
            message.error(v.message);
        }
        return v
    })
}
export default {
    get: function (url, params, conf = {}) {
        return httpBox(url, params, { ...conf, method: 'GET' })
    },
    post: function (url, params, conf = {}) {
        return httpBox(url, params, { ...conf, method: 'POST' })
    }
}
