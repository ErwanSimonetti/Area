import { Box, Button, Grid, TextField, Paper } from '@mui/material';

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
                    <Item>
                        <GitHubLogin
                            clientId={clientidGithub}
                            render={renderProps => (
                                <FacebookLoginButton onClick={renderProps.onClick} disabled={renderProps.disabled} />
                            )}
                            onSuccess={(e) => { console.log("LOGIN SUCCESS! Current user :", e.profileObj) }}
                            onFailure={(e) => { console.log("LOGIN FAILED! ", e); }}
                        />
                    </Item>
                </Grid>
                <Grid item>

                    {/* <Item>
                        <FacebookLogin
                            appId={clientidFacebook}
                            autoLoad={false}
                            fields="name,email,picture"
                            onClick={() => { console.log("clicked") }}
                            callback={() => { console.log("callback ?") }} />
                    </Item> */}
                </Grid>
                <Grid item >
                    <GithubLoginButton />
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