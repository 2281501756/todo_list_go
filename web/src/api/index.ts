import request from '../utils/request'

type User = {
  ID: string
  user_name: string
  password: string
  CreatedAt: string
  UpdatedAt: string
}

type LoginRequestType = {
  user_name: string
  password: string
}
type LoginResponseType = {
  token: string
  user: User
}

type Task = {
  ID: string
  userId: string
  title: string
  content: string
  status: number
  startTime: number
  endTime: number
  user: User
}
export const ApiLogin = (data: LoginRequestType) => {
  return request<LoginResponseType>({
    url: '/v1/user/login',
    method: 'POST',
    data: {
      ...data,
    },
  }).then((res) => {
    if (res.status === 200) {
      localStorage.setItem('token', res.data.token)
    }
    return res
  })
}

export const ApiRegister = (data: LoginRequestType) => {
  return request<LoginResponseType>({
    url: '/v1/user/register',
    method: 'POST',
    data: {
      ...data
    }
  })
}

export const ApiGettask = () => {
  return request<Task[]>({
    url: '/v1/task',
    method: 'GET',
  })
}
