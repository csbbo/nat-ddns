import axios from 'axios'
import qs from 'qs'

const fetchData = (url = '', data = {}, method = 'GET') => {
    if (method === 'GET') {
        let param = qs.stringify(data)
        return new Promise(resolve => {
            axios.get(url+'?'+param).then((resp) => {
                resolve(resp.data)
            })
        })

    } else {
        return new Promise(resolve => {
            axios({
                method: method,
                url: url,
                data: data
            }).then((resp) => {
                resolve(resp.data)
            });
        })

    }
}

export const LoginAPI = data => fetchData('/api/login', data, 'POST')
export const RegistAPI = data => fetchData('/api/regist', data, 'POST')
export const LogoutAPI = () => fetchData('/api/logout', {}, 'POST')
export const CheckAuthAPI = () => fetchData('/api/checkAuth', {}, 'GET')