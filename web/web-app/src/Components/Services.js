import { Typography, Box, Button, Card, CardContent, CardMedia, Grid, ButtonBase } from '@mui/material'
import * as React from 'react'
import axios from 'axios'
import githubImg from '../resources/github.png'
import spotifyImg from '../resources/spotify.png'
import discordImg from '../resources/discord.svg'

export default function Services () {
    const services = [{ name: 'Spotify', token: null, img: spotifyImg }, { name: 'Github', token: null, img: githubImg }, { name: 'Discord', token: null, img: discordImg }]
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

    const getSpotifyToken = () => {
        const headers = { 'Content-Type': 'text/plain' }
        axios.get('http://localhost:8080/spotify/auth/url', { headers })
        .then(function (response) {
            console.log(response.data)
            location.href = response.data
        }).catch(function (error) {
            console.log(error)
        })
    }

    const getGithubToken = () => {
        const headers = { 'Content-Type': 'text/plain' }
        axios.get('http://localhost:8080/github/auth/url', { headers })
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
            case 'Spotify':
                getSpotifyToken()
                break
            case 'Discord':
                getDiscordToken()
                break
            case 'Github':
                getGithubToken()
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
