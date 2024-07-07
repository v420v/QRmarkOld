import { useState, useEffect, useContext } from 'react';
import axios from "../../api/axios";
import SessionContext from '../../context/SessionProvider';

const UseFetch = (url, dep = []) => {
  const [data, setData] = useState(null);
  const [isPending, setIsPending] = useState(true);
  const [error, setError] = useState(null);

  const [session, setSession] = useContext(SessionContext);

    useEffect(() => {
        setIsPending(true);

        setTimeout(async() => {
            try {
                const response = await axios.get(url, {
                    headers: { 'Authorization': `Bearer ${session}`,'Content-Type': 'text/plain; charset=utf-8' },
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
