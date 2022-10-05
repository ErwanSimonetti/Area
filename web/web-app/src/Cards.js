import * as React from 'react';
import Card from '@mui/material/Card';
import CardMedia from '@mui/material/CardMedia';

import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import Grid from '@mui/material/Grid';

// export function BasicCard(ActionName, ActionPlatform, ReactionName, ReactionPlatform) {
//   return (
//     <Card variant='outlined' sx={{ minWidth: 275 }}>
//       <CardContent>
//         <Typography variant="h5" component="div">
//           {ActionName}
//         </Typography>
//         <Typography sx={{ mb: 1.5 }} color="text.secondary">
//           {ActionPlatform}
//         </Typography>
//         <Typography variant="h5" component="div">
//           {ReactionName}
//         </Typography>
//         <Typography sx={{ mb: 1.5 }} color="text.secondary">
//           {ReactionPlatform}
//         </Typography>
//       </CardContent>
//       <CardActions>
//         <Button size="big">Ajouter cette AREAction</Button>
//       </CardActions>
//     </Card>
//   );
// }

export function AREACard({ cards }) {
  return (
    <Grid container spacing={4}>
      {cards.map((card) => (
        <Grid item key={card} xs={12} sm={6} md={4}>
          <Card
            sx={{ height: '100%', display: 'flex', flexDirection: 'column' }}
          >
            <CardMedia
              component="img"
              sx={{
                // 16:9
                pt: '56.25%',
              }}
              image="https://source.unsplash.com/random"
              alt="random"
            />
            <CardContent sx={{ flexGrow: 1 }}>
              <Typography gutterBottom variant="h5" component="h2">
                {card.actionName}
              </Typography>
              <Typography gutterBottom variant="h5" component="h2">
                {card.actionService}
              </Typography>
              <Typography>
                {card.reactionName}
              </Typography>
              <Typography>
                {card.reactionService}
              </Typography>
            </CardContent>
            <CardActions>
              <Button size="small">View</Button>
              <Button size="small">Edit</Button>
            </CardActions>
          </Card>
        </Grid>
      ))}
    </Grid>
  )
}

{/* <Grid container justifyContent="left" spacing={2}>
        <Grid item> */}
{/* <Paper elevation={16}>
        <CardContent>
          <Typography variant="h5" component="div">
            {props.ActionName}
          </Typography>
          <Typography variant="body2"   >
            {bull} {props.ReactionPlatform}
          </Typography>
          <Typography variant="h5" component="div">
            {props.ReactionName}
          </Typography>
          <Typography variant="body2">
            {bull} {props.ReactionPlatform}
          </Typography>
        </CardContent>
        <CardActions>
          <Button size="big">Ajouter cette AREAction</Button>
        </CardActions>
      </Paper> */}
{/* </Grid>
      </Grid> */}