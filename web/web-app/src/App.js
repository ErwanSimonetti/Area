import { BasicCard, ActionAreaCard } from "./Cards";
import { TitlebarBelowImageList } from "./CardsList";
import logoGitHub from "./image/github.png"

const AppDisplay = (
  <div>
    {/* <article>
      <h1>Gradient Text</h1>
    </article> */}
    {/* <div>
      {BasicCard("Nouveau morceau de Mehro", "Spotify", "Liker la chason", "Spotify", 'blue')}
    </div>
    <div>
      {BasicCard("Tananai dépasse 4 000 000 écoutes", "Spotify", "Twitter un lien vers son profil Spotify", "Twitter", 'red')} 
    </div>
    <div>
      {BasicCard("J'update une de mes playlists", "Spotify", "Un lien vers la playlist est envoyé", "Discord", 'yellow')}
    </div>
    <div>
      {BasicCard("Like d'une story avec une musique", "Instagram", "Ajout de la musique dans ma playlist", "Deezer", 'green')}
    </div> */}
    <div>
      {ActionAreaCard()}
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