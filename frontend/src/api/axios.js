import axios from 'axios';

export default axios.create({
    baseURL: "https://ibukiqrmark.com/api",
    withCredentials: true,
});


