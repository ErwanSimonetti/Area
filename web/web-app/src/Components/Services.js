import { Typography, Box, Button, Card, CardContent, CardMedia, Grid, ButtonBase } from '@mui/material'
import * as React from 'react'
import axios from 'axios'

const twitterImg = 'https://www.shareicon.net/data/256x256/2017/06/28/888044_logo_512x512.png'
const spotifyImg = 'https://raw.githubusercontent.com/iobroker-community-adapters/ioBroker.spotify-premium/HEAD/admin/spotify-premium.png'
const discordImg = 'https://www.svgrepo.com/show/353655/discord-icon.svg'
const githubImg = 'https://cdn.iconscout.com/icon/free/png-256/github-163-761603.png'
const deezerImg = 'https://cdn-1.webcatalog.io/catalog/deezer/deezer-icon-filled.png'

export default function Services () {
    const services = [{ name: 'Spotify', token: null, img: spotifyImg }, { name: 'Github', token: null, img: githubImg }, { name: 'Discord', token: null, img: discordImg }, { name: 'Deezer', token: null, img: deezerImg }, { name: 'Twitter', token: null, img: twitterImg }]
    return (
        <React.Fragment>
            <Box sx={{
                marginTop: 8,
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center'
            }}>
                <Typography variant='h2' gutterBottom>Services</Typography>
            </Box>
            <ServicesCard services={ services }/>
        </React.Fragment >
    )
}

function ServicesCard ({ services }) {
    const [service, setService] = React.useState(null)
    const [serviceToken, setServiceToken] = React.useState(null)

    const getDiscordToken = () => {
        const headers = { 'Content-Type': 'text/plain' }
        axios.get('http://localhost:8080/discord/auth/url', { headers })
        .then(function (response) {
            console.log(response.data)
            location.href = response.data
        }).catch(function (error) {
            console.log(error)
        })
    }
    const handleClick = (service) => {
        console.log('clicked ' + service.name)
        setService(service.name)
        switch (service.name) {
            case 'Discord':
                getDiscordToken()
                break
            default:
                break
        }
    }
    return (
        <Grid container spacing={4} sx={{ padding: '0 10%', width: '100%', marginLeft: '0px' }}>
            {services.map((service, index) => (
                < Grid item key={index} xs={12} sm={6} md={4} style={{ paddingRight: '32px' }}>
                    <ButtonBase onClick={e => handleClick(service)} style={{ width: '100%' }}>
                    <Card
                        sx={{ display: 'flex', flexDirection: 'column' }}
                    >
                        <CardContent
                        sx={{ flexGrow: 1 }}
                        style={ service.token === null ? { backgroundColor: '#d5dbe6' } : { backgroundColor: 'white' } }
                        >
                            <Typography gutterBottom variant="h5" component="h2" align='center'>
                                {service.name}
                            </Typography>
                        </CardContent>
                            <CardMedia
                                component="img"
                                image={ service.img }
                                alt="serviceImg"
                            />
                    </Card>
                        </ButtonBase>
                </Grid>
            ))
            }
        </Grid >
    )
}
