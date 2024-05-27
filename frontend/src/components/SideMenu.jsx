import React from 'react'
import { useNavigate } from 'react-router-dom'

const SideMenu = () => {
    const navigate = useNavigate();

    return (
        <div className='w-52 h-screen bg-black text-white p-4'>
            <div className='mb-8'>BNC</div>
            <div className='mb-2 cursor-pointer' onClick={() => navigate('/')}>Home</div>
            {localStorage.getItem('role') === 'MAKER' && 
                <div className='mb-2 cursor-pointer' onClick={() => navigate('/create-transaction')}>Create Transaction</div>
            }
            <div className='cursor-pointer' onClick={() => navigate('/transfer-list')}>Transfer List</div>
        </div>
    )
}

export default SideMenu