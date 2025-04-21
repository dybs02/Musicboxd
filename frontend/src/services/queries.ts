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
        _id
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
  mutation CreateOrUpdateReview($itemId: String!, $itemType: String!, $title: String, $description: String, $value: Int) {
    createOrUpdateReview(itemId: $itemId, itemType: $itemType, title: $title, description: $description, value: $value) {
      value
      title
      description
    }
  }
`;

const ADD_COMMENT = gql`
  mutation AddComment($itemId: String!, $reviewId: String!, $text: String!) {
    addComment(itemId: $itemId, reviewId: $reviewId, text: $text) {
      _id
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

const RECENT_REVIEWS = gql`
query RecentReviews($number: Int!) {
  recentReviews(number: $number) {
    _id
    value
    itemId
    title
    description
    userId
    album {
      name
      images {
        url
      }
      artists {
        name
      }
    }
  }
}
`;

const ALBUMS_BY_IDS = gql`
query AlbumsByIds($ids: [String!]!) {
  albumsByIds(ids: $ids) {
    id
    name
    images {
      url
    }
    external_urls {
      spotify
    }
    artists {
      id
      name
    }
  }
}
`;

const GET_TRACK_BY_ID = gql`
  query Track($id: String!) {
    track(id: $id) {
      duration_ms
      href
      id
      name
      popularity
      track_number
      external_urls {
        spotify
      }
      artists {
        external_urls {
          spotify
        }
        name
        id
      }
      album {
        images {
          url
        }
        name
        id
        total_tracks
        release_date
        external_urls {
          spotify
        }
      }
    }
  }
`;

const REPORT_COMMENT = gql`
  mutation ReportComment($id: String!) {
    reportComment(id: $id)
  }
`;

export { 
  GET_ALBUM_BY_ID,
  GET_REWIEW_BY_ITEM_ID_USER_ID,
  CREATE_UPDATE_REWIEW_BY_ITEM_ID,
  ADD_COMMENT,
  RECENT_REVIEWS,
  ALBUMS_BY_IDS,
  GET_TRACK_BY_ID,
  REPORT_COMMENT,
};