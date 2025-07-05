import {Link} from 'react-router-dom'
import { Button } from 'react-bootstrap'
import {Camera} from 'react-bootstrap-icons'

// item search page
export default function()
{
    return (
        <>
            <h3>Scan an Item</h3>
            <p className="text-muted">Scan an item's Barcode.</p>
            <hr />

            <Button className="">
                <Camera />
                <span> Scan Now</span>
            </Button>
            <br />
            <b>or</b>
            <br />
            <Button className="">
                <Camera />
                <span> Upload Image</span>
            </Button>
            <br />
            <br />
            <Link to="/search">Seach with text →</Link><br />
            <Link to="/create">Enter manually →</Link>
        </>
    )
}