import useApi from "../api/useApi";
import Resource from "./Resource";
import {Route, Link} from 'react-router-dom'
import {Button} from 'react-bootstrap'
import Table from '../components/Table'

export default class Account extends Resource {
    constructor()
    {
        super('accounts', { Email: "", Password: "", Name: "" })
    }

    // page for reading your own account
    My = () =>
    {
        const {data, loading, error} = useApi(`/api/accounts/my`);

        if (loading)
            return <p>Loading...</p>

        return (
            <>
                <h3>My Account</h3>
                <hr />
                <Link to="edit">
                    <Button>
                        Edit
                    </Button>
                </Link>
                <Link to="delete">
                    <Button>
                        Delete
                    </Button>
                </Link>
                <br />
                <br />
                
                <Table data={data} />
            </>
        )
    }

    // register new routes
    Routes() {
        return super.Routes({ children: <Route path='my' element={<this.My />} /> })
    }
}