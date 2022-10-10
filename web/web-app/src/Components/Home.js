import * as React from 'react';
import { Box, Button, Grid, TextField, Paper } from '@mui/material';
import { styled } from '@mui/material/styles';
import { FacebookLoginButton, GoogleLoginButton, GithubLoginButton } from "react-social-login-buttons";

import { GoogleLogin } from 'react-google-login';
import LoginGithub from 'react-login-github';
// import FacebookLogin from 'react-facebook-login';

import NavBar from './Navbar';

const clientidGoogle = "78828642227-b3tlfon89t2j66b2a81c60mu8oe45ijb.apps.googleusercontent.com"
const clientidGithub = "3480f13c0e7f898c17d6"
const gitHubClienSecret = "e019c2e9ae5a8e2a49bcb0904f321f036d1796e2"
const clientidFacebook = "ac56fad434a3a3c1561e"

const Item = styled(Paper)(({ theme }) => ({
    backgroundColor: theme.palette.mode === 'dark' ? '#1A2027' : '#fff',
    ...theme.typography.body2,
    padding: theme.spacing(1),
    textAlign: 'center',
    color: theme.palette.text.secondary,
}));

function Login() {
    return (
        <Grid container style={{ display: 'flex', wrap: 'nowrap', flexGrow: 1 }}>
            <Grid container style={{ display: 'flex', flexDirection: 'column' }}>
                <Grid item>
                    <GoogleLogin
                        clientId={clientidGoogle}
                        render={renderProps => (
                            <GoogleLoginButton onClick={renderProps.onClick} disabled={renderProps.disabled} />
                        )}
                        buttonText="Login"
                        onSuccess={(e) => { console.log("LOGIN SUCCESS! Current user :", e.profileObj) }}
                        onFailure={(e) => { console.log("LOGIN FAILED! ", e); }}
                        cookiePolicy={'single_host_origin'}
                    />

                </Grid>
                <Grid item>
                    <LoginGithub
                        clientId={clientidGithub}
                        // render={renderProps => (
                        //     <GithubLoginButton onClick={renderProps.onClick} disabled={renderProps.disabled} />
                        // )}
                        onSuccess={(e) => { console.log("jaj") }}
                        onFailure={(e) => { console.log("jaj") }}
                    />
                </Grid>
                <Grid item>
                </Grid>
                <Grid item >
                </Grid>
            </Grid>
            <Grid container style={{ display: 'flex', flexDirection: 'column' }}>
                <Grid item>
                    <TextField id="username" label="nom d'utilisateur" />
                </Grid>
                <Grid item>
                    <TextField id="password" type="password" label="mot de passe" />
                </Grid>
                <Grid item>
                    <Button variant='contained'>Valider</Button>
                </Grid>
            </Grid>
        </Grid >
    )
}
export default function Home() {

    return (
        <React.Fragment>
            <NavBar />
            <Box>
                <Login />
            </Box>

        </React.Fragment>
    )
}