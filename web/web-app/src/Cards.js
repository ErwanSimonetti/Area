import * as React from 'react';
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';

export function BasicCard(ActionName, ActionPlatform, ReactionName, ReactionPlatform) {
  return (
    <Card variant='outlined' sx={{ minWidth: 275 }}>
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
