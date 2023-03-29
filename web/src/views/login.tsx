import styled from 'styled-components'
import img from '../assets/4.png'
import { ApiLogin } from '../api/index'
import { useForm } from 'react-hook-form'

type formType = {
  user_name: string
  password: string
}

function Login() {
  const { register, handleSubmit } = useForm<formType>({
    defaultValues: {
      user_name: '',
      password: '',
    },
  })
  const submit = async (data: formType) => {
    let res = await ApiLogin(data)
    console.log(res)
  }

  return (
    <RootStyle>
      <div className="body">
        <div className="title">备忘录登录</div>
        <form
          className="table"
          onSubmit={handleSubmit((data) => {
            submit(data)
          })}
        >
          <div>
            <span>账号</span>
            <input {...register('user_name')}></input>
          </div>
          <div>
            <span>密码</span>
            <input type="password" {...register('password')}></input>
          </div>
          <div>
            <button>登录</button>
            <button>注册</button>
          </div>
        </form>
      </div>
    </RootStyle>
  )
}

export default Login

const RootStyle = styled.div`
  width: 100%;
  height: 100%;
  background: url(${img}) no-repeat center/cover;
  display: flex;
  align-items: center;
  justify-content: center;

  .body {
    width: 400px;
    height: 300px;
    background-color: #fff;
    border-radius: 20px;
    padding: 25px;
    & > div {
      text-align: center;
    }
    .title {
      font-size: 20px;
    }
    .table {
      & > div {
        margin: 10px;
      }
    }
  }
`
