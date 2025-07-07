import BuildingsSticky from '../components/BuildingsSticky'
import PageGradient    from '../components/PageGradient'
import CenterFrostDiv  from '../components/CenterFrostDiv'

import {Form, Button} from 'react-bootstrap'
import { useNavigate } from 'react-router'

import { apiCall } from '../api/useApi'
import { Link } from 'react-router-dom'

export default function()
{
    const navigate = useNavigate();
    const submit   = async (formData) => {
        const partial = {
            Email: formData.get('email'),
            Password: formData.get('password')
        }

        const {data, error} = await apiCall('/api/accounts/authenticate', {
            method: "POST",
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(partial),
        });
        
        if (!data?.error && !error)
        {
            window.sessionStorage.setItem('token', data.token)
            navigate('/')
        }
    }

    return (
        <>
            <BuildingsSticky />
            <PageGradient>
                <CenterFrostDiv>
                    <h1>Sign In</h1>
                    <p className='text-white'>Sign in to your MyPlace account.</p>
                    <hr />
                    <Form action={submit} className='text-white p-5 pt-1'>
                    <Form.Group className="mb-3" controlId="email">
                        <Form.Control name="email" className="p-2" type="email" placeholder="Enter email" />
                    </Form.Group>
                    <Form.Group className="mb-3" controlId="password">
                        <Form.Control name="password" className="p-2" type="password" placeholder="Password" />
                    </Form.Group>
                    <Button className="w-100 rounded m-2" type="submit">
                        Submit
                    </Button>
                    <Link to="/accounts/create">
                        <Button className="text-white w-100 rounded m-2">
                            Or Create an Account
                        </Button>
                    </Link>
                    </Form>
                </CenterFrostDiv>
            </PageGradient>
        </>
    )
}