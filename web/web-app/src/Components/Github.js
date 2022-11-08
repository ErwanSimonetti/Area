/* eslint-disable */
import * as React from 'react';
import Button from '@mui/material/Button';
import axios from 'axios';


export default function Github () {
const githubAuth = (event) => {
    event.preventDefault()
    const headers = {
        'Content-Type': 'text/plain'
    };
    axios.get('http://localhost:8080/github/auth/url', {headers})
    .then(function (response) {
        location.href = response.data
    }).catch(function (error) {
        console.log(error)
    })
    //location.href = "https://github.com/api/oauth2/authorize?client_id=1033382176785432656&redirect_uri=http%3A%2F%2Flocalhost%3A8080%2Fauth%2Fgithub&response_type=code&scope=webhook.incoming&permissions=536870912"
}
return (
    <Button variant='outlined' onClick={githubAuth}>BJR</Button>
)}