import MyForm from '../Form'
import Error from '../Error'
import useApi, { apiCall } from '../../api/useApi'

import {Button, Table} from 'react-bootstrap'
import {Link, useNavigate, useParams} from 'react-router-dom'

// item search page
export default function()
{
    const navigate               = useNavigate()
    const { id }                 = useParams()
    const {data, loading, error} = useApi(`/api/products/${id}`, {});

    if (loading)
        return <p>Loading...</p>

    const onSubmit = async (formData) => {
        // send update
        const {data, error} = await apiCall(`/api/products/${id}`, {
            method: "PUT",
            body: JSON.stringify(formData)
        });

        if (error)
            return console.error(error)

        // view change
        navigate('/products/view/'+id)
    }

    return (
        <>
            <Error data={error}></Error>

            <h3>Product Directory</h3>
            <hr />

            <Link to="/products">
                <Button>
                    All Products
                </Button>
            </Link>
            <br />
            <br />
            
            {
                <MyForm data={data} onSubmit={onSubmit} />
            }
        </>
    );
}