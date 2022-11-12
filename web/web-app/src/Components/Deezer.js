/* eslint-disable */
import * as React from 'react';
import Button from '@mui/material/Button';
import axios from 'axios';


export default function Deezer () {
const deezerAuth = (event) => {
    event.preventDefault()
    const headers = {
        'Content-Type': 'text/plain'
    };
    axios.get('http://localhost:8080/deezer/auth/url', {headers})
    .then(function (response) {
        location.href = response.data
    }).catch(function (error) {
        console.log(error)
    })
}
return (
    <Button variant='outlined' onClick={deezerAuth}>BJR</Button>
)}