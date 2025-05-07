
type ExplicitContentType = {
  filterEnabled: boolean;
  filterLocked: boolean;
};

const emptyExplicitContent: ExplicitContentType = {
  filterEnabled: false,
  filterLocked: false
};


type ExternalUrlsType = {
  spotify: string;
};

const emptyExternalUrls: ExternalUrlsType = {
  spotify: ""
};


type FollowersType = {
  href: string;
  total: number;
};

const emptyFollowers: FollowersType = {
  href: "",
  total: 0
};


type ImageType = {
  url: string;
  height: number;
  width: number;
};

const emptyImage: ImageType = {
  url: "",
  height: 0,
  width: 0
};


type SimplifiedArtistType = {
  external_urls: ExternalUrlsType;
  href: string;
  id: string;
  name: string;
  type: string;
  uri: string;
}

const emptySimplifiedArtist: SimplifiedArtistType = {
  external_urls: emptyExternalUrls,
  href: '',
  id: '',
  name: '',
  type: '',
  uri: ''
};


type ReviewAlbumType = {
  albumId: string;
  name: string;
  images: ImageType[];
  artists: SimplifiedArtistType[];
}

const emptyReviewAlbum: ReviewAlbumType = {
  albumId: "",
  name: "",
  images: [emptyImage],
  artists: [emptySimplifiedArtist],
}



export type {
  ExplicitContentType,
  ExternalUrlsType,
  FollowersType,
  ImageType,
  SimplifiedArtistType,
  ReviewAlbumType,
};

export { 
  emptyExplicitContent,
  emptyExternalUrls,
  emptyFollowers,
  emptyImage,
  emptySimplifiedArtist,
  emptyReviewAlbum,
};