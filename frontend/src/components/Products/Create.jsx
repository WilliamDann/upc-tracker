import MyForm from '../Form'
import Error from '../Error'
import useApi, { apiCall } from '../../api/useApi'

import {Button, Table} from 'react-bootstrap'
import {Link, useNavigate, useParams} from 'react-router-dom'

// item search page
export default function()
{
    const navigate               = useNavigate()

    const onSubmit = async (formData) => {
        // send update
        const {data, error} = await apiCall(`/api/products`, {
            method: "POST",
            body: JSON.stringify(formData)
        });

        if (error)
            return console.error(error)

        // view change
        navigate('/products/view/'+data.id)
    }

    return (
        <>
            <h3>Create Product</h3>
            <hr />

            <Link to="/products">
                <Button>
                    All Products
                </Button>
            </Link>
            <br />
            <br />
            
            {
                <MyForm data={{UPC: "", Name: "", Desc: ""}} onSubmit={onSubmit} />
            }
        </>
    );
}