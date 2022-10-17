import * as React from 'react'
import { Button, Box, Dialog, DialogTitle, List, ListItemText, ListItem, Typography, FormControlLabel, FormGroup, Checkbox } from '@mui/material'
import { AREACard } from './Cards'
import './../App.css'

const services = ['Spotify', 'Twitter', 'Discord', 'Github']
const actions = ['Un artiste poste un nouveau son', "J'ai ajouté une chanson à une playlist", "Un autre option pour laquelle j'ai pas d'idée"]

export function Wallet () {
  const [openDialog, setOpenDialog] = React.useState(false)
  const [cards, setCards] = React.useState([{
    action: "J'update une de mes playlists",
    actionService: 'Spotify',
    reaction: 'Un lien vers la playlist est envoyé',
    reactionService: 'Spotify'
  }])
  const [newCard, setNewCard] = React.useState({
    action: null,
    actionService: null,
    reaction: null,
    reactionService: null
  })

  const handleNewCard = () => {
    setCards([...cards, {
      action: newCard.action,
      actionService: newCard.actionService,
      reaction: newCard.reaction,
      reactionService: newCard.reactionService
    }])
    setNewCard({
      action: null,
      actionService: null,
      reaction: null,
      reactionService: null
    })

    setOpenDialog(false)
  }

  return (
        <React.Fragment>
            <Box sx={{
              marginTop: 8,
              display: 'flex',
              flexDirection: 'column',
              alignItems: 'center'
            }}>
                <Typography variant='h2' gutterBottom>Wallet</Typography>
            </Box>
            <Button size="small" onClick={() => { setOpenDialog(true) }}> nouvelle AREA </Button>
            <AREACard cards={cards} />
            <NewCardDialog onClose={handleNewCard} open={openDialog} newCard={newCard} setNewCard={setNewCard}>
            </NewCardDialog>
        </React.Fragment >
  )
}

export function NewCardDialog ({ setNewCard, newCard, ...props }) {
  const [openServiceActionDialog, setOpenServiceActionDialog] = React.useState(false)
  const [openActionDialog, setOpenActionDialog] = React.useState(false)
  const [openServiceReactionDialog, setOpenServiceReactionDialog] = React.useState(false)
  const [openReactionDialog, setOpenReactionDialog] = React.useState(false)

  return (
        <React.Fragment>
            <Dialog onClose={props.onClose} open={props.open}>
                <DialogTitle>Créer une nouvelle AREA :</DialogTitle>
                <FormGroup>
                    <FormControlLabel disabled control={<Checkbox checked={!!newCard.actionService} />} label={<Button onClick={() => setOpenServiceActionDialog(true)}> {newCard.actionService ? newCard.actionService : "service d'action"}</Button>} />
                    <FormControlLabel disabled control={<Checkbox checked={newCard.action != null} />} label={<Button onClick={() => setOpenActionDialog(true)}> {newCard.action ? newCard.action : 'action'}</Button>} />
                    <FormControlLabel disabled control={<Checkbox checked={newCard.reactionService != null} />} label={<Button onClick={() => setOpenServiceReactionDialog(true)}> {newCard.reactionService ? newCard.reactionService : 'service de réaction'}</Button>} />
                    <FormControlLabel disabled control={<Checkbox checked={newCard.reaction != null} />} label={<Button onClick={() => setOpenReactionDialog(true)}> {newCard.reaction ? newCard.reaction : 'réaction'}</Button>} />
                    <Button variant="outlined" onClick={() => { props.onClose(false) }}>Valider</Button>
                </FormGroup>
            </Dialog>
            {/* Service Action Pick */}
            <Dialog onClose={() => setOpenServiceActionDialog(false)} open={openServiceActionDialog}>
                <DialogTitle>Choisir un Service d'action</DialogTitle>
                <List sx={{ pt: 0 }}>
                    {services.map((service) => (
                        <ListItem button onClick={() => { setNewCard({ ...newCard, actionService: service }); setOpenServiceActionDialog(false) }} key={service}>
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
                        <ListItem button onClick={() => { setNewCard({ ...newCard, action }); setOpenActionDialog(false) }} key={action}>
                            <ListItemText primary={action} />
                        </ListItem>
                    ))}
                </List>
            </Dialog >
            {/* Service Reaction Pick */}
            < Dialog onClose={() => setOpenServiceActionDialog(false)} open={openServiceReactionDialog} >
                <DialogTitle>Choisir un Service de réaction</DialogTitle>
                <List sx={{ pt: 0 }}>
                    {services.map((service) => (
                        <ListItem button onClick={() => { setNewCard({ ...newCard, reactionService: service }); setOpenServiceReactionDialog(false) }} key={service}>
                            <ListItemText primary={service} />
                        </ListItem>
                    ))}
                </List>
            </Dialog >
            {/* Reaction Pick */}
            < Dialog onClose={() => setOpenReactionDialog(false)} open={openReactionDialog} >
                <DialogTitle>Choisir une réaction</DialogTitle>
                <List sx={{ pt: 0 }}>
                    {actions.map((reaction) => (
                        <ListItem button onClick={() => { setNewCard({ ...newCard, reaction }); setOpenReactionDialog(false) }} key={reaction}>
                            <ListItemText primary={reaction} />
                        </ListItem>
                    ))}
                </List>
            </Dialog >
        </React.Fragment >
  )
}

export default Wallet
