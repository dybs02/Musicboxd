import type { CommentType } from "./comments";
import type { ReviewType } from "./review";
import { emptyUser, type UserType } from "./user";

type PostType = {
  _id: string;
  content: string;
  user: UserType;
  createdAt: string;
  updatedAt: string;
  comments: CommentType[];
  commentsNumber: number;
  likesCount: number;
  dislikesCount: number;
  userReaction: string;
  linkedReviewId: string;
  linkedReview?: ReviewType;
}

const emptyPost: PostType = {
  _id: "",
  content: "",
  user: emptyUser,
  createdAt: "",
  updatedAt: "",
  comments: [],
  commentsNumber: 0,
  likesCount: 0,
  dislikesCount: 0,
  userReaction: "",
  linkedReviewId: "",
  linkedReview: undefined,
};

type RecentPostsType = {
  totalPosts: number;
  totalPages: number;
  hasPreviousPage: boolean;
  hasNextPage: boolean;
  posts: PostType[];
}

const emptyRecentPosts: RecentPostsType = {
  totalPosts: 0,
  totalPages: 0,
  hasPreviousPage: false,
  hasNextPage: false,
  posts: [],
};

export type {
  PostType,
  RecentPostsType,
}

export {
  emptyPost,
  emptyRecentPosts,
};