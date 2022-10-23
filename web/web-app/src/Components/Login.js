import * as React from 'react'
import Avatar from '@mui/material/Avatar'
import Button from '@mui/material/Button'
import CssBaseline from '@mui/material/CssBaseline'
import TextField from '@mui/material/TextField'
import Link from '@mui/material/Link'
import Grid from '@mui/material/Grid'
import Box from '@mui/material/Box'
import LockOutlinedIcon from '@mui/icons-material/LockOutlined'
import Typography from '@mui/material/Typography'
import Container from '@mui/material/Container'
import { createTheme, ThemeProvider } from '@mui/material/styles'
import { GoogleLogin } from 'react-google-login'
import { GoogleLoginButton } from 'react-social-login-buttons'
import { gapi } from 'gapi-script'

// const clientidGithub = "ac56fad434a3a3c1561e"

const theme = createTheme()
// const navigate = useNavigate()
const clientId = '78828642227-b3tlfon89t2j66b2a81c60mu8oe45ijb.apps.googleusercontent.com'

export default function SignIn () {
  const [email, setEmail] = React.useState('')
  const [password, setPassword] = React.useState('')

  React.useEffect(() => {
    gapi.load('client:auth2', () => {
      gapi.auth2.init({ clientId: clientId })
    })
  }, [])

  const handleSubmit = (event) => {
    const data = new FormData(event.currentTarget)
    console.log(data)

    const [email, password] = [data.get('email'), data.get('password')]
    fetch('http://localhost:8080/login/', {
      method: 'POST',
      // We convert the React state to JSON and send it as the POST body
      // body: JSON.stringify(this.state)
      body: JSON.stringify({ email, password })
    }).then(function (response) {
      console.log(response)
      return response.json()
    })

    event.preventDefault()
  }

  const googleResponse = (e) => {
    console.log('LOGIN SUCCESS! Current user :', e.profileObj)
  }

  return (
    <ThemeProvider theme={theme}>
      <Container component="main" maxWidth="xs">
        <CssBaseline />
        <Box
          sx={{
            marginTop: 8,
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center'
          }}
        >
          <Avatar sx={{ m: 1, bgcolor: 'secondary.main' }}>
            <LockOutlinedIcon />
          </Avatar>
          <Typography component="h1" variant="h5">
            Sign in
          </Typography>
          {/* <Box component="form" onSubmit={handleSubmit} noValidate sx={{ mt: 1 }}> */}
          <TextField
            autoComplete='email'
            autoFocus
            value={email}
            onChange={e => setEmail(e.target.value)}
            placeholder="Email Address"
            type="text"
            name="email"
            required
          />
          <TextField
            autoComplete="current-password"
            value={password}
            onChange={e => setPassword(e.target.value)}
            placeholder="Password"
            type="password"
            name="password"
            required
          />
          <Button
            type="submit"
            onClick={handleSubmit}
            fullWidth
            variant="contained"
            sx={{ mt: 3, mb: 2 }}
          >
            Sign In
          </Button>
          <GoogleLogin
            clientId={clientId}
            render={renderProps => (
              <GoogleLoginButton onClick={renderProps.onClick} disabled={renderProps.disabled} />
            )}
            buttonText="Login"
            onSuccess={googleResponse}
            onFailure={googleResponse}
            cookiePolicy={'single_host_origin'}
          />
          <Grid container>
            <Grid item>
              <Link href="/register" variant="body2">
                {"Don't have an account? Sign Up"}
              </Link>
            </Grid>
          </Grid>
          {/* </Box> */}
        </Box>
      </Container>
    </ThemeProvider>
  )
}

// function AAA() {

//       return (
//         <ThemeProvider theme={theme}>
//           <Box
//             component="form"
//             sx={{
//               '& > :not(style)': { m: 1, width: '25ch' },
//             }}
//             noValidate
//             autoComplete="off"
//           >
//             {/* <form onSubmit={handleSubmit}> */}

//             <Button type="submit">Submit</Button>
//           </Box>
//           {/* </form> */}

//         </ThemeProvider>
//       )
//     }

// { /* <GitHubLogin
//                             clientId={clientidGithub}
//                             render={renderProps => (
//                                 <GithubLoginButton onClick={renderProps.onClick} disabled={renderProps.disabled} />
//                             )}
//                             onSuccess={(e) => { console.log("LOGIN SUCCESS! Current user :", e.profileObj) }}
//                             onFailure={(e) => { console.log("LOGIN FAILED! ", e) }}
//                         /> */ }
// { /* <Grid container> */ }
// { /* <Grid item xs>
//                                 <Link href="#" variant="body2">
//                                     Forgot password?
//                                 </Link>
//                             </Grid> */ }
// { /* <Grid item>
//                                 <Link href="/register" variant="body2">
//                                     {"Don't have an account? Sign Up"}
//                                 </Link>
//                             </Grid>

//                         </Grid>
//                     </Box>
//                 </Box>
//             </Container>
// ); */ }
