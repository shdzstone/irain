import axios from 'axios'

class HttpRequest {
    constructor(baseUrl = baseUrl) {
        this.baseURL = baseUrl;
        this.queue = {}
    }

    getInsideConfig () {
        const config = {
            baseURL: this.baseURL,
            headers: {
                // header头部信息
            }
        }
        return config
    }

    // 清空请求池
    destroy(url) {
        delete this.queue[url];
        if (!Object.keys(this.queue).length) {
            // 当请求池为空时
        }
    }

    // 拦截器
    interceptors(instance, url) {
        // 请求拦截
        instance.interceptors.request.use(config => {
            if (!Object.keys(this.queue).length) {

            }
            this.queue[url] = true
            return config       // 要将接收的数据返回才能继续下一步请求
        }, error => {  
            return Promise.reject(error)
        })

        // 相应拦截
        instance.interceptors.response.use(res => {
            this.destroy(url)
            const {data, status} = res;
            return {data, status}
        }, error => {
            return Promise.reject(error)
        })
    }



    request (options) {
        const instance = axios.create()
        options = Object.assign(this.getInsideConfig(), options)
        this.interceptors(instance, options.url)
        return instance
    }
}

export default HttpRequest