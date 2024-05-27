import React from 'react'
import router from './router'
import { RouterProvider } from "react-router-dom";
import 'bootstrap/dist/css/bootstrap.min.css';
import 'dotenv/config'


const App = () => {
  return (
    <div className='flex'>
      <RouterProvider router={router} />
    </div>
  )
}

export default App