/* eslint-disable */
import * as React from 'react';
import Button from '@mui/material/Button';
import axios from 'axios';


export default function Spotify () {
const spotifyAuth = (event) => {
    event.preventDefault()
    const headers = {
        'Content-Type': 'text/plain'
    };
    axios.get('http://localhost:8080/spotify/auth/url', {headers})
    .then(function (response) {
        location.href = response.data
    }).catch(function (error) {
        console.log(error)
    })
    //location.href = "https://discord.com/api/oauth2/authorize?client_id=1033382176785432656&redirect_uri=http%3A%2F%2Flocalhost%3A8080%2Fauth%2Fdiscord&response_type=code&scope=webhook.incoming&permissions=536870912"
}
return (
    <Button variant='outlined' onClick={spotifyAuth}>BJR</Button>
)}