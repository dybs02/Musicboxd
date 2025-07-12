import { ApolloClient, ApolloLink, InMemoryCache, concat, createHttpLink } from '@apollo/client/core';
import { GraphQLWsLink } from '@apollo/client/link/subscriptions';
import { createClient } from 'graphql-ws';
import { useAuthStore } from './authStore';
import { HttpLink, split } from '@apollo/client/core'
import { getMainDefinition } from '@apollo/client/utilities'

const httpLink = createHttpLink({
  // @ts-ignore
  uri: import.meta.env.VITE_GRAPHQL_API_URL,
  credentials: 'include',
});

// TODO redirect to login page if not authenticated
const wsLink = new GraphQLWsLink(
  createClient({
    // @ts-ignore
    url: import.meta.env.VITE_GRAPHQL_API_URL_WS,
    connectionParams: () => {
      const store = useAuthStore();
      const token = store.getJWT();
      if (!token) {
        return {};
      }
      return {
        Authorization: `${token}`,
      };
    },
  })
)


const authMiddleware = new ApolloLink((operation, forward) => {
  const store = useAuthStore();
  const token = store.getJWT();
  operation.setContext({
    headers: {
      Authorization: token ? `${token}` : "",
    },
  });
  return forward(operation);
});


const link = split(
  ({ query }) => {
    const definition = getMainDefinition(query)
    return (
      definition.kind === 'OperationDefinition'
      && definition.operation === 'subscription'
    )
  },
  wsLink,
  concat(authMiddleware, httpLink)
)

const apolloClient = new ApolloClient({
  link: link,
  cache: new InMemoryCache(),
  credentials: 'include',
});


export default apolloClient;