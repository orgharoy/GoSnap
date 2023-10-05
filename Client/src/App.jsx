import { useState } from 'react'
import {Route, Routes, useLocation} from 'react-router-dom';
import './App.css'
import Home from './pages/Home';
import Auth from './pages/Auth';

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
      <Routes>
        <Route path="/auth" element={<Auth/>} />  
        <Route path="/" element={<Home/>} />  
      </Routes>
    </>
  )
}

export default App
