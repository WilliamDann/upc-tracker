import ItemTable from '../ItemTable'
import Error from '../Error'
import useApi from '../../api/useApi'

import {Button} from 'react-bootstrap'
import {Link} from 'react-router-dom'

// item search page
export default function()
{
    const {data, loading, error} = useApi('/api/products/all', {});

    if (loading)
        return <p>Loading...</p>

    return (
        <>
            <Error data={error}></Error>

            <h3>Product Directory</h3>
            <hr />

            <Link to="create">
                <Button>
                    Create
                </Button>
            </Link>
            <br />
            <br />
            
            {
                <ItemTable className="w-100" items={data} />
            }
        </>
    );
}