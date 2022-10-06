import { Button, Dialog, DialogTitle, List, ListItemText, ListItem } from '@mui/material';
import * as React from 'react';
import { AREACard } from './Cards';

function RollingCarousel() {
  return (
    <div class="Iam">
      <p>Bienvenue sur</p>
      <b>
        <div class="innerIam">
          votre wallet<br />
          vos actions<br />
          vos réactions<br />
          vos services<br />
          vos AREActions
        </div>
      </b>
    </div>
  )
}
const services = ["Spotify", "Twitter", "Discord", "Github"];
const actions = ["Un artiste poste un nouveau son", "J'ai ajouté une chanson à une playlist", "Un autre option pour laquelle j'ai pas d'idée"];

export function Wallet() {
  const [openDialog, setOpenDialog] = React.useState(false);
  const newCard = null;

  const [cards, setCards] = React.useState([{
    actionName: "J'update une de mes playlists",
    actionService: "Spotify",
    reactionName: "Un lien vers la playlist est envoyé",
    reactionService: "Spotify"
  }]);


  const handleNewAREA = () => {
    setOpenDialog(true);
  };
  const handleClose = () => {
    if (newCard !== null)
      cards.push(newCard);
    setOpenDialog(false);
  };

  return (
    <React.Fragment>
      <RollingCarousel />
      <Button size="small" onClick={() => { setOpenDialog(true); console.log("open") }}> nouvelle AREA </Button>
      <AREACard cards={cards} />
      <NewCardDialog onClose={() => setOpenDialog(false)} open={openDialog} />
    </React.Fragment >
  );
}

export function NewCardDialog(props) {

  const [actionService, setActionService] = React.useState('');
  const [action, setAction] = React.useState('');
  const [reactionService, setReactionService] = React.useState('');
  const [reaction, setReaction] = React.useState('');
  const [openServiceActionDialog, setOpenServiceActionDialog] = React.useState(false);
  const [openActionDialog, setOpenActionDialog] = React.useState(false);
  const [openServiceReactionDialog, setOpenServiceReactionDialog] = React.useState(false);
  const [openReactionDialog, setOpenReactionDialog] = React.useState(false);

  return (
    <React.Fragment>
      {/* Ce fdp de dialogue s'ouvre pas là jsp pourquoi ??? */}
      <Dialog onClose={() => props.setOpenDialog(false)} open={props.openDialog}>
        <DialogTitle>Nouvelle AREA</DialogTitle>
      </Dialog>
      {/* Service Action Pick */}
      <Dialog onClose={() => setOpenServiceActionDialog(false)} open={openServiceActionDialog}>
        <DialogTitle>Choisir un Service d'action</DialogTitle>
        <List sx={{ pt: 0 }}>
          {services.map((service) => (
            <ListItem button onClick={() => setActionService(service)} key={service}>
              <ListItemText primary={service} />
            </ListItem>
          ))}
        </List>
      </Dialog>
      {/* Action Pick */}
      <Dialog onClose={() => setOpenActionDialog(false)} open={openActionDialog}>
        <DialogTitle>Choisir une action</DialogTitle>
        <List sx={{ pt: 0 }}>
          {actions.map((action) => (
            <ListItem button onClick={() => setAction(action)} key={action}>
              <ListItemText primary={action} />
            </ListItem>
          ))}
        </List>
      </Dialog>
      {/* Service Reaction Pick */}
      <Dialog onClose={() => setOpenServiceReactionDialog(false)} open={openServiceReactionDialog}>
        <DialogTitle>Choisir un Service d'action</DialogTitle>
        <List sx={{ pt: 0 }}>
          {services.map((service) => (
            <ListItem button onClick={() => setReactionService(service)} key={service}>
              <ListItemText primary={service} />
            </ListItem>
          ))}
        </List>
      </Dialog>
    </React.Fragment >
  )
}

export default Wallet