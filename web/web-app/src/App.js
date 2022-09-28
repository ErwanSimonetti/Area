import "./App.css"
import { BasicCard } from "./Cards";

const AppDisplay = (
  <div>
    <head>
      <link rel="stylesheet" type="text/css" href="style.css"></link>
      <link href='https://fonts.googleapis.com/css?family=Montserrat:400,700' rel='stylesheet' type='text/css'></link>
    </head>
    <body>
      <div class="Iam">
        <p>Bienvenue sur</p>
        <b>
          <div class="innerIam">
            votre wallet<br/> 
            vos actions<br/>
            vos réactions<br/>
            vos services<br/>
            vos AREActions
          </div>
        </b>
      </div>
    </body>
    <br/>
    <div>
      {BasicCard("Nouveau morceau de Mehro", "Spotify", "Liker la chason", "Spotify")}
    </div>
    <div>
      {BasicCard("Tananai dépasse 4 000 000 écoutes", "Spotify", "Twitter un lien vers son profil Spotify", "Twitter")} 
    </div>
    <div>
      {BasicCard("J'update une de mes playlists", "Spotify", "Un lien vers la playlist est envoyé", "Discord")}
    </div>
    <div>
      {BasicCard("Like d'une story avec une musique", "Instagram", "Ajout de la musique dans ma playlist", "Deezer")}
    </div>
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