import { useEffect } from 'react'
import { ApiGettask } from '../api'

function Main() {
  useEffect(() => {
    ApiGettask().then((res) => {
      console.log(res)
    })
  }, [])
  return <div>首页</div>
}

export default Main
