import * as React from 'react';
import ImageList from '@mui/material/ImageList';
import ImageListItem from '@mui/material/ImageListItem';
import ImageListItemBar from '@mui/material/ImageListItemBar';
import logoGitHub from './image/github.png'

export function TitlebarBelowImageList() {
    return (
    <ImageList sx={{ width: 500, height: 450 }}>
        {itemData.map((item) => (
        <ImageListItem key={item.img}>
          <img
            src={item.img}
            alt={item.title}
            loading="lazy"
          />
          <ImageListItemBar
            title={item.title}
            position="below"
          />
        </ImageListItem>
      ))}
    </ImageList>
  );
}

const itemData = [
    {
    img: {logoGitHub},
    title: 'Breakfast',
    author: '@bkristastucchio',
    },
];