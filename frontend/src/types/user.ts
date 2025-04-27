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
  role: ""
};

export type {
  ExplicitContentType,
  ExternalUrlsType,
  FollowersType,
  ImageType,
  UserType
};

export { 
  emptyExplicitContent,
  emptyExternalUrls,
  emptyFollowers,
  emptyImage,
  emptyUser
};