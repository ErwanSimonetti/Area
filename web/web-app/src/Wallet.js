import * as React from 'react';
import { AREACard } from './Cards';

export function Wallet() {
  const cards = [{
    actionName: "J'update une de mes playlists",
    actionService: "Spotify",
    reactionName: "Un lien vers la playlist est envoyé",
    reactionService: "Spotify"
  },
  {
    actionName: "J'update une de mes playlists",
    actionService: "Spotify",
    reactionName: "Un lien vers la playlist est envoyé",
    reactionService: "Spotify"
  },
  {
    actionName: "J'update une de mes playlists",
    actionService: "Spotify",
    reactionName: "Un lien vers la playlist est envoyé",
    reactionService: "Spotify"
  },
  {
    actionName: "J'update une de mes playlists",
    actionService: "Spotify",
    reactionName: "Un lien vers la playlist est envoyé",
    reactionService: "Spotify"
  }, {
    actionName: "J'update une de mes playlists",
    actionService: "Spotify",
    reactionName: "Un lien vers la playlist est envoyé",
    reactionService: "Spotify"
  },
  {
    actionName: "J'update une de mes playlists",
    actionService: "Spotify",
    reactionName: "Un lien vers la playlist est envoyé",
    reactionService: "Spotify"
  }];
  return (
    <React.Fragment>
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
      <AREACard cards={cards} />
    </React.Fragment>
  );

}
export default Wallet