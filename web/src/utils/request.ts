import axios, { AxiosRequestConfig } from 'axios'

type ResponseSerializer<T = any> = {
  status: number
  data: T
  msg: string
  err: string
}

const axiosInstance = axios.create({
  baseURL: '/api',
  timeout: 5000,
})

const whiteList = ['/v1/user/login', '/v1/user/register']

axiosInstance.interceptors.request.use(
  (config) => {
    console.log(config)
    if (config.url) {
      if (!whiteList.includes(config.url) && localStorage.getItem('token')) {
        config.headers.Authorization = localStorage.getItem('token')
      }
    }
    return config
  },
  () => {}
)

const request = async <T = any>(config: AxiosRequestConfig): Promise<ResponseSerializer<T>> => {
  try {
    const { data } = await axiosInstance.request(config)
    return data
  } catch (e) {
    console.log(e)
    return {
      msg: '请求失败',
      data: null as any,
      status: 400,
      err: '请求失败',
    }
  }
}

export default request
