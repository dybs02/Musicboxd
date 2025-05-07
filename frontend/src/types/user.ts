import { emptyExplicitContent, emptyExternalUrls, emptyFollowers, emptyImage, type ExplicitContentType, type ExternalUrlsType, type FollowersType, type ImageType } from "./common";
import { emptyReviewAlbum, type ReviewAlbumType } from "./review";


type FavouriteAlbumEntryType = {
  key: number;
  album: ReviewAlbumType;
}

const emptyFavouriteAlbumEntry: FavouriteAlbumEntryType = {
  key: -1,
  album: emptyReviewAlbum
};


type UserType = {
  _id: string;
  country: string;
  displayName: string;
  email: string;
  explicitContent: ExplicitContentType;
  externalUrls: ExternalUrlsType;
  followers: FollowersType;
  href: string;
  spotifyId: string;
  images: ImageType[];
  product: string;
  type: string;
  uri: string;
  role: string;
  favouriteAlbums: FavouriteAlbumEntryType[];
};

const emptyUser: UserType = {
  _id: "",
  country: "",
  displayName: "",
  email: "",
  explicitContent: emptyExplicitContent,
  externalUrls: emptyExternalUrls,
  followers: emptyFollowers,
  href: "",
  spotifyId: "",
  images: [emptyImage],
  product: "",
  type: "",
  uri: "",
  role: "",
  favouriteAlbums: [emptyFavouriteAlbumEntry],
};

export type {
  FavouriteAlbumEntryType,
  UserType
};

export {
  emptyFavouriteAlbumEntry,
  emptyUser
};