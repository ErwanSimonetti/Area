/* eslint-disable */
import * as React from 'react'
import { styled } from '@mui/material/styles'
import AREALogo from './Icons/AREALogo';
import NewAreaButton from './Icons/NewAreaButton';
import AccountIcon from './Icons/AccountIcon'
import AccountCircleIcon from '@mui/icons-material/AccountCircle';
import { Toolbar, Typography, AppBar, Button, Link, Grid } from '@mui/material'
import { Box } from '@mui/system'

const StyledAppBar = styled(AppBar)(({theme}) => ({
    backgroundColor:'#262626',
}));



export default function NavBar () {
    return (
        <StyledAppBar position='sticky'>
            <Toolbar>
                <Box className="navItem">
                    <Link href="/">
                        <Box >
                            <AREALogo/>
                        </Box>
                    </Link> 
                    <Box >
                        <AccountIcon/>
                    </Box>
                </Box>
            </Toolbar>
        </StyledAppBar>
    )
}
