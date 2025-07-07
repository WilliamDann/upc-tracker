import useApi from "../api/useApi";
import Resource from "./Resource";
import {Route, Link, useNavigate} from 'react-router-dom'
import {Button} from 'react-bootstrap'
import Table from '../components/Table'
import Signin from "../pages/Signin";

import Error from '../components/Error'

export default class Account extends Resource {
    constructor()
    {
        super('accounts', { Email: "", Password: "", Name: "" })
    }

    // page for reading your own account
    My = () =>
    {
        const navigate                = useNavigate()
        const {data, loading, error} = useApi(`/api/accounts/my`);

        if (loading)
            return <p>Loading...</p>
        if (error)
            navigate('/signin')

        return (
            <>
                <h3>My Account</h3>
                <hr />
                <Table data={data} />
                <br />
                <br />
                <Link className="p-1" to={`/accounts/edit/${data?.id}`}>
                    <Button>
                        Edit
                    </Button>
                </Link>
                <Link className="p-1" to={`/accounts/delete/${data?.id}`}>
                    <Button>
                        Delete
                    </Button>
                </Link>
            </>
        )
    }

    // register new routes
    Routes() {
        return super.Routes({ children: <>
            <Route path='my' element={<this.My />} />
        </> })
    }
}