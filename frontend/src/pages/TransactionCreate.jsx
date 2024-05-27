import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom';
import { Button } from 'react-bootstrap';
import Form from 'react-bootstrap/Form';
import { validateTransferRecord, createTransaction } from '../api/transaction';


const TransactionCreate = () => {
  const [instructionType, setInstructionType] = useState('IMMEDIATE')
  const [transferRecords, setTransferRecords] = useState([])
  const [totalTransferRecord, setTotalTransferRecord] = useState(0)
  const [totalErrorTransferRecord, setTotalErrorTransferRecord] = useState(0)
  const [totalTransferAmount, setTotalTransferAmount] = useState(0)
  const [transferDate, setTransferDate] = useState(null)
  const [transferTime, setTransferTime] = useState(null)
  const [validationState, setValidationState] = useState('')

  const navigation = useNavigate()

  useEffect(() => {
    if (!localStorage.getItem('token')) {
      navigate('/login')
    }
  }, [])

  const validate = async (file) => {
    try {
      const formData = new FormData()
      formData.append('file', file)
      const data = await validateTransferRecord(formData)
      console.log(data)
      if (data.message === 'success') {
        setTransferRecords(data.data.transfer_records)
        setTotalTransferRecord(data.data.total_transfer_record)
        setTotalTransferAmount(data.data.total_transfer_amount)
        setValidationState(data.message)
      } 
    } catch (err) {
      const error = err.response.data
      setTransferRecords(error.data.transfer_records)
      setTotalTransferRecord(error.data.total_transfer_record)
      setTotalErrorTransferRecord(error.data.errors.length)
      setTotalTransferAmount(error.data.total_transfer_amount)
      setValidationState(error.message)
    }
  }

  const submit = async () => {
    try {
      await createTransaction({
        instruction_type: instructionType,
        transfer_records: transferRecords,
        transfer_date: transferDate,
        total_transfer_amount: totalTransferAmount,
        total_transfer_record: totalTransferRecord,
      })
      navigation('/')
    } catch (err) {
      console.log(err)
    }
  }

  return (
    <div className='px-4'>
      <div className='font-bold'>Please enter transfer information</div>
      <Form.Group controlId="formFile" className="mb-3">
        <Form.Label>Choose your template</Form.Label>
        <Form.Control type="file" onChange={() => {
          validate(event.target.files[0])
        }} />

      </Form.Group>

      {validationState === 'success' && (
        <div className='bg-yellow-100 px-3 py-1 border border-yellow-400 rounded-lg'>
          After detection, there are <b>{totalTransferRecord} transfer</b> record, the total transfer amount is <b>Rp{totalTransferAmount}</b>
        </div>
      )}
      
      {validationState === 'success with errors' && (
        <div className='bg-red-100 px-3 py-1 border border-red-400 rounded-lg'>
          After detection, there are <b>{totalTransferRecord} transfer</b> record and <b>{totalErrorTransferRecord} record that are error</b>, Please reupload your template
        </div>
      )}

      <Form.Label>Instruction Type</Form.Label>
      <Form.Check // prettier-ignore
        type="radio"
        label="Immediate"
        value="IMMEDIATE"
        onClick={() => {
          setInstructionType('IMMEDIATE')
        }}
        checked={instructionType === 'IMMEDIATE'}
      />

      <Form.Check
        type="radio"
        label="Standing Instruction"
        value="STANDING_INSTRUCTION"
        onClick={() => {
          setInstructionType('STANDING_INSTRUCTION')
        }}
        checked={instructionType === 'STANDING_INSTRUCTION'}
      />

      {instructionType === 'STANDING_INSTRUCTION' && (
        <>
          <Form.Label>Transfer Date</Form.Label>
          <Form.Control
            type="date"
            value={transferDate}
            onChange={() => {
              setTransferDate(event.target.value)
            }}
          />
          
          <Form.Label>Transfer Time</Form.Label>
          <Form.Control
            type="time"
            value={transferTime}
            onChange={() => {
              setTransferTime(event.target.value)
            }}
          />
        </>
      )}

      <Button
        variant="primary"
        onClick={submit}
      >
        Next
      </Button>
    </div>
  )
}

export default TransactionCreate