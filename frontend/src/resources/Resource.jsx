import {Route, useNavigate} from 'react-router-dom'

import {Link, useParams}    from 'react-router-dom'
import {Button}             from 'react-bootstrap'

import Form     from '../components/Form';
import Table    from '../components/Table'
import useApi, { apiCall }   from '../api/useApi';
import ItemTable from '../components/ItemTable';

export default class Resource
{
    constructor(resourceName, resourceModel) {
        this.resourceName  = resourceName;
        this.resourceModel = resourceModel;

        this.navigate = null;
    }

    // helper for title case text
    toTitleCase(str) {
        return str
            .split('_')
            .map(word => word.charAt(0).toUpperCase() + word.slice(1).toLowerCase())
            .join(' ');
    }

    // handle create form submission
    onCreate = async (formData) =>
    {
        // send update
        const {data, error} = await apiCall(`/api/${this.resourceName}`, {
            method: "POST",
            body: JSON.stringify(formData)
        });

        if (error)
            return console.error(error)

        // view change
        this.navigate(`/${this.resourceName}`)
    }

    // handle update form action
    onUpdate = async (id, formData) =>
    {
        // send update
        const {data, error} = await apiCall(`/api/${this.resourceName}/${id}`, {
            method: "PUT",
            body: JSON.stringify(formData)
        });

        if (error)
            return console.error(error)

        // view change
        this.navigate(`/${this.resourceName}`)
    }

    // handle delete form action
    // TODO
    onDelete = async (formdata) =>
    {
        throw new Error("NOT IMPLEMENTED")
    }

    // component for create page
    Create = () => {
        this.navigate = useNavigate();

        return (
            <>
                <h3>Create {this.toTitleCase(this.resourceName)}</h3>
                <hr />

                <Link to={`/${this.resourceName}`}>
                    <Button>
                        Button
                    </Button>
                </Link>
                <br />
                <br />

                <Form data={this.resourceModel} onSubmit={this.onCreate} />
            </>
        );
    }

    // component for read page
    Read = () => 
    {
        const {id} = useParams();
        const {data, loading, error} = useApi(`/api/${this.resourceName}/${id}`);

        if (loading)
            return <p>Loading...</p>

        return (
            <>
                <h3>Read {this.toTitleCase(this.resourceName)}</h3>
                <hr />
                <Link to={`/${this.resourceName}`}>
                    <Button>
                        Back
                    </Button>
                </Link>
                <br />
                <br />
                
                {
                    <Table data={data} />
                }
            </>
        )
    }

    // component for update page
    Update = () => {
        const { id }                 = useParams()
        const {data, loading, error} = useApi(`/api/${this.resourceName}/${id}`, {});
        this.navigate = useNavigate();

        if (loading)
            return <p>Loading...</p>

        return (
            <>
                <h3>Edit {this.toTitleCase(this.resourceName)}</h3>
                <hr />

                <Link to={`/${this.resourceName}`}>
                    <Button>
                        Back
                    </Button>
                </Link>
                <br />
                <br />
                
                {
                    <Form data={data} onSubmit={(data) => this.onUpdate(id, data) } />
                }
            </>
        );
    }

    // component for delete page
    Delete = () => {
        return <p>TODO</p>
    }

    // component for listing records
    List = () => {
        const {data, loading, error} = useApi(`/api/${this.resourceName}/all`, {});

        if (loading)
            return <p>Loading...</p>

        return (
            <>
                <h3>{this.toTitleCase(this.resourceName)} Directory</h3>
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

    // component for index page
    Index = () => {
        return <this.List />
    }

    // routes for this resource
    Routes = () => {
        return (
        <Route path={`/${this.resourceName}`}>
          <Route index                                     element={ <this.Index /> } />
          <Route path={`/${this.resourceName}/create/`}    element={ <this.Create /> } />
          <Route path={`/${this.resourceName}/view/:id`}   element={ <this.Read /> } />
          <Route path={`/${this.resourceName}/edit/:id`}   element={ <this.Update /> } />
          <Route path={`/${this.resourceName}/delete/:id`} element={ <this.Delete /> } />
        </Route>
        )
    }
}