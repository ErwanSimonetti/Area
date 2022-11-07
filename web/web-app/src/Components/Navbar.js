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
export default function NavBar ({ setLoggedIn }) {
    React.useEffect(() => {
        console.log('je fais un truc')
        const cookie = document.cookie.indexOf('jwt')
        console.log(cookie)
        if (cookie !== -1) {
            setLoggedIn(true)
        }
    }, [])

    const handleLogout = (event) => {
        event.preventDefault()
        axios.get('http://localhost:8080/logout/', { withCredentials: true })
        .then(function (response) {
            localStorage.setItem('loggedIn', false)
        }).catch(function (error) {
            console.log(error)
        })
    }

    return (
        <StyledAppBar position='sticky'>
            <Box style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                <Box >
                    <Link href='/wallet'>
                        <AccountIcon/>
                    </Link>
                </Box>
                <Box >
                    <Link href="/">
                        <AREALogo/>
                    </Link>
                </Box>
                <Box >
                    <Button onClick={handleLogout}>
                        <LogoutIcon/>
                    </Button>
                </Box>
            </Box>
        </StyledAppBar>
    )
}
