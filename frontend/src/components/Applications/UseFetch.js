import { useState, useEffect } from 'react';
import axios from "../../api/axios";

const UseFetch = (url, dep = []) => {
  const [data, setData] = useState(null);
  const [isPending, setIsPending] = useState(true);
  const [error, setError] = useState(null);

    useEffect(() => {
        setIsPending(true);

        setTimeout(async() => {
            try {
                const response = await axios.get(url, {
                    headers: { 'Content-Type': 'text/plain; charset=utf-8' },
                    withCredentials: true,
                });
                setData(response.data);
            } catch (err) {
                setError(err);
            } finally {
                setIsPending(false);
            }
        }, 300);

    }, dep);

  return { data, isPending, error };
}
 
export default UseFetch;
