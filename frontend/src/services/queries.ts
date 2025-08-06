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
      averageRating
      ratingCount
    }
  }
`;

const SEARCH = gql`
  query Search($type: String, $query: String!) {
    search(type: $type, query: $query) {
      tracks {
        items {
          album {
            images {
              url
            }
          }
          artists {
            name
          }
          name
          id
        }
      }
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
`;

const GET_REWIEW_ID_BY_ITEM_ID_USER_ID = gql`
  query Rewiew($itemId: String!, $userId: String!) {
    review(itemId: $itemId, userId: $userId) {
      _id
    }
  }
`;

const CREATE_UPDATE_REWIEW_BY_ITEM_ID = gql`
  mutation CreateOrUpdateReview($itemId: String!, $itemType: String!, $title: String, $description: String, $value: Int) {
    createOrUpdateReview(itemId: $itemId, itemType: $itemType, title: $title, description: $description, value: $value) {
      _id
      value
      title
      description
    }
  }
`;

const ADD_COMMENT = gql`
  mutation AddComment($itemId: String!, $itemType: String!, $text: String!, $replyingToId: String) {
    addComment(itemId: $itemId, itemType: $itemType, text: $text, replyingToId: $replyingToId) {
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
    track {
      name
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
      averageRating
      ratingCount
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
      trackReviewsNumber
      albumReviewsNumber
      isFollowing
      isFollower
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


const ADD_LIKE_OR_DISLIKE_REVIEW = gql`
  mutation AddLikeDislikeReview($reviewId: String!, $action: String!) {
    addLikeDislikeReview(reviewId: $reviewId, action: $action) {
      likesCount
      dislikesCount
      userReaction
    }
  }
`


const ADD_LIKE_OR_DISLIKE_COMMENT = gql`
  mutation AddLikeDislikeComment($commentId: String!, $action: String!) {
    addLikeDislikeComment(commentId: $commentId, action: $action) {
      likesCount
      dislikesCount
      userReaction
    }
  }
`


const ADD_LIKE_OR_DISLIKE_POST = gql`
  mutation addLikeDislikePost($postId: String!, $action: String!) {
    addLikeDislikePost(postId: $postId, action: $action) {
      likesCount
      dislikesCount
      userReaction
    }
  }
`


const GET_COMMENTS_BY_ITEM_ID = gql`
  query commentsPage(
    $itemId: String!
    $pageSize: Int
    $page: Int!
  ) {
    commentsPage(
      itemId: $itemId
      pageSize: $pageSize
      page: $page
    ) {
      totalComments
      totalPages
      hasPreviousPage
      hasNextPage
      comments {
        _id
        text
        itemId
        createdAt
        updatedAt
        user {
          _id
          displayName
          images {
            url
          }
        }
        likes
        likesCount
        dislikes
        dislikesCount
        userReaction
        repliesCount
      }
    }
  }
`


const GET_REWIEW_BY_ID = gql`
  query ReviewById($reviewId: String!) {
    reviewById(reviewId: $reviewId) {
      _id
      itemId
      value
      title
      itemType
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
    }
  }
`;

const GET_REWIEW_BY_ID_POST_LINK = gql`
  query ReviewById($reviewId: String!) {
    reviewById(reviewId: $reviewId) {
      _id
      value
      title
      itemType
      description
      user {
        _id
        displayName
        images {
          url
        }
      }
      album {
        images {
          url
        }
      }
    }
  }
`;


const GET_COMMENT_REPLIES = gql`
  query Replies($commentId: String!, $repliesLength: Int) {
    replies(commentId: $commentId, repliesLength: $repliesLength) {
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
      likes
      likesCount
      dislikes
      dislikesCount
      userReaction
      repliesCount
    }
}
`;


const ADD_POST = gql`
  mutation CreatePost($content: String!, $linkedReviewId: String) {
    createPost(content: $content, linkedReviewId: $linkedReviewId) {
      _id
      content
      createdAt
      updatedAt
      user {
        _id
        displayName
        images {
          url
        }
      }
      likesCount
      dislikesCount
      userReaction
    }
  }
`;


const GET_RECENT_POSTS = gql`
  query GetRecentPost($pageSize: Int, $page: Int!, $type: String) {
    getRecentPost(pageSize: $pageSize, page: $page, type: $type) {
      totalPosts
      totalPages
      hasPreviousPage
      hasNextPage
      posts {
        _id
        content
        createdAt
        updatedAt
        user {
          _id
          displayName
          images {
            url
          }
        }
        commentsNumber
        likesCount
        dislikesCount
        userReaction
        linkedReview {
          _id
          value
          title
          itemType
          description
          user {
            _id
            displayName
            images {
              url
            }
          }
          album {
            images {
              url
            }
          }
        }
      }
    }
  }
`;


const FOLLOW_USER = gql`
  mutation FollowUser($userId: String!) {
    followUser(userId: $userId) {
      isFollowing
      isFollower
    }
  }
`;


const UNFOLLOW_USER = gql`
  mutation UnfollowUser($userId: String!) {
    unfollowUser(userId: $userId) {
      isFollowing
      isFollower
    }
  }
`;


const SEND_MESSAGE = gql`
  mutation SendMessage($chatId: ID!, $content: String!) {
    sendMessage(chatId: $chatId, content: $content) {
      _id
      content
      senderId
      createdAt
    }
  }
`;


const CHAT_BY_PARTICIPANT = gql`
  query Chat($participantId: ID!) {
    chat(participantId: $participantId) {
      _id
      name
      createdAt
      participants {
        _id
        displayName
        images {
          url
        }
      }
    }
  }
`;


const CHAT_BY_ID = gql`
  query Chat($chatId: ID!) {
    chat(chatId: $chatId) {
      _id
      name
      createdAt
      participants {
        _id
        displayName
        images {
          url
        }
      }
    }
  }
`;


const CHAT_SUBSCRIPTION = gql`
  subscription MessageAdded($chatId: ID!) {
    messageAdded(chatId: $chatId) {
      _id
      content
      senderId
      createdAt
    }
  }
`;


const GET_MESSAGES_BY_CHAT_ID_PAGE = gql`
  query messagesPage(
    $chatId: ID!
    $pageSize: Int
    $page: Int!
  ) {
    messagesPage(
      chatId: $chatId
      pageSize: $pageSize
      page: $page
    ) {
      totalMessages
      totalPages
      hasPreviousPage
      hasNextPage
      messages {
        _id
        content
        senderId
        createdAt
      }
    }
  }
`


const NOTIFICATIONS_PAGE = gql`
  query NotificationsPage($pageSize: Int, $page: Int!) {
    notificationsPage(pageSize: $pageSize, page: $page) {
      totalNotifications
      totalPages
      hasPreviousPage
      hasNextPage
      notifications {
        _id
        message
        notifyingUserId
        notifiedUserId
        isRead
        createdAt
      }
    }
  }
`;


const MARK_NOTIFICATION_AS_READ = gql`
  mutation MarkNotificationAsRead($notificationId: ID!) {
    markNotificationAsRead(notificationId: $notificationId) {
      _id
    }
  }
`;


export { 
  GET_ALBUM_BY_ID,
  SEARCH,
  GET_REWIEW_ID_BY_ITEM_ID_USER_ID,
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
  ADD_LIKE_OR_DISLIKE_REVIEW,
  ADD_LIKE_OR_DISLIKE_COMMENT,
  ADD_LIKE_OR_DISLIKE_POST,
  GET_COMMENTS_BY_ITEM_ID,
  GET_REWIEW_BY_ID,
  GET_REWIEW_BY_ID_POST_LINK,
  GET_COMMENT_REPLIES,
  ADD_POST,
  GET_RECENT_POSTS,
  FOLLOW_USER,
  UNFOLLOW_USER,
  SEND_MESSAGE,
  CHAT_BY_PARTICIPANT,
  CHAT_SUBSCRIPTION,
  GET_MESSAGES_BY_CHAT_ID_PAGE,
  NOTIFICATIONS_PAGE,
  MARK_NOTIFICATION_AS_READ,
};