import { Button, Dialog, DialogTitle, List, ListItemText, ListItem } from '@mui/material';
import { FormControl, FormControlLabel, FormGroup, Checkbox } from '@mui/material';
import * as React from 'react';
import { AREACard } from './Cards';

const services = ["Spotify", "Twitter", "Discord", "Github"];
const actions = ["Un artiste poste un nouveau son", "J'ai ajouté une chanson à une playlist", "Un autre option pour laquelle j'ai pas d'idée"];


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

export function Wallet() {
  const [openDialog, setOpenDialog] = React.useState(false);
  const [newCard, setNewCard] = React.useState({
    action: null,
    actionService: null,
    reaction: null,
    reactionService: null
  });

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
      <NewCardDialog onClose={() => setOpenDialog(false)} open={openDialog} newCard={newCard} setNewCard={setNewCard} />
    </React.Fragment >
  );
}

export function NewCardDialog(props) {

  const [actionService, setActionService] = React.useState("Spotify");
  const [action, setAction] = React.useState(null);
  const [reactionService, setReactionService] = React.useState(null);
  const [reaction, setReaction] = React.useState(null);
  const [openServiceActionDialog, setOpenServiceActionDialog] = React.useState(false);
  const [openActionDialog, setOpenActionDialog] = React.useState(false);
  const [openServiceReactionDialog, setOpenServiceReactionDialog] = React.useState(false);
  const [openReactionDialog, setOpenReactionDialog] = React.useState(false);

  let updatedValuetempCard = { action, actionService, reaction, reactionService };
  // Trouver comment modifier l'objet newCard dans le component enfant ici
  const handleNewCard = () => {
    // props.setNewCard(props.newCard => ({
    //   ...props.newCard,
    //   ...tempCard
    // });
  }
  // props.setOpenDialog(false);

  return (
    <React.Fragment>
      <Dialog onClose={props.onClose} open={props.open}>
        <DialogTitle>Créer une nouvelle AREA :</DialogTitle>
        <FormGroup>
          <FormControlLabel disabled control={<Checkbox checked={actionService ? true : false} />} label={<Button onClick={() => setOpenServiceActionDialog(true)}> {actionService ? actionService : "service d'action"}</Button>} />
          <FormControlLabel disabled control={<Checkbox checked={action != null} />} label={<Button onClick={() => setOpenActionDialog(true)}> {action ? action : "action"}</Button>} />
          <FormControlLabel disabled control={<Checkbox checked={reactionService != null} />} label={<Button onClick={() => setOpenServiceReactionDialog(true)}> {reactionService ? reactionService : "service de réaction"}</Button>} />
          <FormControlLabel disabled control={<Checkbox checked={reaction != null} />} label={<Button onClick={() => setOpenReactionDialog(true)}> {reaction ? reaction : "réaction"}</Button>} />
          <Button variant="outlined" onClick={handleNewCard()}>Valider</Button>
        </FormGroup>
      </Dialog>
      {/* Service Action Pick */}
      <Dialog onClose={() => setOpenServiceActionDialog(false)} open={openServiceActionDialog}>
        <DialogTitle>Choisir un Service d'action</DialogTitle>
        <List sx={{ pt: 0 }}>
          {services.map((service) => (
            <ListItem button onClick={() => { setActionService(service); setOpenServiceActionDialog(false) }} key={service}>
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
            <ListItem button onClick={() => { setAction(action); setOpenActionDialog(false) }} key={action}>
              <ListItemText primary={action} />
            </ListItem>
          ))}
        </List>
      </Dialog>
      {/* Service Reaction Pick */}
      <Dialog onClose={() => setOpenServiceActionDialog(false)} open={openServiceReactionDialog}>
        <DialogTitle>Choisir un Service de réaction</DialogTitle>
        <List sx={{ pt: 0 }}>
          {services.map((service) => (
            <ListItem button onClick={() => { setReactionService(service); setOpenServiceReactionDialog(false) }} key={service}>
              <ListItemText primary={service} />
            </ListItem>
          ))}
        </List>
      </Dialog>
      {/* Reaction Pick */}
      <Dialog onClose={() => setOpenReactionDialog(false)} open={openReactionDialog}>
        <DialogTitle>Choisir une réaction</DialogTitle>
        <List sx={{ pt: 0 }}>
          {actions.map((reaction) => (
            <ListItem button onClick={() => { setReaction(reaction); setOpenReactionDialog(false) }} key={reaction}>
              <ListItemText primary={reaction} />
            </ListItem>
          ))}
        </List>
      </Dialog>
    </React.Fragment >
  )
}

export default Wallet