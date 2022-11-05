import React from 'react'
import Home from './Components/Home'
import Login from './Components/Login'
import Wallet from './Components/Wallet'
import NavBar from './Components/Navbar'
import { Route, Routes } from 'react-router-dom'
import Register from './Components/Register'
import Services from './Components/Services'


export default function App () {
  return (
    <div className='App'>
      <NavBar />
      <Routes>
        <Route path='/' element={<Home />} />
        <Route path='/register' element={<Register />} />
        <Route path='/login' element={<Login />} />
        <Route path='/wallet' element={<Wallet />} />
        <Route path='/user/services' element={<Services />} />
      </Routes>
    </div>
  )
}
