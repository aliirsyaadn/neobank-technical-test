import { sendOTPCode } from "../api/user";

export default {
    login: '/user/login',
    register: '/user/register',
    sendOTPCode: '/user/send-otp',
    transaction: {
        list: '/transaction',
        validate: '/transaction/validate-transfer-record',
        detail: (id) => `/transaction/${id}`,
    }
}