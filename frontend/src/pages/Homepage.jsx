import React, { useEffect } from 'react'
import Button from 'react-bootstrap/Button';
import { useNavigate } from 'react-router-dom';
import TransactionList from '../components/TransactionList';

const Homepage = () => {
  const navigate = useNavigate();

  useEffect(() => {
    if (!localStorage.getItem('token')) {
      navigate('/login')
    }
  }, [])

  return (
    <div className='px-4'>
      <div className='w-full p-3 flex justify-between mb-3 bg-white rounded-lg'>
        Last login time: {localStorage.getItem('lastLoginTime')}
      </div>
      <div className='bg-white w-full p-6 rounded-lg'>
        <div className='border-gray-500 border-solid border rounded-lg p-3 mb-5'>
          <div className='font-semibold mb-2'>
            Transaction Overview
          </div>
          <div className='flex justify-between gap-3'>
            <div className='bg-gray-400 flex-1'>
              Awaiting Approval
            </div>
            <div className='bg-gray-400 flex-1'>
              Successfully
            </div>
            <div className='bg-gray-400 flex-1'>
              Rejected
            </div>
          </div>
        </div>
        <TransactionList />
      </div>
    </div>
  )
}

export default Homepage