import { emptyExternalUrls, emptyImage, emptySimplifiedArtist, type ExternalUrlsType, type ImageType, type SimplifiedArtistType } from "./common";

type TracksType = {
  href: string;
  limit: number;
  next: string;
  offset: number;
  previous: string;
  total: number;
  items: TrackType[];
}

const emptyTracks: TracksType = {
  href: '',
  limit: 0,
  next: '',
  offset: 0,
  previous: '',
  total: 0,
  items: []
};


type TrackType = {
  album: AlbumType;
  artists: SimplifiedArtistType[];
  available_markets: string[];
  disc_number: number;
  duration_ms: number;
  explicit: boolean;
  external_urls: ExternalUrlsType;
  href: string;
  id: string;
  is_playable: boolean;
  name: string;
  popularity: number;
  preview_url: string | null;
  track_number: number;
  type: string;
  uri: string;
  is_local: boolean;

  averageRating: number;
  ratingCount: number;
}

const emptyTrack: TrackType = {
  album: {} as AlbumType,
  artists: [],
  available_markets: [],
  disc_number: 0,
  duration_ms: 0,
  explicit: false,
  external_urls: emptyExternalUrls,
  href: '',
  id: '',
  is_playable: false,
  name: '',
  popularity: 0,
  preview_url: null,
  track_number: 0,
  type: '',
  uri: '',
  is_local: false,

  averageRating: 0,
  ratingCount: 0
};


type AlbumsType = {
  href: string;
  limit: number;
  next: string;
  offset: number;
  previous: string;
  total: number;
  items: AlbumType[];
}

const emptyAlbums: AlbumsType = {
  href: '',
  limit: 0,
  next: '',
  offset: 0,
  previous: '',
  total: 0,
  items: []
};


type AlbumType = {
  album_type: string;
  total_tracks: number;
  available_markets: string[];
  external_urls: ExternalUrlsType;
  href: string;
  id: string;
  images: ImageType[];
  name: string;
  release_date: string;
  release_date_precision: string;
  type: string;
  uri: string;
  artists: SimplifiedArtistType[];
  tracks: TracksType;

  averageRating: number;
  ratingCount: number;
};

const emptyAlbum: AlbumType = {
  album_type: '',
  total_tracks: 0,
  available_markets: [""],
  external_urls: emptyExternalUrls,
  href: '',
  id: '',
  images: [emptyImage],
  name: '',
  release_date: '',
  release_date_precision: '',
  type: '',
  uri: '',
  artists: [emptySimplifiedArtist],
  tracks: emptyTracks,

  averageRating: 0,
  ratingCount: 0
};



export type {
  AlbumsType,
  AlbumType,
  TracksType,
  TrackType,
  SimplifiedArtistType,
}

export {
  emptyAlbums,
  emptyAlbum,
  emptyTracks,
  emptyTrack,
  emptySimplifiedArtist,
}