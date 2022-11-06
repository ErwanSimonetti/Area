import * as React from 'react'
import { styled } from '@mui/material/styles'
import AREALogo from './Icons/AREALogo'
import AccountIcon from './Icons/AccountIcon'
import { Toolbar, AppBar, Link, Button } from '@mui/material'
import { Box } from '@mui/system'
import axios from 'axios'
import LogoutIcon from './Icons/LogoutIcon'

const StyledAppBar = styled(AppBar)(({ theme }) => ({
    backgroundColor: '#262626'
}))

export default function NavBar () {
    const [disabledWallet, setDisabledWallet] = React.useState(false)

    React.useEffect(() => {
        console.log('je fais un truc')
        if (localStorage.getItem('loggedIn') !== true) {
            setDisabledWallet(true)
        }
    })

    const handleLogout = (event) => {
        event.preventDefault()
        axios.get('http://localhost:8080/logout/', { withCredentials: true })
        .then(function (response) {
            localStorage.setItem('loggedIn', false)
            setDisabledWallet(false)
        }).catch(function (error) {
            console.log(error)
        })
    }

    return (
        <StyledAppBar position='sticky'>
            <Toolbar>
                <Box style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                    <Box className="navItem">
                        <Button component={Link} to="/wallet" >
                            <Link href='/wallet'>
                                <Box >
                                    <AccountIcon/>
                                </Box>
                            </Link>
                        </Button>
                        <Link href="/">
                            <Box >
                                <AREALogo/>
                            </Box>
                        </Link>
                        <Button onClick={handleLogout}>
                            <Box >
                                <LogoutIcon/>
                            </Box>
                        </Button>
                    </Box>
                </Box>
            </Toolbar>
        </StyledAppBar>
    )
}
