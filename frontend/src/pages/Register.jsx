import React, { useState } from 'react'
import { useNavigate } from "react-router-dom";
import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import Toast from 'react-bootstrap/Toast';
import InputGroup from 'react-bootstrap/InputGroup';


import {
    login,
    register,
    sendOTPCode,
} from '../api/user';



const Login = () => {

	const [corpAccNo, setCorpAccNo] = useState('')
	const [corpName, setCorpName] = useState('')
	const [userId, setUserId] = useState('')
	const [userName, setUserName] = useState('')
	const [role, setRole] = useState('')
    const [phoneNo, setPhoneNo] = useState('')
    const [email, setEmail] = useState('')
    const [verificationCode, setVerificationCode] = useState('')
	const [loginPassword, setLoginPassword] = useState('')
	const [visibleToast, setVisibleToast] = useState(false)
	
	const navigate = useNavigate();

	const Register = async () => {
		try {
			await register({
				id: userId,
                corporate_account_no: corpAccNo,
                corporate_name: corpName,
                name: userName,
                role: role,
                phone: '+62' + phoneNo,
                password: loginPassword,
                email: email,
                otp: verificationCode
			})
			// redirect to homepage
			navigate('/login')

		} catch (error) {
			setVisibleToast(true)
		}
	}

    const SendOTPCode = async () => {
        try {
            await sendOTPCode({
                email: email
            })
            alert('OTP Code sent to your email')

        } catch (error) {
            setVisibleToast(true)
        }
    }

	return (
		<div className='flex flex-col gap-10 justify-center items-center w-screen h-screen'>
			<span>
				Register
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
				
                <Form.Label>Corporate Name</Form.Label>
				<Form.Control
					type="text"
					placeholder="Corporate Name"
					value={corpName}
					onInput={() => {
						setCorpName(event.target.value)
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

                <Form.Label>User Name</Form.Label>
                <Form.Control
                    type="text"
                    placeholder="User Name"
                    value={userName}
                    onInput={() => {
                        setUserName(event.target.value)
                    }}
                />

                <Form.Label>Role</Form.Label>
                <Form.Select
                    onChange={() => {
                        setRole(event.target.value)
                    }}
                >
                    <option value="APPROVER">Approver</option>
                    <option value="MAKER">Maker</option>
                </Form.Select>

                <Form.Label>Phone No</Form.Label>
                <InputGroup className="mb-3">
                    <InputGroup.Text id="basic-addon1">+62</InputGroup.Text>
                    <Form.Control
                        type="text"
                        placeholder="Phone No"
                        value={phoneNo}
                        onInput={() => {
                            setPhoneNo(event.target.value)
                        }}
                    />
                </InputGroup>
                

                <Form.Label>Email</Form.Label>
                <Form.Control
                    type="text"
                    placeholder="Email"
                    value={email}
                    onInput={() => {
                        setEmail(event.target.value)
                    }}
                />

                <Form.Label>Verifiication Code</Form.Label>
                <InputGroup className="mb-3">
                    <Form.Control
                        type="text"
                        placeholder="Verifiication Code"
                        value={verificationCode}
                        onInput={() => {
                            setVerificationCode(event.target.value)
                        }}
                    />
                    <InputGroup.Text>
                        <Button
                            disabled={email === ''}
                            variant="secondary"
                            onClick={SendOTPCode}
                        >
                            Send OTP Code
                        </Button>
                    </InputGroup.Text>
                </InputGroup>
                
				
				<Form.Label>Password</Form.Label>
				<Form.Control
					type="password"
					placeholder="Password"
					value={loginPassword}
					onInput={() => {
						setLoginPassword(event.target.value)
					}}
				/>
				<Button
					variant="primary"
					onClick={Register}
				>
					Register
				</Button>

			</div>
			{visibleToast && (
			<Toast >
				<Toast.Body>Register error!</Toast.Body>
			</Toast>	
			)}
		</div>
  )
}

export default Login