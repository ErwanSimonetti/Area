import { BasicCard } from "./Cards";

const AppDisplay = (
  <div>
    <article>
      <h1>Gradient Text</h1>
    </article>
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