import React, { useEffect, useState } from 'react'
import Table from 'react-bootstrap/Table';
import { Button } from 'react-bootstrap';

import { fetchTransactionList, updateTransaction } from '../api/transaction';

const TransactionList = () => {

  const [transactionList, setTransactionList] = useState([])

  useEffect(() => {
    getTransactionList()
  }, [])

  const getTransactionList = async () => {
    try {
      const { data } = await fetchTransactionList()
      setTransactionList(data)
    } catch (error) {
      alert(error)
    }
  }

  const updateApprovalTransaction = async (id, status) => {
    try {
      await updateTransaction(id, { status })
      getTransactionList()
    } catch (error) {
      alert(error)
    }
  }

  return (
    <div>
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>Reference No</th>
            <th>Total Transfer Amount (Rp)</th>
            <th>Total Transfer Record</th>
            <th>From Account No.</th>
            <th>Maker</th>
            <th>Transfer Date</th>
            <th>Status</th>
            <th>Operation</th>
          </tr>
        </thead>
        <tbody>
          { transactionList.map((transaction, index) => (
            <tr key={index}>
              <td>{transaction.reference_no}</td>
              <td>{transaction.total_transfer_amount}</td>
              <td>{transaction.total_transfer_record}</td>
              <td>{transaction.from_account_no}</td>
              <td>{transaction.maker_user_id}</td>
              <td>{transaction.transfer_date}</td>
              <td>{transaction.status}</td>
              <td>
                { localStorage.getItem('role') === 'APPROVER' && transaction.status === 'AWAITING_APPROVAL' && (
                  <>
                    <Button variant="danger" onClick={() => updateApprovalTransaction(transaction.reference_no, 'REJECTED')} className='mr-1'>
                      Reject
                    </Button>
                    <Button variant="success" onClick={() => updateApprovalTransaction(transaction.reference_no, 'APPROVED')}>
                      Approve
                    </Button>
                  </>
                )}
              </td>
            </tr>
          ))}

        </tbody>
      </Table>
    </div>
  )
}

export default TransactionList