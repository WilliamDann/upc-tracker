import {Link} from 'react-router-dom'
import {Button, Form} from 'react-bootstrap'
import { useState } from 'react'

import { apiCall } from '../api/useApi'
import ItemTable from '../components/ItemTable'
import Error from '../components/Error'

// item search page
export default function()
{
    const [searchData, setSearchData] = useState([]);
    const [error, setError]           = useState(null);

    // called when submit is pressed
    const submit = async (event) => {
        event.preventDefault();

        // set url params
        // const params = new URLSearchParams();
        // params.set('query', formData.get('query'));

        // send search query
        let {data, error} = await apiCall('/api/products/all', { });
        
        console.log(data)
        console.log(error)

        if (error)
            return setError(error);

        // update state with search results
        setSearchData(data);
    }

    return (
        <>
            <Error data={error}></Error>

            <h3>Search For an Item</h3>
            <p className="text-muted">Search for an item by parameters.</p>
            <hr />

            <Form onSubmit={submit}>
                <Form.Group className="mb-3" controlId="formBasicEmail">
                    <Form.Label>What are you looking for?</Form.Label>
                    <Form.Control name="query" type="text" placeholder="Enter your search..." />
                </Form.Group>
                
                <Link to="/scan" className='m-2'>Search by scanning item →</Link>
                <br />

                <Button className='m-2' variant="primary" type="submit">
                    Search
                </Button>
            </Form>
            <hr />

            {
                <ItemTable items={searchData} />
            }
        </>
    )
}