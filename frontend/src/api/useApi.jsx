import { useEffect, useState } from "react";

export async function apiCall(endpoint, opts = {})
{
    const baseApi = import.meta.env.VITE_BASE_API; 
    let data    = null;
    let error   = null;
    
    const token = window.sessionStorage.getItem("token")
    if (token)
    {
        if (!opts['haders'])
            opts['headers'] = {}
        opts.headers.Authorization = `Bearer ${token}`
    }

    console.log(opts)

    try {
        const response = await fetch(baseApi + endpoint, opts)
        data           = await response.json();;
    } catch (err) {
        error = err;
    }
    
    return {data, error};
}

export default function useApi(endpoint, opts)
{
    const [data, setData]        = useState(null);
    const [loading, setLoading]  = useState(true);
    const [error, setError]      = useState(null);
    
    const fetchData = async () => {
        setLoading(true);
        
        try {
            let {data, error} = await apiCall(endpoint, opts);

            if (error)
                return setError(error);
            setData(data);
        } catch (err) {
            setError(err);
        } finally {
            setLoading(false);
        }
    };
    
    useEffect(() => {
        (async () => { await fetchData() })();
    }, []);
    
    return {data, loading, error};
}