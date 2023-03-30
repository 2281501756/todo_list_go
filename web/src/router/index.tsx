import { BrowserRouter, Route, Routes } from 'react-router-dom'
import Login from '../views/login'
import Main from '../views/main'
import Register from '../views/register'

function RouterConfig() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/login" element={<Login />}></Route>
        <Route path="/register" element={<Register />}></Route>
        <Route path="/home" element={<Main />}></Route>
      </Routes>
    </BrowserRouter>
  )
}

export default RouterConfig
