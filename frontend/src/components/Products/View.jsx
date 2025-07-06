import MyTable from '../Table'
import Error from '../Error'
import useApi from '../../api/useApi'

import {Button, Table} from 'react-bootstrap'
import {Link, useParams} from 'react-router-dom'

// item search page
export default function()
{
    const { id }                 = useParams()
    const {data, loading, error} = useApi(`/api/products/${id}`, {});

    if (loading)
        return <p>Loading...</p>

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
                <MyTable data={data} />
            }
        </>
    );
}