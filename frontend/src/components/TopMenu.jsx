import React from 'react'
import { useNavigate } from 'react-router-dom'

const TopMenu = () => {
  const navigate = useNavigate()

  return (
    <div className='w-full flex justify-end gap-3 p-4'>
        <div>User ID</div>
        <div className='cursor-pointer' onClick={() => {
          localStorage.removeItem('token')
          localStorage.removeItem('role')
          navigate('/login')
        }}>Logout</div>
    </div>
  )
}

export default TopMenu