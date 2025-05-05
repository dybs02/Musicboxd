import { emptyReviewAlbum, type ReviewAlbumType } from "./common";
import { emptyUser, type UserType } from "./user";

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
  itemType: string;
  title: string;
  description: string;
  userId: string;
  comments: CommentType[];
  album: ReviewAlbumType;
}

const emptyReview: ReviewType = {
  _id: "",
  value: 0,
  itemId: "",
  itemType: "",
  title: "",
  description: "",
  userId: "",
  comments: [emptyComment],
  album: emptyReviewAlbum,
}

export type {
  ReviewType,
  ReviewAlbumType,
  CommentType,
};

export { 
  emptyReview,
  emptyReviewAlbum,
  emptyComment,
};