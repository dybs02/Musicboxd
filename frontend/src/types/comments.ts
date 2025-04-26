type CommentType = {
  _id: string;
  text: string;
  createdAt: string;
  updatedAt: string;
  user: {
    _id: string;
    displayName: string;
    images: {
      url: string;
    }[];
  };
};

type ReportedCommentType = {
  _id: string;
  commentId: string;
  comment: CommentType;
  reportedBy: string;
  // TODO add types for user and more
  reportedByUser: {
    _id: string;
    displayName: string;
    images: {
      url: string;
    }[];
  };
  status: string;
  createdAt: string;
  resolvedAt: string;
  moderatorId: string;
  moderatorNotes: string;
}


export type { 
  CommentType,
  ReportedCommentType,
};