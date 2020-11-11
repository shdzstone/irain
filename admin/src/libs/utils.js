import Cookies from 'js-cookie';
import config from '@/config'
export const TOKEN_KEY = 'token'
const { title, cookieExpires } = config

/**
 * 设置和读取token
 */

export const setToken = (token) => {
    Cookies.set(TOKEN_KEY, token, { expires: cookieExpires || 1 })
}

export const getToken = () => {
    const token = Cookies.get(TOKEN_KEY)
    if (token) return token
    else return false
}