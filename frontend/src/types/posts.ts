import { emptyUser, type UserType } from "./user";

type PostType = {
  _id: string;
  content: string;
  user: UserType;
  createdAt: string;
  updatedAt: string;
  likesCount: number;
  dislikesCount: number;
  userReaction: string;
}

const emptyPost: PostType = {
  _id: "",
  content: "",
  user: emptyUser,
  createdAt: "",
  updatedAt: "",
  likesCount: 0,
  dislikesCount: 0,
  userReaction: "",
};

// const emptyPost: PostType = {
//   _id: "y4787y82w37gh8",
//   content: "Yooo just listend to this new track, it's fire!",
//   user: emptyUser,
//   createdAt: "2023-10-01T12:00:00Z",
//   updatedAt: "2023-10-01T12:00:00Z",
//   likesCount: 2,
//   dislikesCount: 1,
//   userReaction: "like",
// };

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