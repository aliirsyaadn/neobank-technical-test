import axios from "axios";

const baseURL = "http://localhost:3000";

export const getRequest = async (url, query) => {
  const config = {
    ...query,
    headers: {
      "Authorization": "Bearer " + localStorage.getItem("token"),
    }
  }
  const response = await axios.get(baseURL + url, config);
  return response.data;
}

export const postRequest = async (url, data) => {
  const config = {
    headers: {
      "Authorization": "Bearer " + localStorage.getItem("token"),
    }
  }
  const response = await axios.post(baseURL + url, data, config);
  return response.data;
}


export const putRequest = async (url, data) => {
  const config = {
    headers: {
      "Authorization": "Bearer " + localStorage.getItem("token"),
    }
  }
  const response = await axios.put(baseURL + url, data, config);
  return response.data;
}