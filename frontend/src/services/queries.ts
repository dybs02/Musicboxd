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
      tracks {
        total
        items {
          name
          duration_ms
          id
          external_urls {
            spotify
          }
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
      likesCount
      dislikesCount
      userReaction
      user {
        _id
        displayName
        images {
          url
        }
      }
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
query RecentReviews($number: Int!, $itemType: String!) {
  recentReviews(number: $number, itemType: $itemType) {
    _id
    value
    itemId
    itemType
    title
    description
    userId
    user {
      images {
        url
      }
    }
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

const GET_USER_BY_ID = gql`
  query UserById($id: String!) {
    userById(id: $id) {
      _id
      country
      displayName
      href
      spotifyId
      role
      images {
        url
      }
      favouriteAlbums {
        key
        album {
          albumId
          name
          images {
            url
          }
          artists {
            id
            name
          }
        }
        }
    }
  }
`;

const SEARCH_FOR_ALBUMS = gql`
  query SearchForAlbum($query: String!) {
    search(query: $query, type: "album") {
      albums {
        items {
          images {
            url
          }
          artists {
            name
          }
          name
          id
        }
      }
    }
  }
`


const UPDATE_CURRENT_USER_FAVOURITE_ALBUM = gql`
  mutation UpdateCurrentUser($favouriteAlbum: FavouriteAlbumEntryInput) {
    updateCurrentUser(favouriteAlbum: $favouriteAlbum) {
      favouriteAlbums {
        key
        album {
          albumId
          name
          images {
            url
          }
          artists {
            id
            name
          }
        }
      }
    }
  }
`


const GET_RECENT_USER_REVIEWS_PAGINATION = gql`
  query RecentUserReviews($pageSize: Int, $page: Int!, $itemType: String!, $userId: String!) {
    recentUserReviews(page: $page, itemType: $itemType, userId: $userId, pageSize: $pageSize) {
      totalReviews
      totalPages
      hasPreviousPage
      hasNextPage
      reviews {
        value
        _id
        itemId
        itemType
        title
        userId
        createdAt
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
  }
`


const ADD_LIKE_OR_DISLIKE = gql`
  mutation AddLikeDislike($itemId: String!, $action: String!) {
    addLikeDislike(itemId: $itemId, action: $action) {
      likesCount
      dislikesCount
      userReaction
    }
  }
`


export { 
  GET_ALBUM_BY_ID,
  GET_REWIEW_BY_ITEM_ID_USER_ID,
  CREATE_UPDATE_REWIEW_BY_ITEM_ID,
  ADD_COMMENT,
  RECENT_REVIEWS,
  ALBUMS_BY_IDS,
  GET_TRACK_BY_ID,
  REPORT_COMMENT,
  GET_USER_BY_ID,
  SEARCH_FOR_ALBUMS,
  UPDATE_CURRENT_USER_FAVOURITE_ALBUM,
  GET_RECENT_USER_REVIEWS_PAGINATION,
  ADD_LIKE_OR_DISLIKE,
};