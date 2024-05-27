import { getRequest, postRequest, putRequest } from "./axios";
import apiUrls from "../data/api-urls";

export const fetchTransactionList = async (body) => {
    return getRequest(apiUrls.transaction.list, body)
}

export const validateTransferRecord = async (body) => {
    return postRequest(apiUrls.transaction.validate, body)
}

export const createTransaction = async (body) => {
    return postRequest(apiUrls.transaction.list, body)
}

export const updateTransaction = async (id, body) => {
    return putRequest(apiUrls.transaction.detail(id), body)
}