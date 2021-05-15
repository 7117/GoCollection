import axios from "axios";

const instance = axios.create({
    baseURL: 'http://127.0.0.1:8090',
});


export default instance
