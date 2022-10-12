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
import { GoogleLogin } from 'react-google-login';
import { GoogleLoginButton } from 'react-social-login-buttons';

const clientidGoogle = "78828642227-b3tlfon89t2j66b2a81c60mu8oe45ijb.apps.googleusercontent.com"
// const clientidGithub = "ac56fad434a3a3c1561e"

const theme = createTheme();

export default function SignIn() {

    const handleSubmit = (event) => {
        event.preventDefault();
        const data = new FormData(event.currentTarget);

        console.log({
            email: data.get('email'),
            password: data.get('password'),
        });
    };

    return (
        <ThemeProvider theme={theme}>
            <Container component="main" maxWidth="xs">
                <CssBaseline />
                <Box
                    sx={{
                        marginTop: 8,
                        display: 'flex',
                        flexDirection: 'column',
                        alignItems: 'center',
                    }}
                >
                    <Typography component="h1" variant="h5">
                        Sign in
                    </Typography>
                    <Box component="form" onSubmit={handleSubmit} noValidate sx={{ mt: 1 }}>
                        <TextField
                            margin="normal"
                            required
                            fullWidth
                            id="email"
                            label="Email Address"
                            name="email"
                            autoComplete="email"
                            autoFocus
                        />
                        <TextField
                            margin="normal"
                            required
                            fullWidth
                            name="password"
                            label="Password"
                            type="password"
                            id="password"
                            autoComplete="current-password"
                        />
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
                        {/* <GitHubLogin
                            clientId={clientidGithub}
                            render={renderProps => (
                                <GithubLoginButton onClick={renderProps.onClick} disabled={renderProps.disabled} />
                            )}
                            onSuccess={(e) => { console.log("LOGIN SUCCESS! Current user :", e.profileObj) }}
                            onFailure={(e) => { console.log("LOGIN FAILED! ", e); }}
                        /> */}
                        <Button
                            onClick={handleSubmit}
                            // type="submit"
                            fullWidth
                            variant="contained"
                            sx={{ mt: 3, mb: 2 }}
                        >
                            Sign In
                        </Button>
                        <Grid container>
                            {/* <Grid item xs>
                                <Link href="#" variant="body2">
                                    Forgot password?
                                </Link>
                            </Grid> */}
                            <Grid item>
                                <Link href="/register" variant="body2">
                                    {"Don't have an account? Sign Up"}
                                </Link>
                            </Grid>

                        </Grid>
                    </Box>
                </Box>
            </Container>
        </ThemeProvider >
    );
}