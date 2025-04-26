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

export { 
  GET_REPORTED_COMMENTS
};