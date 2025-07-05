import { Link } from "react-router"

export default function()
{
    return (
        <>
            <h3>MyPlace Item Tracker</h3>
            <hr />

            <h4>About</h4>
            <p>The MyPlace item tracker is designed to store information about items by their <a href="https://en.wikipedia.org/wiki/Universal_Product_Code">UPC Code</a>. You can scan items to enter / update them, or you can search existing items.</p>
        
            <h4>Actions</h4>
            <p>You can preform two major actions:</p>

            <Link to="search">Search for an item →</Link><br />
            <Link to="scan">Scan a UPC Code →</Link><br />

        </>
    )
}