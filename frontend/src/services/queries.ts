import { gql } from "@apollo/client/core";


const GET_REWIEW_BY_ITEM_ID_USER_ID = gql`
  query Rewiew($itemId: String!, $userId: String!) {
    review(itemId: $itemId, userId: $userId) {
      value
      title
      description
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


export { 
  GET_REWIEW_BY_ITEM_ID_USER_ID,
  CREATE_UPDATE_REWIEW_BY_ITEM_ID,
};