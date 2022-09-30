import * as React from 'react';
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardMedia from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import { CardActionArea } from '@mui/material';

export function ActionAreaCard(path) {
  return (
    <Card sx={{ maxWidth: 345 }}>
      <CardActionArea>
        <body>
          <img src={require("./image/github.png")}/>
        </body>
        <CardContent>
          <Typography gutterBottom variant="h5" component="div">
            Lizard
          </Typography>
          <Typography variant="body2" color="text.secondary">
            Lizards are a widespread group of squamate reptiles, with over 6,000
            species, ranging across all continents except Antarctica
          </Typography>
        </CardContent>
      </CardActionArea>
    </Card>
  );
}

export function BasicCard(ActionName, ActionPlatform, ReactionName, ReactionPlatform, color) {
  return (
    <Card variant='outlined' sx={{ minWidth: 275 }} style={{flex:1, backgroundColor:color}}>
      <CardContent>
        <Typography variant="h5" component="div">
          {ActionName}
        </Typography>
        <Typography sx={{ mb: 1.5 }} color="text.secondary">
          {ActionPlatform}
        </Typography>
        <Typography variant="h5" component="div">
          {ReactionName}
        </Typography>
        <Typography sx={{ mb: 1.5 }} color="text.secondary">
          {ReactionPlatform}
        </Typography>
      </CardContent>
      <CardActions>
        <Button size="big">Ajouter cette AREAction</Button>
      </CardActions>
    </Card>
  );
}
