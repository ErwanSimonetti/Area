import * as React from 'react'
import propTypes from 'prop-types'
import { Button, Box, Dialog, Grid, DialogTitle, List, ListItemText, ListItem, FormControlLabel, FormGroup, Checkbox } from '@mui/material'
import { AREACard } from './Cards'
import './../App.css'
import NewAreaButton from './Icons/NewAreaButton'
import { createTheme, ThemeProvider, Typography } from '@material-ui/core'
import axios from 'axios'

const services = ['Spotify', 'Twitter', 'Discord', 'Github']
const actions = ['Un artiste poste un nouveau son', "J'ai ajouté une chanson à une playlist", "Un autre option pour laquelle j'ai pas d'idée"]

export default function Wallet () {
    const [openDialog, setOpenDialog] = React.useState(false)
    const [singleCard, setSingleCard] = React.useState(false)
    const [areaCards, setAreaCards] = React.useState([])
    const [newCard, setNewCard] = React.useState({
        ID: null,
        action: null,
        actionService: null,
        // actionTitle:
        reaction: null,
        // reactionTitle
        reactionService: null
    })
    const cards = []

    React.useEffect(() => {
        axios.get('http://localhost:8080/area/get', { withCredentials: true })
        .then(function (response) {
            const areas = response.data
            areas.forEach(area => {
                const formattedArea = {
                    ID: area.ID,
                    action: area.action_func,
                    actionService: area.action_service,
                    // actionTitle:
                    reaction: area.reaction_func,
                    // reactionTitle
                    reactionService: area.reaction_service
                }
                cards.push(formattedArea)
            })
            setAreaCards(cards)
        }).catch(function (error) {
            console.log(error)
        })
    }, [])

    const requestAREAS = (event) => {
        event.preventDefault()
    }

    const handleNewCard = () => {
        if (singleCard) {
            const addedCard = {
                action: newCard.action,
                actionService: newCard.actionService,
                // actionTitle:
                reaction: newCard.reaction,
                // reactionTitle
                reactionService: newCard.reactionService
            }
            areaCards.push(addedCard)
            setNewCard({
                ID: null,
                action: null,
                actionService: null,
                // actionTitle:
                reaction: null,
                // reactionTitle
                reactionService: null
            })
            setSingleCard(false)
        }
        setOpenDialog(false)
    }

    const handleOpenDialog = () => {
        setOpenDialog(true)
    }
    const theme = createTheme({
        typography: {
          fontFamily: ['Titan One', 'cursive'].join(',')
        }
      })
    return (
        <React.Fragment>
                <Box sx={{
                    marginTop: 5,
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center'
                }}>
                    <ThemeProvider theme={theme}>
                        <Typography variant='h2' gutterBottom> Mon Wallet</Typography>
                    </ThemeProvider>
                    <Button size="small" onClick={() => { setOpenDialog(true) }} className="newAreaButton">
                        <NewAreaButton/>
                    </Button>
                </Box>
                <AREACard cards={areaCards} />
                <NewCardDialog onClose={handleNewCard} setSingleCard={setSingleCard} singleCard={singleCard} open={openDialog} newCard={newCard} setNewCard={setNewCard} />
        </React.Fragment >
    )
}

function NewCardDialog ({ setNewCard, newCard, ...props }) {
    const [openServiceActionDialog, setOpenServiceActionDialog] = React.useState(false)
    const [openActionDialog, setOpenActionDialog] = React.useState(false)
    const [openServiceReactionDialog, setOpenServiceReactionDialog] = React.useState(false)
    const [openReactionDialog, setOpenReactionDialog] = React.useState(false)

    React.useEffect(() => {
        if (newCard.action != null && newCard.actionService != null && newCard.reaction != null && newCard.reactionService != null) {
            props.setSingleCard(true)
        }
    })
jhào
    return (
        <React.Fragment>
            <Dialog onClose={props.onClose} open={props.open}>
                <DialogTitle>Créer une nouvelle AREA :</DialogTitle>
                <FormGroup>
                    <FormControlLabel disabled control={<Checkbox checked={newCard.actionService !== null} />} label={<Button onClick={() => setOpenServiceActionDialog(true)}> {newCard.actionService ? newCard.actionService : "service d'action"}</Button>} />
                    <FormControlLabel disabled control={<Checkbox checked={newCard.action !== null} />} label={<Button onClick={() => setOpenActionDialog(true)}> {newCard.action ? newCard.action : 'action'}</Button>} />
                    <FormControlLabel disabled control={<Checkbox checked={newCard.reactionService !== null} />} label={<Button onClick={() => setOpenServiceReactionDialog(true)}> {newCard.reactionService ? newCard.reactionService : 'service de réaction'}</Button>} />
                    <FormControlLabel disabled control={<Checkbox checked={newCard.reaction !== null} />} label={<Button onClick={() => setOpenReactionDialog(true)}> {newCard.reaction ? newCard.reaction : 'réaction'}</Button>} />
                    <Button variant="outlined" disabled={!props.singleCard} onClick={() => { props.onClose(false); console.log('ici') }}>Valider</Button>
                </FormGroup>
            </Dialog>
            {/* Service Action Pick */}
            <Dialog onClose={() => setOpenServiceActionDialog(false)} open={openServiceActionDialog}>
                <DialogTitle>Choisir un Service d&apos;action</DialogTitle>
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
