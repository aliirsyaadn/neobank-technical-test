import React, { useState } from 'react'
import { useNavigate } from "react-router-dom";
import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import Toast from 'react-bootstrap/Toast';

import { login } from '../api/user';

const Login = () => {

	const [corpAccNo, setCorpAccNo] = useState('')
	const [userId, setUserId] = useState('')
	const [loginPassword, setLoginPassword] = useState('')
	const [visibleToast, setVisibleToast] = useState(false)
	
	const navigate = useNavigate();

	const Login = async () => {
		try {
			const { data } = await login({
				corporate_account_no: corpAccNo,
				id: userId,
				password: loginPassword
			})
			// redirect to homepage
			localStorage.setItem('token', data.token)
			localStorage.setItem('role', data.role)
			navigate('/')

		} catch (error) {
			setVisibleToast(true)
		}
	}

	return (
		<div className='flex flex-col gap-10 justify-center items-center w-screen h-screen'>
			<span>
				Login
			</span>
			<div>
				<Form.Label>Corporate Account No.</Form.Label>
				<Form.Control
					type="text"
					placeholder="Corporate Account No."
					value={corpAccNo}
					onInput={() => {
						setCorpAccNo(event.target.value)
					}}
				/>
				
				<Form.Label>User ID</Form.Label>
				<Form.Control
					type="text"
					placeholder="User ID"
					value={userId}
					onInput={() => {
						setUserId(event.target.value)
					}}
				/>
				
				<Form.Label>Login password</Form.Label>
				<Form.Control
					type="password"
					placeholder="Login password"
					value={loginPassword}
					onInput={() => {
						setLoginPassword(event.target.value)
					}}
				/>
				<Button
					variant="primary"
					onClick={Login}
				>
					Login
				</Button>

			</div>
			{visibleToast && (
			<Toast >
				<Toast.Body>Login error!</Toast.Body>
			</Toast>	
			)}
		</div>
  )
}

export default Login