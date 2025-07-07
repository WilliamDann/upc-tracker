import {Container} from "react-bootstrap"
import { Link } from "react-router";

export default function() {
    return (
        <Container className='p-3 m-0'>
            <b>This page was not found!</b>
            <br />
            <Link to="/">Back Home</Link>
        </Container>
    )
}