import * as React from 'react';
import Box from '@mui/material/Box';
import Link from '@mui/material/Link';
import List from '@mui/material/List';
import Avatar from '@mui/material/Avatar';
import ListItem from '@mui/material/ListItem';
import LoginIcon from '@mui/icons-material/Login';
import ListItemText from '@mui/material/ListItemText';
import ListItemAvatar from '@mui/material/ListItemAvatar';
import AppRegistrationIcon from '@mui/icons-material/AppRegistration';
import AccountBalanceWalletIcon from '@mui/icons-material/AccountBalanceWallet';

function RollingCarousel() {
    return (
        <div className="Iam">
            <p>Bienvenue sur</p>
            <b>
                <div className="innerIam">
                    votre wallet<br />
                    vos actions<br />
                    vos réactions<br />
                    vos services<br />
                    vos AREActions
                </div>
            </b >
        </div >
    )
}

export default function FolderList() {
    return (
        <React.Fragment>
            <Box sx={{ display: 'flex', alignItems: 'center' }}>
                <RollingCarousel />
            </Box>
            <Box sx={{
                marginTop: 8,
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
            }}>
                <List sx={{ width: '100%', maxWidth: 360, bgcolor: 'background.paper' }}>
                    <ListItem>
                        <ListItemAvatar>
                            <Avatar>
                                <AccountBalanceWalletIcon />
                            </Avatar>
                        </ListItemAvatar>
                        <Link href="/wallet">
                            <ListItemText primary="Wallet" />
                        </Link>
                    </ListItem>
                    <ListItem>
                        <ListItemAvatar>
                            <Avatar>
                                <LoginIcon />
                            </Avatar>
                        </ListItemAvatar>
                        <Link href="/login">
                            <ListItemText primary="Login" />
                        </Link>
                    </ListItem>
                    <ListItem>
                        <ListItemAvatar>
                            <Avatar>
                                <AppRegistrationIcon />
                            </Avatar>
                        </ListItemAvatar>
                        <Link href="/register">
                            <ListItemText primary="Register" />
                        </Link>
                    </ListItem>
                </List>
            </Box>
        </React.Fragment>
    );
}
