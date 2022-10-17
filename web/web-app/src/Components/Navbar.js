// import "./App.css"
import MenuIcon from '@mui/icons-material/Menu'
import { Toolbar, Typography, AppBar } from '@mui/material'

export default function NavBar () {
  return (
        <div>
            <AppBar position="relative">
                <Toolbar>
                    <MenuIcon sx={{ mr: 2 }} />
                    <Typography variant="h6" color="inherit" noWrap>
                        AREA
                    </Typography>
                </Toolbar>
            </AppBar>
        </div>
  )
}
