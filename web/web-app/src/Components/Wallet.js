/*eslint-disable*/
import * as React from 'react'
import propTypes from 'prop-types'
import { Button, Box, Dialog, Grid, DialogTitle, List, ListItemText, ListItem, FormControlLabel, FormGroup, Checkbox } from '@mui/material'
import { AREACard } from './Cards'
import './../App.css'
import NewAreaButton from './Icons/NewAreaButton'
import { createTheme, ThemeProvider, Typography } from '@material-ui/core'
import axios from 'axios'

export default function Wallet () {
    const [openDialog, setOpenDialog] = React.useState(false)
    const [singleCard, setSingleCard] = React.useState(false)
    const [areaCards, setAreaCards] = React.useState([])
    const [serviceArray, setServiceArray] = React.useState([])
    const [newCard, setNewCard] = React.useState({
        ID: null,
        action: null,
        actionService: null,
        reaction: null,
        reactionService: null
    })
    const cards = []
    const servicesData = []

    React.useEffect(() => {
        axios.get('http://localhost:8080/area/user/areas', { withCredentials: true })
        .then(function (response) {
            const areas = response.data
            areas.forEach(area => {
                const formattedArea = {
                    ID: area.ID,
                    action: area.action_func,
                    actionService: area.action_service,
                    reaction: area.reaction_func,
                    reactionService: area.reaction_service
                }
                cards.unshift(formattedArea)
            })
            setAreaCards(cards)
        }).catch(function (error) {
            console.log(error)
        })
    }, [])

    React.useEffect(() => {
        axios.get('http://localhost:8080/area/user/propositions', { withCredentials: true })
        .then(function (response) {
            console.log(response)
            const services = response.data
            services.forEach(service => {
                const actionsArray = []
                const reactionsArray = []
                service.actions.forEach((action) => {
                    actionsArray.push(action.description)
                })
                service.reactions.forEach((reaction) => {
                    reactionsArray.push(reaction.description)
                })
                const fService = {
                    service: service.name,
                    actions: actionsArray,
                    reactions: reactionsArray
                }
                servicesData.push(fService)
            })
            setServiceArray(servicesData)
        })
        .catch (function (error){
            console.log(error)
        })
    }, [])

    const handleNewCard = () => {
        if (singleCard) {
            const headers = {
                'Content-Type': 'text/plain'
            }
            const sentCard = {
                action: newCard.action,
                actionService: newCard.actionService,
                reaction: newCard.reaction,
                reactionService: newCard.reactionService
            }
            axios.post('http://localhost:8080/area/create', {
                action_service: sentCard.actionService,
                action_func: sentCard.action,
                action_func_params: 'action de B 3',
                reaction_service: sentCard.reactionService,
                reaction_func: sentCard.reaction,
                reaction_func_params: 'reaction de B 2'
            }, { headers: { 'Content-Type': 'text/plain' }, withCredentials: true })
            .then(function (response) {
                sentCard.ID = response.data
                window.location.reload(true)
            })
            .catch(function (error) {
                console.log(error)
            })
            setNewCard({
                ID: null,
                action: null,
                actionService: null,
                reaction: null,
                reactionService: null
            })
            setSingleCard(false)
        }
        setOpenDialog(false)
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
                    <Button size="small" onClick={ () => { setOpenDialog(true) } } className="newAreaButton">
                        <NewAreaButton/>
                    </Button>
                </Box>
                <AREACard cards={areaCards} />
                <NewCardDialog
                    onClose={handleNewCard}
                    setSingleCard={setSingleCard}
                    singleCard={singleCard}
                    open={openDialog}
                    newCard={newCard}
                    setNewCard={setNewCard}
                    serviceArray={serviceArray} />
        </React.Fragment >
    )
}

function NewCardDialog ({ setNewCard, newCard, serviceArray, ...props }) {
    const [openServiceActionDialog, setOpenServiceActionDialog] = React.useState(false)
    const [openActionDialog, setOpenActionDialog] = React.useState(false)
    const [openServiceReactionDialog, setOpenServiceReactionDialog] = React.useState(false)
    const [openReactionDialog, setOpenReactionDialog] = React.useState(false)
    const [currentService, setCurrentService] = React.useState({
        service: null,
        actions: null,
        reactions: null
    })
    React.useEffect(() => {
        if (newCard.action != null && newCard.actionService != null && newCard.reaction != null && newCard.reactionService != null) {
            props.setSingleCard(true)
        }
    })
    React.useEffect(() => {
        console.log(currentService)
    }, [currentService])
    
    const handleClickService = React.useCallback((service) => {
        setNewCard({ ...newCard, actionService: service.service })
        console.log(service.actions)
        setCurrentService(service)
        setOpenServiceActionDialog(false)
    }, [])

    return (
        <React.Fragment>
            <Dialog onClose={props.onClose} open={props.open}>
                <DialogTitle>Créer une nouvelle AREA :</DialogTitle>
                <FormGroup>
                    <FormControlLabel disabled control={<Checkbox checked={newCard.actionService !== null} />} label={<Button onClick={() => setOpenServiceActionDialog(true)}> {newCard.actionService ? newCard.actionService : "service d'action"}</Button>} />
                    <FormControlLabel disabled control={<Checkbox checked={newCard.action !== null} />} label={<Button onClick={() => setOpenActionDialog(true)}> {newCard.action ? newCard.action : 'action'}</Button>} />
                    {/* <FormControlLabel disabled control={<Checkbox checked={newCard.reactionService !== null} />} label={<Button onClick={() => setOpenServiceReactionDialog(true)}> {newCard.reactionService ? newCard.reactionService : 'service de réaction'}</Button>} /> */}
                    <FormControlLabel disabled control={<Checkbox checked={newCard.reaction !== null} />} label={<Button onClick={() => setOpenReactionDialog(true)}> {newCard.reaction ? newCard.reaction : 'réaction'}</Button>} />
                    <Button variant="outlined" disabled={!props.singleCard} onClick={() => { props.onClose(false) }}>Valider</Button>
                </FormGroup>
            </Dialog>
            {/* Service Action Pick */}
            <Dialog onClose={() => setOpenServiceActionDialog(false)} open={openServiceActionDialog}>
                <DialogTitle>Choisir un Service d&apos;action</DialogTitle>
                    <List sx={{ pt: 0 }}>
                    {serviceArray.map((service, index) => (
                        <ListItem button onClick={() => handleClickService(service) } key={index}>
                            <ListItemText primary={service.service} />
                    </ListItem>
                    ))}
                    </List>
            </Dialog>
            <Dialog onClose={() => setOpenActionDialog(false)} open={openActionDialog}>
                <DialogTitle>Choisir une action</DialogTitle>
                <List sx={{ pt: 0 }}>
                    { currentService?.actions && currentService.actions.map((action, index) => (
                        <ListItem button onClick={() => { setNewCard({ ...newCard, action }); setOpenActionDialog(false) }} key={index}>
                            <ListItemText primary={action} />
                        </ListItem>
                    ))}
                </List>
            </Dialog >
            <Dialog onClose={() => setOpenReactionDialog(false)} open={openReactionDialog}>
                <DialogTitle>Choisir une action</DialogTitle>
                <List sx={{ pt: 0 }}>
                    { currentService?.reactions && currentService.reactions.map((reaction, index) => (
                        <ListItem button onClick={() => { setNewCard({ ...newCard, reaction }); setOpenReactionDialog(false) }} key={index}>
                            <ListItemText primary={reaction} />
                        </ListItem>
                    ))}
                </List>
            </Dialog >
        </React.Fragment >
    )
}
