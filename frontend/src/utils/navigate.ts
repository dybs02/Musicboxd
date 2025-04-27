import type { Router } from 'vue-router';


const navigateToUser = (router: Router, userId: string) => {
  router.push({ 
    name: 'user', 
    params: { id: userId }
  });
}


const navigateToAlbum = (router: Router, albumId: string, userId: string | null) => {
  if (userId === null) {
    // TODO navigate to error page
    console.error('User ID is required to navigate to album');
    return;
  }

  router.push({ 
    name: 'album',
    params: { 
      albumId: albumId,
      userId: userId
    }
  });
}

const navigateToTrack = (router: Router, trackId: string, userId: string | null) => {
  if (userId === null) {
    // TODO navigate to error page
    console.error('User ID is required to navigate to album');
    return;
  }

  router.push({ 
    name: 'track',
    params: { 
      trackId: trackId,
      userId: userId
    }
  });
}

export {
  navigateToUser,
  navigateToAlbum,
  navigateToTrack
}