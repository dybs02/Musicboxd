import { emptyUser, type UserType } from "./user";


type CommentType = {
  _id: string;
  itemId: string;
  reviewId: string;
  userId: string;
  user: UserType;
  text: string;
  createdAt: string;
  updatedAt: string;
  likesCount: number;
  dislikesCount: number;
  userReaction: string;
}

const emptyComment: CommentType = {
  _id: "",
  itemId: "",
  reviewId: "",
  userId: "",
  user: emptyUser,
  text: "",
  createdAt: "",
  updatedAt: "",
  likesCount: 0,
  dislikesCount: 0,
  userReaction: "",
}

type CommentsPageType = {
  totalComments: number;
  totalPages: number;
  hasPreviousPage: boolean;
  hasNextPage: boolean;
  comments: CommentType[];
}

const emptyCommentsPage: CommentsPageType = {
  totalComments: 0,
  totalPages: 0,
  hasPreviousPage: false,
  hasNextPage: false,
  comments: [],
}

export type {
  CommentType,
  CommentsPageType,
};

export { 
  emptyComment,
  emptyCommentsPage,
};
