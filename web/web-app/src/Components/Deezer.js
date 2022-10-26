/* eslint-disable */
import * as React from 'react';
import Button from '@mui/material/Button';

export default function Deezer() {
const DeezerAuth = () => {
    console.log("attempting Deezer auth...")
    location.href = "https://connect.deezer.com/oauth/auth.php?app_id=564742&redirect_uri=http://localhost:8080/deezer/auth&perms=basic_access,email,offline_access,manage_library,listening_history"
}
return (
    <Button variant='outlined' onClick={DeezerAuth}>Deezer auth</Button>
)}