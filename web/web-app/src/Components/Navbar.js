// import "./App.css"
import * as React from 'react'
import MenuIcon from '@mui/icons-material/Menu'
import { Toolbar, Typography, AppBar, Button, Link } from '@mui/material'

export default function NavBar () {
  return (
        <div>
            <AppBar position='relative'>
                <Toolbar>
                    <MenuIcon sx={{ mr: 2 }} />
                        <Link href="/">
                            <Button >
                                <Typography variant='h6' color='white' noWrap>
                                    AREA
                                </Typography>
                            </Button>
                        </Link>
                </Toolbar>
            </AppBar>
        </div>
  )
}
