import {Link} from 'react-router-dom'
import {Button, Form} from 'react-bootstrap'

// item search page
export default function()
{
    return (
        <>
            <h3>Search For an Item</h3>
            <p className="text-muted">Search for an item by parameters.</p>
            <hr />

            <Form>
                <Form.Group className="mb-3" controlId="formBasicEmail">
                    <Form.Label>What are you looking for?</Form.Label>
                    <Form.Control type="text" placeholder="Enter your search..." />
                </Form.Group>
                
                <Link to="/scan" className='m-2'>Search by scanning item →</Link>
                <br />

                <Button className='m-2' variant="primary" type="submit">
                    Search
                </Button>
            </Form>
        </>
    )
}