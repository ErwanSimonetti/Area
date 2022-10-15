import * as React from 'react';
import Button from '@mui/material/Button';
import CssBaseline from '@mui/material/CssBaseline';
import TextField from '@mui/material/TextField';
import Link from '@mui/material/Link';
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import GitHubLogin from 'react-login-github';
import { GoogleLogin } from 'react-google-login';
import { GoogleLoginButton, FacebookLoginButton, GithubLoginButton } from 'react-social-login-buttons';

const clientidGoogle = "78828642227-b3tlfon89t2j66b2a81c60mu8oe45ijb.apps.googleusercontent.com"
// const clientidGithub = "ac56fad434a3a3c1561e"

const theme = createTheme();


export default function SignIn() {

    const handleSubmit = (event) => {
        alert('A form was submitted: ');
        const data = new FormData(event.currentTarget);
        const [email, password] = [data.get('email'), data.get('password')]
        fetch('http://localhost:8080/login/', {
            method: 'POST',
            // We convert the React state to JSON and send it as the POST body
            // body: JSON.stringify(this.state)
            body: JSON.stringify({email, password })
        }).then(function (response) {
            console.log(response)
            return response.json();
        });

        event.preventDefault();
    }

    const [email, setEmail] = React.useState("");
    const [password, setPassword] = React.useState("");

    return (
        <ThemeProvider theme={theme}>
            <form onSubmit={handleSubmit}>
                <input
                    autoComplete='email'
                    autoFocus
                    value={email}
                    onChange={e => setEmail(e.target.value)}
                    placeholder="Email Address"
                    type="text"
                    name="email"
                    required
                />
                <input
                    autoComplete="current-password"
                    value={password}
                    onChange={e => setPassword(e.target.value)}
                    placeholder="Password"
                    type="password"
                    name="password"
                    required
                />
                <button type="submit">Submit</button>
            </form>
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
        </ThemeProvider>
    );
}
                        {/* <GitHubLogin
                            clientId={clientidGithub}
                            render={renderProps => (
                                <GithubLoginButton onClick={renderProps.onClick} disabled={renderProps.disabled} />
                            )}
                            onSuccess={(e) => { console.log("LOGIN SUCCESS! Current user :", e.profileObj) }}
                            onFailure={(e) => { console.log("LOGIN FAILED! ", e); }}
                        /> */}
                        {/* <Grid container> */}
                            {/* <Grid item xs>
                                <Link href="#" variant="body2">
                                    Forgot password?
                                </Link>
                            </Grid> */}
                            {/* <Grid item>
                                <Link href="/register" variant="body2">
                                    {"Don't have an account? Sign Up"}
                                </Link>
                            </Grid>

                        </Grid>
                    </Box>
                </Box>
            </Container>
); */}
