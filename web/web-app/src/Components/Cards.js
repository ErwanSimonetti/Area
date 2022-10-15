import * as React from 'react';
import Card from '@mui/material/Card';
import CardMedia from '@mui/material/CardMedia';

import { CardContent, Typography, Grid } from '@mui/material';

export function AREACard({ cards }) {
    return (
        <Grid container spacing={4}>
            {cards.map((card) => (
                <Grid item key={card} xs={12} sm={6} md={4}>
                    <Card
                        sx={{ height: '100%', display: 'flex', flexDirection: 'column' }}
                    >
                        <CardContent sx={{ flexGrow: 1 }}>
                            <Typography gutterBottom variant="h5" component="h2">
                                {card.actionName}
                            </Typography>
                            <Typography>
                                {card.actionService}
                            </Typography>
                            <Typography gutterBottom variant="h5" component="h2">
                                {card.reactionName}
                            </Typography>
                            <Typography>
                                {card.reactionService}
                            </Typography>
                        </CardContent>
                        <CardMedia
                            component="img"
                            image="https://storage.googleapis.com/pr-newsroom-wp/1/2018/11/folder_920_201707260845-1.png"
                            alt="random"
                        />
                    </Card>
                </Grid>
            ))}
        </Grid>
    )
}
