import "./App.css"
import { AREACard, BasicCard } from "./Cards";
import Wallet from "./Wallet"
import Grid from '@mui/material/Grid';

// const Item = styled(Paper)(({ theme }) => ({
//   ...theme.typography.body2,
//   textAlign: 'center',
//   color: theme.palette.text.secondary,
//   height: 60,
//   lineHeight: '60px',
// }));
// const cards = { [key = 1, "actionName" = "J'update une de mes playlists"] [] }
const AppDisplay = (
  <div>
    <head>
      <link rel="stylesheet" type="text/css" href="style.css"></link>
      <link href='https://fonts.googleapis.com/css?family=Montserrat:400,700' rel='stylesheet' type='text/css'></link>
    </head>
    <br />
    <div>
      <Grid container rowSpacing={1} columnSpacing={{ xs: 1, sm: 2, md: 3 }}>
        <Wallet />;
        {/* <Grid item >
          <Item>
            <AREACard ActionName={"Nouveau morceau de Mehro"} ActionPlatform={"Spotify"} ReactionName={"Liker la chanson"} ReactionPlatform={"Spotify"} />
          </Item>
        </Grid>
        <Grid item >
          <Item>
            <AREACard ActionName={"Tananai dépasse 4 000 000 écoutes"} ActionPlatform={"Spotify"} ReactionName={"Twitter un lien vers son profil Spotify"} ReactionPlatform={"Twitter"} />
          </Item>
        </Grid> */}
        <Grid item >
          {/* <Item> */}
          <AREACard ActionName={"J'update une de mes playlists"} ActionPlatform={"Spotify"} ReactionName={"Un lien vers la playlist est envoyé"} ReactionPlatform={"Discord"} />
          {/* </Item> */}
        </Grid>
      </Grid>
    </div>
    {/* 
    <div>
      {BasicCard("J'update une de mes playlists", "Spotify", "Un lien vers la playlist est envoyé", "Discord")}
    </div>
    <div>
      {BasicCard("Like d'une story avec une musique", "Instagram", "Ajout de la musique dans ma playlist", "Deezer")}
    </div> */}
  </div>
)

export function App() {
  return (
    <div>
      {AppDisplay}
    </div>
  );
}

export default App