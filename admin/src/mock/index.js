import Mock from 'mockjs'
import { login } from './login'


Mock.setup({
    timeout: 1000
})

Mock.mock(/\/login/, login)


export default Mock