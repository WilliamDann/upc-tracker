import { Link }                  from "react-router-dom"
import { PersonFillGear, Camera, Box2, PinMap } from "react-bootstrap-icons"

export default function()
{
    return (
        <>
            <div className="d-flex flex-column flex-shrink-0 p-2 text-white bg-dark h-100" style={{width: '180px'}}>
                <Link to="/" className="d-flex text-white text-decoration-none">
                    <svg className="bi me-2" width="40" height="32"><use xlinkHref="#bootstrap"></use></svg>
                    <span className="fs-4">MyPlace</span>
                </Link>
                <br />
                <ul className="nav nav-pills flex-column mb-auto">


                <li>
                    <Link to="accounts" className="nav-link text-white">
                        <PersonFillGear />
                        <span> Accounts</span>
                    </Link>
                </li>
                <li>
                    <Link to="products" className="nav-link text-white">
                        <Box2 />
                        <span> Products</span>
                    </Link>
                </li>
                <li>
                    <Link to="places" className="nav-link text-white">
                        <PinMap />
                        <span> Places</span>
                    </Link>
                </li>
                <li>
                    <Link to="scan" className="nav-link text-white inline">
                       <Camera />
                        <span> Scan Item</span>
                    </Link>
                </li>
                </ul>
                <hr />

                <Link to="accounts/my" className="d-flex align-items-center text-white text-decoration-none">
                    <strong>My Account</strong>
                </Link>
            </div>
        </>
    )
}