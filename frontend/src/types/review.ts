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
  user: UserType;
  createdAt: string;
  updatedAt: string;
  comments: CommentType[];
  album: ReviewAlbumType;
  likesCount: number;
  dislikesCount: number;
  userReaction: string;
}

const emptyReview: ReviewType = {
  _id: "",
  value: 0,
  itemId: "",
  itemType: "",
  title: "",
  description: "",
  userId: "",
  user: emptyUser,
  createdAt: "",
  updatedAt: "",
  comments: [emptyComment],
  album: emptyReviewAlbum,
  likesCount: 0,
  dislikesCount: 0,
  userReaction: "",
}

type RecentUserReviewsType = {
  totalReviews: number;
  totalPages: number;
  hasPreviousPage: boolean;
  hasNextPage: boolean;
  reviews: ReviewType[];
}

const emptyRecentUserReviews: RecentUserReviewsType = {
  totalReviews: 0,
  totalPages: 0,
  hasPreviousPage: false,
  hasNextPage: false,
  reviews: [],
}

export type {
  ReviewType,
  ReviewAlbumType,
  CommentType,
  RecentUserReviewsType,
};

export { 
  emptyReview,
  emptyReviewAlbum,
  emptyComment,
  emptyRecentUserReviews,
};