import { gql } from "@apollo/client/core";


const GET_REWIEW_BY_ITEM_ID_USER_ID = gql`
  query Rewiew($itemId: String!, $userId: String!) {
    review(itemId: $itemId, userId: $userId) {
      value
      title
      description
    }
  }
`


export { GET_REWIEW_BY_ITEM_ID_USER_ID };