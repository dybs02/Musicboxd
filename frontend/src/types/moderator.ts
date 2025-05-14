import { emptyComment, type CommentType } from "./comments"
import { emptyUser, type UserType } from "./user"


type ReportedCommentType = {
  _id: string,
  commentId: string
  comment: CommentType
  reportedBy: string
  reportedByUser: UserType
  status: string
  createdAt: string
  resolvedAt: string
  moderatorId: string
  moderatorNotes: string
}

const emptyReportedComment: ReportedCommentType = {
  _id: "",
  commentId: "",
  comment: emptyComment,
  reportedBy: "",
  reportedByUser: emptyUser,
  status: "",
  createdAt: "",
  resolvedAt: "",
  moderatorId: "",
  moderatorNotes: ""
}

export type {
  ReportedCommentType,
};

export {
  emptyReportedComment,
};