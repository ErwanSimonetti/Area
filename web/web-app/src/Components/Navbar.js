/* eslint-disable */
import * as React from 'react'
import { styled } from '@mui/material/styles'
import MenuIcon from '@mui/icons-material/Menu'
import AccountCircleIcon from '@mui/icons-material/AccountCircle';
import { Toolbar, Typography, AppBar, Button, Link } from '@mui/material'
import { Box } from '@mui/system'

const StyledAppBar = styled(AppBar)(({theme}) => ({
    backgroundColor:'#3A3A3A',
    height:'190px'
}))

export default function NavBar () {
  return (
        <StyledAppBar position='relative'>
            <Toolbar>
                <Link href="/">
                    <Box style={{display:'flex', alignItems:'center'}}>
                        <Box>
                            <MenuIcon fontSize='large' color='white'/>
                        </Box>
                        <Box>
                            <Button >
                                <Typography variant='h2' color='white' noWrap>
                                    AREA
                                </Typography>
                            </Button>
                        </Box>
                        <Box>
                            <AccountCircleIcon fontSize='large' color='white'/>
                        </Box>
                    </Box>
                </Link>
            </Toolbar>
        </StyledAppBar>
    )
}
