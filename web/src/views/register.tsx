import { useForm } from "react-hook-form"
import { useNavigate, useRoutes } from "react-router-dom"
import styled from "styled-components"
import { ApiRegister } from "../api"
import img from '../assets/2.png'

type FormType = {
  user_name: string
  password: string
  confirmPassword: string
}

function Register() {
  const navigate = useNavigate()
  const { register, handleSubmit } = useForm<FormType>({
    defaultValues: {
      user_name: '',
      password: '',
      confirmPassword: '',
    }
  })
  const submit = async (data: FormType) => {
    let res = await ApiRegister(data)
    console.log(res)
  }
  const handleToLogin = () => {
    navigate('/login')
  }

  return <RootStyle>
    <div className="body">
      <div className="title">备忘录注册</div>
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
          <span>确认密码</span>
          <input type="password" {...register('confirmPassword')}></input>
        </div>
        <div className='body-box'>
          <button>注册</button>
          <button onClick={handleToLogin}>切换到登录</button>
        </div>
      </form>
    </div>
  </RootStyle>
}

export default Register

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
        display: flex;
        justify-content: center;
      }
    }
  }

  .body-box {
    text-align: center;
    & > button {
      margin: 10px;
    }
}
`
