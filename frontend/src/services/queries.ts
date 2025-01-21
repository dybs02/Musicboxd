import { gql } from "@apollo/client/core";


const GET_ALBUM_BY_ID = gql`
  query Album($id: String!) {
    album(id: $id) {
      total_tracks
      name
      release_date
      images {
        url
      }
      external_urls {
        spotify
      }
      artists {
        external_urls {
          spotify
        }
        name
        id
        href
        type
      }
      track_list {
        name
        duration_ms
        id
        external_urls {
          spotify
        }
      }
    }
  }
`;

const GET_REWIEW_BY_ITEM_ID_USER_ID = gql`
  query Rewiew($itemId: String!, $userId: String!) {
    review(itemId: $itemId, userId: $userId) {
      value
      title
      description
      comments {
        text
        createdAt
        updatedAt
        user {
          _id
          displayName
          images {
            url
          }
        }
      }
    }
  }
`;

const CREATE_UPDATE_REWIEW_BY_ITEM_ID = gql`
  mutation CreateOrUpdateReview($itemId: String!, $title: String, $description: String, $value: Int) {
    createOrUpdateReview(itemId: $itemId, title: $title, description: $description, value: $value) {
      value
      title
      description
    }
  }
`;

const ADD_COMMENT = gql`
  mutation AddComment($itemId: String!, $reviewId: String!, $text: String!) {
    addComment(itemId: $itemId, reviewId: $reviewId, text: $text) {
      text
      createdAt
      updatedAt
      user {
        _id
        displayName
        images {
          url
        }
      }
    }
  }
`;


export { 
  GET_ALBUM_BY_ID,
  GET_REWIEW_BY_ITEM_ID_USER_ID,
  CREATE_UPDATE_REWIEW_BY_ITEM_ID,
  ADD_COMMENT,
};