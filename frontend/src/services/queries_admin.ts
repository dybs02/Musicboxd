import { gql } from "@apollo/client/core";


const GET_REPORTED_COMMENTS = gql`
  query ReportedComments($number: Int) {
    reportedComments(number: $number) {
      _id
      comment {
        reviewId,
        user {
          _id
          displayName
          role
          images {
            url
          }
        }
        text,
        createdAt,
        updatedAt,
      }
      reportedByUser {
        _id
        displayName
        role
        images {
          url
        }
      }
      createdAt
    }
  }
`;

const RESOLE_REPORTED_COMMENT = gql`
  mutation ResolveComment($id: String!, $status: String!, $notes: String) {
    resolveComment(id: $id, status: $status, notes: $notes)
  }
`;

export { 
  GET_REPORTED_COMMENTS,
  RESOLE_REPORTED_COMMENT,
};