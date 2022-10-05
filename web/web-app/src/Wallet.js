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

function SimpleDialog({ onClose, selectedValue, open, title }) {
  const handleClose = () => {
    onClose(selectedValue);
  };

  const handleListItemClick = (value) => {
    // onClose(value);
  };
  return (
    <Dialog onClose={handleClose} open={open}>
      <DialogTitle>{title}</DialogTitle>
      <List sx={{ pt: 0 }}>
        {services.map((service) => (
          <ListItem button onClick={() => handleListItemClick(service)} key={service}>
            <ListItemText primary={service} />
          </ListItem>
        ))}
      </List>
    </Dialog>
  );
}
export function Wallet() {
  const [open, setOpen] = React.useState(false);


  const [cards, setCards] = React.useState([{
    actionName: "J'update une de mes playlists",
    actionService: "Spotify",
    reactionName: "Un lien vers la playlist est envoyé",
    reactionService: "Spotify"
  }]);


  const handleNewAREA = () => {
    setOpen(true);
  };
  const handleClose = () => {
    setOpen(false);
  };

  return (
    <React.Fragment>
      <RollingCarousel />
      <Button size="small" onClick={handleNewAREA}> nouvelle AREA </Button>
      <AREACard cards={cards} />
      <SimpleDialog
        open={open}
        onClose={handleClose}
        title={"Nouvelle AREA"}
      />
    </React.Fragment>
  );

}
export default Wallet