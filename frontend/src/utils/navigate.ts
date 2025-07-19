import type { Router } from 'vue-router';


const navigateToUser = (router: Router, userId: string) => {
  router.push({ 
    name: 'user', 
    params: { id: userId }
  });
}


const navigateToAlbum = (router: Router, albumId: string) => {
  if (albumId === null) {
    // TODO navigate to error page
    console.error('User ID is required to navigate to album');
    return;
  }

  router.push({ 
    name: 'album',
    params: { 
      id: albumId,
    }
  });
}

const navigateToTrack = (router: Router, trackId: string) => {
  if (trackId === null) {
    // TODO navigate to error page
    console.error('User ID is required to navigate to album');
    return;
  }

  router.push({ 
    name: 'track',
    params: { 
      id: trackId,
    }
  });
}

const navigateToReview = (router: Router, reviewId: string) => {
  if (reviewId === null) {
    // TODO navigate to error page
    console.error('User ID is required to navigate to review');
    return;
  }

  router.push({ 
    name: 'review',
    params: { 
      id: reviewId
    }
  });
}

const navigateToChat = (router: Router, userId: string) => {
  if (userId === null) {
    // TODO navigate to error page
    console.error('User ID is required to navigate to chat');
    return;
  }

  router.push({ 
    name: 'chat',
    params: { 
      userId: userId
    }
  });
}

export {
  navigateToUser,
  navigateToAlbum,
  navigateToTrack,
  navigateToReview,
  navigateToChat,
}