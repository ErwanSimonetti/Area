/* eslint-disable */
import * as React from 'react';
import Card from '@mui/material/Card';
import Box from '@mui/material/Box';
import CardMedia from '@mui/material/CardMedia';

import { CardContent, Typography, Grid } from '@mui/material';

const twitterImg = "https://www.1min30.com/wp-content/uploads/2017/05/Embl%C3%A8me-Twitter.jpg"
const spotifyImg = "https://storage.googleapis.com/pr-newsroom-wp/1/2018/11/folder_920_201707260845-1.png"
const discordImg = "https://logo-marque.com/wp-content/uploads/2020/12/Discord-Logo.png"
const githubImg = "https://logos-marques.com/wp-content/uploads/2021/03/GitHub-Logo.png"

export function AREACard({ cards }) {
    return (
        <Grid container spacing={4}>
            {cards.map((card, index) => (
                < Grid item key={index} xs={12} sm={6} md={4} >
                    <Card
                        sx={{ height: '100%', display: 'flex', flexDirection: 'column' }}
                    >
                        <CardContent sx={{ flexGrow: 1 }}>
                            <Typography gutterBottom variant="h5" component="h2">
                                {card.action}
                            </Typography>
                            <Typography>
                                {card.actionService}
                            </Typography>
                            <Typography gutterBottom variant="h5" component="h2">
                                {card.reaction}
                            </Typography>
                            <Typography>
                                {card.reactionService}
                            </Typography>
                        </CardContent>
                        <Box sx={{ display: 'flex', flexDirection: 'column' }}>
                            <CardMedia
                                component="img"
                                image={(card.actionService === "Spotify") ? spotifyImg : (card.actionService === "Twitter") ? twitterImg : (card.actionService === "Discord") ? discordImg : (card.actionService === "Github") ? githubImg : null}
                                alt="serviceImg"
                            />
                            <CardMedia
                                component="img"
                                image={(card.reactionService === "Spotify") ? spotifyImg : (card.reactionService === "Twitter") ? twitterImg : (card.reactionService === "Discord") ? discordImg : (card.reactionService === "Github") ? githubImg : null}
                                alt="serviceImage"
                            />
                        </Box>
                    </Card>
                </Grid>
            ))
            }
        </Grid >
    )
}
