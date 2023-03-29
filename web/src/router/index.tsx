import { BrowserRouter, Route, Routes } from 'react-router-dom'
import Login from '../views/login'
import Main from '../views/main'

function RouterConfig() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Login />}></Route>
        <Route path="/home" element={<Main />}></Route>
      </Routes>
    </BrowserRouter>
  )
}

export default RouterConfig
