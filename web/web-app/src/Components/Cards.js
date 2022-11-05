import * as React from 'react'
import Card from '@mui/material/Card'
import Box from '@mui/material/Box'
import CardMedia from '@mui/material/CardMedia'
import './style.css'
import { CardContent, Typography, Grid } from '@mui/material'

const commonStyles = {
    bgcolor: 'background.paper',
    borderColor: 'text.primary',
    m: 3,
    background: 'linear-gradient(180deg, #5CCCE2 50%, #3A3A3A 50%)',
    borderRadius: '30px'
}

export function AREACard ({ cards }) {
    return (
        <Box sx={{ display: 'flex', alignItems: 'center' }} className="cardContainer">
            <Grid container spacing={{ xs: 2, sm: 2, md: 2 }} className="cardContainer" sx={{ marginTop: '5px', margin: '10px' }}>
                {cards.map((card, index) => (
                    <Grid item key={index} xs={12} sm={3} md={2.5} sx={{ ...commonStyles }}>
                        <Typography gutterBottom variant="h5">
                            {card.action}
                        </Typography>
                        <Typography gutterBottom variant="h5" color={'white'}>
                            {card.reaction}
                        </Typography>
                    </Grid>
                ))
            }
        </Grid>
        </Box>
    )
}
