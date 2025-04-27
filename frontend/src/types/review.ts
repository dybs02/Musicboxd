import { emptySimplifiedArtist, type SimplifiedArtistType } from "./spotify";
import { emptyImage, emptyUser, type ImageType, type UserType } from "./user";


type ReviewAlbum = {
  name: string;
  images: ImageType[];
  artists: SimplifiedArtistType[];
}

const emptyReviewAlbum: ReviewAlbum = {
  name: "",
  images: [emptyImage],
  artists: [emptySimplifiedArtist],
}


type CommentType = {
  _id: string;
  reviewId: string;
  userId: string;
  user: UserType;
  text: string;
  createdAt: string;
  updatedAt: string;
}

const emptyComment: CommentType = {
  _id: "",
  reviewId: "",
  userId: "",
  user: emptyUser,
  text: "",
  createdAt: "",
  updatedAt: "",
}


type ReviewType = {
  _id: string;
  value: number;
  itemId: string;
  title: string;
  description: string;
  userId: string;
  comments: CommentType[];
  album: ReviewAlbum;
}

const emptyReview: ReviewType = {
  _id: "",
  value: 0,
  itemId: "",
  title: "",
  description: "",
  userId: "",
  comments: [emptyComment],
  album: emptyReviewAlbum,
}

export type {
  ReviewType,
  ReviewAlbum,
  CommentType,
};

export { 
  emptyReview,
  emptyReviewAlbum,
  emptyComment,
};