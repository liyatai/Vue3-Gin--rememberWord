import axios from "axios";
const _axios = axios.create({
    baseURL:"http://localhost:8090"
})
export default _axios