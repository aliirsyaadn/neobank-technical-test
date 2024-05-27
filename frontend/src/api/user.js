import { postRequest } from "./axios";
import apiUrls from "../data/api-urls";

export const login = async (body) => {
    return postRequest(apiUrls.login, body)
}

export const register = async (body) => {
    return postRequest(apiUrls.register, body)
}

export const sendOTPCode = async (body) => {
    return postRequest(apiUrls.sendOTPCode, body)
}