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

const navigateToReview = (router: Router, itemType: string, itemId: string, userId: string | null) => {
  if (userId === null) {
    // TODO navigate to error page
    console.error('User ID is required to navigate to review');
    return;
  }

  if (itemType !== 'album' && itemType !== 'track') {
    // TODO navigate to error page
    console.error('Item type must be either album or track');
    return;
  }

  if (itemType === 'album') {
    console.log('navigating to album');
    navigateToAlbum(router, itemId, userId);
    return;
  }

  if (itemType === 'track') {
    navigateToTrack(router, itemId, userId);
    return;
  }
}

export {
  navigateToUser,
  navigateToAlbum,
  navigateToTrack,
  navigateToReview,
}